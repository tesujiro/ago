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
	"strings"
	"syscall"
	"time"

	"github.com/tesujiro/ago/ast"
	"github.com/tesujiro/ago/vm"
)

func toStr(v reflect.Value) string {
	return vm.ToString(v.Interface()).(string)
}

func regexpToStr(v reflect.Value) string {
	//fmt.Printf("v=%#v", v.Elem().Interface())
	if v.Type().Kind() == reflect.String {
		return v.Interface().(string)
	}
	switch v.Elem().Interface().(type) {
	case ast.ConstExpr, ast.StringExpr:
		//return v.Elem().Interface().(string)
		return v.Elem().FieldByName("Literal").String()
	default:
		return ""
	}
}

func toInt(v reflect.Value) int {
	return vm.ToInt(v.Interface()).(int)
}

func toInt64(v reflect.Value) int64 {
	switch v.Type().Kind() {
	case reflect.Int64:
		return v.Interface().(int64)
	default:
		return int64(toInt(v))
	}
}

func toFloat64(v reflect.Value) float64 {
	return vm.ToFloat64(v.Interface()).(float64)
}

func updateArgs(format string, a ...interface{}) []interface{} {
	fmtSpec := `%\d*(\.\d+)?[d|e|f|g|o|x|c|s]`
	re := regexp.MustCompile(fmtSpec)
	specifiers := re.FindAllString(format, -1)
	//fmt.Printf("specifiers=%v\n", specifiers)
	//fmt.Printf("a=%v\n", a)
	for i, spec := range specifiers {
		if i > len(a)-1 {
			break
		}
		switch spec[len(spec)-1] {
		case 'd':
			a[i] = vm.ToInt(a[i].(reflect.Value).Interface())
		case 'e', 'f', 'g':
			a[i] = vm.ToFloat64(a[i].(reflect.Value).Interface())
		case 's':
			a[i] = vm.ToString(a[i].(reflect.Value).Interface())
		}
	}
	return a
}

func importPrintf(env *vm.Env) {
	printf := func(format string, a ...interface{}) (n int, err error) {
		return fmt.Printf(format, updateArgs(format, a...)...)
	}
	env.Define("printf", reflect.ValueOf(printf))

	sprintf := func(format string, a ...interface{}) string {
		return fmt.Sprintf(format, updateArgs(format, a...)...)
	}
	env.Define("sprintf", reflect.ValueOf(sprintf))
}

func importClose(env *vm.Env) {
	close := func(file string) int {
		err := env.CloseFile(file)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return 1
		}
		return 0
	}
	env.Define("close", reflect.ValueOf(close))
}

func sum(args ...int) int {
	var result int
	for _, v := range args {
		result += v
	}
	return result
}

func cat(args ...string) string {
	var result string
	for _, v := range args {
		result += v
	}
	return result
}

func system(command string) int {
	re := regexp.MustCompile("[ \t]+")
	cmdArray := re.Split(command, -1)
	cmd := exec.Command(cmdArray[0], cmdArray[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("%v\n", err)
		return 1
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("%v\n", err)
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
			fmt.Printf("%v\n", err)
			return 1
		}
	}
	return 0
}

func substr(str, begin reflect.Value, endArgs ...reflect.Value) string { // TODO: reflect.Value => string
	var end reflect.Value
	if len(endArgs) > 0 {
		end = endArgs[0] // arg[0] is a pointer to var name
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

func importLength(env *vm.Env) {
	length := func(v_args ...reflect.Value) int { // TODO: reflect.Value => string
		var v reflect.Value
		if len(v_args) > 0 {
			v = v_args[0] // arg[0] is a pointer to var name
		} else {
			vString, _ := env.GetFieldZero()
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
}

func index(v1, v2 reflect.Value) int {
	s := toStr(v1)
	substr := toStr(v2)
	if len(s) == 0 {
		return 0
	}
	return strings.Index(s, substr) + 1
}

func tolower(v1 reflect.Value) string {
	return strings.ToLower(toStr(v1))
}

func toupper(v1 reflect.Value) string {
	return strings.ToUpper(toStr(v1))
}

func importSubGsub(env *vm.Env) {
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
	regexReplace := func(regexStr string, after string, args ...*string) int {
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
			vVal, _ = env.GetFieldZero()
		}
		// MAIN
		//regexStr := "(" + regexpToStr(before) + ")" // add parenthes for '&' meta chars
		re := regexp.MustCompile(regexStr)
		match := re.FindAllString(vVal.(string), -1)
		//after = strings.ReplaceAll(after, "&", "${1}") // replace '&' meta char
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
		return len(match)
	}

	gsub := func(before reflect.Value, after string, args ...*string) int {
		regexStr := "(" + regexpToStr(before) + ")" // add parenthes for '&' meta chars
		//after = strings.Replace(after, "&", "${1}", -1) // replace '&' meta char
		re := regexp.MustCompile(`(^|[^\\])\&`)
		after = re.ReplaceAllString(after, "${1}$${1}") // replace '&' meta char
		after = strings.Replace(after, "\\&", "&", -1)
		return regexReplace(regexStr, after, args...)
	}
	env.Define("gsub", reflect.ValueOf(gsub))

	sub := func(before reflect.Value, after string, args ...*string) int {
		regexStr := "^(.*?)(" + regexpToStr(before) + ")(.*)$"
		after = "${1}" + after + "${3}"
		//after = strings.Replace(after, "&", "${2}", -1) // replace '&' meta char
		re := regexp.MustCompile(`(^|[^\\])\&`)
		after = re.ReplaceAllString(after, "${1}$${2}") // replace '&' meta char
		after = strings.Replace(after, "\\&", "&", -1)
		return regexReplace(regexStr, after, args...)
	}
	env.Define("sub", reflect.ValueOf(sub))
}

func importMatch(env *vm.Env) {
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
}

func importSplit(env *vm.Env) {
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
}

func strftime(format, timestamp reflect.Value) string {
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

func mktime(datespec reflect.Value) int64 {
	//loc, _ := time.LoadLocation("Asia/Tokyo")
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006 01 02 15 04 05", toStr(datespec), loc)
	if err != nil {
		return 0
	}
	return t.Unix()

}

// Import imports standard library.
func Import(env *vm.Env) *vm.Env {
	env.Define("println", reflect.ValueOf(fmt.Println))
	importPrintf(env)

	importClose(env)
	env.Define("sum", reflect.ValueOf(sum))
	env.Define("cat", reflect.ValueOf(cat))
	env.Define("system", reflect.ValueOf(system))

	env.Define("substr", reflect.ValueOf(substr))
	importLength(env)
	env.Define("index", reflect.ValueOf(index))
	env.Define("tolower", reflect.ValueOf(tolower))
	env.Define("toupper", reflect.ValueOf(toupper))

	importSubGsub(env)
	importMatch(env)
	importSplit(env)

	systime := func() int64 {
		return time.Now().Unix()
	}
	env.Define("systime", reflect.ValueOf(systime))
	env.Define("strftime", reflect.ValueOf(strftime))
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
