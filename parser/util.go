package parser

import (
	"fmt"
	"reflect"

	"github.com/tesujiro/ago/debug"
)

// Dump provide dump AST function.
func Dump(obj interface{}) {
	p := func(indent string, obj interface{}) {
		fmt.Printf("%s%T\t%#v\n", indent, obj, obj)
	}
	pf := func(indent, name string, obj interface{}) {
		fmt.Printf("%s%s\t%T\t%#v\n", indent, name, obj, obj)
	}

	var dumpHelper func(string, interface{})
	dumpHelper = func(indent string, obj interface{}) {
		nextIndent := indent + "\t"
		t := reflect.TypeOf(obj)
		v := reflect.ValueOf(obj)
		switch t.Kind() {
		case reflect.Ptr:
			debug.Println(indent, "pointer!!")
			dumpHelper(indent, v.Elem().Interface())
		//case reflect.Interface:
		//debug.Println(indent, "interface")
		//p(indent, v.Elem().Interface())
		case reflect.Slice | reflect.Array:
			debug.Println(indent, "slice|array")
			for i := 0; i < v.Len(); i++ {
				dumpHelper(indent, v.Index(i).Interface())
			}
		case reflect.Struct:
			debug.Println(indent, "struct")
			pf(indent, t.String(), v.Interface())
			for i := 0; i < v.NumField(); i++ {
				pf(indent+"\t", t.Field(i).Name, v.Field(i).Interface())
				if v.Field(i).Kind() != reflect.String &&
					v.Field(i).Kind() != reflect.Struct &&
					v.Field(i).Kind() != reflect.Bool &&
					!v.Field(i).IsNil() {
					dumpHelper(nextIndent, v.Field(i).Interface())
				}
			}
		default:
			debug.Println(indent, "default Kind():", t.Kind())
			//p(indent, obj)
		}
	}
	//t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.Len(); i++ {
		p("", v.Index(i).Interface())
		dumpHelper("\t", v.Index(i).Interface())
	}
}
