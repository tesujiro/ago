package ast

import "reflect"

// Expr provides all of interfaces for expression.
type Expr interface{}

// IdentExpr provide identity expression.
type IdentExpr struct {
	Literal string
}

// FieldExpr provide field expression. ex: $1, $NF
type FieldExpr struct {
	Expr Expr
}

// NumExpr provide number expression.
type NumExpr struct {
	Literal string
}

// StringExpr provide string expression.
type StringExpr struct {
	Literal string
}

// ConstExpr provide constant expression.
type ConstExpr struct {
	Literal string
}

// ParentExpr provide parenthesis expression.
type ParentExpr struct {
	SubExpr Expr
}

// UnaryExpr provide unary operator expression or minus expression.
type UnaryExpr struct {
	Operator string
	Expr     Expr
}

// BinOpExpr provide binary operator expression.
type BinOpExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

// TriOpExpr provide if/then/else expression.
type TriOpExpr struct {
	Cond Expr
	Then Expr
	Else Expr
}

// AssExpr provide assignment expression.
type AssExpr struct {
	Left  []Expr
	Right []Expr
}

// CompExpr provide composite expression. ex: --1, ++1, 1--, 1++
type CompExpr struct {
	Left     Expr
	Operator string
	Right    Expr
	After    bool
}

// FuncExpr provide function definition expression.
type FuncExpr struct {
	Name  string
	Args  []string
	Stmts []Stmt
}

// CallExpr provide function call expression.
type CallExpr struct {
	Name     string
	Func     reflect.Value
	SubExprs []Expr
}

// AnonymousCallExpr provide anonymous function call expression.
type AnonymousCallExpr struct {
	Expr     Expr
	SubExprs []Expr
}

// ItemExpr provide expression to refer array item.
type ItemExpr struct {
	Expr  Expr
	Index []Expr
}

// RegExpr provide regular expression.
type RegExpr struct {
	Literal string
}

// MatchExpr provide match expression. ex: $1 ~ /AAA/
type MatchExpr struct {
	Expr    Expr
	RegExpr Expr
}

// ContainKeyExpr provide contain key expression
type ContainKeyExpr struct {
	KeyExpr Expr
	MapID   string
}

// GetlineExpr provide 'getline' expression
type GetlineExpr struct {
	Command Expr
	Var     Expr
	Redir   Expr
}
