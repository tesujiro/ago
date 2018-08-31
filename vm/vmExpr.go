package vm

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/tesujiro/goa/ast"
)

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
		//fmt.Printf("CallExpr env:%v builtin.field:%#v\n", env, env.builtin.field)
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
		var value, index interface{}
		var err error
		// index
		index, err = getHashIndex(env, expr.(*ast.ItemExpr).Index)
		if err != nil {
			return nil, err
		}
		// value
		e := expr.(*ast.ItemExpr).Expr
		ie, ok := e.(*ast.IdentExpr)
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
			value, err = evalExpr(e, env)
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
	case *ast.MapExpr:
		exprs := expr.(*ast.MapExpr).MapExpr
		m := make(map[interface{}]interface{}, len(exprs))
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
	case *ast.AssExpr:
		assExpr := expr.(*ast.AssExpr)
		left, right := assExpr.Left, assExpr.Right

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
			if reflect.ValueOf(val).Kind() == reflect.Interface {
				val = reflect.ValueOf(val).Elem().Interface()
			}
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
		left := expr.(*ast.CompExpr).Left
		right := expr.(*ast.CompExpr).Right
		operator := expr.(*ast.CompExpr).Operator
		if expr.(*ast.CompExpr).After {
			afterCompExpr := &ast.CompExpr{Left: left, Right: right, Operator: operator, After: false} //After: false!
			afterStmts = append(afterStmts, &ast.ExprStmt{Expr: afterCompExpr})
			return evalExpr(left, env)
		}
		switch operator {
		case "++", "--":
			right = &ast.NumExpr{Literal: "1"}
		}
		result, err := evalExpr(&ast.BinOpExpr{Left: left, Operator: operator[0:1], Right: right}, env)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("ast.CompExpr: result:%v\n", result)
		return evalAssExpr(left, result, env)

	case *ast.TriOpExpr:
		condExpr := expr.(*ast.TriOpExpr).Cond
		thenExpr := expr.(*ast.TriOpExpr).Then
		elseExpr := expr.(*ast.TriOpExpr).Else
		cond, err := evalExpr(condExpr, env)
		if err != nil {
			return nil, err
		}
		cond_b, err := strictToBool(cond, "ternary operator")
		if err != nil {
			return nil, err
		}
		if cond_b {
			return evalExpr(thenExpr, env)
		} else {
			return evalExpr(elseExpr, env)
		}

	case *ast.BinOpExpr:
		var left, right interface{}
		var err error
		if left, err = evalExpr(expr.(*ast.BinOpExpr).Left, env); err != nil {
			return nil, err
		}
		if right, err = evalExpr(expr.(*ast.BinOpExpr).Right, env); err != nil {
			return nil, err
		}
		switch expr.(*ast.BinOpExpr).Operator {
		case "||":
			left_b, err := strictToBool(left, "left expression of OR operator")
			if err != nil {
				return nil, err
			}
			if left_b {
				return true, nil
			}
			right_b, err := strictToBool(right, "right expression of OR operator")
			if err != nil {
				return nil, err
			}
			if right_b {
				return true, nil
			}
			return false, nil
		case "&&":
			left_b, err := strictToBool(left, "left expression of AND operator")
			if err != nil {
				return nil, err
			}
			if !left_b {
				return false, nil
			}
			right_b, err := strictToBool(right, "right expression of AND operator")
			if err != nil {
				return nil, err
			}
			if right_b {
				return true, nil
			}
			return false, nil
		case "==":
			return left == right, nil
		case "!=":
			return left != right, nil
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
			case (l_kind == reflect.Slice || l_kind == reflect.Array) && (r_kind == reflect.Slice || r_kind == reflect.Array):
				return reflect.AppendSlice(reflect.ValueOf(left), reflect.ValueOf(right)).Interface(), nil
			case l_kind == reflect.Slice || l_kind == reflect.Array:
				return reflect.Append(reflect.ValueOf(left), reflect.ValueOf(right)).Interface(), nil
			case r_kind == reflect.Slice || r_kind == reflect.Array:
				right = reflect.ValueOf(right).Index(0).Interface()
				fallthrough
			case l_kind == reflect.String || r_kind == reflect.String:
				return toString(left) + toString(right), nil
			case l_kind == reflect.Float64 || r_kind == reflect.Float64:
				return toFloat64(left) + toFloat64(right), nil
			case l_kind == reflect.Int || r_kind == reflect.Int:
				return toInt(left) + toInt(right), nil
			default:
				return toString(left) + toString(right), nil
			}
		case "+":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case (l_kind == reflect.Slice || l_kind == reflect.Array) && (r_kind == reflect.Slice || r_kind == reflect.Array):
				return reflect.AppendSlice(reflect.ValueOf(left), reflect.ValueOf(right)).Interface(), nil
			case l_kind == reflect.Slice || l_kind == reflect.Array:
				return reflect.Append(reflect.ValueOf(left), reflect.ValueOf(right)).Interface(), nil
			case r_kind == reflect.Slice || r_kind == reflect.Array:
				right = reflect.ValueOf(right).Index(0).Interface()
				fallthrough
			case l_kind == reflect.Float64 || r_kind == reflect.Float64:
				return toFloat64(left) + toFloat64(right), nil
			case l_kind == reflect.Int || r_kind == reflect.Int:
				return toInt(left) + toInt(right), nil
			case l_kind == reflect.String || r_kind == reflect.String:
				return toString(left) + toString(right), nil
			default:
				return toFloat64(left) + toFloat64(right), nil
			}
		case "-":
			//TODO difference from "+"
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.Slice || l_kind == reflect.Array:
				return 0, nil
			case r_kind == reflect.Slice || r_kind == reflect.Array:
				right = reflect.ValueOf(right).Index(0).Interface()
				fallthrough
			case l_kind == reflect.Float64 || r_kind == reflect.Float64:
				return toFloat64(left) - toFloat64(right), nil
			case l_kind == reflect.Int && r_kind == reflect.Int:
				return toInt(left) - toInt(right), nil
			default:
				return toFloat64(left) - toFloat64(right), nil
			}
		case "*":
			l_kind := reflect.ValueOf(left).Kind()
			r_kind := reflect.ValueOf(right).Kind()
			switch {
			case l_kind == reflect.Slice || l_kind == reflect.Array:
				return 0, nil
			case r_kind == reflect.Slice || r_kind == reflect.Array:
				right = reflect.ValueOf(right).Index(0).Interface()
				fallthrough
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
			case l_kind == reflect.Slice || l_kind == reflect.Array:
				return 0, nil
			case r_kind == reflect.Slice || r_kind == reflect.Array:
				right = reflect.ValueOf(right).Index(0).Interface()
				fallthrough
			case l_kind == reflect.Int && r_kind == reflect.Int:
				return toInt(left) / toInt(right), nil
			default:
				return toFloat64(left) / toFloat64(right), nil
			}
		case "%":
			return toInt(left) % toInt(right), nil
		}
	case *ast.MatchExpr:
		val, err := evalExpr(expr.(*ast.MatchExpr).Expr, env)
		if err != nil {
			return nil, err
		}
		s := toString(val)
		re := expr.(*ast.MatchExpr).RegExpr
		return regexp.MatchString(re, s)
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
			//fmt.Println("fieldExpr index error") //TODO
			return nil, err
		}
		index, ok := i_val.(int)
		if !ok {
			return nil, fmt.Errorf("field index not int :%v", reflect.TypeOf(i_val))
		}
		switch val.(type) {
		case []interface{}:
			val = reflect.ValueOf(val).Index(0).Interface()
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
		e := lexp.(*ast.ItemExpr).Expr
		ie, ok := e.(*ast.IdentExpr)
		if !ok {
			return nil, errors.New("invalid assignment")
		}
		id := ie.Literal

		index, err := getHashIndex(env, lexp.(*ast.ItemExpr).Index)
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
	return val, nil
}
