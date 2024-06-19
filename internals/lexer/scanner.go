package lexer

import (
	"bufio"
	"io"
	"log"
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
	case '\\':
		return Token{ESCAPE, ch}
	case '\n':
		return Token{NEWLINE, ch}
	case '\t':
		return Token{TAB, ch}
	case '\r':
		return Token{CARRIAGE_RETURN, ch}
	case '+':
		return Token{ONE_OR_MORE, ch}
	case '*':
		return Token{ZERO_OR_MORE, ch}
	case '.':
		return Token{DOT, ch}
	case '^':
		return Token{START_OF_LINE, ch}
	case '$':
		return Token{END_OF_LINE, ch}
	case '\b':
		return Token{WORD_BOUNDARY, ch}
	default:
		return Token{LITERAL, ch}
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
	return cleanEscapeSequence(tokens)
}

func cleanEscapeSequence(tokens []Token) []Token {
	var result []Token
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Kind == ESCAPE {
			if tokens[i+1].Kind == EOF {
				log.Fatal("error parsing regex: trailing backslash at end of expression")
			}
			sequence := tokenToEscapeSequence(tokens[i+1])
			result = append(result, sequence)
			i++
		} else {
			result = append(result, tokens[i])
		}		
	}
	return result
}

func tokenToEscapeSequence(tok Token) Token {
	switch tok.Lit{
	case '\\':
		return Token{LITERAL, tok.Lit}
	case 'n':
		return Token{NEWLINE, '\n'}
	case '\n':
		return Token{NEWLINE, '\n'}
	case 't':
		return Token{TAB, '\t'}
	case '\t':
		return Token{TAB, '\t'}
	case 'r':
		return Token{CARRIAGE_RETURN, '\r'}
	case '\r':
		return Token{CARRIAGE_RETURN, '\r'}
	case 'd':
		return Token{ANY_ONE_DIGIT, 0}
	case 'D':
		return Token{ANY_NONE_DIGIT, 0}
	case 'w':
		return Token{ANY_ONE_WORD, 0}
	case 'W':
		return Token{ANY_NONE_WORD, 0}
	case 's':
		return Token{ANY_ONE_SPACE, 0}
	case 'S':
		return Token{ANY_NONE_SPACE, 0}
	case 'b':
		return Token{WORD_BOUNDARY, 0}
	case 'B':
		return Token{NON_WORD_BOUNDARY, 0}
	case '<':
		return Token{WORD_START, 0}
	case '>':
		return Token{WORD_END, 0}
	default:
		if isMetaCharacter(tok.Lit) {
			return Token{LITERAL, tok.Lit}
		}
		log.Fatalf("error parsing regexp: invalid escape sequence: '%q'", tok.Lit)
		return Token{}
	}
}

func isMetaCharacter(lit rune) bool {
	metaCharacters := []rune{'.', '|', '+', '*', '?', '^', '$', '{', '(', '['}
	for _, v := range metaCharacters {
		if v == lit {
			return true
		}
	}
	return false
}

