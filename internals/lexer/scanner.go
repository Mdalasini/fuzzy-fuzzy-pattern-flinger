package lexer

import (
	"bufio"
	"io"
	"strings"
)

type scanner struct {
	r *bufio.Reader
}

func newScanner(r io.Reader) *scanner {
	return &scanner{bufio.NewReader(r)}
}

func (s *scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// * Scan returns the next token and literal value.
func (s *scanner) scan() (tok Token) {
	ch := s.read()

	if isLiteral(ch) {
		return Token{LITERAL, ch}
	} 
	// TODO: deal with escape

	switch ch {
	case eof:
		return Token{EOF, ch}
	case '?':
		return Token{OPTIONAL, ch}
	case '{':
		return Token{RANGE_OPENING, ch}
	case '(':
		return Token{CAPTURING_OPENING, ch}
	case '[':
		return Token{MATCHING_OPENING, ch}
	default:
		return Token{ILLEGAL, ch}
	}
}

func Tokenise(source string) []Token {
	scanner := newScanner(strings.NewReader(source))

	tokens := []Token{}
	for {
		if tok:= scanner.scan(); tok.Kind != EOF {
			tokens = append(tokens, tok)
		} else {
			tokens = append(tokens, tok)
			break
		}
	}
	return tokens
}