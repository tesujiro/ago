package vm

import (
	"testing"
)

func TestSetGet(t *testing.T) {
	cases := []struct {
		testSet    bool
		testDefine bool
		k          string
		v          interface{}
		errMessage string
	}{
		{testSet: true, testDefine: false, k: "NF", v: "abc", errMessage: "type of NF must be int ,not string."},
		{testSet: true, testDefine: false, k: "NF", v: 123},
		{testSet: true, testDefine: false, k: "KEY", v: 123, errMessage: "unknown symbol 'KEY'"},
		{testSet: false, testDefine: true, k: "KEY", v: 123},
		{testSet: false, testDefine: true, k: "NF", v: 123, errMessage: "cannot define builtin variable 'NF'"},
	}
	for _, c := range cases {
		e := NewEnv()
		if c.testSet {
			if err := e.Set(c.k, c.v); err != nil {
				if err.Error() != c.errMessage {
					t.Errorf("env.Set() got %v\nwant %v", err, c.errMessage)
				}
				continue
			}
		}
		if c.testDefine {
			if err := e.Define(c.k, c.v); err != nil {
				if err.Error() != c.errMessage {
					t.Errorf("env.Define() got %v\nwant %v", err, c.errMessage)
				}
				continue
			}
		}
		if actual, err := e.Get(c.k); err != nil && err.Error() != c.errMessage {
			t.Errorf("env.Get() got %v\nwant %v", err, c.errMessage)
			continue
		} else if actual != c.v {
			t.Errorf("env.Get() got %v\nwant %v", actual, c.v)
			continue
		}
		//e.Dump()
	}
}

func TestChildEnv(t *testing.T) {
	tests := []struct {
		k string
		v interface{}
	}{
		{k: "int", v: 123},
		{k: "float", v: 1.1},
	}

	root := NewEnv()
	for _, test := range tests {
		if err := root.Define(test.k, test.v); err != nil {
			t.Errorf("Env.Set error :%v", err)
			return
		} else if actual, err := root.Get(test.k); err != nil {
			t.Errorf("Env.Get error :%v", err)
			return
		} else if actual != test.v {
			t.Errorf("got %v\nwant %v", actual, test.v)
			return
		}
	}

	child := root.NewEnv()
	for _, test := range tests {
		if err := child.Define(test.k, test.v); err != nil {
			t.Errorf("Env.Set error :%v", err)
			return
		} else if actual, err := child.Get(test.k); err != nil {
			t.Errorf("Env.Get error :%v", err)
			return
		} else if actual != test.v {
			t.Errorf("got %v\nwant %v", actual, test.v)
			return
		}
	}

}
