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
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", path.Base(os.Args[0]))
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
	/*
		if len(files) == 0 {
			files = []string{""} // STDIN
		}
	*/

	if dbg {
		fmt.Println("Start debug mode.")
		debug.On()
	} else {
		debug.Off()
	}
	if cpuProf {
		defer profile.Start(profile.ProfilePath(".")).Stop()
	}
	if memProf {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}

	var ret int
	//runFile := func(env *vm.Env, file string) int {
	run := func(env *vm.Env) int {
		var scriptReader io.Reader
		if programFile != "" {
			//fmt.Println("read from programFile:", programFile)
			fp, err := os.Open(programFile)
			if err != nil {
				fmt.Println("script file open error:", err)
				return 1
			}
			defer fp.Close()
			scriptReader = bufio.NewReader(fp)
		} else {
			//fmt.Println("read script:", script)
			scriptReader = strings.NewReader(script)
		}

		/*
			//TODO: refact: open file in runScript
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
			return runScript(env, scriptReader, file, fileReader)
		*/
		return runScript(env, scriptReader)
	}

	env := initEnv(files)
	if dbg {
		env.Dump()
	}

	/*
		for _, file := range files {
			ret = runFile(env, file)
			if ret != 0 {
				return ret
			}
		}
	*/
	ret = run(env)
	if ret != 0 {
		return ret
	}
	return 0
}

func initEnv(files []string) *vm.Env {
	env := vm.NewEnv(files)
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

//func runScript(env *vm.Env, scriptReader io.Reader, file string, fileReader *os.File) int {
func runScript(env *vm.Env, scriptReader io.Reader) int {

	//fmt.Println("runScript")
	bytes, err := ioutil.ReadAll(scriptReader)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
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

	parser.EnableErrorVerbose()
	ast, parseError := parser.ParseSrc(source)
	if parseError != nil {
		fmt.Printf("Syntax error: %v\n", parseError)
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

	/*
		var fileScanner *bufio.Scanner
		rc := io.ReadCloser(fileReader)
		fileScanner, err = env.SetFile(file, &rc)
		if err != nil {
			fmt.Printf("env error: %v \n", err)
			return 1
		}
		env.SetFILENAME(file)
	*/

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
		/*
			err = env.CloseFile(file)
			if err != nil {
				fmt.Printf("close file error: %v \n", err)
				return 1
			}
		*/

		return 0
	}

	// reset variable
	env.SetNF()

	// MAIN
	/*
		var number int
		for fileScanner.Scan() {
			number++
			fileLine := fileScanner.Text()
			env.IncNR()
			env.SetFNR(number)
	*/
	//var err error
	for {
		//fmt.Println("MAINLOOP")
		fileLine, err := env.GetLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error:%v\n", err)
			return 1
		}
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

	/*
		err = env.CloseFile(file)
		if err != nil {
			fmt.Printf("close file error: %v \n", err)
			return 1
		}
	*/

	return 0
}
