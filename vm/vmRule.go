package vm

import (
	"fmt"
	"reflect"

	"github.com/tesujiro/ago/ast"
	"github.com/tesujiro/ago/debug"
)

// SeparateRules classifies rules to func, begin, main and end rules.
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

// RunFuncRules executes func rules with a specified env.
func RunFuncRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println("FUNC")

		funcExpr := &ast.FuncExpr{Name: rule.Pattern.(*ast.FuncPattern).Name, Args: rule.Pattern.(*ast.FuncPattern).Args, Stmts: rule.Action}
		result, err = evalExpr(funcExpr, env)
		if err != nil {
			return toInt(result), err
		}
	}
	return toInt(result), nil
}

// RunBeginRules executes begin rules with a specified env.
func RunBeginRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println("BEGIN")
		childEnv := env.NewEnv()
		result, err = runStmts(rule.Action, childEnv)
		if err != nil {
			return toInt(result), err
		}
	}
	return toInt(result), err
}

// RunMainRules executes main rules with a specified env.
func RunMainRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println(env.builtin.NR, ":MAIN")
		childEnv := env.NewEnv()
		switch pattern := rule.Pattern.(type) {
		case *ast.ExprPattern:
			expr := pattern.Expr
			if expr != nil {
				result, err := evalExpr(expr, childEnv)
				if err != nil {
					return toInt(result), err
				}
				b, err := strictToBool(result)
				if err != nil {
					return nil, fmt.Errorf("convert rule expression:%v", err)
				}
				//fmt.Printf("vmRule ExprPattern result:%#v bool:%v\n", result, b)
				if !b {
					debug.Printf("Line: %v skipped\n", childEnv.builtin.NR)
					continue
				}
			}

			result, err = runStmts(rule.Action, childEnv)
			if err != nil {
				return toInt(result), err
			}
		case *ast.StartStopPattern:
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
	return toInt(result), err
}

// RunEndRules executes end rules with a specified env.
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
