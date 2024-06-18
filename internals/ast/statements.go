package ast

type PatternStmt struct {
	Body []Expr
}

func (p PatternStmt) stmt() {}
