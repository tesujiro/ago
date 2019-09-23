package ast

// Stmt provides all of interfaces for statement.
type Stmt interface{}

// ExprStmt provide expression statement.
type ExprStmt struct {
	Expr Expr
}

// DelStmt provides "delete" statement.
type DelStmt struct {
	Expr Expr
}

// IfStmt provides "if/else/else if" statement.
type IfStmt struct {
	If     Expr
	Then   []Stmt
	Else   []Stmt
	ElseIf []Stmt
}

// ReturnStmt provides "return" statement.
type ReturnStmt struct {
	Exprs []Expr
}

// ExitStmt provides "exit" statement.
type ExitStmt struct {
	Expr Expr
}

// LoopStmt provides loop statement. ex: for{}, while{},
type LoopStmt struct {
	Expr  Expr
	Stmts []Stmt
}

// CForLoopStmt provides C-style "for (;;)" statement.
type CForLoopStmt struct {
	Stmt1 Stmt
	Expr2 Expr
	Expr3 Expr
	Stmts []Stmt
}

// DoLoopStmt provides "do {}" statement.
type DoLoopStmt struct {
	Expr  Expr
	Stmts []Stmt
}

// BreakStmt provides "break" statement.
type BreakStmt struct {
}

// ContinueStmt provides "continue" statement.
type ContinueStmt struct {
}

// NextStmt provides "next" statement.
type NextStmt struct {
}

// MapLoopStmt provides "for ( id in MAP ) {}" statement.
type MapLoopStmt struct {
	KeyID string
	MapID string
	Stmts []Stmt
}

// PrintStmt provides "print" statement.
type PrintStmt struct {
	Exprs []Expr
}
