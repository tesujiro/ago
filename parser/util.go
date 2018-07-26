package parser

import (
	"fmt"
	"reflect"

	"github.com/tesujiro/exp_yacc/debug"
)

func Dump(obj interface{}) {
	p := func(indent string, obj interface{}) {
		fmt.Printf("%s%T\t%#v\n", indent, obj, obj)
	}
	pf := func(indent, name string, obj interface{}) {
		fmt.Printf("%s%s\t%T\t%#v\n", indent, name, obj, obj)
	}

	var dump_helper func(string, interface{})
	dump_helper = func(indent string, obj interface{}) {
		next_indent := indent + "\t"
		t := reflect.TypeOf(obj)
		v := reflect.ValueOf(obj)
		switch t.Kind() {
		case reflect.Ptr:
			debug.Println(indent, "pointer!!")
			dump_helper(indent, v.Elem().Interface())
			//p(indent, v.Elem().Interface()) //same
			//dump_helper(next_indent, v.Elem().Interface())
		case reflect.Interface:
			debug.Println(indent, "interface")
			p(indent, v.Elem().Interface())
		case reflect.Slice | reflect.Array:
			//case reflect.Array:
			debug.Println(indent, "slice|array")
			//p(indent, obj)
			for i := 0; i < v.Len(); i++ {
				//dump_helper(next_indent, v.Index(i).Interface())
				dump_helper(indent, v.Index(i).Interface())
			}
		case reflect.Struct:
			debug.Println(indent, "struct")
			//v = v.Elem()
			for i := 0; i < v.NumField(); i++ {
				pf(indent, t.Field(i).Name, v.Field(i).Interface())
				//if !v.Elem().IsNil() {
				if v.Field(i).Kind() != reflect.String && !v.Field(i).IsNil() {
					dump_helper(next_indent, v.Field(i).Interface())
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
		dump_helper("\t", v.Index(i).Interface())
	}
}
