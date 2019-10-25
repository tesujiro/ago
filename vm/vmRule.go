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
				if !b {
					debug.Printf("Line: %v skipped\n", childEnv.builtin.NR)
					continue
				}
			}
		case *ast.StartStopPattern:
			isMatch := func() (bool, error) {
				var b interface{}
				if !env.GetLoop() {
					//fmt.Printf("pattern=%#v\n", pattern)
					//fmt.Printf("Start=%#v\n", pattern.Start)
					b, err = evalExpr(pattern.Start, childEnv)
				} else {
					b, err = evalExpr(pattern.Stop, childEnv)
				}
				if err != nil {
					return false, err
				}
				switch reflect.ValueOf(b).Kind() {
				case reflect.Bool:
					return reflect.ValueOf(b).Interface().(bool), nil
				case reflect.Int, reflect.Int64, reflect.Float64:
					return b != 0, nil
				default:
					return false, nil

				}
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
			} else if !env.GetLoop() {
				debug.Printf("Line: %v skipped\n", childEnv.builtin.NR)
				continue
			}
		}
		result, err = runStmts(rule.Action, childEnv)
		if err != nil {
			return toInt(result), err
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
