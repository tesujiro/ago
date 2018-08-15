package lib

import (
	"reflect"

	"github.com/tesujiro/goa/vm"
)

func Import(env *vm.Env) *vm.Env {
	// length cannot return error -> TODO:change length to embedded func
	//length := func(val interface{}) int {
	//v := val.(reflect.Value)
	length := func(v reflect.Value) int {
		switch v.Type().Kind() {
		case reflect.String:
			s := v.Interface().(string)
			return len(s)
		case reflect.Map:
			s := v.Interface().(map[interface{}]interface{})
			return len(s)
		case reflect.Slice:
			s := v.Interface().([]interface{})
			return len(s)
		default:
			// fmt.Errorf("invalid argument %v (type %v) for len",s,reflect.TypeOf(s))
			return 0
		}
	}
	//env.Define("length", length)
	env.Define("length", reflect.ValueOf(length))
	//env.Define("len", length)
	env.Define("len", reflect.ValueOf(length))

	//substr := func(str, start, end interface{}) {
	//}

	return env
}
