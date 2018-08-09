package vm

import (
	"errors"
	"fmt"
	"reflect"
)

const defaultValue = ""

var ErrUnknownSymbol = errors.New("unknown symbol")

type Env struct {
	env     map[string]interface{}
	parent  *Env
	builtin *builtin
	global  map[string]interface{}
}

// Global Scope
func NewEnv() *Env {
	return &Env{
		env:     make(map[string]interface{}),
		parent:  nil,
		builtin: NewBuiltIn(),
		global:  make(map[string]interface{}),
	}
}

func (e *Env) NewEnv() *Env {
	return &Env{
		env:     make(map[string]interface{}),
		parent:  e,
		builtin: e.builtin,
		global:  e.global,
	}
}

func isGlobalVarName(s string) bool {
	if len(s) == 0 {
		return false
	}
	r := s[0]
	return ('A' <= r && r <= 'Z')
}

func (e *Env) GetDefaultValue() interface{} {
	return defaultValue
}

func (e *Env) Set(k string, v interface{}) error {
	//fmt.Printf("Set(%v,%v)\n", k, v)
	// BuiltIn variable
	bv := reflect.ValueOf(e.builtin).Elem()
	bt := reflect.TypeOf(e.builtin).Elem()
	if f, ok := bt.FieldByName(k); ok {
		if f.Type != reflect.TypeOf(v) {
			return fmt.Errorf("type of %v must be %v ,not %v.", f.Name, f.Type, reflect.TypeOf(v))
		}
		fv := bv.FieldByName(k)
		if !fv.CanSet() {
			return fmt.Errorf("cannot update %v", f.Name)
		}
		fv.Set(reflect.ValueOf(v))
		return nil
	}

	// global variable
	if isGlobalVarName(k) {
		//fmt.Printf("==>global var\n")
		if _, ok := e.global[k]; ok {
			e.global[k] = v
			return nil
		}
		return e.Define(k, v)
	}

	// local variable
	if err := e.setLocalVar(k, v); err != nil {
		return ErrUnknownSymbol
	}
	return nil
}

func (e *Env) setLocalVar(k string, v interface{}) error {
	if _, ok := e.env[k]; ok {
		e.env[k] = v
		return nil
	}
	if e.parent == nil {
		return ErrUnknownSymbol
	}
	return e.parent.setLocalVar(k, v)
}

func (e *Env) DefineDefaultValue(k string) (interface{}, error) {
	v := defaultValue
	return v, e.Define(k, v)
}

// TODO: DefineDefaultMapValue should be implemented in vmExpr using Env.GetDefaultValue()
func (e *Env) DefineDefaultMapValue(k string, idx interface{}) (interface{}, error) {
	v := make(map[interface{}]interface{})
	v[idx] = defaultValue
	return v, e.Define(k, v)
}

func (e *Env) Define(k string, v interface{}) error {
	// builtin
	bt := reflect.TypeOf(e.builtin).Elem()
	if _, ok := bt.FieldByName(k); ok {
		return fmt.Errorf("cannot define builtin variable '%v'", k)
	}
	if isGlobalVarName(k) {
		// global var
		e.global[k] = v
	} else {
		// local var
		e.env[k] = v
	}
	return nil
}

func (e *Env) Get(k string) (interface{}, error) {
	// Builtin
	bv := reflect.ValueOf(e.builtin).Elem()
	bt := reflect.TypeOf(e.builtin).Elem()
	if _, ok := bt.FieldByName(k); ok {
		fv := bv.FieldByName(k)
		return fv.Interface(), nil
	}

	// global variable
	if v, ok := e.global[k]; ok {
		return v, nil
	}

	// local variable
	if v, err := e.getLocalVar(k); err != nil {
		return nil, ErrUnknownSymbol
	} else {
		return v, nil
	}
}

func (e *Env) getLocalVar(k string) (interface{}, error) {

	if v, ok := e.env[k]; ok {
		return v, nil
	}
	if e.parent == nil {
		return nil, ErrUnknownSymbol
	}
	return e.parent.getLocalVar(k)
}

func (e *Env) Dump() {
	var dump_helper func(*Env) string
	dump_helper = func(e *Env) string {
		var indent string
		if e.parent == nil {
			indent = ""
		} else {
			indent = dump_helper(e.parent)
		}
		for k, v := range e.env {
			fmt.Println(indent, k, ":", v)
		}
		fmt.Println("builtin:", e.builtin)
		fmt.Println("global:", e.global)
		return indent + "\t"
	}
	dump_helper(e)
	return
}
