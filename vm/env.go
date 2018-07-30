package vm

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
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

func (e *Env) setNR(i int) {
	e.builtin.NR = i
}

func (e *Env) setNF(i int) {
	e.builtin.NF = i
}

func (e *Env) SetFS(fs string) {
	e.builtin.FS = fs
	//e.Dump()
}

func (e *Env) GetField(i int) (string, error) {
	// TODO: out of index
	if i < 0 || i >= len(e.builtin.field) {
		return "", nil
	}
	return e.builtin.field[i], nil
}

func (e *Env) SetFieldZero() error {
	//fmt.Println("SetFieldZero:", e.builtin.field)
	if len(e.builtin.field) <= 1 {
		e.builtin.field[0] = ""
		return nil
	}
	str := e.builtin.field[1]
	//fmt.Println("len:", len(e.builtin.field))
	for i := 2; i < len(e.builtin.field); i++ {
		str += e.builtin.FS + e.builtin.field[i]
	}
	e.builtin.field[0] = str
	return nil
}

func (e *Env) SetField(index int, str string) error {
	if index < 0 {
		return fmt.Errorf("Field Index Out of Range:%v\n", index)
	}
	if index > len(e.builtin.field) {
		//TODO
	} else {
		e.builtin.field[index] = str
	}
	e.SetFieldZero()
	return nil
}

var re_org_awk_truncate = regexp.MustCompile("^[ \t]*([^ \t].*[^ \t])[ \t]*$")

func (e *Env) SetFieldFromLine(line string) error {
	split := func(regex, line string) {
		re := regexp.MustCompile(regex) //TODO: STORE PRE COMPILED VALUE TO ENV FOR PERFORMANCE
		result := re.Split(line, -1)
		e.builtin.field = make([]string, len(result)+1)
		for i, f := range result {
			e.builtin.field[i+1] = f
		}
	}
	switch e.builtin.FS {
	case "":
		return errors.New("Field Seaparotor not set")
	case " ":
		//THIS IS SPECIAL CASE FOR ORIGINAL AWK
		//fmt.Printf("line %v:[%v]\n", e.builtin.NR, line)
		line = re_org_awk_truncate.ReplaceAllString(line, "$1")
		//fmt.Printf("line %v:[%v]\n", e.builtin.NR, line)
		split("[ \t]+", line)
	default:
		fmt.Printf("line %v:FS[%v]\n", e.builtin.NR, e.builtin.FS)
		split(e.builtin.FS, line)
	}
	//e.builtin.field[0] = line
	e.SetFieldZero()
	e.setNF(len(e.builtin.field) - 1)

	return nil
}
