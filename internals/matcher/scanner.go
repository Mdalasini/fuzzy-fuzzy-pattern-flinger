package matcher

import (
	"bufio"
	"strings"
)

type scanner struct {
	r *bufio.Reader
}

func newScanner(source string) *scanner {
	return &scanner{bufio.NewReader(strings.NewReader(source))}
}