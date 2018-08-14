package ast

type Rule struct {
	Pattern Pattern
	Action  []Stmt
}

type Pattern interface{}

type BeginPattern struct {
}

type EndPattern struct {
}

type ExprPattern struct {
	Expr Expr
}

type FuncPattern struct {
	Name string
	Args []string
}
