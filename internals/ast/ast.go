package ast

type Expr interface {
	expr()
}

type Stmt interface {
	stmt()
}