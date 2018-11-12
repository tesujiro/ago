package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/profile"
	"github.com/tesujiro/ago/debug"
	"github.com/tesujiro/ago/lib"
	"github.com/tesujiro/ago/parser"
	"github.com/tesujiro/ago/vm"
)

type hash map[string]string

func (kvs hash) String() string {
	var s string
	for k, v := range kvs {
		s = fmt.Sprintf("%s %s=%s", s, k, v)
	}
	return s
}

func (kvs hash) Set(s string) error {
	z := strings.SplitN(s, "=", 2)
	if len(z) < 2 {
		return fmt.Errorf("parameter must be KEY=VALUE format :%v", s)
	}
	key := z[0]
	value := z[1]
	_, ok := kvs[key]
	if ok {
		kvs[key] = value
	} else {
		kvs[key] = value
	}
	return nil
}

var FS = flag.String("F", " ", "Field separator")
var program_file = flag.String("f", "", "Program file")
var dbg = flag.Bool("d", false, "debug option")
var globalVar = flag.Bool("g", false, "global variable option")
var dbglexer = flag.Bool("l", false, "debug lexer option")
var ast_dump = flag.Bool("a", false, "AST dump option")
var mem_prof = flag.Bool("m", false, "Memory Profile")
var cpu_prof = flag.Bool("c", false, "CPU Profile")
var ver = flag.Bool("version", false, "version")
var variables hash = hash{}

const version = "0.0.0"

func main() {
	flag.Var(&variables, "v", "followed by var=value, assign variable before execution")
	flag.Parse()
	os.Exit(_main())
}

func _main() int {
	var file, script string
	switch len(flag.Args()) {
	case 1:
		if *program_file != "" {
			file = flag.Arg(0)
		} else {
			script = flag.Arg(0)
		}
	case 2:
		script = flag.Arg(0)
		file = flag.Arg(1)
	}

	if *ver {
		fmt.Println("Version:", version)
		return 0
	}

	if *dbg {
		debug.On()
	}
	if *cpu_prof {
		defer profile.Start().Stop()
	}
	if *mem_prof {
		defer profile.Start(profile.MemProfile).Stop()
	}

	var script_reader io.Reader
	if *program_file != "" {
		fmt.Println("program_file=", *program_file)
		fp, err := os.Open(*program_file)
		if err != nil {
			fmt.Println("script file open error:", err)
			return 1
		}
		defer fp.Close()
		script_reader = bufio.NewReader(fp)
	} else {
		script_reader = strings.NewReader(script)
	}

	var file_reader io.Reader
	if file != "" {
		file_reader, err := os.Open(file)
		if err != nil {
			fmt.Println("input file open error:", err)
			return 1
		}
		defer file_reader.Close()
	} else {
		file_reader = os.Stdin
	}

	return runScript(script_reader, file_reader)
}

func initEnv() *vm.Env {
	env := vm.NewEnv()
	env = lib.Import(env)
	env.SetFS(*FS)

	if *globalVar {
		vm.SetGlobalVariables()
	}

	for k, v := range variables {
		env.Set(k, v)
	}

	return env
}

func runScript(script_reader io.Reader, file_reader io.Reader) int {

	env := initEnv()

	bytes, err := ioutil.ReadAll(script_reader)
	if err != nil {
		fmt.Printf("Read error: %v \n", err)
		return 1
	}
	source := string(bytes)
	debug.Println("script:", source)

	if *dbglexer {
		parser.TraceLexer()
	}
	ast, parseError := parser.ParseSrc(source)
	if parseError != nil {
		fmt.Printf("Syntax error: %v \n", parseError)
		return 1
	}
	if *ast_dump {
		parser.Dump(ast)
	}

	var result interface{}

	funcRules, beginRules, mainRules, endRules := vm.SeparateRules(ast)

	// FUNC DEFINITION
	if len(funcRules) > 0 {
		result, err = vm.RunFuncRules(funcRules, env)
		debug.Printf("%#v\n", result)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return 1
		}
	}

	// BEGIN
	result, err = vm.RunBeginRules(beginRules, env)
	debug.Printf("%#v\n", result)
	if err == vm.ErrExit {
		v, ok := result.(int)
		if ok {
			return v
		} else {
			return 1
		}
	}
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return 1
	}
	if *dbg {
		env.Dump()
	}

	if len(mainRules) == 0 && len(endRules) == 0 {
		return 0
	}

	// MAIN
	file_scanner := bufio.NewScanner(file_reader)
	var number int
	for file_scanner.Scan() {
		number++
		file_line := file_scanner.Text()
		env.SetNR(number)
		if err := env.SetFieldFromLine(file_line); err != nil {
			fmt.Printf("error:%v\n", err)
			return 1
		}
		if len(mainRules) > 0 {
			result, err := vm.RunMainRules(mainRules, env)
			if err == vm.ErrNext {
				continue
			}
			if err == vm.ErrExit {
				return result.(int)
			}
			if err != nil {
				fmt.Printf("error:%v\n", err)
				return 1
			}
			//debug.Printf("ENV=%#v\n", env)
			//debug.Printf("%#v\n", res)
			if *dbg {
				env.Dump()
			}
			debug.Printf("%#v\n", result)
		}
	}

	// END
	result, err = vm.RunEndRules(endRules, env)
	debug.Printf("%#v\n", result)
	if err == vm.ErrExit {
		return result.(int)
	}
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return 1
	}

	return 0
}
