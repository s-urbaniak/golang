// This package represents
package lexer

import (
	"states"
)

type tokenType int

type token struct {
	typ tokenType
	val string
}

const (
	tokenError tokenType = iota
	tokenComment
	tokenTypeMatrix
	tokenTypeScalar
	tokenName
	tokenRows
	tokenCols
	tokenValue
)

var keywords = map[string]tokenType{
	"# name:":             tokenName,
	"# type: scalar":      tokenTypeScalar,
	"# type: matrix":      tokenTypeMatrix,
	"# type: bool matrix": tokenTypeMatrix,
	"# rows:":             tokenRows,
	"# columns:":          tokenCols,
	"# ":                  tokenComment,
}

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

type lexer struct {
	line  string
	state stateFn
}

func (lex *lexer) run() {
	for lex.state = stateStart; lex.state != nil; {
		lex.state = lex.state(lex)
	}
}
