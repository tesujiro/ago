package vm

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/tesujiro/ago/ast"
)

func toInt(val interface{}) int {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Float64, reflect.Float32:
		return int(val.(float64))
	case reflect.String:
		if i, err := strconv.Atoi(val.(string)); err != nil {
			return 0
		} else {
			return i
		}
	}
	i, _ := val.(int)
	return i
}

func strictToInt(val interface{}) (int, error) {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return val.(int), nil
	case reflect.Float64, reflect.Float32:
		return int(val.(float64)), nil
	case reflect.String:
		// "1.1" -> 1
		// "1.xx" -> 1
		// "1e1.xx" -> 10 //TODO
		// "011.xx" -> 9  //TODO
		re := regexp.MustCompile(`\d+`)
		num := re.FindString(val.(string))
		if len(num) == 0 {
			return 0, fmt.Errorf("cannot convert to int:%v", reflect.ValueOf(val).Kind())
		}

		if i, err := strconv.Atoi(num); err != nil {
			return 0, err
		} else {
			return i, err
		}
	default:
		return 0, fmt.Errorf("cannot convert to int:%v", reflect.ValueOf(val).Kind())
	}
}

func toFloat64(val interface{}) float64 {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return float64(val.(int))
	case reflect.String:
		if f, err := strconv.ParseFloat(val.(string), 64); err != nil {
			return 0
		} else {
			return f
		}
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

func strictToBool(val interface{}, operation string) (bool, error) {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Bool:
		return val.(bool), nil
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return val.(int) != 0, nil
	case reflect.Float32, reflect.Float64:
		return val.(float64) != 0, nil
	case reflect.String:
		return val.(string) != "", nil
	default:
		return false, fmt.Errorf("convert to bool failed in %v", operation)
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
}

func getHashIndex(env *Env, exprs []ast.Expr) (string, error) {
	var index string
	for i, expr := range exprs {
		//fmt.Printf("Index[%v]:%v\n", k, expr)
		val, err := evalExpr(expr, env)
		if err != nil {
			return "", err
		}
		if i == 0 {
			index = fmt.Sprintf("%v", val)
		} else {
			index = fmt.Sprintf("%v%v%v", index, env.builtin.SUBSEP, val)
		}
	}
	return index, nil
}
