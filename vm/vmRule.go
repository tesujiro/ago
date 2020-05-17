package vm

import (
	"fmt"
	"io"
	"reflect"

	"github.com/tesujiro/ago/ast"
	"github.com/tesujiro/ago/debug"
)

// Run executes rules in the specified environment.
func Run(rules []ast.Rule, env *Env) (interface{}, error) {
	var result interface{}
	var err error

	funcRules, beginRules, mainRules, endRules := separateRules(rules)

	// FUNC DEFINITION
	if len(funcRules) > 0 {
		result, err = runFuncRules(funcRules, env)
		if err != nil {
			return result, err
		}
	}

	// BEGIN
	result, err = runBeginRules(beginRules, env)
	debug.Printf("%#v\n", result)
	if err == ErrExit {
		v, ok := result.(int)
		if ok {
			return v, nil
		}
		return v, err
	}
	if err != nil {
		return result, err
	}

	if len(mainRules) == 0 && len(endRules) == 0 {
		return result, nil
	}

	// reset variable
	env.SetNF()

	// MAIN
	for {
		//fmt.Println("MAINLOOP")
		fileLine, err := env.GetLine()
		if err == io.EOF {
			break
		} else if err != nil {
			//fmt.Printf("error:%v\n", err)
			return nil, err
		}
		env.SetFieldFromLine(fileLine)
		if len(mainRules) > 0 {
			result, err := runMainRules(mainRules, env)
			if err == ErrNext {
				continue
			}
			if err != nil {
				return result, err
			}
		}
	}

	// END
	result, err = runEndRules(endRules, env)
	return result, err
}

// separateRules classifies rules to func, begin, main and end rules.
func separateRules(rules []ast.Rule) (Func, Begin, Main, End []ast.Rule) {
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

// runFuncRules executes func rules with a specified env.
func runFuncRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println("FUNC")

		funcExpr := &ast.FuncExpr{Name: rule.Pattern.(*ast.FuncPattern).Name, Args: rule.Pattern.(*ast.FuncPattern).Args, Stmts: rule.Action}
		result, err = evalExpr(funcExpr, env)
		if err != nil {
			return env.toInt(result), err
		}
	}
	return env.toInt(result), nil
}

// runBeginRules executes begin rules with a specified env.
func runBeginRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println("BEGIN")
		childEnv := env.NewEnv()
		result, err = runStmts(rule.Action, childEnv)
		if err != nil {
			return env.toInt(result), err
		}
	}
	return env.toInt(result), err
}

// runMainRules executes main rules with a specified env.
func runMainRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
	for _, rule := range rules {
		debug.Println(env.builtin.NR, ":MAIN")
		childEnv := env.NewEnv()
		switch pattern := rule.Pattern.(type) {
		case *ast.ExprPattern:
			expr := pattern.Expr
			if expr != nil {
				result, err := evalExpr(expr, childEnv)
				if err != nil {
					return env.toInt(result), err
				}
				b, err := env.strictToBool(result)
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
				case reflect.String:
					return true, nil
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
			return env.toInt(result), err
		}
	}
	return env.toInt(result), err
}

// runEndRules executes end rules with a specified env.
func runEndRules(rules []ast.Rule, env *Env) (result interface{}, err error) {
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
