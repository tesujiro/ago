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

	/*
		var sourceBytes []byte
		var err error
		if file != "" {
			sourceBytes, err = ioutil.ReadFile(file)
			if err != nil {
				fmt.Println("ReadFile error:", err)
				return
			}
		} else {
			sourceBytes, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("Read Stdin error:", err)
				return
			}
		}
		source := string(sourceBytes)
	*/
	runScriptFile(script, file)
}

func runScriptFile(source, file string) {

	env := vm.NewEnv()
	_ = env

	fmt.Println("script:", source)
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
	/*
		if res, err := vm.Run(ast, env); err != nil {
			fmt.Printf("Eval error:%v\n", err)
			return
		} else {
			debug.Printf("ENV=%#v\n", env)
			fmt.Printf("%#v\n", res)
		}
	*/
	return
}

func run() {
	env := vm.NewEnv()
	_ = env
	line_scanner := bufio.NewScanner(os.Stdin) // This Scanner
	var source string

	for line_scanner.Scan() {
		source += line_scanner.Text()
		//if source == "" {
		//continue
		//}
		if source == "exit" || source == "quit" {
			break
		}

		l := new(parser.Lexer)
		//l.Init(strings.NewReader(source))

		fset := token.NewFileSet()                         // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(source)) // register input "file"
		l.Init(file, []byte(source), nil /* no error handler */, scanner.ScanComments)

		ast, parseError := parser.Parse(l)
		/*
			for _, stmt := range ast {
				debug.Printf("%#v\n", stmt)
			}
		*/
		if *ast_dump {
			parser.Dump(ast)
		}
		if parseError != nil {
			debug.Println("[", parseError.Error(), "]")
			//if parseError.Error() == "unexpected $end" || parseError.Error() == "comment not terminated" { //Does not work
			if parseError.Error() == "unexpected $end" {
				// note: scanner.Scan() does not return "end of line" ,
				// this is just for separating tokens
				source += "\n"
				//fmt.Println("source;[" + source + "]")
				continue
			} else {
				fmt.Printf("Syntax error: %v \n", parseError)
				//fmt.Printf("Syntax error: %v at %v\n", e, l.Position) //TODO
			}
		}
		/*
			if res, err := vm.Run(ast, env); err != nil {
				fmt.Printf("Eval error:%v\n", err)
			} else {
				debug.Printf("ENV=%#v\n", env)
				fmt.Printf("%#v\n", res)
			}
		*/
		source = ""
	}
}
