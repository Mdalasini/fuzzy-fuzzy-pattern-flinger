package parser

import (
	"github.com/mdalasini/regex-engine/internals/ast"
	"github.com/mdalasini/regex-engine/internals/lexer"
)

type bindingPower int

const (
	defaultBp bindingPower = iota
	relational
)

type nudHandler func(p *parser) ast.Expr
type ledHandler func(p *parser, left ast.Expr, bp bindingPower) ast.Expr

type nudLookup  map[lexer.TokenKind]nudHandler
type ledLookup  map[lexer.TokenKind]ledHandler
type bpLookup   map[lexer.TokenKind]bindingPower

var bpLu = bpLookup{}
var nudLu = nudLookup{}
var ledLu = ledLookup{}

func led(kind lexer.TokenKind, bp bindingPower, ledFn ledHandler) {
	bpLu[kind] = bp
	ledLu[kind] = ledFn
}

func nud(kind lexer.TokenKind, nudFn nudHandler) {
	nudLu[kind] = nudFn
}

func createTokenLookups() {
	bpLu[lexer.LITERAL] = defaultBp // * Literals have the lowest binding power
	
	// * relational
	led(lexer.RANGE_OPENING, relational, parseRangeQuantifierExpr)

	// * Literals
	nud(lexer.LITERAL, parsePrimaryExpr)
}