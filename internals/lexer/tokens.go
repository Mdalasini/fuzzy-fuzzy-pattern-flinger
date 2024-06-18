package lexer

type Token struct {
	Kind TokenKind
	Lit  rune
}

type TokenKind int

const (
	ILLEGAL TokenKind = iota
	EOF

	// * Escape sequences
	ESCAPE // * '\' makes special characters get treated as theit literal form
	BACKSLASH // * '\\'
	NEWLINE // * '\n'
	TAB // * '\t'
	CARRIAGE_RETURN // * '\r'
	// ! ESCAPE works differently in side the bracket list

	// * OR operator
	OR 	  // * '|'

	// * Occurence indicators (or Repetition operators)
	ONE_OR_MORE // * '+'
	ZERO_OR_MORE // * '*'
	OPTIONAL // * '?'

	// * Meta-charcters
	DOR // * '.'
	ANY_ONE_DIGIT // * '\d' any digit chatcter (digit 0-9)
	ANY_NONE_DIGIT // * '\D' any non-digit character
	ANY_ONE_WORD  // * '\w' any word character (letters, digits, and underscores)
	ANY_NONE_WORD // * '\W' any non-word character
	ANY_ONE_SPACE // * '\s' any whitespace character (spaces, tabs, newlines, etc.)
	ANY_NONE_SPACE // * '\S' any non-whitespace character

	// * Position Anchors
	START_OF_LINE // * '^'
	END_OF_LINE // * '$'
	WORD_BOUNDARY // * '\b' where a word character is not followed or preceded by another word character
	NON_WORD_BOUNDARY // * '\B'
	WORD_START // * '\<' next character is the beginning of a word
	WORD_END // * '\>' next character is the end of a word

	// * Grouping literals
	RANGE_OPENING // * '{'
	CAPTURING_OPENING // * '('
	MATCHING_OPENING // * '['

	LITERAL
)

var eof = rune(0)

func contains(s []rune, e rune) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func isLiteral(ch rune) bool {
	literalSymbols := []rune{'_', ')', '}', ']', ',', '-',}
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || contains(literalSymbols, ch)
}