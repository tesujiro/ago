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

	runScript(script_reader, file_reader)
}

func runScript(script_reader io.Reader, file_reader io.Reader) {

	env := vm.NewEnv()

	bytes, err := ioutil.ReadAll(script_reader)
	if err != nil {
		os.Exit(1)
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
		return
	}
	if *ast_dump {
		parser.Dump(ast)
	}

	//vm.Init() // TODO: NR=0
	env.SetFS(*FS)

	var result interface{}

	beginRules, mainRules, endRules := vm.SeparateRules(ast)

	// BEGIN
	result, err = vm.RunBeginRules(beginRules, env)
	debug.Printf("%#v\n", result)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	if len(mainRules) > 0 {
		// MAIN
		file_scanner := bufio.NewScanner(file_reader)
		var number int
		for file_scanner.Scan() {
			number++
			file_line := file_scanner.Text()
			result, err := vm.RunMainRules(mainRules, env, file_line, number)
			if err != nil {
				fmt.Printf("error:%v\n", err)
				return
			}
			//debug.Printf("ENV=%#v\n", env)
			//debug.Printf("%#v\n", res)
			if *dbg {
				env.Dump()
			}
			debug.Printf("%#v\n", result)
			/*
				for k, v := range env.GetField() {
					debug.Println("Field[", k, "]=\t", v)
				}
			*/
		}
	}

	// END
	result, err = vm.RunEndRules(endRules, env)
	debug.Printf("%#v\n", result)
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}

	return
}
