package main

import (
	"fmt"
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

var rightPairTokens = map[int]bool{
	TOKEN_RIGHT_BRACKET: true,
	TOKEN_RIGHT_BRACE:   true,
}

var mirrorTokenMap = map[int]int{
	TOKEN_LEFT_BRACKET: TOKEN_RIGHT_BRACKET,
	TOKEN_LEFT_BRACE:   TOKEN_RIGHT_BRACE,
	TOKEN_QUOTE:        TOKEN_QUOTE,
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

func (lexer *Lexer) getTopTokenOnMirrorStack() int {
	mirrotTokenStackLen := len(lexer.MirrorTokenStack)
	if mirrotTokenStackLen == 0 {
		return TOKEN_EOF
	}
	return lexer.TokenStack[mirrotTokenStackLen-1]
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

func (lexer *Lexer) dumpMirrorTokenStackToString() string {
	var stackInString strings.Builder
	for _, token := range lexer.MirrorTokenStack {
		stackInString.WriteString(tokenNameMap[token])
	}
	return stackInString.String()
}

func (lexer *Lexer) isLeftPairToken(token int) bool {
	if token == TOKEN_QUOTE {
		return lexer.getTopTokenOnMirrorStack() != TOKEN_QUOTE
	} else {
		itIs, hit := leftPairTokens[token]
		return itIs && hit
	}
}

func (lexer *Lexer) isRightPairToken(token int) bool {
	if token == TOKEN_QUOTE {
		return lexer.getTopTokenOnMirrorStack() == TOKEN_QUOTE
	} else {
		itIs, hit := rightPairTokens[token]
		return itIs && hit
	}
}

func (lexer *Lexer) skipJSONSegment(n int) {
	lexer.JSONSegment = lexer.JSONSegment[n:]
}

func (lexer *Lexer) matchToken() (int, byte) {
	// finish
	fmt.Printf("[DUMP] len(lexer.JSONSegment): %d\n", len(lexer.JSONSegment))
	fmt.Printf("[DUMP] lexer.JSONSegment: '%s'\n", lexer.JSONSegment)
	if len(lexer.JSONSegment) == 0 {
		return TOKEN_EOF, byte(0)
	}
	tokenSynbol := lexer.JSONSegment[0]

	// check token
	switch tokenSynbol {
	case TOKEN_LEFT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACKET, tokenSynbol
	case TOKEN_RIGHT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACKET, tokenSynbol
	case TOKEN_LEFT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACE, tokenSynbol
	case TOKEN_RIGHT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACE, tokenSynbol
	case TOKEN_COLON_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COLON, tokenSynbol
	case TOKEN_DOT_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_DOT, tokenSynbol
	case TOKEN_COMMA_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COMMA, tokenSynbol
	case TOKEN_QUOTE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_QUOTE, tokenSynbol
	default:
		lexer.skipJSONSegment(1)
		return TOKEN_OTHERS, tokenSynbol
	}
}

func (lexer *Lexer) AppendString(str string) error {
	lexer.JSONSegment = str
	for {
		token, tokenSymbol := lexer.matchToken()
		switch token {
		case TOKEN_EOF:
			// nothing to do with TOKEN_EOF
		case TOKEN_OTHERS:
			lexer.JSONContent.WriteByte(tokenSymbol)
		default:
			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.isLeftPairToken(token) {
				lexer.pushTokenStack(token)
				lexer.pushMirrorTokenStack(mirrorTokenMap[token])
			} else if lexer.isRightPairToken(token) {
				lexer.pushTokenStack(token)
				lexer.popMirrorTokenStack()
			}
		}

		// check if end
		if token == TOKEN_EOF {
			break
		}
	}
	return nil
}

func (lexer *Lexer) CompleteJSON() string {
	return lexer.JSONContent.String() + lexer.dumpMirrorTokenStackToString()
}
