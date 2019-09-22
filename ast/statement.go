package ast

type Stmt interface{}

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

type ExitStmt struct {
	Expr Expr
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

type NextStmt struct {
}

type MapLoopStmt struct {
	KeyID string
	MapID string
	Stmts []Stmt
}

type PrintStmt struct {
	Exprs []Expr
}
