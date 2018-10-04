package vm

import (
	"fmt"
	"reflect"

	"github.com/tesujiro/ago/ast"
	"github.com/tesujiro/ago/debug"
)

func SeparateRules(rules []ast.Rule) (Func, Begin, Main, End []ast.Rule) {
	for _, rule := range rules {
		switch rule.Pattern.(type) {
		case *ast.FuncPattern:
			Func = append(Func, rule)
		case *ast.BeginPattern:
			Begin = append(Begin, rule)
		case *ast.ExprPattern, *ast.StartStopPattern:
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

func RunMainRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println(env.builtin.NR, ":MAIN")
		childEnv := env.NewEnv()
		switch rule.Pattern.(type) {
		case *ast.ExprPattern:
			expr := rule.Pattern.(*ast.ExprPattern).Expr
			if expr != nil {
				if result, err := evalExpr(expr, childEnv); err != nil {
					return result, err
				} else {
					b, err := strictToBool(result, "rule expression")
					if err != nil {
						return nil, err
					}
					//fmt.Printf("vmRule ExprPattern result:%#v bool:%v\n", result, b)
					if !b {
						debug.Printf("Line: %v skipped\n", childEnv.builtin.NR)
						continue
					}
				}
			}

			result, err = runStmts(rule.Action, childEnv)
			if err != nil {
				return
			}
		case *ast.StartStopPattern:
			pattern := rule.Pattern.(*ast.StartStopPattern)
			var b interface{}
			isMatch := func() (bool, error) {
				if !env.GetLoop() {
					b, err = evalExpr(pattern.Start, childEnv)
				} else {
					b, err = evalExpr(pattern.Stop, childEnv)
				}
				if err != nil {
					return false, err
				}
				if reflect.ValueOf(b).Kind() != reflect.Bool {
					err = fmt.Errorf("pattern is not bool: %v %v", reflect.ValueOf(b).Kind(), b)
					return false, err
				}
				return reflect.ValueOf(b).Interface().(bool), nil
			}
			match, err := isMatch()
			if err != nil {
				return nil, err
			}
			if match {
				env.SetLoop(!env.GetLoop())
				match, err := isMatch()
				if err != nil {
					return nil, err
				}
				if match {
					env.SetLoop(!env.GetLoop())
				}
			} else if !env.GetLoop() && !reflect.ValueOf(b).Interface().(bool) {
				debug.Printf("Line: %v skipped\n", childEnv.builtin.NR)
				continue
			}
			result, err = runStmts(rule.Action, childEnv)
			if err != nil {
				return nil, err
			}
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
