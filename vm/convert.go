package vm

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/tesujiro/ago/ast"
)

// ToInt convert type to int
func (env *Env) ToInt(val interface{}) interface{} {
	i, err := env.strictToInt(val)
	if err != nil {
		return val
	}
	return i
}

func (env *Env) toInt(val interface{}) int {
	i, err := env.strictToInt(val)
	if err != nil {
		return 0
	}
	return i
}

func (env *Env) strictToInt(val interface{}) (int, error) {
	f, err := env.strictToFloat(val)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

// ToFloat convert type to float64
func (env *Env) ToFloat64(val interface{}) interface{} {
	i, err := env.strictToFloat(val)
	if err != nil {
		return val
	}
	return i
}

func (env *Env) toFloat64(val interface{}) float64 {
	f, _ := env.strictToFloat(val)
	return f
}

func (env *Env) strictToFloat(val interface{}) (float64, error) {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Int64:
		return float64(val.(int64)), nil
	//case reflect.Int32:
	//return float64(val.(int32)), nil
	//case reflect.Int16:
	//return float64(val.(int16)), nil
	//case reflect.Int8:
	//return float64(val.(int8)), nil
	case reflect.Int:
		return float64(val.(int)), nil
	case reflect.Float64, reflect.Float32:
		return val.(float64), nil
	case reflect.String:
		// "1.1" -> 1.1
		// ".123" -> 0.123
		// "1.xx" -> 1
		// "1e1.xx" -> 10
		// "0x11.xx" -> 17  //TODO
		//digit := `(\-|\+)?\d+(\.\d*)?`
		//digit := `(\-|\+)?\d+(\.\d*)?|(\-|\+)?\.\d+`
		digit := `(\-|\+)?\d+(\.\d*)?((e|E)(\-|\+)?\d+)?|(\-|\+)?\.\d+((e|E)(\-|\+)?\d+)?`
		re := regexp.MustCompile(`^` + digit)
		numStr := re.FindString(val.(string))
		if len(numStr) == 0 {
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

func (env *Env) toBool(val interface{}) bool {
	b, err := env.strictToBool(val)
	if err != nil {
		return false
	}
	return b
}

func (env *Env) strictToBool(val interface{}) (bool, error) {
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

// ToFloat convert type to string
func (env *Env) ToString(val interface{}) interface{} {
	return env.toString(val)
}

func (env *Env) toString(val interface{}) string {
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

func (env *Env) getHashIndex(exprs []ast.Expr) (string, error) {
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
