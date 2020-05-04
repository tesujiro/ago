// Package vm implements virtual-machine for ago.
package vm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

const defaultValue = ""
const defaultNumberValue = 0

// ErrUnknownSymbol provides unknown symbol error.
var ErrUnknownSymbol = errors.New("unknown symbol")

// ErrAlreadyKnownSymbol provides already known symbol error.
var ErrAlreadyKnownSymbol = errors.New("already known symbol")

// Env provides an environment.
type Env struct {
	env     map[string]interface{}
	parent  *Env
	builtin *builtin
	global  map[string]interface{}
	funcArg map[string]interface{}
	//importFunc map[string]func(*Env) (reflect.Value, error)
	readCloser     map[string]*io.ReadCloser
	scanner        map[string]*bufio.Scanner
	files          []string
	curFileIndex   int
	curFileCloser  *io.ReadCloser
	curFileScanner *bufio.Scanner
}

// NewEnv is a Env constructor,
func NewEnv(files []string) *Env {
	global := make(map[string]interface{})
	global["ENVIRON"] = getEnvVars()

	return &Env{
		env:     make(map[string]interface{}),
		parent:  nil,
		builtin: newBuiltIn(files),
		global:  global,
		funcArg: make(map[string]interface{}),
		//importFunc: make(map[string]func(*Env) (reflect.Value, error)),
		readCloser:   make(map[string]*io.ReadCloser),
		scanner:      make(map[string]*bufio.Scanner),
		files:        files,
		curFileIndex: 0,
	}
}

// NewEnv provides a new child Env,
func (e *Env) NewEnv() *Env {
	return &Env{
		env:     make(map[string]interface{}),
		parent:  e,
		builtin: e.builtin,
		global:  e.global,
		funcArg: make(map[string]interface{}),
		//importFunc: e.importFunc,
		readCloser: e.readCloser,
		scanner:    e.scanner,
	}
}

func getEnvVars() map[interface{}]interface{} {
	envs := make(map[interface{}]interface{})
	for _, v := range os.Environ() {
		x := strings.SplitN(v, "=", 2)
		envs[x[0]] = x[1]
	}
	return envs
}

var globalVars bool

// SetGlobalVariables set all variables to be global.
func SetGlobalVariables() {
	globalVars = true
}

func isGlobalVarName(s string) bool {
	if globalVars {
		return true
	}
	return len(s) > 0 && 'A' <= s[0] && s[0] <= 'Z'
}

// GetDefaultValue provides a default vavlue.
func (e *Env) GetDefaultValue() interface{} {
	return defaultValue
}

// GetDefaultValue provides a default vavlue.
func (e *Env) GetDefaultNumberValue() interface{} {
	return defaultNumberValue
}

// Set provides a setter for an environment variable.
func (e *Env) Set(k string, v interface{}) error {
	//fmt.Printf("Set(%v,%v)\n", k, v)
	// BuiltIn variable
	if e.isBuiltin(k) {
		bv := reflect.ValueOf(e.builtin).Elem()
		bt := reflect.TypeOf(e.builtin).Elem()
		if f, ok := bt.FieldByName(k); ok {
			/*
				if f.Type != reflect.TypeOf(v) {
					return fmt.Errorf("type of %v must be %v ,not %v.", f.Name, f.Type, reflect.TypeOf(v))
				}
			*/
			switch reflect.TypeOf(v).Kind() {
			case f.Type.Kind():
				break
			case reflect.Float64:
				v = int(v.(float64))
			default:
				return fmt.Errorf("type of %v must be %v ,not %v", f.Name, f.Type, reflect.TypeOf(v))
			}
			fv := bv.FieldByName(k)
			if !fv.CanSet() {
				return fmt.Errorf("cannot update %v", f.Name)
			}
			fv.Set(reflect.ValueOf(v))
			if k == "RS" {
				err := e.setScannerSplit("")
				if err != nil {
					return err
				}
			}
			return nil
		}
	}

	// global variable
	if isGlobalVarName(k) {
		if _, ok := e.funcArg[k]; ok {
			e.funcArg[k] = v
			return nil
		}
		if _, ok := e.global[k]; ok {
			e.global[k] = v
			return nil
		}
		return e.Define(k, v)
	}

	// local variable
	return e.setLocalVar(k, v)
}

// setLocalVar sets a local variable.
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

// DefineDefaultValue sets a single value variable with a default value.
func (e *Env) DefineDefaultValue(k string) (interface{}, error) {
	v := defaultValue
	return v, e.Define(k, v)
}

// DefineDefaultMap sets a map variable with default values.
func (e *Env) DefineDefaultMap(k string) (interface{}, error) {
	v := make(map[interface{}]interface{})
	return v, e.Define(k, v)
}

// DefineDefaultMapValue sets a map element with a default value.
// TODO: DefineDefaultMapValue should be implemented in vmExpr using Env.GetDefaultValue()
func (e *Env) DefineDefaultMapValue(k string, idx interface{}) (interface{}, error) {
	v := make(map[interface{}]interface{})
	v[idx] = defaultValue
	return v, e.Define(k, v)
}

// Define sets a variable to a value.
func (e *Env) Define(k string, v interface{}) error {
	// builtin
	bt := reflect.TypeOf(e.builtin).Elem()
	if _, ok := bt.FieldByName(k); ok && k == strings.ToUpper(k) {
		//if _, ok := bt.FieldByName(k); ok {
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

func (e *Env) DefineFuncArg(k string, v interface{}) error {
	// builtin
	bt := reflect.TypeOf(e.builtin).Elem()
	if _, ok := bt.FieldByName(k); ok && k == strings.ToUpper(k) {
		return fmt.Errorf("cannot define builtin variable '%v'", k)
	}
	if isGlobalVarName(k) {
		// function arg var
		e.funcArg[k] = v
	} else {
		// local var
		e.env[k] = v
	}
	return nil
}

// Get gets a value of a variable.
func (e *Env) Get(k string) (interface{}, error) {
	// Builtin
	if e.isBuiltin(k) {
		bv := reflect.ValueOf(e.builtin).Elem()
		bt := reflect.TypeOf(e.builtin).Elem()
		if _, ok := bt.FieldByName(k); ok {
			fv := bv.FieldByName(k)
			return fv.Interface(), nil
		}
	}

	// function arg variable
	if v, ok := e.funcArg[k]; ok {
		return v, nil
	}

	// global variable
	if v, ok := e.global[k]; ok {
		return v, nil
	}

	// local variable
	return e.getLocalVar(k)
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

/*
func (e *Env) DefineImportFunc(k string, f func(*Env) (reflect.Value, error)) error {
	e.importFunc[k] = f
	return nil
}

func (e *Env) GetDynamicFunc(k string) (interface{}, error) {
	impf, ok := e.importFunc[k]
	if !ok {
		return nil, ErrUnknownSymbol
	}
	fn, err := impf(e)
	if err != nil {
		return nil, err
	}
	e.Define(k, fn) // for cache
	return fn, nil
}
*/

// Dump dumps the environment.
func (e *Env) Dump() {
	var dumpHelper func(*Env) string
	dumpHelper = func(e *Env) string {
		var indent string
		if e.parent == nil {
			indent = ""
		} else {
			indent = dumpHelper(e.parent)
		}
		for k, v := range e.env {
			fmt.Println(indent, k, ":", v)
		}
		fmt.Println("builtin:", e.builtin)
		fmt.Println("global:", e.global)
		return indent + "\t"
	}
	dumpHelper(e)
	return
}
