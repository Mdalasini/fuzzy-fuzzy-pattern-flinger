package parser

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/mdalasini/regex-engine/internals/ast"
)

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	tokenKind := p.at().Kind
	nudFn, exists := nudLu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD HANDLER EXPECTED FOR THE TOKEN %c\n", p.at().Lit))
	}

	left := nudFn(p)
	// * While we have a led and the current BP < BP of current token
	for bpLu[p.at().Kind] > bp {
		tokenKind = p.at().Kind
		ledFn, exists := ledLu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED HANDLER EXPECTED FOR THE TOKEN %c\n", p.at().Lit))
		}

		left = ledFn(p, left, bpLu[p.at().Kind])
	}

	return left
}

func parsePrimaryExpr(p *parser) ast.Expr {
	return ast.LiteralExpr{Value: p.advance().Lit}
}

func parseRangeQuantifierExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	if isValidRangeQuantifier(p) {
		r := ast.RangeQuantifier{Element: left}
		// * Get lowerbound
		var buf bytes.Buffer
		p.advance() // * skip '{'
		currentToken := p.advance() 
		for {
			if !isDigit(currentToken.Lit) {break}
			buf.WriteRune(currentToken.Lit)
			currentToken = p.advance()
		}
		lb, _ := bufferToInt(buf)
		r.LowerBound = lb
		buf.Reset()
		if currentToken.Lit == ',' {
			r.Atleast = true
			currentToken = p.advance()
			if isDigit(currentToken.Lit) {
				for {
					if !isDigit(currentToken.Lit) {break}
					buf.WriteRune(currentToken.Lit)
					currentToken = p.advance()
				}
				ub, _ := bufferToInt(buf)
				r.UpperBound = ub
				r.Atleast = false
			}
		}
		return r
	}
	return ast.LiteralExpr{Value: p.advance().Lit}
}



// * HELPER FUNCTIONS
func isDigit(lit rune) bool {
	return (lit >= '0' && lit <= '9')
}

func bufferToInt(buf bytes.Buffer) (int, error) {
	str := buf.String()
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// * isValidRangeQuantifier checks if '{' does in fact mean the start of a range quantifier
func isValidRangeQuantifier(p *parser) bool {
	pos := p.pos
	// * checking for lowerbound
	pos++
	for {
		if pos < len(p.tokens) && isDigit(p.tokens[pos].Lit) {
			pos++
		} else {
			break
		}
	}

	// * checking for at least
	if pos < len(p.tokens) && p.tokens[pos].Lit == ',' {
		pos ++ 
		// * checking for upperbound
		if isDigit(p.tokens[pos].Lit) {
			for {
				if pos < len(p.tokens) && isDigit(p.tokens[pos].Lit) {
					pos++
				} else {
					break
				}
			}
		}
	}
	return pos < len(p.tokens) && p.tokens[pos].Lit == '}'
}