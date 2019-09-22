package vm

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/tesujiro/ago/ast"
)

func toInt(val interface{}) int {
	i, err := strictToInt(val)
	if err != nil {
		return 0
	}
	return i
}

func strictToInt(val interface{}) (int, error) {
	f, err := strictToFloat(val)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

func toFloat64(val interface{}) float64 {
	f, _ := strictToFloat(val)
	return f
}

func strictToFloat(val interface{}) (float64, error) {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Int64:
		return float64(val.(int64)), nil
	case reflect.Int32:
		return float64(val.(int32)), nil
	case reflect.Int16:
		return float64(val.(int16)), nil
	case reflect.Int8:
		return float64(val.(int8)), nil
	case reflect.Int:
		return float64(val.(int)), nil
	case reflect.Float64, reflect.Float32:
		return val.(float64), nil
	case reflect.String:
		// "1.1" -> 1
		// "1.xx" -> 1
		// "1e1.xx" -> 10 //TODO
		// "0x11.xx" -> 17  //TODO
		digit := `(\-|\+)?\d+(\.\d*)?`
		re := regexp.MustCompile(`^` + digit)
		numStr := re.FindString(val.(string))
		if len(numStr) == 0 {
			re = regexp.MustCompile(`^` + digit + `(e|E)` + digit)
			return 0, fmt.Errorf("cannot convert to float:%v", reflect.ValueOf(val).Kind())
		}

		var num float64
		var err error
		if num, err = strconv.ParseFloat(numStr, 64); err != nil {
			return 0, err
		}
		return num, err
	default:
		return 0, fmt.Errorf("cannot convert to float:%v", reflect.ValueOf(val).Kind())
	}
}

func toBool(val interface{}) bool {
	b, err := strictToBool(val)
	if err != nil {
		return false
	}
	return b
}

func strictToBool(val interface{}) (bool, error) {
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
		return false, fmt.Errorf("convert interface{} to bool failed")
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
