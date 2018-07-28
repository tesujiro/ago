package vm

import (
	"errors"
	"reflect"

	"github.com/tesujiro/goa/ast"
	"github.com/tesujiro/goa/debug"
)

func RunBeginRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.BeginPattern:
			debug.Println("BEGIN")
			result, err = runStmts(rule.Action, env)
			if err != nil {
				return
			}
		}
	}
	return
}

func RunMainRules(rules []ast.Rule, env *Env, line string, line_number int) (result interface{}, err error) {
	env.setNR(line_number)
	if err := env.SetField(line); err != nil {
		return nil, err
	}
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.ExprPattern:
			debug.Println(env.builtin.NR, ":MAIN")

			expr := rule.Pattern.(*ast.ExprPattern).Expr
			if expr != nil {
				if b, err := evalExpr(expr, env); err != nil {
					return result, err
				} else {
					if reflect.ValueOf(b).Kind() != reflect.Bool {
						err = errors.New("pattern is not bool")
						return result, err
					}

					if reflect.ValueOf(b).Interface() != true {
						debug.Printf("Line: %v skipped\n", env.builtin.NR)
						continue
					}
				}
			}

			result, err = runStmts(rule.Action, env)
			if err != nil {
				return
			}
		}
	}
	return
}

func RunEndRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.EndPattern:
			debug.Println("END")
			result, err = runStmts(rule.Action, env)
			if err != nil {
				return
			}
		}
	}
	return
}
