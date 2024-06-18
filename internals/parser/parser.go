package parser

import (
	"github.com/mdalasini/regex-engine/internals/ast"
	"github.com/mdalasini/regex-engine/internals/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func newParser(tokens []lexer.Token) *parser {
	return &parser{tokens: tokens}
}

func Parse(tokens []lexer.Token) ast.PatternStmt{
	createTokenLookups()
	pattern := ast.PatternStmt{Body: make([]ast.Expr, 0)}
	p := newParser(tokens)

	for p.hasTokens() {
		pattern.Body = append(pattern.Body, parseExpr(p, defaultBp))
	}

	return pattern
}

// * HELPER FUNCTIONS

// * at gets the current token
func (p *parser) at() lexer.Token {
	return p.tokens[p.pos]
}

// * advance moves parser position by one token
// * Returns the previous token
func (p *parser) advance() lexer.Token {
	currentToken := p.at()
	p.pos++
	return currentToken
}

// * hasTokens checks if the parser has any tokens left or is at the end of the file
func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.at().Kind != lexer.EOF
}
