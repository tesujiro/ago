package vm

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/tesujiro/goa/ast"
)

func toInt(val interface{}) int {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Float64, reflect.Float32:
		return int(val.(float64))
	}
	i, _ := val.(int)
	return i
}

func toFloat64(val interface{}) float64 {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return float64(val.(int))
	}
	f, _ := val.(float64)
	return f
}

func toBool(val interface{}) bool {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Bool:
		return val.(bool)
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return val.(int) != 0
	case reflect.Float32, reflect.Float64:
		return val.(float64) != 0
	case reflect.String:
		return val.(string) != ""
	default:
		return true
	}
}

func toString(val interface{}) string {
	switch reflect.ValueOf(val).Kind() {
	case reflect.String:
		return val.(string)
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return fmt.Sprintf("%v", val)
	case reflect.Float64, reflect.Float32:
		return fmt.Sprintf("%v", val)
	default:
		return ""
	}
	s, _ := val.(string)
	return s
}

func evalExpr(expr ast.Expr, env *Env) (interface{}, error) {
	switch expr.(type) {
	case *ast.IdentExpr:
		id := expr.(*ast.IdentExpr).Literal
		v, err := env.Get(id)
		if err == ErrUnknownSymbol {
			val, err := env.DefineDefaultValue(id)
			if err != nil {
				return nil, err
			}
			return val, nil
		} else if err != nil {
			return nil, err
		}
		return v, nil
	case *ast.FieldExpr:
		expr := expr.(*ast.FieldExpr).Expr
		index, err := evalExpr(expr, env)
		if err != nil {
			return nil, err
		}
		i, _ := index.(int)
		if field, err := env.GetField(i); err != nil {
			return nil, err
		} else {
			return field, nil
		}
	case *ast.NumExpr:
		lit := expr.(*ast.NumExpr).Literal
		//fmt.Println("NumExpr:", lit)
		if strings.Contains(lit, ".") {
			if f, err := strconv.ParseFloat(lit, 64); err != nil {
				return 0.0, err
			} else {
				return f, nil
			}
		}
		//if i, err := strconv.ParseInt(lit, 10, 64); err != nil {
		if i, err := strconv.ParseInt(lit, 10, 0); err != nil {
			return 0, err
		} else {
			//fmt.Println("==> return :", int(i))
			return int(i), nil
		}
	case *ast.StringExpr:
		str := expr.(*ast.StringExpr).Literal
		return str, nil
	case *ast.ConstExpr:
		switch expr.(*ast.ConstExpr).Literal {
		case "true":
			return true, nil
		case "false":
			return false, nil
		case "nil":
			return nil, nil
		}
	case *ast.FuncExpr:
		return (defineFunc(expr.(*ast.FuncExpr), env))
	case *ast.CallExpr:
		return (callFunc(expr.(*ast.CallExpr), env))
	case *ast.LenExpr:
		sub := expr.(*ast.LenExpr).Expr
		result, err := evalExpr(sub, env)
		if err != nil {
			return nil, err
		}
		switch reflect.ValueOf(result).Kind() {
		case reflect.Slice, reflect.Array, reflect.String, reflect.Map:
			return reflect.ValueOf(result).Len(), nil
		default:
			return nil, fmt.Errorf("type %s does not support len operation", reflect.ValueOf(result).Kind().String())
		}
	case *ast.AnonymousCallExpr:
		return (callAnonymousFunc(expr.(*ast.AnonymousCallExpr), env))
	case *ast.ParentExpr:
		sub := expr.(*ast.ParentExpr).SubExpr
		return evalExpr(sub, env)
	case *ast.ArrayExpr:
		exprs := expr.(*ast.ArrayExpr).Exprs
		array := make([]interface{}, len(exprs))
		for i, expr := range exprs {
			if val, err := evalExpr(expr, env); err != nil {
				fmt.Println("ArrayExpr error at index:", i) // TODO:
				return nil, err
			} else {
				array[i] = val
			}
		}
		return array, nil
	case *ast.ItemExpr:
		var index string
		for i, expr := range expr.(*ast.ItemExpr).Index {
			//fmt.Printf("Index[%v]:%v\n", k, expr)
			val, err := evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
			if i == 0 {
				index = fmt.Sprintf("%v", val)
			} else {
				index = fmt.Sprintf("%v%v%v", index, env.builtin.SUBSEP, val)
			}
		}
		id := expr.(*ast.ItemExpr).Literal
		//fmt.Printf("ItemExpr\tid:%v\tindex:%v\n", id, index)
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
				//return nil, nil
				defaultValue := env.GetDefaultValue()
				newMap := make(map[interface{}]interface{}, len(m)+1)
				newMap[index] = defaultValue
				for k, v := range m {
					newMap[k] = v
				}

				err := env.Set(id, newMap)
				if err != nil {
					return nil, err
				}
				return defaultValue, nil
			}
			return v, nil

		default:
			return nil, errors.New("type " + reflect.ValueOf(value).Kind().String() + " does not support index operation")
		}
	case *ast.MapExpr:
		exprs := expr.(*ast.MapExpr).MapExpr
		m := make(map[interface{}]interface{}, len(exprs))
		//fmt.Println("map len:", len(exprs))
		var keyResult, valResult interface{}
		var err error
		for keyExpr, valExpr := range exprs {
			keyResult, err = evalExpr(keyExpr, env)
			if err != nil {
				return nil, err
			}
			valResult, err = evalExpr(valExpr, env)
			if err != nil {
				return nil, err
			}
			m[keyResult] = valResult
		}
		return m, nil

	case *ast.UnaryExpr:
		var val interface{}
		var err error
		if val, err = evalExpr(expr.(*ast.UnaryExpr).Expr, env); err != nil {
			return nil, err
		}
		switch expr.(*ast.UnaryExpr).Operator {
		case "+":
			return val, nil
		case "-":
			kind := reflect.ValueOf(val).Kind()
			switch kind {
			case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
				return -1 * toInt(val), nil
			case reflect.Float64, reflect.Float32:
				return -1 * toFloat64(val), nil
			}
		case "!":
			return !toBool(val), nil
		}
	case *ast.CompExpr:
		/*
			var left, right interface{}
			var err error
			if left, err = evalExpr(expr.(*ast.BinOpExpr).Left, env); err != nil {
				return nil, err
			}
			ltype := reflect.TypeOf(left)
			rtype := reflect.TypeOf(right)
			lvalue := reflect.ValueOf(left)
			rvalue := reflect.ValueOf(right)
		*/
		left := expr.(*ast.CompExpr).Left
		right := expr.(*ast.CompExpr).Right
		operator := expr.(*ast.CompExpr).Operator
		switch operator {
		case "++":
			if ident, ok := left.(*ast.IdentExpr); ok {
				v, err := env.Get(ident.Literal)
				if err == ErrUnknownSymbol {
					val, err := env.DefineDefaultValue(ident.Literal)
					if err != nil {
						return nil, err
					}
					v = val
				} else if err != nil {
					return nil, err
				}
				switch reflect.TypeOf(v).Kind() {
				case reflect.Int, reflect.Int32, reflect.Int64:
					v = toInt(v) + 1
				case reflect.Float32, reflect.Float64:
					v = toFloat64(v) + 1.0
				case reflect.String:
					v = 1
				default:
					return nil, errors.New("Invalid operation")
				}
				if err := env.Set(ident.Literal, v); err == ErrUnknownSymbol {
					if err := env.Define(ident.Literal, v); err != nil {
						return nil, err
					}
				} else if err != nil {
					return nil, err
				}
				return v, nil

			} else {
				return nil, errors.New("Invalid operation")
			}
		case "--":
			if ident, ok := left.(*ast.IdentExpr); ok {
				v, err := env.Get(ident.Literal)
				if err == ErrUnknownSymbol {
					val, err := env.DefineDefaultValue(ident.Literal)
					if err != nil {
						return nil, err
					}
					v = val
				} else if err != nil {
					return nil, err
				}
				switch reflect.TypeOf(v).Kind() {
				case reflect.Int, reflect.Int32, reflect.Int64:
					v = toInt(v) - 1
				case reflect.Float32, reflect.Float64:
					v = toFloat64(v) - 1.0
				case reflect.String:
					v = -1
				default:
					return nil, errors.New("Invalid operation")
				}
				if err := env.Set(ident.Literal, v); err == ErrUnknownSymbol {
					if err := env.Define(ident.Literal, v); err != nil {
						return nil, err
					}
				} else if err != nil {
					return nil, err
				}
				return v, nil

			} else {
				return nil, errors.New("Invalid operation")
			}
		}
		/*
			if right == nil {
				right = &ast.NumExpr{Literal: "1"}
			}
		*/
		result, err := evalExpr(&ast.BinOpExpr{Left: left, Operator: operator[0:1], Right: right}, env)
		if err != nil {
			return nil, err
		}
		return evalAssExpr(left, result, env)

	case *ast.BinOpExpr:
		var left, right interface{}
		var err error
		if left, err = evalExpr(expr.(*ast.BinOpExpr).Left, env); err != nil {
			return nil, err
		}
		if right, err = evalExpr(expr.(*ast.BinOpExpr).Right, env); err != nil {
			return nil, err
		}
		/*
			//TODO: checking type  "a==1"
			if left == nil && right == nil {
				return 0, nil
			} else if left == nil {
				return right, nil
			} else if right == nil {
				return left, nil
			}
		*/
		switch expr.(*ast.BinOpExpr).Operator {
		case "||":
			if l, ok := left.(bool); !ok {
				return nil, errors.New("cannot convert to bool")
			} else {
				if l {
					return true, nil
				}
				if r, ok := right.(bool); !ok {
					return nil, errors.New("cannot convert to bool")
				} else {
					return r, nil
				}
			}
		case "&&":
			if l, ok := left.(bool); !ok {
				return nil, errors.New("cannot convert to bool")
			} else {
				if !l {
					return false, nil
				}
				if r, ok := right.(bool); !ok {
					return nil, errors.New("cannot convert to bool")
				} else {
					return r, nil
				}
			}
		case "==":
			return left == right, nil
		case "!=":
			return left != right, nil
		case ">":
			//fmt.Printf("toFloat(left)=%#v\n", toFloat64(left))
			//fmt.Printf("toFloat(right)=%#v\n", toFloat64(right))
			return toFloat64(left) > toFloat64(right), nil
		case ">=":
			return toFloat64(left) >= toFloat64(right), nil
		case "<":
			return toFloat64(left) < toFloat64(right), nil
		case "<=":
			return toFloat64(left) <= toFloat64(right), nil
		case "+":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case (l_kind == reflect.Slice || l_kind == reflect.Array) && (r_kind == reflect.Slice || r_kind == reflect.Array):
				return reflect.AppendSlice(reflect.ValueOf(left), reflect.ValueOf(right)).Interface(), nil
			case l_kind == reflect.Slice || l_kind == reflect.Array:
				return reflect.Append(reflect.ValueOf(left), reflect.ValueOf(right)).Interface(), nil
			case l_kind == reflect.Int && r_kind == reflect.Int:
				return toInt(left) + toInt(right), nil
			case l_kind == reflect.String || r_kind == reflect.String:
				return toString(left) + toString(right), nil
			default:
				return toFloat64(left) + toFloat64(right), nil
			}
		case "-":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.Int && r_kind == reflect.Int:
				return toInt(left) - toInt(right), nil
			default:
				return toFloat64(left) - toFloat64(right), nil
			}
		case "*":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.Int && r_kind == reflect.Int:
				return toInt(left) * toInt(right), nil
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
			case l_kind == reflect.Int && r_kind == reflect.Int:
				return toInt(left) / toInt(right), nil
			default:
				return toFloat64(left) / toFloat64(right), nil
			}
		case "%":
			return toInt(left) % toInt(right), nil
		}
	}
	return 0, nil
}

func evalAssExpr(lexp ast.Expr, val interface{}, env *Env) (interface{}, error) {
	switch lexp.(type) {
	case *ast.IdentExpr:
		id := lexp.(*ast.IdentExpr).Literal
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
		expr := lexp.(*ast.FieldExpr).Expr
		i_val, err := evalExpr(expr, env)
		if err != nil {
			//fmt.Println("FieldExpr index error") //TODO
			return nil, err
		}
		index, ok := i_val.(int)
		if !ok {
			return nil, fmt.Errorf("Field index not int :%v", reflect.TypeOf(i_val))
		}
		val_string, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("Field value is not string :%v", reflect.TypeOf(val))
		}
		err = env.SetField(index, val_string)
		if err != nil {
			//fmt.Println("FieldExpr SetField error") //TODO
			return nil, err
		}
		return nil, nil
	case *ast.ItemExpr:
		var index string
		for i, expr := range lexp.(*ast.ItemExpr).Index {
			//fmt.Printf("Index[%v]:%v\n", k, expr)
			val, err := evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
			if i == 0 {
				index = fmt.Sprintf("%v", val)
			} else {
				index = fmt.Sprintf("%v%v%v", index, env.builtin.SUBSEP, val)
			}
		}
		id := lexp.(*ast.ItemExpr).Literal
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
						return evalAssExpr(lexp.(*ast.ItemExpr).Value, ar, env)
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
			_, ok = m[index]
			if ok {
				m[index] = val
				return val, nil
			} else {
				newMap := make(map[interface{}]interface{}, len(m)+1)
				newMap[index] = val
				for k, v := range m {
					newMap[k] = v
				}

				err := env.Set(id, newMap)
				if err != nil {
					return nil, err
				}
				return newMap, nil
			}
		default:
			return nil, errors.New("type " + reflect.TypeOf(value).Kind().String() + " does not support index operation")
		}

	default:
		// TODO:?
		return nil, errors.New("Invalid Operation")
	}
	return val, nil
}
