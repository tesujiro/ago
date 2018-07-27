package vm

import (
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

func RunMainRules(rules []ast.Rule, env *Env, line string) (result interface{}, err error) {
	env.incNR()
	if err := env.SetField(line); err != nil {
		return nil, err
	}
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.ExprPattern:
			debug.Println(env.builtin.NR, ":MAIN:")
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
