package ast

import "fmt"

// * -------------------
// * LITERAL EXPRESSION
// * -------------------

type LiteralExpr struct {
	Value rune
}

func (l LiteralExpr) expr() {}

func (l LiteralExpr) String() string {
	return fmt.Sprintf("Literal(%c)", l.Value)
}

type SymbolExpr struct {
	Value rune
}

func (s SymbolExpr) expr() {}

func (s SymbolExpr) String() string {
	return fmt.Sprintf("Symbol(%c)", s.Value)
}

// * -------------------
// * COMPLEX EXPRESSION
// * -------------------

type RangeQuantifier struct {
	LowerBound int
	Atleast    bool
	UpperBound int
	Element    Expr
}

func (r RangeQuantifier) expr() {}

func (r RangeQuantifier) String() string {
	return fmt.Sprintf("RangeQunatifier(%v, %d, %d)", r.Element, r.LowerBound, r.UpperBound)
}