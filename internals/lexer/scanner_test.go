package lexer

import "testing"

// func TestTokenise(t *testing.T) {
// 	result := Tokenise(`ab\`)
// 	expected := []Token{{LITERAL, 'a'}, {LITERAL, 'b'}, {ESCAPE, '\\'}, {EOF, eof}}
// 	if !tokenSliceAreSame(result, expected) {
// 		t.Errorf("Results incorrect")
// 	}
// }

func TestTokenise(t *testing.T) {
	testCases := []struct {
		desc	string
		input	string
		want    []Token
	}{
		{
			desc: "testing literals",
			input: "ab12",
			want: []Token{{LITERAL, 'a'}, {LITERAL, 'b'}, {LITERAL, '1'}, {LITERAL, '2'}, {EOF, eof}},
		},
		{
			desc: "testing escape sequences",
			input: "\\\n\t",
			want: []Token{{NEWLINE, '\n'}, {TAB, '\t'}, {EOF, eof}},
		},
		{
			desc: "testing escape sequences with backticks",
			input: `\n\t`,
			want: []Token{{NEWLINE, '\n'}, {TAB, '\t'}, {EOF, eof}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ans := Tokenise(tC.input)
			if !tokenSliceAreSame(ans, tC.want) {
				t.Errorf("Tokenization incorrect")
			}
		})
	}
}

// * HELPER FUNCTION
func tokenSliceAreSame(toks1, toks2 []Token) bool {
	if len(toks1) != len(toks2) {
		return false
	}
	for i, tok := range toks1 {
		if tok != toks2[i] {
			return false
		}
	}
	return true
}

