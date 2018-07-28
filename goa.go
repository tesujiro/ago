package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"os"

	"github.com/pkg/profile"
	"github.com/tesujiro/goa/debug"
	"github.com/tesujiro/goa/parser"
	"github.com/tesujiro/goa/vm"
)

var FS = flag.String("f", " ", "Field separator") //TODO: REGEX
var dbg = flag.Bool("d", false, "debug option")
var ast_dump = flag.Bool("a", false, "AST dump option")
var mem_prof = flag.Bool("m", false, "Memory Profile")
var cpu_prof = flag.Bool("c", false, "CPU Profile")
var ver = flag.Bool("v", false, "version")

const version = "0.0.0"

func main() {
	flag.Parse()
	script := flag.Arg(0)
	file := flag.Arg(1)

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

	runScript(script, file)
}

func openInputFile(f string) (fp *os.File, err error) {
	if f != "" {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println("file open error:", err)
			return nil, err
		}
		return fp, nil
	} else {
		return os.Stdin, nil
	}
}

func runScript(source string, file string) {

	env := vm.NewEnv()

	debug.Println("script:", source)
	l := new(parser.Lexer)
	//l.Init(strings.NewReader(source))

	fset := token.NewFileSet()                      // positions are relative to fset
	f := fset.AddFile("", fset.Base(), len(source)) // register input "file"
	l.Init(f, []byte(source), nil /* no error handler */, scanner.ScanComments)

	ast, parseError := parser.Parse(l)
	if parseError != nil {
		fmt.Printf("Syntax error: %v \n", parseError)
		//fmt.Printf("Syntax error: %v at %v\n", e, l.Position) //TODO
		return
	}
	if *ast_dump {
		parser.Dump(ast)
	}

	//vm.Init() // TODO: NR=0
	env.SetFS(*FS)

	var result interface{}
	var err error

	// BEGIN
	result, err = vm.RunBeginRules(ast, env)
	debug.Printf("%#v\n", result)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	//TODO:Check There is main rules
	fp, err := openInputFile(file)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	if fp != os.Stdin {
		defer fp.Close()
	}

	// MAIN
	file_scanner := bufio.NewScanner(fp)
	var number int
	for file_scanner.Scan() {
		number++
		file_line := file_scanner.Text()
		result, err := vm.RunMainRules(ast, env, file_line, number)
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

	// END
	result, err = vm.RunEndRules(ast, env)
	debug.Printf("%#v\n", result)
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}

	return
}
