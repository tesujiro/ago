package lib

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tesujiro/goa/vm"
)

func Import(env *vm.Env) *vm.Env {
	toStr := func(v reflect.Value) string {
		switch v.Type().Kind() {
		case reflect.String:
			return v.Interface().(string)
		case reflect.Int:
			return fmt.Sprintf("%v", v.Interface().(int))
		case reflect.Float64:
			return fmt.Sprintf("%v", v.Interface().(float64))
		default:
			return ""
		}
	}

	toInt := func(v reflect.Value) int {
		switch v.Type().Kind() {
		case reflect.String:
			i, err := strconv.Atoi(v.Interface().(string))
			if err != nil {
				return 0
			} else {
				return i
			}
		case reflect.Int:
			return v.Interface().(int)
		default:
			return 0
		}
	}

	toInt64 := func(v reflect.Value) int64 {
		switch v.Type().Kind() {
		case reflect.Int64:
			return v.Interface().(int64)
		default:
			return int64(toInt(v))
		}
	}

	length := func(v reflect.Value) int {
		switch v.Type().Kind() {
		case reflect.Int:
			return len(toStr(v))
		case reflect.String:
			return len(toStr(v))
		case reflect.Map:
			s := v.Interface().(map[interface{}]interface{})
			return len(s)
		case reflect.Slice:
			s := v.Interface().([]interface{})
			return len(s)
		default:
			// fmt.Errorf("invalid argument %v (type %v) for len",s,reflect.TypeOf(s))
			return 0
		}
	}
	env.Define("length", reflect.ValueOf(length))
	env.Define("len", reflect.ValueOf(length))

	substr := func(str, begin, end reflect.Value) string {
		s := toStr(str)
		b := toInt(begin)
		e := toInt(end)
		var from, to int
		if b > 0 {
			from = b
		} else {
			from = 1
		}
		if from+e < len(s)+1 {
			//fmt.Printf("path1:")
			to = from + e
		} else {
			//fmt.Printf("path2:")
			to = len(s) + 1
		}
		if len(s) == 0 || from >= to {
			return ""
		}
		return s[from-1 : to-1]
	}
	env.Define("substr", reflect.ValueOf(substr))

	index := func(v1, v2 reflect.Value) int {
		s := toStr(v1)
		substr := toStr(v2)
		if len(s) == 0 {
			return 0
		}
		return strings.Index(s, substr) + 1
	}
	env.Define("index", reflect.ValueOf(index))

	tolower := func(v1 reflect.Value) string {
		return strings.ToLower(toStr(v1))
	}
	env.Define("tolower", reflect.ValueOf(tolower))

	toupper := func(v1 reflect.Value) string {
		return strings.ToUpper(toStr(v1))
	}
	env.Define("toupper", reflect.ValueOf(toupper))

	// TODO: NOT SAME SPEC AS AWK gsub
	// AWK : function call args by reference
	gsub := func(v1, v2, v3 reflect.Value) string {
		re := regexp.MustCompile(toStr(v1))
		result := re.ReplaceAllString(toStr(v3), toStr(v2))
		//fmt.Printf("lib:gsub v1:%v v2:%v v3:%v\tresult:%#v\n", v1, v2, v3, result)
		return result
	}
	env.Define("gsub", reflect.ValueOf(gsub))

	sub := func(s, g, fs reflect.Value) string {
		return gsub(reflect.ValueOf("^(.*?)"+toStr(s)+"(.*)$"), reflect.ValueOf("${1}"+toStr(g)+"${2}"), fs)
	}
	env.Define("sub", reflect.ValueOf(sub))

	//TODO: how can i set RSTART,RLENGTH -> builtin function
	match := func(s, r reflect.Value) int {
		re := regexp.MustCompile(toStr(r))
		loc := re.FindStringIndex(toStr(s))
		if len(loc) > 0 {
			return loc[0] + 1
		} else {
			return 0
		}
	}
	env.Define("match", reflect.ValueOf(match))

	/*
		split := func(s, g, fs reflect.Value) int {
		}
		env.Define("split", reflect.ValueOf(split))
	*/

	systime := func() int64 {
		return time.Now().Unix()
	}
	env.Define("systime", reflect.ValueOf(systime))

	strftime := func(format, timestamp reflect.Value) string {
		table := map[string]string{
			"%Y": "2006", "%y": "06",
			"%m": "01",
			"%d": "02",
			"%H": "15",
			"%M": "04",
			"%S": "05",
		}
		f := toStr(format)
		for k, v := range table {
			f = strings.Replace(f, k, v, -1)
		}

		//fmt.Printf("timestamp=%#v\ntimestamp.Kind=%#v\n", timestamp, timestamp.Kind().String())
		t64 := toInt64(timestamp)
		u := time.Unix(t64, 0)
		return u.Format(f)
	}
	env.Define("strftime", reflect.ValueOf(strftime))

	/*
		mktime := func(datespec reflect.Value) int {

		}
		env.Define("mktime", reflect.ValueOf(mktime))
	*/

	return env
}
