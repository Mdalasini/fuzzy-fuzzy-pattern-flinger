package matcher

import (
	"github.com/mdalasini/regex-engine/internals/ast"
)

type matcher interface {
	matches(scanner *scanner) bool
	isEpsilon() bool
}

type EpsilonMatcher struct{}

func (e EpsilonMatcher) matches(scanner *scanner) bool {
	return true
}

func (e EpsilonMatcher) isEpsilon() bool {
	return true
}

type LiteralMatcher struct {
	Expr ast.LiteralExpr
}

func (l LiteralMatcher) matches(scanner *scanner) bool {
	bytes, err := scanner.r.Peek(1)
	if err != nil {
		return false // * means no bytes left
	}
	if l.Expr.Value == rune(bytes[0]) {
		scanner.r.ReadRune()
		return true
	} else {
		return false
	}
}

func (l LiteralMatcher) isEpsilon() bool {
	return false
}

type RangeMatcher struct{
	Expr ast.RangeQuantifier
}

func (r RangeMatcher) matches(scanner *scanner) bool {
	minLen := r.Expr.LowerBound
	for i := 0; i < r.Expr.LowerBound; i++ {
		if !CreateMatcher(r.Expr.Element).matches(scanner) {
			return false
		}
	}
	scanner.r.Discard(minLen)

	if r.Expr.UpperBound > r.Expr.LowerBound {
		for i := r.Expr.LowerBound; i < r.Expr.UpperBound; i++ {
			CreateMatcher(r.Expr.Element).matches(scanner)
		}
		return true
	}

	if r.Expr.Atleast {
		for{
			if !CreateMatcher(r.Expr.Element).matches(scanner) {
				break
			}
		}		 
	}
	return true
}

func (r RangeMatcher) isEpsilon() bool {
	return false
}

// * HELPER FUNCTIONS
func CreateMatcher(expr ast.Expr) matcher {
	switch e := expr.(type) {
	case ast.LiteralExpr:
		return LiteralMatcher{Expr: e}
	case ast.RangeQuantifier:
		return RangeMatcher{Expr: e}
	default:
		return nil
	}
}