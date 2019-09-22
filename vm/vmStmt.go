package vm

import (
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/tesujiro/ago/ast"
)

var (
	ErrBreak    = errors.New("unexpected break")
	ErrContinue = errors.New("unexpected continue")
	ErrNext     = errors.New("unexpected next")
	ErrReturn   = errors.New("unexpected return")
	ErrExit     = errors.New("unexpected exit")
)

func runStmts(stmts []ast.Stmt, env *Env) (interface{}, error) {

	if result, err := run(stmts, env); err == ErrReturn {
		return result, nil
	} else {
		return result, err
	}
}

var afterStmts []ast.Stmt

func run(stmts []ast.Stmt, env *Env) (interface{}, error) {
	//fmt.Println("run -> env.Dump()")
	//env.Dump()
	var result interface{}
	var err error
	for _, stmt := range stmts {
		switch stmt.(type) {
		case *ast.BreakStmt:
			return nil, ErrBreak
		case *ast.ContinueStmt:
			return nil, ErrContinue
		case *ast.NextStmt:
			return result, ErrNext
		case *ast.ReturnStmt:
			//fmt.Println("Return1")
			result, err = runSingleStmt(stmt, env)
			if err != nil && err != ErrReturn {
				//fmt.Printf("Return1 error !! err:%v\n", err)
				return nil, err
			}
			return result, ErrReturn
		case *ast.ExitStmt:
			result, err = runSingleStmt(stmt, env)
			//fmt.Printf("vmStmt ExitStmt result:%#v err:%v\n", result, err)
			if err != nil && err != ErrExit {
				return nil, err
			}
			return result, ErrExit
		default:
			result, err = runSingleStmt(stmt, env)
			//fmt.Printf("vmStmt default result:%#v err:%v\n", result, err)
			switch err {
			case nil, ErrReturn:
			case ErrExit:
				return result, err
			default:
				return nil, err
			}
		}
		for {
			if len(afterStmts) == 0 {
				break
			}
			stmt := afterStmts[0]
			afterStmts = afterStmts[1:]
			result, err = run([]ast.Stmt{stmt}, env)
			if err != nil {
				return result, err
			}
			//fmt.Println("afterStmt result:", result)
		}
	}
	return result, err
}

func runSingleStmt(stmt ast.Stmt, env *Env) (interface{}, error) {
	switch stmt := stmt.(type) {
	/*
		case *ast.AssStmt:
			left, right := stmt.Left, stmt.Right

			// evaluate right expressions
			right_values := make([]interface{}, len(right))
			var err error
			for i, expr := range right {
				right_values[i], err = evalExpr(expr, env)
				if err != nil {
					return nil, err
				}
			}

			// evaluate assExpr
			switch {
			case len(left) == 1 && len(right) == 1:
				return evalAssExpr(left[0], right_values[0], env)
			case len(left) > 1 && len(right) == 1:
				val := right_values[0]
				if reflect.ValueOf(val).Kind() == reflect.Interface {
					val = reflect.ValueOf(val).Elem().Interface()
				}
				if reflect.ValueOf(val).Kind() != reflect.Slice {
					return nil, errors.New("single value assign to multi values")
				} else {
					elements := reflect.ValueOf(val)
					right_values = make([]interface{}, elements.Len())
					for i := 0; i < elements.Len(); i++ {
						right_values[i] = elements.Index(i).Interface()
					}
				}
				fallthrough
			default:
				for i, expr := range left {
					if i >= len(right_values) {
						return right_values[len(right_values)-1], nil
					}
					if _, err := evalAssExpr(expr, right_values[i], env); err != nil {
						return nil, err
					}
				}
				return right_values[len(left)-1], nil
			}
	*/
	case *ast.ExprStmt:
		return evalExpr(stmt.Expr, env)
	case *ast.DelStmt:
		expr := stmt.Expr
		//fmt.Println("TypeOf(Expr):", reflect.TypeOf(expr))
		var id string
		var index string
		switch expr := expr.(type) {
		case *ast.IdentExpr:
			id = expr.Literal
		case *ast.ItemExpr:
			ie, ok := expr.Expr.(*ast.IdentExpr)
			if !ok {
				return nil, errors.New("non variable does not support delete operation")
			}
			id = ie.Literal
			var err error
			index, err = getHashIndex(env, expr.Index)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("type %s does not support delete operation", reflect.TypeOf(expr))
		}
		val, err := env.Get(id)
		if err == ErrUnknownSymbol {
			// Set Default Map to env
			val, err = env.DefineDefaultMapValue(id, index)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		if reflect.ValueOf(val).Kind() == reflect.Map {
			m := val.(map[interface{}]interface{})
			if index == "" {
				// Delete All Map Elements
				_, err := env.DefineDefaultMapValue(id, index)
				if err != nil {
					return nil, err
				}
			} else {
				// Delete Map Element
				delete(m, index)
			}
		} else {
			// Error
			return nil, fmt.Errorf("type %s does not support delete operation", reflect.ValueOf(val).Kind())
		}
		return nil, nil
	case *ast.PrintStmt:
		for i, expr := range stmt.Exprs {
			result, err := evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
			if 0 < i && i < len(stmt.Exprs) {
				fmt.Printf("%v", env.builtin.OFS)
			}
			//fmt.Printf("%v", result)
			switch reflect.ValueOf(result).Kind() {
			case reflect.Float64:
				fmt.Printf("%.6g", result)
			case reflect.Int, reflect.Int64, reflect.Bool, reflect.String:
				fmt.Printf("%v", result)
			case reflect.Slice:
				len := reflect.ValueOf(result).Len()
				for i := 0; i < len; i++ {
					if i > 0 {
						fmt.Printf("%v", env.builtin.OFS)
					}
					fmt.Printf("%v", reflect.ValueOf(result).Index(i))
				}
			case reflect.Invalid:
				fmt.Printf("")
			default:
				return nil, fmt.Errorf("type %s does not support print operation", reflect.ValueOf(result).Kind().String())
			}
		}
		fmt.Printf("%v", env.builtin.ORS)
	case *ast.IfStmt:
		child := env.NewEnv()
		result, err := evalExpr(stmt.If, child)
		if err != nil {
			return nil, err
		}
		b, err := strictToBool(result)
		if err != nil {
			return nil, fmt.Errorf("convert if condition:%v", err)
		}
		if b {
			//fmt.Println("If then -> env.Dump()")
			//child.Dump()
			result, err = run(stmt.Then, child)
			if err != nil {
				return result, err
			}
			return result, nil
		}
		for _, stmt := range stmt.ElseIf {
			result, err := evalExpr(stmt.(*ast.IfStmt).If, child)
			if err != nil {
				return nil, err
			}
			b, err := strictToBool(result)
			if err != nil {
				return nil, fmt.Errorf("convert else if condition:%v", err)
			}
			if b {
				result, err = run(stmt.(*ast.IfStmt).Then, child)
				if err != nil {
					return result, err
				}
				return result, nil
			}
		}

		if len(stmt.Else) > 0 {
			result, err = run(stmt.Else, child)
			if err != nil {
				return result, err
			}
		}
		return result, nil
	case *ast.ReturnStmt:
		//fmt.Println("Return2")
		length := len(stmt.Exprs)

		resultExpr := make([]interface{}, length)
		var err error
		for i, expr := range stmt.Exprs {
			resultExpr[i], err = evalExpr(expr, env)
			if err != nil {
				//fmt.Printf("Return2 error!  err;%v\n", err)
				return nil, err
			}
		}

		switch length {
		case 0:
			return nil, nil
		case 1:
			return resultExpr[0], nil
		default:
			return resultExpr, nil
		}
	case *ast.ExitStmt:
		result, err := evalExpr(stmt.Expr, env)
		if err != nil {
			return nil, err
		}
		return result, nil
	case *ast.LoopStmt:
		newEnv := env.NewEnv()
		for {
			exp := stmt.Expr
			if exp != nil {
				result, err := evalExpr(exp, newEnv)
				if err != nil {
					return nil, err
				}
				b, err := strictToBool(result)
				if err != nil {
					return nil, fmt.Errorf("convert while condition:%v", err)
				}
				if !b {
					break
				}
			}

			ret, err := run(stmt.Stmts, newEnv)
			if err == ErrReturn {
				return ret, nil
			}
			if err == ErrBreak {
				break
			}
			if err == ErrContinue {
				continue
			}
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case *ast.CForLoopStmt:
		stmt1 := stmt.Stmt1
		expr2 := stmt.Expr2
		expr3 := stmt.Expr3
		stmts := stmt.Stmts
		newEnv := env.NewEnv()
		if stmt1 != nil {
			_, err := run([]ast.Stmt{stmt1}, newEnv)
			if err != nil {
				return nil, err
			}
		}
		for {
			if expr2 != nil {
				result, err := evalExpr(expr2, newEnv)
				if err != nil {
					return nil, err
				}
				b, err := strictToBool(result)
				if err != nil {
					return nil, fmt.Errorf("convert for loop condition:%v", err)
				}
				if !b {
					break
				}
			}
			ret, err := run(stmts, newEnv)
			if err == ErrReturn {
				return ret, nil
			}
			if err == ErrBreak {
				break
			}
			if err == ErrContinue {
				//continue
				err = nil
			}
			if err != nil {
				return nil, err
			}
			if expr3 != nil {
				_, err := evalExpr(expr3, newEnv)
				if err != nil {
					return nil, err
				}
			}
		}
		return nil, nil
	case *ast.DoLoopStmt:
		newEnv := env.NewEnv()
		for {
			ret, err := run(stmt.Stmts, newEnv)
			if err == ErrReturn {
				return ret, nil
			}
			if err == ErrBreak {
				break
			}
			if err == ErrContinue {
				continue
			}
			if err != nil {
				return nil, err
			}
			exp := stmt.Expr
			if exp != nil {
				result, err := evalExpr(exp, newEnv)
				if err != nil {
					return nil, err
				}
				b, err := strictToBool(result)
				if err != nil {
					return nil, fmt.Errorf("convert do loop condition:%v", err)
				}
				if !b {
					break
				}
			}
		}
		return nil, nil
	case *ast.MapLoopStmt:
		keyID := stmt.KeyID
		mapID := stmt.MapID
		stmts := stmt.Stmts
		v, err := env.Get(mapID)
		if err == ErrUnknownSymbol {
			val, err := env.DefineDefaultMap(mapID)
			if err != nil {
				return nil, err
			}
			v = val
		} else if err != nil {
			return nil, err
		}
		if reflect.TypeOf(v).Kind() != reflect.Map {
			return nil, fmt.Errorf("for key loop not in associated array,%s", reflect.TypeOf(v).Kind())
		}
		// sort hash keys
		m := v.(map[interface{}]interface{})
		indecies := make([]string, len(m))
		i := 0
		for k, _ := range m {
			indecies[i] = k.(string)
			i++
		}
		sort.Strings(indecies)

		newEnv := env.NewEnv()
		for _, index := range indecies {
			if err := newEnv.Set(keyID, index); err == ErrUnknownSymbol {
				if err := newEnv.Define(keyID, index); err != nil {
					return nil, err
				}
			} else if err != nil {
				return nil, err
			}
			ret, err := run(stmts, newEnv)

			if err == ErrReturn {
				return ret, nil
			}
			if err == ErrBreak {
				break
			}
			if err == ErrContinue {
				continue
			}
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}
	return nil, nil
}
