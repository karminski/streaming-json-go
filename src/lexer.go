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

var leftPairTokens = map[int]bool{
	TOKEN_LEFT_BRACKET: true,
	TOKEN_LEFT_BRACE:   true,
}

type Lexer struct {
	JSONContent      strings.Builder
	JSONSegment      string
	TokenStack       []int
	MirrorTokenStack []int
}

func NewLexer() *Lexer {
	return &Lexer{}
}

func (lexer *Lexer) getTopTokenOnStack() int {
	tokenStackLen := len(lexer.TokenStack)
	if tokenStackLen == 0 {
		return TOKEN_EOF
	}
	return lexer.TokenStack[tokenStackLen-1]
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

func (lexer *Lexer) popMirrorTokenStack() int {
	mirrorTokenStackLen := len(lexer.MirrorTokenStack)
	if mirrorTokenStackLen == 0 {
		return TOKEN_EOF
	}
	token := lexer.MirrorTokenStack[mirrorTokenStackLen-1]
	lexer.MirrorTokenStack = lexer.MirrorTokenStack[:mirrorTokenStackLen-1]
	return token
}

func (lexer *Lexer) pushTokenStack(token int) {
	lexer.TokenStack = append(lexer.TokenStack, token)
}

func (lexer *Lexer) pushMirrorTokenStack(token int) {
	lexer.MirrorTokenStack = append(lexer.MirrorTokenStack, token)
}

func (lexer *Lexer) isLeftPairToken(token int) bool {
	if token == TOKEN_QUOTE {
		return lexer.getTopTokenOnStack() == TOKEN_QUOTE
	} else {
		itIs, hit := leftPairTokens[token]
		return itIs && hit
	}
}

func (lexer *Lexer) skipJSONSegment(n int) {
	lexer.JSONSegment = lexer.JSONSegment[n:]
}

func (lexer *Lexer) matchToken() (int, byte) {
	// finish
	if len(lexer.JSONSegment) == 0 {
		return TOKEN_EOF, byte(0)
	}

	// check token
	switch lexer.JSONSegment[0] {
	case TOKEN_LEFT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACKET, lexer.JSONSegment[0]
	case TOKEN_RIGHT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACKET, lexer.JSONSegment[0]
	case TOKEN_LEFT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACE, lexer.JSONSegment[0]
	case TOKEN_RIGHT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACE, lexer.JSONSegment[0]
	case TOKEN_COLON_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COLON, lexer.JSONSegment[0]
	case TOKEN_DOT_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_DOT, lexer.JSONSegment[0]
	case TOKEN_COMMA_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COMMA, lexer.JSONSegment[0]
	case TOKEN_QUOTE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_QUOTE, lexer.JSONSegment[0]
	default:
		return TOKEN_OTHERS, lexer.JSONSegment[0]

	}
}

func (lexer *Lexer) AppendString(str string) {
	lexer.JSONSegment = str
	for {
		token, tokenSymbol := lexer.matchToken()
		switch token {
		case TOKEN_EOF:
			break
		case TOKEN_OTHERS:
			lexer.JSONContent.WriteByte(tokenSymbol)
		default:
			lexer.JSONContent.WriteByte(tokenSymbol)
			lexer.pushTokenStack(token)
		}

	}
}
