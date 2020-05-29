package vm

import (
	"fmt"
	"math"
	"reflect"
)

func evalBoolOp(env *Env, op string, left, right interface{}) (interface{}, error) {
	boolLeft, err := env.strictToBool(left)
	if err != nil {
		return nil, fmt.Errorf("convert left expression of %v operator:%v", op, err)
	}
	if op == "||" && boolLeft {
		return true, nil
	} else if op == "&&" && !boolLeft {
		return false, nil
	}
	boolRight, err := env.strictToBool(right)
	if err != nil {
		return nil, fmt.Errorf("convert right expression of %v operator:%v", op, err)
	}
	return boolRight, nil
}

/*
func isString(val interface{}) bool {
	kind := reflect.ValueOf(val).Kind()
	if kind == reflect.String {
		return true
	}
	if kind == reflect.Struct {
		if f, ok := val.(Field); ok && f.Type == StringType {
			return true
		}
	}
	return false
}
*/

func isNumber(val interface{}) bool {
	kind := reflect.ValueOf(val).Kind()
	if kind == reflect.Float64 || kind == reflect.Int {
		return true
	}
	if kind == reflect.Struct {
		if f, ok := val.(Field); ok && f.Type == NumberType {
			//debug.Printf("isNumber(%T)-->true\n", val)
			return true
		}
	}
	//debug.Printf("isNumber(%T)-->false\n", val)
	return false
}

func isArray(val interface{}) bool {
	kind := reflect.ValueOf(val).Kind()
	if kind == reflect.Map {
		return true
	}
	return false
}

func compareEqual(env *Env, op string, left, right interface{}) (interface{}, error) {
	compEq := func(op string, l, r interface{}) bool {
		if op == "==" {
			return l == r
		}
		return l != r
	}
	switch {
	case isArray(left) || isArray(right):
		return nil, fmt.Errorf("can't read value of array")
	case isNumber(left) || isNumber(right):
		return compEq(op, env.toFloat64(left), env.toFloat64(right)), nil
	default:
		return compEq(op, env.toString(left), env.toString(right)), nil
	}
}

func compareInequal(env *Env, op string, left, right interface{}) (interface{}, error) {
	compNumber := func(op string, l, r float64) bool {
		//fmt.Printf("op=%v\tl=%v\tr=%v\n", op, l, r)
		switch op {
		case ">":
			return l > r
		case "<":
			return l < r
		case ">=":
			return l >= r
		default:
			return l <= r
		}
	}
	compString := func(op string, l, r string) bool {
		switch op {
		case ">":
			return l > r
		case "<":
			return l < r
		case ">=":
			return l >= r
		default:
			return l <= r
		}
	}
	switch {
	case isArray(left) || isArray(right):
		return nil, fmt.Errorf("can't read value of array")
	case isNumber(left) && isNumber(right):
		return compNumber(op, env.toFloat64(left), env.toFloat64(right)), nil
	default:
		return compString(op, env.toString(left), env.toString(right)), nil
	}
}

func evalArithOp(env *Env, op string, left, right interface{}) (interface{}, error) {
	lKind := reflect.ValueOf(left).Kind()
	rKind := reflect.ValueOf(right).Kind()
	if lKind == reflect.Map || rKind == reflect.Map {
		return nil, fmt.Errorf("can't read value of array")
	}
	if (op == "/" || op == "%") && right == 0 {
		return nil, fmt.Errorf("devision by zero")
	}
	normNumber := func(f float64) interface{} {
		if math.Round(f) == f {
			return int(f)
		}
		return f
	}
	switch op {
	case "+":
		//return toFloat64(left) + toFloat64(right), nil
		return normNumber(env.toFloat64(left) + env.toFloat64(right)), nil
	case "-":
		//return toFloat64(left) - toFloat64(right), nil
		return normNumber(env.toFloat64(left) - env.toFloat64(right)), nil
	case "*":
		//return toFloat64(left) * toFloat64(right), nil
		return normNumber(env.toFloat64(left) * env.toFloat64(right)), nil
	case "/":
		//return toFloat64(left) / toFloat64(right), nil
		return normNumber(env.toFloat64(left) / env.toFloat64(right)), nil
	case "^":
		//return math.Pow(toFloat64(left), toFloat64(right)), nil
		return normNumber(math.Pow(env.toFloat64(left), env.toFloat64(right))), nil
	default:
		q := int(env.toFloat64(left) / env.toFloat64(right))
		//return toFloat64(left) - toFloat64(right)*float64(q), nil
		return normNumber(env.toFloat64(left) - env.toFloat64(right)*float64(q)), nil
	}
}
