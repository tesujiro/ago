package ast

type Rule struct {
	Pattern Pattern
	Action  []Stmt
}

type Pattern struct {

	//Expr Expr
}

//type Stmt interface{}
type Stmt struct {
	Message string
}
