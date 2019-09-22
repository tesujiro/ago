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
		findex, err := strictToInt(index)
		if err != nil {
			return nil, fmt.Errorf("field index can not convert to int :%v", err)
		}
		if field, err := env.GetField(findex); err != nil {
			return nil, err
		} else {
			return field, nil
		}
	case *ast.NumExpr:
		lit := expr.Literal
		if strings.Contains(lit, ".") || strings.Contains(lit, "e") {
			if f, err := strconv.ParseFloat(lit, 64); err != nil {
				return 0.0, err
			} else {
				return f, nil
			}
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
		index, err = getHashIndex(env, expr.Index)
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
				return -1 * toFloat64(val), nil
			case reflect.Float64, reflect.Float32:
				return -1 * toFloat64(val), nil
			}
		case "!":
			return !toBool(val), nil
		}
	case *ast.AssExpr:
		left, right := expr.Left, expr.Right

		// evaluate right expressions
		right_values := make([]interface{}, len(right))
		var err error
		for i, expr := range right {
			right_values[i], err = evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
		}

		// evaluate assExpr
		switch {
		case len(left) == 1 && len(right) == 1:
			return evalAssExpr(left[0], right_values[0], env)
		case len(left) > 1 && len(right) == 1:
			val := right_values[0]
			//if reflect.ValueOf(val).Kind() == reflect.Interface {
			//val = reflect.ValueOf(val).Elem().Interface()
			//}
			if reflect.ValueOf(val).Kind() != reflect.Slice {
				return nil, errors.New("single value assign to multi values")
			} else {
				elements := reflect.ValueOf(val)
				right_values = make([]interface{}, elements.Len())
				for i := 0; i < elements.Len(); i++ {
					right_values[i] = elements.Index(i).Interface()
				}
			}
			fallthrough
		default:
			for i, expr := range left {
				if i >= len(right_values) {
					return right_values[len(right_values)-1], nil
				}
				if _, err := evalAssExpr(expr, right_values[i], env); err != nil {
					return nil, err
				}
			}
			return right_values[len(left)-1], nil
		}
	case *ast.CompExpr:
		left := expr.Left
		right := expr.Right
		operator := expr.Operator
		before_val, err := evalExpr(left, env)
		if err != nil {
			return nil, err
		}

		if operator == "++" || operator == "--" {
			right = &ast.NumExpr{Literal: "1"}
		}
		result, err := evalExpr(&ast.BinOpExpr{Left: left, Operator: operator[0:1], Right: right}, env)
		if err != nil {
			return nil, err
		}

		after_val, err := evalAssExpr(left, result, env)
		if err != nil {
			return nil, err
		}
		if expr.After {
			return before_val, nil
		} else {
			return after_val, nil
		}

	case *ast.TriOpExpr:
		cond, err := evalExpr(expr.Cond, env)
		if err != nil {
			return nil, err
		}
		cond_b, err := strictToBool(cond)
		if err != nil {
			return nil, fmt.Errorf("convert ternary operator:%v", err)
		}
		if cond_b {
			return evalExpr(expr.Then, env)
		} else {
			return evalExpr(expr.Else, env)
		}
	case *ast.ContainKeyExpr:
		key, err := evalExpr(expr.KeyExpr, env)
		if err != nil {
			return nil, err
		}
		k := toString(key)

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
		case "||":
			left_b, err := strictToBool(left)
			if err != nil {
				return nil, fmt.Errorf("convert left expression of OR perator:%v", err)
			}
			if left_b {
				return true, nil
			}
			right_b, err := strictToBool(right)
			if err != nil {
				return nil, fmt.Errorf("convert right expression of OR perator:%v", err)
			}
			if right_b {
				return true, nil
			}
			return false, nil
		case "&&":
			left_b, err := strictToBool(left)
			if err != nil {
				return nil, fmt.Errorf("convert left expression of AND perator:%v", err)
			}
			if !left_b {
				return false, nil
			}
			right_b, err := strictToBool(right)
			if err != nil {
				return nil, fmt.Errorf("convert right expression of AND perator:%v", err)
			}
			if right_b {
				return true, nil
			}
			return false, nil
		case "==":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.String && r_kind == reflect.String:
				return left == right, nil
			case l_kind == reflect.Float64 || r_kind == reflect.Float64:
				return toFloat64(left) == toFloat64(right), nil
			case l_kind == reflect.Int || r_kind == reflect.Int:
				return toString(left) == toString(right), nil
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				return toString(left) == toString(right), nil
			}
		case "!=":
			//return left != right, nil
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.String && r_kind == reflect.String:
				return left != right, nil
			case l_kind == reflect.Float64 || r_kind == reflect.Float64:
				return toFloat64(left) != toFloat64(right), nil
			case l_kind == reflect.Int || r_kind == reflect.Int:
				return toString(left) != toString(right), nil
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				return toString(left) != toString(right), nil
			}
		case ">":
			return toFloat64(left) > toFloat64(right), nil
		case ">=":
			return toFloat64(left) >= toFloat64(right), nil
		case "<":
			return toFloat64(left) < toFloat64(right), nil
		case "<=":
			return toFloat64(left) <= toFloat64(right), nil
		case "CAT":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.String || r_kind == reflect.String:
				return toString(left) + toString(right), nil
			case l_kind == reflect.Int || r_kind == reflect.Int:
				return toString(left) + toString(right), nil
			case l_kind == reflect.Float64 || r_kind == reflect.Float64:
				return toFloat64(left) + toFloat64(right), nil
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				return toString(left) + toString(right), nil
			}
		case "+":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				return toFloat64(left) + toFloat64(right), nil
			}
		case "-":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				return toFloat64(left) - toFloat64(right), nil
			}
		case "*":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				return toFloat64(left) * toFloat64(right), nil
			}
		case "/":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			if right == 0 {
				return nil, fmt.Errorf("devision by zero")
			}
			switch {
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				return toFloat64(left) / toFloat64(right), nil
			}
		case "%":
			//return toInt(left) % toInt(right), nil
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			if right == 0 {
				return nil, fmt.Errorf("devision by zero")
			}
			switch {
			case l_kind == reflect.Map || r_kind == reflect.Map:
				return nil, fmt.Errorf("can't read value of array")
			default:
				q := int(toFloat64(left) / toFloat64(right))
				return toFloat64(left) - toFloat64(right)*float64(q), nil
			}
		}
	case *ast.MatchExpr:
		val, err := evalExpr(expr.Expr, env)
		if err != nil {
			return nil, err
		}
		s := toString(val)
		re := expr.RegExpr.(*ast.RegExpr).Literal
		return regexp.MatchString(re, s)
	case *ast.GetlineExpr:
		var redir string
		if expr.Command != nil {
			command_interface, err := evalExpr(expr.Command, env)
			if err != nil {
				return nil, err
			}
			commandLine := command_interface.(string)
			redir = commandLine
			//fmt.Println("command=", commandLine)
			_, err = env.GetScanner(redir)
			if err == ErrUnknownSymbol {
				re := regexp.MustCompile("[ \t]+")
				cmd_array := re.Split(commandLine, -1)
				cmd := exec.Command(cmd_array[0], cmd_array[1:]...)
				if stdout, err := cmd.StdoutPipe(); err != nil {
					return nil, err
				} else {
					_, err = env.SetFile(redir, &stdout)
					if err != nil {
						return 0, err
					}
				}
				if err := cmd.Start(); err != nil {
					return nil, err
				}
			} else if err != nil {
				return nil, err
			}
		} else {
			if expr.Redir != nil {
				redir_interface, err := evalExpr(expr.Redir, env)
				if err != nil {
					return nil, err
				}
				redir = (redir_interface).(string)
			} else {
				redir = "-" // Stdin
			}
		}

		var scanner *bufio.Scanner
		var err error
		scanner, err = env.GetScanner(redir)
		if err == ErrUnknownSymbol {
			// Open File if not opened yet.
			if f, err := os.Open(redir); err != nil {
				return 0, err
			} else {
				rc := io.ReadCloser(f)
				scanner, err = env.SetFile(redir, &rc)
				if err != nil {
					return 0, err
				}
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
			if err := env.SetFieldFromLine(line); err != nil {
				fmt.Printf("error:%v\n", err)
				return 0, err
			}
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
		if env_val, err := env.Get(id); err == nil {
			if reflect.TypeOf(env_val).Kind() == reflect.Map {
				return nil, fmt.Errorf("can't assign to %v; it's an associated array name.", id)
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
		i_val, err := evalExpr(lexp.Expr, env)
		if err != nil {
			//fmt.Println("fieldExpr index error") //TODO
			return nil, err
		}
		/*
			index, ok := i_val.(int)
			if !ok {
				return nil, fmt.Errorf("field index not int :%v", reflect.TypeOf(i_val))
			}
		*/
		index_f, err := strictToFloat(i_val)
		if err != nil {
			return nil, fmt.Errorf("field index cannot convert to int :%v", err)
		}
		index := int(index_f)

		if val_int, ok := val.(int); ok {
			val = fmt.Sprintf("%v", val_int)
		}
		//fmt.Printf("evalAssExpr FieldExpr: index:%v \tval:%v\n", index, val) //TODO
		val_string, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("field value is not string :%v", reflect.TypeOf(val))
		}

		err = env.SetField(index, val_string)
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

		index, err := getHashIndex(env, lexp.Index)
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
