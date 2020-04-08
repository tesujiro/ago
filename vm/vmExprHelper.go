package vm

import (
	"fmt"
	"math"
	"reflect"
)

func evalBoolOp(op string, left, right interface{}) (interface{}, error) {
	boolLeft, err := strictToBool(left)
	if err != nil {
		return nil, fmt.Errorf("convert left expression of %v operator:%v", op, err)
	}
	if op == "||" && boolLeft {
		return true, nil
	} else if op == "&&" && !boolLeft {
		return false, nil
	}
	boolRight, err := strictToBool(right)
	if err != nil {
		return nil, fmt.Errorf("convert right expression of %v operator:%v", op, err)
	}
	return boolRight, nil
}

func compareEqual(op string, left, right interface{}) (interface{}, error) {
	lKind := reflect.ValueOf(left).Kind()
	rKind := reflect.ValueOf(right).Kind()
	compEq := func(op string, l, r interface{}) bool {
		if op == "==" {
			return l == r
		}
		return l != r
	}
	switch {
	case lKind == reflect.String && rKind == reflect.String:
		return compEq(op, left, right), nil
	case lKind == reflect.Float64 || rKind == reflect.Float64:
		return compEq(op, toFloat64(left), toFloat64(right)), nil
	case lKind == reflect.Int || rKind == reflect.Int:
		return compEq(op, toString(left), toString(right)), nil
	case lKind == reflect.Map || rKind == reflect.Map:
		return nil, fmt.Errorf("can't read value of array")
	default:
		return compEq(op, toString(left), toString(right)), nil
	}
}

func compareInequal(op string, left, right interface{}) (interface{}, error) {
	lKind := reflect.ValueOf(left).Kind()
	rKind := reflect.ValueOf(right).Kind()
	compNumber := func(op string, l, r float64) bool {
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
	if lKind == reflect.Map || rKind == reflect.Map {
		return nil, fmt.Errorf("can't read value of array")
	} else if lKind == reflect.String && rKind == reflect.String {
		return compString(op, toString(left), toString(right)), nil
	} else {
		return compNumber(op, toFloat64(left), toFloat64(right)), nil
	}
}

func evalArithOp(op string, left, right interface{}) (interface{}, error) {
	lKind := reflect.ValueOf(left).Kind()
	rKind := reflect.ValueOf(right).Kind()
	if lKind == reflect.Map || rKind == reflect.Map {
		return nil, fmt.Errorf("can't read value of array")
	}
	if (op == "/" || op == "%") && right == 0 {
		return nil, fmt.Errorf("devision by zero")
	}
	switch op {
	case "+":
		return toFloat64(left) + toFloat64(right), nil
	case "-":
		return toFloat64(left) - toFloat64(right), nil
	case "*":
		return toFloat64(left) * toFloat64(right), nil
	case "/":
		return toFloat64(left) / toFloat64(right), nil
	case "^":
		return math.Pow(toFloat64(left), toFloat64(right)), nil
	default:
		q := int(toFloat64(left) / toFloat64(right))
		return toFloat64(left) - toFloat64(right)*float64(q), nil
	}
}
