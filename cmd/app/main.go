package main

import (
	"fmt"

	"github.com/mdalasini/regex-engine/internals/lexer"
	"github.com/mdalasini/regex-engine/internals/matcher"
	"github.com/mdalasini/regex-engine/internals/parser"
)

func main() {
	pattern := "ab{2,4}"
	compile(pattern, "abbbb")
	compile(pattern, "abcdabbbb")
}

func compile(pattern, input string) {
	tokens := lexer.Tokenise(pattern)
	ast := parser.Parse(tokens)
	states := getStates(len(ast.Body)+1)
	engine :=matcher.NewEngineNFA()
	engine.DeclareStates(states...)
	engine.SetInitialState(states[0])
	engine.SetEndingStates([]string{states[len(states)-1]})
	for i, expr := range ast.Body {
		engine.AddTransition(states[i], states[i+1], matcher.CreateMatcher(expr))
	}
	fmt.Println(engine.Compute(input))
}

func getStates(count int) []string {
    states := make([]string, count)
    for i := 0; i < count; i++ {
        states[i] = fmt.Sprintf("q%d", i)
    }
    return states
}