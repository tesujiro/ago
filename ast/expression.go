package ast

import "reflect"

type Expr interface{}

type IdentExpr struct {
	Literal string
}

type FieldExpr struct {
	Expr Expr
}

type NumExpr struct {
	Literal string
}

type StringExpr struct {
	Literal string
}

type ConstExpr struct {
	Literal string
}

type ParentExpr struct {
	SubExpr Expr
}

type UnaryExpr struct {
	Operator string
	Expr     Expr
}

type BinOpExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

type TriOpExpr struct {
	Cond Expr
	Then Expr
	Else Expr
}

type AssExpr struct {
	Left  []Expr
	Right []Expr
}

type CompExpr struct {
	Left     Expr
	Operator string
	Right    Expr
	After    bool
}

type FuncExpr struct {
	Name  string
	Args  []string
	Stmts []Stmt
}

type CallExpr struct {
	Name     string
	Func     reflect.Value
	SubExprs []Expr
}

type AnonymousCallExpr struct {
	Expr     Expr
	SubExprs []Expr
}

type ItemExpr struct {
	Expr  Expr
	Index []Expr
}

type RegExpr struct {
	Literal string
}

type MatchExpr struct {
	Expr    Expr
	RegExpr Expr
}

type ContainKeyExpr struct {
	KeyExpr Expr
	MapId   string
}

type GetlineExpr struct {
	Command Expr
	Var     Expr
	Redir   Expr
}
