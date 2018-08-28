package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/profile"
	"github.com/tesujiro/goa/debug"
	"github.com/tesujiro/goa/lib"
	"github.com/tesujiro/goa/parser"
	"github.com/tesujiro/goa/vm"
)

var FS = flag.String("F", " ", "Field separator")
var program_file = flag.String("f", "", "Program file")
var dbg = flag.Bool("d", false, "debug option")
var ast_dump = flag.Bool("a", false, "AST dump option")
var mem_prof = flag.Bool("m", false, "Memory Profile")
var cpu_prof = flag.Bool("c", false, "CPU Profile")
var ver = flag.Bool("v", false, "version")

const version = "0.0.0"

func main() {
	var file, script string
	flag.Parse()
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
		os.Exit(0)
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
			os.Exit(1)
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
			os.Exit(1)
		}
		defer file_reader.Close()
	} else {
		file_reader = os.Stdin
	}

	os.Exit(runScript(script_reader, file_reader))
}

func runScript(script_reader io.Reader, file_reader io.Reader) int {

	env := vm.NewEnv()
	env = lib.Import(env)
	if *dbg {
		env.Dump()
	}

	bytes, err := ioutil.ReadAll(script_reader)
	if err != nil {
		fmt.Printf("Read error: %v \n", err)
		//os.Exit(1)
		return 1
	}
	source := string(bytes)
	debug.Println("script:", source)
	l := new(parser.Lexer)

	fset := token.NewFileSet()                      // positions are relative to fset
	f := fset.AddFile("", fset.Base(), len(source)) // register input "file"
	l.Init(f, []byte(source), nil, scanner.ScanComments)

	ast, parseError := parser.Parse(l)
	if parseError != nil {
		fmt.Printf("Syntax error: %v \n", parseError)
		return 1
	}
	if *ast_dump {
		parser.Dump(ast)
	}

	//vm.Init() // TODO: NR=0
	env.SetFS(*FS)

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
		if *dbg {
			env.Dump()
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
