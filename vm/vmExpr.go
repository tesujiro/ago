package vm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/tesujiro/ago/ast"
)

func evalExpr(expr ast.Expr, env *Env) (interface{}, error) {
	switch expr := expr.(type) {
	case *ast.IdentExpr:
		v, err := env.Get(expr.Literal)
		if err == ErrUnknownSymbol {
			val, err := env.DefineDefaultValue(expr.Literal)
			if err != nil {
				return nil, err
			}
			return val, nil
		} else if err != nil {
			return nil, err
		}
		return v, nil
	case *ast.FieldExpr:
		index, err := evalExpr(expr.Expr, env)
		if err != nil {
			return nil, err
		}
		findex, err := env.strictToInt(index)
		if err != nil {
			return nil, fmt.Errorf("field index cannot convert to int :%v", err)
		}
		field, err := env.GetField(findex)
		if err != nil {
			return nil, err
		}
		return field, nil
	case *ast.NumExpr:
		lit := expr.Literal
		lit = strings.Replace(lit, "E", "e", 1)
		if strings.Contains(lit, ".") || strings.Contains(lit, "e") {
			f, err := strconv.ParseFloat(lit, 64)
			if err != nil {
				return 0.0, err
			}
			return f, nil
		}
		var i int64
		var err error
		if strings.HasPrefix(lit, "0x") {
			i, err = strconv.ParseInt(lit[2:], 16, 0)
		} else {
			i, err = strconv.ParseInt(lit, 10, 0)
		}
		if err != nil {
			return 0, err
		}
		return int(i), nil

	case *ast.StringExpr:
		return expr.Literal, nil
	case *ast.ConstExpr:
		switch expr.Literal {
		case "true":
			return true, nil
		case "false":
			return false, nil
		case "nil":
			return nil, nil
		}
	case *ast.FuncExpr:
		return (defineFunc(expr, env))
	case *ast.CallExpr:
		//fmt.Printf("CallExpr env:%v builtin.field:%#v\n", env, env.builtin.field)
		return (callFunc(expr, env))
	case *ast.AnonymousCallExpr:
		return (callAnonymousFunc(expr, env))
	case *ast.ParentExpr:
		return evalExpr(expr.SubExpr, env)
	case *ast.ItemExpr:
		var value, index interface{}
		var err error
		// index
		index, err = env.getHashIndex(expr.Index)
		if err != nil {
			return nil, err
		}
		// value
		ie, ok := expr.Expr.(*ast.IdentExpr)
		if ok {
			id := ie.Literal
			value, err = env.Get(id)
			if err == ErrUnknownSymbol {
				v, err := env.DefineDefaultMapValue(id, index)
				if err != nil {
					return nil, err
				}
				value = v
			} else if err != nil {
				return nil, err
			}
		} else {
			value, err = evalExpr(expr.Expr, env)
			if err != nil {
				return nil, err
			}
		}

		// TODO:Elem()

		switch reflect.ValueOf(value).Kind() {
		/*
			case reflect.Slice, reflect.Array:
				// index change to int
				if i, ok := index.(int); !ok {
					return nil, errors.New("index cannot convert to int")
				} else {
					if i < 0 || reflect.ValueOf(value).Len() <= i {
						return nil, errors.New("index out of range")
					}
					return reflect.ValueOf(value).Index(i).Interface(), nil
				}
		*/
		case reflect.Map:
			m, ok := value.(map[interface{}]interface{})
			if !ok {
				return nil, errors.New("value cannot convert to map")
			}
			v, ok := m[index]
			if !ok {
				defaultValue := env.GetDefaultValue()
				m[index] = defaultValue
				return defaultValue, nil
			}
			return v, nil

		default:
			return nil, errors.New("type " + reflect.ValueOf(value).Kind().String() + " does not support index operation")
		}
	case *ast.UnaryExpr:
		var val interface{}
		var err error
		if val, err = evalExpr(expr.Expr, env); err != nil {
			return nil, err
		}
		switch expr.Operator {
		case "+":
			return val, nil
		case "-":
			kind := reflect.ValueOf(val).Kind()
			switch kind {
			case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
				return -1 * env.toFloat64(val), nil
			case reflect.Float64, reflect.Float32:
				return -1 * env.toFloat64(val), nil
			}
		case "!":
			return !env.toBool(val), nil
		}
	case *ast.AssExpr:
		left, right := expr.Left, expr.Right

		// evaluate right expressions
		rightValues := make([]interface{}, len(right))
		var err error
		for i, expr := range right {
			rightValues[i], err = evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
		}

		// evaluate assExpr
		switch {
		case len(left) == 1 && len(right) == 1:
			return evalAssExpr(left[0], rightValues[0], env)
		case len(left) > 1 && len(right) == 1:
			val := rightValues[0]
			//if reflect.ValueOf(val).Kind() == reflect.Interface {
			//val = reflect.ValueOf(val).Elem().Interface()
			//}
			if reflect.ValueOf(val).Kind() != reflect.Slice {
				return nil, errors.New("single value assign to multi values")
			}
			elements := reflect.ValueOf(val)
			rightValues = make([]interface{}, elements.Len())
			for i := 0; i < elements.Len(); i++ {
				rightValues[i] = elements.Index(i).Interface()
			}
			fallthrough
		default:
			for i, expr := range left {
				if i >= len(rightValues) {
					return rightValues[len(rightValues)-1], nil
				}
				if _, err := evalAssExpr(expr, rightValues[i], env); err != nil {
					return nil, err
				}
			}
			return rightValues[len(left)-1], nil
		}
	case *ast.CompExpr:
		left := expr.Left
		right := expr.Right
		operator := expr.Operator
		beforeVal, err := evalExpr(left, env)
		if err != nil {
			return nil, err
		}
		// Check beforeValue is Number
		kind := reflect.ValueOf(beforeVal).Kind()
		if kind != reflect.Int && kind != reflect.Float64 {
			beforeVal = env.GetDefaultNumberValue()
		}

		if operator == "++" || operator == "--" {
			right = &ast.NumExpr{Literal: "1"}
		}
		result, err := evalExpr(&ast.BinOpExpr{Left: left, Operator: operator[0:1], Right: right}, env)
		if err != nil {
			return nil, err
		}

		afterVal, err := evalAssExpr(left, result, env)
		if err != nil {
			return nil, err
		}
		if expr.After {
			return beforeVal, nil
		}
		return afterVal, nil

	case *ast.TriOpExpr:
		cond, err := evalExpr(expr.Cond, env)
		if err != nil {
			return nil, err
		}
		boolCond, err := env.strictToBool(cond)
		if err != nil {
			return nil, fmt.Errorf("convert ternary operator:%v", err)
		}
		if boolCond {
			return evalExpr(expr.Then, env)
		}
		return evalExpr(expr.Else, env)
	case *ast.ContainKeyExpr:
		key, err := evalExpr(expr.KeyExpr, env)
		if err != nil {
			return nil, err
		}
		k := env.toString(key)

		mapID := expr.MapID
		mapInterface, err := env.Get(mapID)
		if err == ErrUnknownSymbol {
			v, err := env.DefineDefaultMap(mapID)
			if err != nil {
				return nil, err
			}
			mapInterface = v
		} else if err != nil {
			return nil, err
		}
		m, _ := mapInterface.(map[interface{}]interface{})

		_, ok := m[k]
		return ok, nil
	case *ast.BinOpExpr:
		var left, right interface{}
		var err error
		if left, err = evalExpr(expr.Left, env); err != nil {
			return nil, err
		}
		if right, err = evalExpr(expr.Right, env); err != nil {
			return nil, err
		}
		switch expr.Operator {
		case "||", "&&":
			return evalBoolOp(env, expr.Operator, left, right)
		case "==", "!=":
			return compareEqual(env, expr.Operator, left, right)
		case ">", ">=", "<", "<=":
			return compareInequal(env, expr.Operator, left, right)
		case "CAT":
			lKind := reflect.ValueOf(left).Kind()
			rKind := reflect.ValueOf(right).Kind()
			if lKind == reflect.Map || rKind == reflect.Map {
				return nil, fmt.Errorf("can't read value of array")
			}
			return env.toString(left) + env.toString(right), nil
		case "+", "-", "*", "/", "%", "^":
			return evalArithOp(env, expr.Operator, left, right)
		}
	case *ast.MatchExpr:
		val, err := evalExpr(expr.Expr, env)
		if err != nil {
			return nil, err
		}
		s := env.toString(val)
		//re := expr.RegExpr.(*ast.ConstExpr).Literal
		re, err := evalExpr(expr.RegExpr, env)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("re=%#v\n", re)
		//fmt.Printf("re=%v\n", toString(re))
		return regexp.MatchString(env.toString(re), s)
	case *ast.GetlineExpr:
		var redir string
		if expr.Command != nil {
			cmdInterface, err := evalExpr(expr.Command, env)
			if err != nil {
				return nil, err
			}
			commandLine := cmdInterface.(string)
			redir = commandLine
			//fmt.Println("command=", commandLine)
			_, err = env.GetScanner(redir)
			if err == ErrUnknownSymbol {
				re := regexp.MustCompile("[ \t]+")
				cmdArray := re.Split(commandLine, -1)
				cmd := exec.Command(cmdArray[0], cmdArray[1:]...)
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					return nil, err
				}
				_, err = env.SetFile(redir, &stdout)
				if err != nil {
					return nil, err
				}
				err = cmd.Start()
				if err != nil {
					return nil, err
				}
			} else if err != nil {
				return nil, err
			}
		} else {
			if expr.Redir != nil {
				redirInterface, err := evalExpr(expr.Redir, env)
				if err != nil {
					return nil, err
				}
				redir = (redirInterface).(string)
			} else {
				redir = "-" // Stdin
			}
		}

		var scanner *bufio.Scanner
		var err error
		scanner, err = env.GetScanner(redir)
		if err == ErrUnknownSymbol {
			// Open File if not opened yet.
			f, err := os.Open(redir)
			if err != nil {
				return 0, err
			}
			rc := io.ReadCloser(f)
			scanner, err = env.SetFile(redir, &rc)
			if err != nil {
				return 0, err
			}
		} else if err != nil {
			return 0, err
		}
		b := scanner.Scan()
		if !b {
			return 0, nil
		}
		line := scanner.Text()
		if expr.Var == nil {
			env.SetFieldFromLine(line)
		} else {
			evalAssExpr(expr.Var, (interface{})(line), env)
		}
		return 1, nil
	}
	return 0, nil
}

func evalAssExpr(lexp ast.Expr, val interface{}, env *Env) (interface{}, error) {
	switch lexp := lexp.(type) {
	case *ast.IdentExpr:
		id := lexp.Literal
		// Check the type of id in env for Safety
		if envVal, err := env.Get(id); err == nil {
			if reflect.TypeOf(envVal).Kind() == reflect.Map {
				return nil, fmt.Errorf("can't assign to %v; it's an associated array name", id)
			}
		}
		if err := env.Set(id, val); err == ErrUnknownSymbol {
			if err := env.Define(id, val); err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		return val, nil
	case *ast.FieldExpr:
		iVal, err := evalExpr(lexp.Expr, env)
		if err != nil {
			//fmt.Println("fieldExpr index error") //TODO
			return nil, err
		}
		/*
			index, ok := iVal.(int)
			if !ok {
				return nil, fmt.Errorf("field index not int :%v", reflect.TypeOf(iVal))
			}
		*/
		indexF, err := env.strictToFloat(iVal)
		if err != nil {
			return nil, fmt.Errorf("field index cannot convert to int :%v", err)
		}
		index := int(indexF)

		if valInt, ok := val.(int); ok {
			val = fmt.Sprintf("%v", valInt)
		} else if valFloat, ok := val.(float64); ok {
			ofmt, _ := env.Get("OFMT")
			val = fmt.Sprintf(ofmt.(string), valFloat)
		}
		//fmt.Printf("evalAssExpr FieldExpr: index:%v \tval:%v\n", index, val) //TODO
		valString, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("field value is not string :%v", reflect.TypeOf(val))
		}

		err = env.SetField(index, valString)
		if err != nil {
			//fmt.Println("fieldExpr SetField error") //TODO
			return nil, err
		}
		return nil, nil
	case *ast.ItemExpr:
		ie, ok := lexp.Expr.(*ast.IdentExpr)
		if !ok {
			return nil, errors.New("invalid assignment")
		}
		id := ie.Literal

		index, err := env.getHashIndex(lexp.Index)
		if err != nil {
			return nil, err
		}

		value, err := env.Get(id)
		if err == ErrUnknownSymbol {
			v, err := env.DefineDefaultMapValue(id, index)
			if err != nil {
				return nil, err
			}
			value = v
		} else if err != nil {
			return nil, err
		}

		switch reflect.TypeOf(value).Kind() {
		/*
			case reflect.Slice | reflect.Array:
				if i, ok := index.(int); !ok {
					return nil, errors.New("index cannot convert to int")
				} else {
					if i < 0 || reflect.ValueOf(value).Len() < i {
						return nil, errors.New("index out of range")
					}
					if i == reflect.ValueOf(value).Len() {
						// append val to array
						ar := reflect.Append(reflect.ValueOf(value), reflect.ValueOf(val)).Interface()
						return evalAssExpr(lexp.Value, ar, env)
					}

					// Set Val To Array
					reflect.ValueOf(value).Index(i).Set(reflect.ValueOf(val))
					return val, nil
				}
		*/
		case reflect.Map:
			m, ok := value.(map[interface{}]interface{})
			if !ok {
				return nil, errors.New("value cannot convert to map")
			}
			m[index] = val
			//fmt.Printf("vmExpr evalAssExpr ItemExpr reflect.Map index:%#v val:%#v\n", index, val)
			return val, nil
		default:
			return nil, errors.New("type " + reflect.TypeOf(value).Kind().String() + " does not support index operation")
		}

	default:
		// TODO:?
		return nil, errors.New("invalid operation")
	}
	//return val, nil
}
