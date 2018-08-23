package ast

type Stmt interface{}

type AssStmt struct {
	Left  []Expr
	Right []Expr
}

type ExprStmt struct {
	Expr Expr
}

type DelStmt struct {
	Expr Expr
}

type IfStmt struct {
	If     Expr
	Then   []Stmt
	Else   []Stmt
	ElseIf []Stmt
}

type ReturnStmt struct {
	Exprs []Expr
}

type LoopStmt struct {
	Expr  Expr
	Stmts []Stmt
}

type CForLoopStmt struct {
	Stmt1 Stmt
	Expr2 Expr
	Expr3 Expr
	Stmts []Stmt
}

type DoLoopStmt struct {
	Expr  Expr
	Stmts []Stmt
}

type BreakStmt struct {
}

type ContinueStmt struct {
}

type HashLoopStmt struct {
	Key   string
	Hash  string
	Stmts []Stmt
}

type PrintStmt struct {
	Exprs []Expr
}
