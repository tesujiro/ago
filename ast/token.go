package ast

// Token is used in the lexer to split characters into a string called a token
type Token struct {
	Token   int
	Literal string
	PosImpl // PosImpl provides get/set the position function.
}

// Position provides interface to store code locations.
type Position struct {
	Line   int
	Column int
}

// Pos interface provides two functions to get/set the position for expression or statement.
type Pos interface {
	Position() Position
	SetPosition(Position)
}

// PosImpl provides commonly implementations for Pos.
type PosImpl struct {
	pos Position
}

// Position return the position of the expression or statement.
/*
func (x *PosImpl) Position() Position {
	return x.pos
}
*/

// SetPosition is a function to specify position of the expression or statement.
func (x *PosImpl) SetPosition(pos Position) {
	x.pos = pos
}
