package vm

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/tesujiro/ago/ast"
	"github.com/tesujiro/ago/debug"
)

func defineFunc(funcExpr *ast.FuncExpr, env *Env) (interface{}, error) {
	// FuncType
	inType := make([]reflect.Type, len(funcExpr.Args))
	for i := 0; i < len(inType); i++ {
		inType[i] = reflect.TypeOf(reflect.Value{})
	}

	outType := []reflect.Type{reflect.TypeOf(reflect.Value{}), reflect.TypeOf(reflect.Value{})}
	isVariadic := false //TODO: variadic
	funcType := reflect.FuncOf(inType, outType, isVariadic)

	// FuncDefinition
	//runVmFunction := func(in []interface{}) (interface{}, error) {
	runVmFunction := func(in []reflect.Value) []reflect.Value {
		newEnv := env.NewEnv()
		//defer newEnv.Destroy()  // Do not delete this line because higer order function

		for i, arg := range funcExpr.Args {
			//val := reflect.ValueOf(in[i]).Interface().(reflect.Value).Interface().(reflect.Value).Interface()
			val := in[i].Interface().(reflect.Value).Interface()
			debug.Printf("arg[%v]: %#v\tType:%v\tValue:%v\n", i, in[i], reflect.TypeOf(val), reflect.ValueOf(val))
			//if err := newEnv.Set(arg, val); err != nil {
			if err := newEnv.Define(arg, val); err != nil {
				return []reflect.Value{reflect.ValueOf(interface{}(nil)), reflect.ValueOf(err)} // TODO:buggy
			}
			//}
		}
		debug.Printf("Env: %#v\n", *env)
		debug.Printf("newEnv: %#v\n", *newEnv)

		if rv, err := runStmts(funcExpr.Stmts, newEnv); err != nil && err != ErrReturn {
			errv := reflect.ValueOf(reflect.ValueOf(&err).Elem())
			debug.Println("errv:\t", errv)
			debug.Println("errv.Type:\t", errv.Type())
			debug.Println("errv.Int():\t", errv.Interface())
			debug.Println("TypeOf(errv.Int()):\t", reflect.TypeOf(errv.Interface()))
			nilValue := reflect.New(reflect.TypeOf((*interface{})(nil)).Elem()).Elem()
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(nilValue)), reflect.ValueOf(errv)}
		} else {
			var errorType = reflect.ValueOf([]error{nil}).Index(0).Type()
			var reflectValueErrorNilValue = reflect.ValueOf(reflect.New(errorType).Elem())
			debug.Println("return value:\t", rv)
			debug.Println("return value Type:\t", reflect.TypeOf(rv))
			debug.Println("return value Value:\t", reflect.ValueOf(rv))

			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(rv)), reflect.ValueOf(reflectValueErrorNilValue)}
		}
	}

	debug.Printf("MakeFunc: funcType %v\n", funcType)
	fn := reflect.MakeFunc(funcType, runVmFunction)

	if funcExpr.Name != "" {
		if err := env.Define(funcExpr.Name, fn); err != nil {
			return nil, err
		}
	}
	return fn, nil
}

func callAnonymousFunc(anonymousCallExpr *ast.AnonymousCallExpr, env *Env) (interface{}, error) {
	ace := anonymousCallExpr
	//fmt.Printf("anonCallExpr:%#v\n", ace)
	debug.Printf("anonCallExpr:%#v\n", ace)
	if result, err := evalExpr(ace.Expr, env); err != nil {
		return nil, err
	} else {
		if rv, ok := result.(reflect.Value); !ok {
			return nil, errors.New("cannot call type " + reflect.TypeOf(result).String())
		} else {
			if rv.Type().Kind() != reflect.Func {
				return nil, errors.New("cannot call type " + reflect.TypeOf(result).String())
			}
			//return callFunc(&ast.CallExpr{Func: rv, SubExprs: ace.SubExprs}, env)
			return evalExpr(&ast.CallExpr{Func: rv, SubExprs: ace.SubExprs}, env)
		}
	}
}

func callFunc(callExpr *ast.CallExpr, env *Env) (interface{}, error) {
	var f reflect.Value
	if callExpr.Name != "" {
		fn, err := env.Get(callExpr.Name) // fn: interface{}
		if err != nil {
			//return nil, err
			fn, err = env.GetDynamicFunc(callExpr.Name)
			if err != nil {
				return nil, err
			}
		}
		var ok bool
		f, ok = fn.(reflect.Value) // interface{} ==> reflect.Value
		if !ok {
			return nil, errors.New("cannot call type " + reflect.TypeOf(fn).String())
		}
	} else {
		f = callExpr.Func
	}
	//debug.Println("func kind:", f.Kind())
	if f.Kind() != reflect.Func {
		return nil, errors.New("cannot call type " + f.Type().String())
	}
	if f.Kind() == reflect.Interface && !f.IsNil() {
		f = f.Elem()
	}
	if args, err := callArgs(f, callExpr, env); err != nil {
		return nil, err
	} else {
		debug.Printf("args: %v\n", args)
		//for i, arg := range args {
		//debug.Printf("=>arg[%d]: %v\n", i, arg.Interface())
		//}

		//fmt.Println("f.Call Start")
		// Call Function
		refvals := f.Call(args)
		//fmt.Println("f.Call End")

		return makeResult(refvals, isGoFunc(f.Type()))
	}
	return nil, nil
}

func isGoFunc(rt reflect.Type) bool {
	if rt.NumOut() != 2 || rt.Out(0) != reflect.TypeOf(reflect.Value{}) || rt.Out(1) != reflect.TypeOf(reflect.Value{}) {
		return true
	}
	return false
}

func callArgs(f reflect.Value, callExpr *ast.CallExpr, env *Env) ([]reflect.Value, error) {
	if f.Type().NumIn() < 1 {
		return []reflect.Value{}, nil
	}
	if !f.Type().IsVariadic() && f.Type().NumIn() != len(callExpr.SubExprs) ||
		f.Type().IsVariadic() && f.Type().NumIn()-1 > len(callExpr.SubExprs) {
		return []reflect.Value{}, fmt.Errorf("function wants %v arguments but received %v", f.Type().NumIn(), len(callExpr.SubExprs))
	}
	var args []reflect.Value
	args = make([]reflect.Value, len(callExpr.SubExprs), len(callExpr.SubExprs))
	/*
		if f.Type().NumIn() > len(callExpr.SubExprs) {
			args = make([]reflect.Value, f.Type().NumIn(), f.Type().NumIn())
		} else {
			args = make([]reflect.Value, len(callExpr.SubExprs), len(callExpr.SubExprs))
		}
	*/
	// func has variadic args
	//hasVariadicArgs := f.Type().In(f.Type().NumIn()-1).Kind() == reflect.Slice
	hasVariadicArgs := f.Type().IsVariadic()

	funcArgType := func(n int) reflect.Type {
		if hasVariadicArgs && f.Type().NumIn()-1 <= n {
			return f.Type().In(f.Type().NumIn() - 1).Elem()
		}
		return f.Type().In(n)
	}

	for k, subExpr := range callExpr.SubExprs {
		// User Defined Funcion
		var arg interface{}
		switch subExpr.(type) {
		case *ast.MatchExpr:
			arg = subExpr.(*ast.MatchExpr).RegExpr
			debug.Println("call parameter contains REGEXP:", arg)
		case *ast.IdentExpr:
			//fmt.Printf("Ident:%v   Func arg type:%v\n", subExpr.(*ast.IdentExpr).Literal, funcArgType(k).Kind().String())
			switch funcArgType(k).Kind() {
			case reflect.Map:
				var err error
				arg, err = env.Get(subExpr.(*ast.IdentExpr).Literal)
				if err != nil {
					return nil, err
				}
			case reflect.Ptr:
				// set variable name to arg
				arg = subExpr.(*ast.IdentExpr).Literal
				//fmt.Printf("arg:%#v   type:%v\n", arg, reflect.TypeOf(arg))
			default:
				var err error
				arg, err = evalExpr(subExpr, env)
				if err != nil {
					return nil, err
				}
			}
		default:
			var err error
			arg, err = evalExpr(subExpr, env)
			if err != nil {
				return nil, err
			}
			debug.Printf("callArg[%v]:%v %v\n", k, arg, reflect.TypeOf(arg))
		}
		if isGoFunc(f.Type()) {
			debug.Printf("call arg[%v]\t%v\ttype:%v\n", k, arg, reflect.TypeOf(arg))
			if k < f.Type().NumIn() {
				debug.Printf("func arg[%v]\ttype:%v\n", k, f.Type().In(k))
			}

			if hasVariadicArgs && f.Type().NumIn()-1 <= k {
				// variadic arg
				variadicArgType := f.Type().In(f.Type().NumIn() - 1).Elem()
				//fmt.Printf("variadicArgType :%v\n", variadicArgType)
				if variadicArgType.Kind() == reflect.Ptr {
					//fmt.Printf("POINTER to %v\n", reflect.PtrTo(variadicArgType))
					//fmt.Printf("arg:%#v   type:%v\n", arg, reflect.TypeOf(arg))
					s := arg.(string) //variable name
					args[k] = reflect.ValueOf(&s)
				} else if reflect.TypeOf(arg) == variadicArgType {
					args[k] = reflect.ValueOf(arg)
				} else {
					//fmt.Printf("func arg[%v] :%v\n", k, f.Type().In(k))
					args[k] = reflect.ValueOf(reflect.ValueOf(arg))
					//args[k] = reflect.ValueOf(&arg.(string))
				}
			} else if reflect.TypeOf(arg) == f.Type().In(k) {
				// not valiadic and same as func arg type
				args[k] = reflect.ValueOf(arg)
			} else if f.Type().In(k).Kind() == reflect.Ptr {
				// TODO
				//args[k] = reflect.ValueOf(reflect.ValueOf(arg))
				s := arg.(string) //variable name
				args[k] = reflect.ValueOf(&s)

			} else {
				// not valiadic and not same as func arg type
				args[k] = reflect.ValueOf(reflect.ValueOf(arg))
			}
		} else {
			args[k] = reflect.ValueOf(reflect.ValueOf(arg))
		}
	}
	debug.Printf("len(args):%v\n", len(args))
	return args, nil
}

func makeResult(ret []reflect.Value, isGoFunction bool) (interface{}, error) {
	debug.Println("ret length:", len(ret))
	// FOR DEBUB
	for i, _ := range ret {
		a := ret[i]
		debug.Printf("ret[%d]           : \tType:%v\tValue:%v\tKind():%v\n", i, reflect.TypeOf(a), reflect.ValueOf(a), reflect.ValueOf(a).Kind())
		b := a.Interface()
		debug.Printf("->Interface()    : \tType:%v\tValue:%v\tKind():%v\n", reflect.TypeOf(b), reflect.ValueOf(b), reflect.ValueOf(b).Kind())
		if c, ok := b.(reflect.Value); ok {
			if c.IsValid() {
				d := c.Interface()
				debug.Printf("->(reflect.Value)    : \tType:%v\tValue:%v\tKind():%v\n", reflect.TypeOf(d), reflect.ValueOf(d), reflect.ValueOf(d).Kind())
			}
		}
	}
	if isGoFunction {
		debug.Println("Go Function")
		// Golang Pacakage Funcion
		if len(ret) == 1 {
			return ret[0].Interface(), nil
		}

		result := make([]interface{}, len(ret))
		for k, v := range ret {
			debug.Printf("ret[%d]           : \tType:%v\tValue:%v\tKind():%v\n", k, reflect.TypeOf(v), reflect.ValueOf(v), reflect.ValueOf(v).Kind())
			result[k] = v.Interface()
		}
		debug.Printf("ret : %v\n", result)
		return result, nil
	}

	debug.Println("User Defined Function")
	if len(ret) != 2 {
		return nil, fmt.Errorf("user defined function did not return 2 values but returned %v values", len(ret))
	}
	if !ret[0].IsValid() {
		return nil, fmt.Errorf("user defined function value 1 did not return reflect value type but returned invalid type")
	}
	if !ret[1].IsValid() {
		return nil, fmt.Errorf("user defined function value 2 did not return reflect value type but returned invalid type")
	}
	if !ret[0].Interface().(reflect.Value).IsValid() {
		ret[0] = reflect.ValueOf(reflect.ValueOf(""))
	}
	if !ret[1].Interface().(reflect.Value).IsValid() {
		ret[1] = reflect.ValueOf(reflect.ValueOf(errors.New("")))
	}
	if ret[0].Type() != reflect.TypeOf(reflect.Value{}) {
		return nil, fmt.Errorf("user defined function value 1 did not return reflect value type but returned %v type", ret[0].Type().String())
	}
	if ret[1].Type() != reflect.TypeOf(reflect.Value{}) {
		return nil, fmt.Errorf("user defined function value 2 did not return reflect value type but returned %v type", ret[1].Type().String())
	}
	// User Defined Funcion
	result := ret[0].Interface().(reflect.Value).Interface()
	rvError := ret[1].Interface().(reflect.Value).Interface().(reflect.Value)

	if !rvError.IsValid() {
		return nil, fmt.Errorf("VM function error type is invalid")
	}

	if rvError.IsNil() {
		// RETURN RESULT
		return result, nil
	}

	var errorType = reflect.ValueOf([]error{nil}).Index(0).Type()
	if rvError.Type() == errorType {
		// RETURN error
		return nil, rvError.Interface().(error)
	}

	return result, nil
}
