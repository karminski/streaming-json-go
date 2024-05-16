package main

import (
	"strings"
)

// token const
const (
	TOKEN_EOF              = iota // end-of-file
	TOKEN_LEFT_BRACKET            // [
	TOKEN_RIGHT_BRACKET           // ]
	TOKEN_LEFT_BRACE              // {
	TOKEN_RIGHT_BRACE             // }
	TOKEN_COLON                   // :
	TOKEN_DOT                     // .
	TOKEN_COMMA                   // ,
	TOKEN_QUOTE                   // "
	TOKEN_ESCAPE_CHARACTER        // \
	TOKEN_OTHERS                  // anything else in json
)

// token symbol const
const (
	TOKEN_LEFT_BRACKET_SYMBOL     = '['
	TOKEN_RIGHT_BRACKET_SYMBOL    = ']'
	TOKEN_LEFT_BRACE_SYMBOL       = '{'
	TOKEN_RIGHT_BRACE_SYMBOL      = '}'
	TOKEN_COLON_SYMBOL            = ':'
	TOKEN_DOT_SYMBOL              = '.'
	TOKEN_COMMA_SYMBOL            = ','
	TOKEN_QUOTE_SYMBOL            = '"'
	TOKEN_ESCAPE_CHARACTER_SYMBOL = '\\'
)

var tokenNameMap = map[int]string{
	TOKEN_EOF:              "EOF",
	TOKEN_LEFT_BRACKET:     "[",
	TOKEN_RIGHT_BRACKET:    "]",
	TOKEN_LEFT_BRACE:       "{",
	TOKEN_RIGHT_BRACE:      "}",
	TOKEN_COLON:            ":",
	TOKEN_DOT:              ".",
	TOKEN_COMMA:            ",",
	TOKEN_QUOTE:            "\"",
	TOKEN_ESCAPE_CHARACTER: "\\",
}

type Lexer struct {
	JSONContent strings.Builder
	JSONSegment string
	TokenStack  []int
}

func NewLexer() *Lexer {
	return &Lexer{}
}

func (lexer *Lexer) popTokenStack() int {
	tokenStackLen := len(lexer.TokenStack)
	if tokenStackLen == 0 {
		return TOKEN_EOF
	}
	token := lexer.TokenStack[tokenStackLen-1]
	lexer.TokenStack = lexer.TokenStack[:tokenStackLen-1]
	return token
}

func (lexer *Lexer) pushTokenStack(token int) {
	lexer.TokenStack = append(lexer.TokenStack, token)
}

func (lexer *Lexer) skipJSONSegment(n int) {
	lexer.JSONSegment = lexer.JSONSegment[n:]
}

func (lexer *Lexer) matchToken() int {
	// finish
	if len(lexer.JSONSegment) == 0 {
		return TOKEN_EOF
	}

	// check token
	switch lexer.JSONSegment[0] {
	case TOKEN_LEFT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACKET
	case TOKEN_RIGHT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACKET
	case TOKEN_LEFT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACE
	case TOKEN_RIGHT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACE
	case TOKEN_COLON_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COLON
	case TOKEN_DOT_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_DOT
	case TOKEN_COMMA_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COMMA
	case TOKEN_QUOTE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_QUOTE
	}

	// other tokens that we does not care about
	return TOKEN_OTHERS
}

func (lexer *Lexer) AppendString(str string) {

}
