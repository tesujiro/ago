// Package ast implements abstract syntax tree for ago.
package ast

// Rule provide an awk rule.Each rule specifies one pattern to search for, and one action to perform when that pattern is found.
type Rule struct {
	Pattern Pattern
	Action  []Stmt
}

// Pattern is an awk pattern. ex: $1~"XYZ", $NR>10, BEGIN, END, func Print(x).
type Pattern interface{}

// BeginPattern is an awk BEGIN pattern.
type BeginPattern struct {
}

// EndPattern is an awk END pattern.
type EndPattern struct {
}

// ExprPattern is an awk expression pattern. ex: $1~"XYZ", $NR>10.
type ExprPattern struct {
	Expr Expr
}

// FuncPattern is an awk function definition pattern. ex: func Print(x).
type FuncPattern struct {
	Name string
	Args []string
}

// StartStopPattern is an awk start stop line pattern. ex: "/^BEGIN$/,/^END$/".
type StartStopPattern struct {
	Start Expr
	Stop  Expr
}
