package vm

import (
	"fmt"
	"testing"
)

func TestSetGet(t *testing.T) {
	cases := []struct {
		checkSet    bool
		checkDefine bool
		key         string
		value       interface{}
		message     string
	}{
		{checkSet: true, checkDefine: false, key: "NF", value: "abc", message: "type of NF must be int ,not string."},
		{checkSet: true, checkDefine: false, key: "NF", value: 123},
		{checkSet: false, checkDefine: true, key: "NF", value: 123, message: "cannot define builtin variable 'NF'"},
		{checkSet: true, checkDefine: false, key: "KEY", value: 123},
		{checkSet: false, checkDefine: true, key: "KEY", value: 123},
	}
	fmt.Printf("TESTSetGet\n")
	for _, c := range cases {
		e := NewEnv()
		if c.checkSet {
			if err := e.Set(c.key, c.value); err != nil {
				if err.Error() != c.message {
					t.Errorf("env.Set() got %v\nwant %v", err, c.message)
				}
				continue
			}
		}
		if c.checkDefine {
			if err := e.Define(c.key, c.value); err != nil {
				if err.Error() != c.message {
					t.Errorf("env.Define() got %v\nwant %v", err, c.message)
				}
				continue
			}
		}
		if actual, err := e.Get(c.key); err != nil && err.Error() != c.message {
			t.Errorf("env.Get() got %v\nwant %v", err, c.message)
			continue
		} else if actual != c.value {
			t.Errorf("env.Get() got %v\nwant %v", actual, c.value)
			continue
		}
		if c.message != "" {
			t.Errorf("no error message want %v", c.message)
		}
		//e.Dump()
	}
}

func TestChildEnv(t *testing.T) {
	tests := []struct {
		key   string
		value interface{}
	}{
		{key: "int", value: 123},
		{key: "float", value: 1.1},
	}

	root := NewEnv()
	for _, test := range tests {
		if err := root.Define(test.key, test.value); err != nil {
			t.Errorf("Env.Set error :%v", err)
			return
		} else if actual, err := root.Get(test.key); err != nil {
			t.Errorf("Env.Get error :%v", err)
			return
		} else if actual != test.value {
			t.Errorf("got %v\nwant %v", actual, test.value)
			return
		}
	}

	child := root.NewEnv()
	for _, test := range tests {
		if err := child.Define(test.key, test.value); err != nil {
			t.Errorf("Env.Set error :%v", err)
			return
		} else if actual, err := child.Get(test.key); err != nil {
			t.Errorf("Env.Get error :%v", err)
			return
		} else if actual != test.value {
			t.Errorf("got %v\nwant %v", actual, test.value)
			return
		}
	}

}
