package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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
		return fmt.Errorf("parameter must be KEY=VALUE format")
	}
	key := z[0]
	value := z[1]
	kvs[key] = value
	return nil
}

var (
	fs, programFile                   string
	dbg, globalVar, dbglexer, astDump bool
	memProf, cpuProf, ver             bool
)
var variables hash = hash{}

const version = "0.0.0"

func init() {
	//flag.Var(&variables, "v", "followed by var=value, assign variable before execution")
}

func main() {
	os.Exit(_main())
}

func _main() int {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.StringVar(&fs, "F", " ", "Field separator")
	f.StringVar(&programFile, "f", "", "Program file")
	f.BoolVar(&dbg, "d", false, "debug option")
	f.BoolVar(&globalVar, "g", false, "global variable option")
	f.BoolVar(&dbglexer, "l", false, "debug lexer option")
	f.BoolVar(&astDump, "a", false, "AST dump option")
	f.BoolVar(&memProf, "m", false, "Memory Profile")
	f.BoolVar(&cpuProf, "c", false, "CPU Profile")
	f.BoolVar(&ver, "version", false, "version")
	f.Var(&variables, "v", "followed by var=value, assign variable before execution")
	err := f.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("argument parse err:%v\n", err)
		return 1
	}
	args := f.Args()

	if ver {
		fmt.Println("Version:", version)
		return 0
	}

	if len(args) == 0 && programFile == "" {
		fmt.Printf("Usage of %s:\n", path.Base(os.Args[0]))
		f.PrintDefaults()
		return 1
	}
	var script string
	var files []string
	if len(args) > 0 {
		if programFile == "" {
			script = args[0]
			files = args[1:]
		} else {
			files = args
		}
	}
	if len(files) == 0 {
		files = []string{""} // STDIN
	}

	if dbg {
		fmt.Println("Start debug mode.")
		debug.On()
	} else {
		debug.Off()
	}
	if cpuProf {
		defer profile.Start().Stop()
	}
	if memProf {
		defer profile.Start(profile.MemProfile).Stop()
	}

	var ret int
	runFile := func(file string) int {
		var sriptReader io.Reader
		if programFile != "" {
			//fmt.Println("read from programFile:", programFile)
			fp, err := os.Open(programFile)
			if err != nil {
				fmt.Println("script file open error:", err)
				return 1
			}
			defer fp.Close()
			sriptReader = bufio.NewReader(fp)
		} else {
			//fmt.Println("read script:", script)
			sriptReader = strings.NewReader(script)
		}
		var fileReader *os.File
		if file != "" {
			fileReader, err = os.Open(file)
			if err != nil {
				fmt.Println("input file open error:", err)
				return 1
			}
			defer fileReader.Close()
		} else {
			fileReader = os.Stdin
		}
		return runScript(sriptReader, fileReader)
	}

	for _, file := range files {
		ret = runFile(file)
		if ret != 0 {
			return ret
		}
	}
	return 0
}

func initEnv() *vm.Env {
	env := vm.NewEnv()
	env = lib.Import(env)
	env.SetFS(fs)

	if globalVar {
		vm.SetGlobalVariables()
	}

	for k, v := range variables {
		env.Set(k, v)
	}

	return env
}

func runScript(sriptReader io.Reader, fileReader *os.File) int {

	env := initEnv()
	if dbg {
		env.Dump()
	}

	bytes, err := ioutil.ReadAll(sriptReader)
	if err != nil {
		fmt.Printf("Read error: %v \n", err)
		return 1
	}
	source := string(bytes)
	debug.Println("script:", source)

	if dbglexer {
		fmt.Println("Start lexer debug mode.")
		parser.TraceLexer()
	} else {
		parser.TraceOffLexer()
	}

	ast, parseError := parser.ParseSrc(source)
	if parseError != nil {
		fmt.Printf("Syntax error: %v \n", parseError)
		if e, ok := parseError.(*parser.Error); ok {
			fmt.Printf("at Line %v Column %v\n", e.Pos.Line, e.Pos.Column)
			line := strings.Split(source, "\n")[e.Pos.Line-1]
			fmt.Println(line)
			for i := 1; i < e.Pos.Column; i++ {
				fmt.Printf(" ")
			}
			fmt.Println("^")
		}
		//e := parseError.Error()
		return 1
	}
	if astDump {
		parser.Dump(ast)
	}

	var fileScanner *bufio.Scanner
	redir := "-" // a kind of stdin
	rc := io.ReadCloser(fileReader)
	fileScanner, err = env.SetFile(redir, &rc)
	if err != nil {
		fmt.Printf("env error: %v \n", err)
		return 1
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
		}
		return 1
	}
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return 1
	}
	if dbg {
		env.Dump()
	}

	if len(mainRules) == 0 && len(endRules) == 0 {
		return 0
	}

	// reset variable
	env.SetNF()

	// MAIN
	//fileScanner := bufio.NewScanner(fileReader)
	var number int
	for fileScanner.Scan() {
		number++
		fileLine := fileScanner.Text()
		env.SetNR(number)
		env.SetFieldFromLine(fileLine)
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
			if dbg {
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
