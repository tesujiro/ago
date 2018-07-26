package vm

import (
	"fmt"

	"github.com/tesujiro/goa/ast"
)

func RunBeginRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.BeginPattern:
			fmt.Println("BEGIN")
			result, err = runStmts(rule.Action, env)
			if err != nil {
				return
			}
		}
	}
	return
}

func RunMainRules(rules []ast.Rule, env *Env, line string) (result interface{}, err error) {
	env.incNR()
	if err := env.setFIELD(line); err != nil {
		return nil, err
	}
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.ExprPattern:
			fmt.Println("MAIN:", env.NR)
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
			fmt.Println("END")
			result, err = runStmts(rule.Action, env)
			if err != nil {
				return
			}
		}
	}
	return
}
