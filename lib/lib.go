// Package lib is a standard library for ago.
package lib

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/tesujiro/ago/ast"
	"github.com/tesujiro/ago/vm"
)

// Import imports standard library.
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

	regexpToStr := func(v reflect.Value) string {
		//fmt.Printf("v=%#v", v.Elem().Interface())
		if v.Type().Kind() == reflect.String {
			return v.Interface().(string)
		}
		switch v.Elem().Interface().(type) {
		case ast.RegExpr:
			//return v.Elem().Interface().(string)
			return v.Elem().FieldByName("Literal").String()
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
			}
			return i
		case reflect.Int:
			return v.Interface().(int)
		case reflect.Float64, reflect.Float32:
			return int(v.Interface().(float64))
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

	toFloat64 := func(v reflect.Value) float64 {
		switch v.Type().Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			return float64(v.Interface().(int))
		case reflect.Float64, reflect.Float32:
			return v.Interface().(float64)
		default:
			return float64(toInt(v))
		}
	}

	env.Define("println", reflect.ValueOf(fmt.Println))
	env.Define("printf", reflect.ValueOf(fmt.Printf))
	env.Define("sprintf", reflect.ValueOf(fmt.Sprintf))

	close := func(file string) int {

		err := env.CloseFile(file)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return 1
		}
		return 0
	}
	env.Define("close", reflect.ValueOf(close))

	sum := func(args ...int) int {
		var result int
		for _, v := range args {
			result += v
		}
		return result
	}
	env.Define("sum", reflect.ValueOf(sum))

	cat := func(args ...string) string {
		var result string
		for _, v := range args {
			result += v
		}
		return result
	}
	env.Define("cat", reflect.ValueOf(cat))

	system := func(command string) int {
		re := regexp.MustCompile("[ \t]+")
		cmdArray := re.Split(command, -1)
		cmd := exec.Command(cmdArray[0], cmdArray[1:]...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Printf("%v", err)
			return 1
		}
		if err := cmd.Start(); err != nil {
			fmt.Printf("%v", err)
			return 1
		}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := cmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				// This works on both Unix and Windows. Although package
				// syscall is generally platform dependent, WaitStatus is
				// defined for both Unix and Windows and in both cases has
				// an ExitStatus() method with the same signature.
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					return status.ExitStatus()
				}
			} else {
				fmt.Printf("%v", err)
				return 1
			}
		}
		return 0
	}
	env.Define("system", reflect.ValueOf(system))

	length := func(v_args ...reflect.Value) int { // TODO: reflect.Value => string
		var v reflect.Value
		if len(v_args) > 0 {
			v = v_args[0] // arg[0] is a pointer to var name
		} else {
			vString, _ := env.GetField(0)
			v = reflect.ValueOf(vString)
		}
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

	substr := func(str, begin reflect.Value, end_args ...reflect.Value) string { // TODO: reflect.Value => string
		var end reflect.Value
		if len(end_args) > 0 {
			end = end_args[0] // arg[0] is a pointer to var name
		} else {
			end = reflect.ValueOf(len(toStr(str)))
		}
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
	/*
		gsub := func(v1, v2, v3 reflect.Value) string {
			re := regexp.MustCompile(toStr(v1))
			result := re.ReplaceAllString(toStr(v3), toStr(v2))
			//fmt.Printf("lib:gsub v1:%v v2:%v v3:%v\tresult:%#v\n", v1, v2, v3, result)
			return result
		}
	*/
	gsub := func(before reflect.Value, after string, args ...*string) int {
		// PARSE ARGS
		var vName string
		var vVal interface{}
		var err error
		if len(args) > 0 {
			vName = *args[0] // arg[0] is a pointer to var name
			vVal, err = env.Get(vName)
			if err == vm.ErrUnknownSymbol {
				vVal = ""
			} else if err != nil { // TODO: unknown symbol
				fmt.Printf("err=%v\n", err)
				return 0
			}

		} else {
			//TODO: error
			vVal, _ = env.GetField(0)
		}
		// MAIN
		regexStr := regexpToStr(before)
		re := regexp.MustCompile(regexStr)
		result := re.ReplaceAllString(vVal.(string), after)
		if len(args) > 0 {
			vName = *args[0]
			err = env.Set(vName, result)
			if err != nil {
				return 0
			}
		} else {
			//TODO: error
			_ = env.SetField(0, result)
		}
		return len(re.FindAllString(vVal.(string), -1))
	}
	env.Define("gsub", reflect.ValueOf(gsub))

	sub := func(before reflect.Value, after string, args ...*string) int {
		regexStr := regexpToStr(before)
		regexStr = "^(.*?)" + regexStr + "(.*)$"
		return gsub(reflect.ValueOf(regexStr), "${1}"+after+"${2}", args...)
	}
	env.Define("sub", reflect.ValueOf(sub))

	match := func(s, r reflect.Value) int {
		//fmt.Printf("s=%v r=%#v\n", toStr(s), r)
		re := regexp.MustCompile(regexpToStr(r))
		loc := re.FindStringIndex(toStr(s))
		result := re.FindString(toStr(s))
		var retloc, retlen int
		if len(loc) > 0 {
			retloc = loc[0] + 1
			retlen = len(result)
		} else {
			retloc = 0
			retlen = -1
		}
		env.SetRSTART(retloc)
		env.SetRLENGTH(retlen)
		return retloc
	}
	env.Define("match", reflect.ValueOf(match))

	split := func(str string, array map[interface{}]interface{}, vars ...string) int {
		var sep string
		if len(vars) > 0 {
			sep = vars[0]
		} else {
			val, _ := env.Get("FS")
			sep = val.(string)
		}

		re := regexp.MustCompile(sep)
		result := re.Split(str, -1)
		for k, v := range result {
			array[fmt.Sprintf("%d", k+1)] = v
		}
		return len(result)
	}
	env.Define("split", reflect.ValueOf(split))

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

	mktime := func(datespec reflect.Value) int64 {
		//loc, _ := time.LoadLocation("Asia/Tokyo")
		loc, _ := time.LoadLocation("Local")
		t, err := time.ParseInLocation("2006 01 02 15 04 05", toStr(datespec), loc)
		if err != nil {
			return 0
		}
		return t.Unix()

	}
	env.Define("mktime", reflect.ValueOf(mktime))

	toInteger := func(v reflect.Value) int {
		return toInt(v)
	}
	env.Define("int", reflect.ValueOf(toInteger))

	random := func() float64 {
		return rand.Float64()
	}
	env.Define("rand", reflect.ValueOf(random))

	srandom := func() {
		rand.Seed(time.Now().UnixNano())
	}
	env.Define("srand", reflect.ValueOf(srandom))

	sqrt := func(arg reflect.Value) float64 {
		return math.Sqrt(toFloat64(arg))
	}
	env.Define("sqrt", reflect.ValueOf(sqrt))

	exp := func(arg reflect.Value) float64 {
		return math.Exp(toFloat64(arg))
	}
	env.Define("exp", reflect.ValueOf(exp))

	log := func(arg reflect.Value) float64 {
		return math.Log(toFloat64(arg))
	}
	env.Define("log", reflect.ValueOf(log))

	sin := func(arg reflect.Value) float64 {
		return math.Sin(toFloat64(arg))
	}
	env.Define("sin", reflect.ValueOf(sin))

	cos := func(arg reflect.Value) float64 {
		return math.Cos(toFloat64(arg))
	}
	env.Define("cos", reflect.ValueOf(cos))

	atan2 := func(arg1, arg2 reflect.Value) float64 {
		return math.Atan2(toFloat64(arg1), toFloat64(arg2))
	}
	env.Define("atan2", reflect.ValueOf(atan2))

	// Dynamic Func : use env in the func

	/*
		importShowEnv := func(env *vm.Env) (reflect.Value, error) {
			f := func() {
				env.Dump()
			}
			return reflect.ValueOf(f), nil
		}
		env.DefineImportFunc("env", importShowEnv)
	*/

	return env
}
