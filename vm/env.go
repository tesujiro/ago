// Package vm implements virtual-machine for ago.
package vm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

const defaultValue = ""

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
	//importFunc map[string]func(*Env) (reflect.Value, error)
	readCloser map[string]*io.ReadCloser
	scanner    map[string]*bufio.Scanner
}

// NewEnv is a Env constructor,
func NewEnv() *Env {
	return &Env{
		env:     make(map[string]interface{}),
		parent:  nil,
		builtin: newBuiltIn(),
		global:  make(map[string]interface{}),
		//importFunc: make(map[string]func(*Env) (reflect.Value, error)),
		readCloser: make(map[string]*io.ReadCloser),
		scanner:    make(map[string]*bufio.Scanner),
	}
}

// NewEnv provides a new child Env,
func (e *Env) NewEnv() *Env {
	return &Env{
		env:     make(map[string]interface{}),
		parent:  e,
		builtin: e.builtin,
		global:  e.global,
		//importFunc: e.importFunc,
		readCloser: e.readCloser,
		scanner:    e.scanner,
	}
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
	if len(s) == 0 {
		return false
	}
	r := s[0]
	return ('A' <= r && r <= 'Z')
}

// GetDefaultValue provides a default vavlue.
func (e *Env) GetDefaultValue() interface{} {
	return defaultValue
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

	// global variable
	if isGlobalVarName(k) {
		if v, ok := e.global[k]; ok {
			return v, nil
		}
	}

	// local variable
	v, err := e.getLocalVar(k)
	if err != nil {
		return nil, ErrUnknownSymbol
	}
	return v, nil
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

// SetFile defines a new file.
func (e *Env) SetFile(k string, f *io.ReadCloser) (*bufio.Scanner, error) {
	_, ok := e.readCloser[k]
	if ok {
		return nil, ErrAlreadyKnownSymbol
	}
	scanner := bufio.NewScanner(io.Reader(*f))
	e.readCloser[k] = f
	e.scanner[k] = scanner
	err := e.setScannerSplit(k)
	if err != nil {
		return nil, err
	}
	return scanner, nil
}

func (e *Env) setScannerSplit(key string) error {
	rs, err := e.Get("RS") // Record Separater
	if err == ErrUnknownSymbol {
		return nil
	} else if err != nil {
		return err
	}
	if rs == "" {
		return nil
	}

	var splitHelper func(int, []byte, []byte, []byte) (int, []byte, error)
	splitHelper = func(advance int, token []byte, data []byte, pat []byte) (int, []byte, error) {
		if len(pat) == 0 {
			return advance, token, nil
		}
		if len(data) == 0 {
			return advance, token, bufio.ErrFinalToken
		}
		/*
			if len(pat) == 0 || len(data) == 0 {
				return advance, token, nil
			}
		*/
		if data[0] == pat[0] {
			return splitHelper(advance+1, append(token, data[0]), data[1:], pat[1:])
		}
		return splitHelper(advance+1, append(token, data[0]), data[1:], pat)
	}
	split := func(data []byte, atEOF bool) (int, []byte, error) {
		i, bs, err := splitHelper(0, []byte{}, data, []byte(rs.(string)))
		if err != nil {
			return i, bs, err
		} else if len(data) == len(bs) {
			return i, bs[:len(bs)-len(rs.(string))], bufio.ErrFinalToken
		} else {
			//fmt.Printf("data=%s\tbs=[%s]\n", data, bs[:len(bs)-len(rs.(string))])
			return i, bs[:len(bs)-len(rs.(string))], nil
		}
	}
	if key == "" {
		// set split func to all the scanners
		for _, scanner := range e.scanner {
			// scanner.Split() panics when used after Scan()
			// No interface to check Scan() is called .
			if len(scanner.Text()) == 0 {
				scanner.Split(split)
			}
		}
	} else {
		// set split func to speified  scanner
		scanner, ok := e.scanner[key]
		if !ok {
			return fmt.Errorf("file key %v not found", key)
		}
		scanner.Split(split)
	}
	return nil
}

// GetScanner returns the scanner with a specified name.
func (e *Env) GetScanner(k string) (*bufio.Scanner, error) {
	s, ok := e.scanner[k]
	if !ok {
		return nil, ErrUnknownSymbol
	}
	return s, nil
}

// CloseFile close a file.
func (e *Env) CloseFile(k string) error {
	f, ok := e.readCloser[k]
	if !ok {
		return ErrUnknownSymbol
	}
	stdin := io.ReadCloser(os.Stdin)
	if f != &stdin {
		if e := (*f).Close(); e != nil {
			return e
		}
	}
	delete(e.readCloser, k)
	delete(e.scanner, k)
	return nil
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
