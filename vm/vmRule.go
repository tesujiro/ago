package vm

import (
	"fmt"
	"reflect"

	"github.com/tesujiro/goa/ast"
	"github.com/tesujiro/goa/debug"
)

func SeparateRules(rules []ast.Rule) (Func, Begin, Main, End []ast.Rule) {
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.FuncPattern:
			Func = append(Func, rule)
		case *ast.BeginPattern:
			Begin = append(Begin, rule)
		case *ast.ExprPattern:
			Main = append(Main, rule)
		case *ast.EndPattern:
			End = append(End, rule)
		}
	}
	return
}

func RunFuncRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println("FUNC")

		funcExpr := &ast.FuncExpr{Name: rule.Pattern.(*ast.FuncPattern).Name, Args: rule.Pattern.(*ast.FuncPattern).Args, Stmts: rule.Action}
		result, err = evalExpr(funcExpr, env)
		if err != nil {
			return
		}
	}
	return
}

func RunBeginRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println("BEGIN")
		childEnv := env.NewEnv()
		result, err = runStmts(rule.Action, childEnv)
		if err != nil {
			return
		}
	}
	return
}

func RunMainRules(rules []ast.Rule, env *Env, line string, line_number int) (result interface{}, err error) {
	env.setNR(line_number)
	if err := env.SetFieldFromLine(line); err != nil {
		return nil, err
	}
	for _, rule := range rules {
		debug.Println(env.builtin.NR, ":MAIN")
		childEnv := env.NewEnv()
		expr := rule.Pattern.(*ast.ExprPattern).Expr
		if expr != nil {
			if b, err := evalExpr(expr, childEnv); err != nil {
				return result, err
			} else {
				if reflect.ValueOf(b).Kind() != reflect.Bool {
					err = fmt.Errorf("pattern is not bool: %v %v", reflect.ValueOf(b).Kind(), b)
					return result, err
				}

				if reflect.ValueOf(b).Interface() != true {
					debug.Printf("Line: %v skipped\n", childEnv.builtin.NR)
					continue
				}
			}
		}

		result, err = runStmts(rule.Action, childEnv)
		if err != nil {
			return
		}
	}
	return
}

func RunEndRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println("END")
		childEnv := env.NewEnv()
		result, err = runStmts(rule.Action, childEnv)
		if err != nil {
			return
		}
	}
	return
}
