package vm

import (
	"fmt"
	"reflect"
)

type Env struct {
	env     map[string]interface{}
	parent  *Env
	builtin *builtin
}

type builtin struct {
	NF, NR int
	FS     string
	field  []string
}

// Global Scope
func NewEnv() *Env {
	return &Env{
		env:     make(map[string]interface{}),
		parent:  nil,
		builtin: &builtin{},
	}
}

func (e *Env) NewEnv() *Env {
	return &Env{
		env:     make(map[string]interface{}),
		parent:  e,
		builtin: e.builtin,
	}
}

func (e *Env) Set(k string, v interface{}) error {
	// BuiltIn
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

	// not Builtin
	if _, ok := e.env[k]; ok {
		e.env[k] = v
		return nil
	}
	if e.parent == nil {
		return fmt.Errorf("unknown symbol '%s'", k)
	}
	return e.parent.Set(k, v)
}

func (e *Env) Define(k string, v interface{}) error {
	bt := reflect.TypeOf(e.builtin).Elem()
	if _, ok := bt.FieldByName(k); ok {
		return fmt.Errorf("cannot define builtin variable '%v'", k)
	}
	e.env[k] = v
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

	// Not Builtin
	if v, ok := e.env[k]; ok {
		return v, nil
	}
	if e.parent == nil {
		return nil, fmt.Errorf("unknown symbol '%s'", k)
	}
	return e.parent.Get(k)
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
		return indent + "\t"
	}
	dump_helper(e)
	return
}
