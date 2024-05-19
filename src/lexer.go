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
	TOKEN_NULL                    // null
	TOKEN_TRUE                    // true
	TOKEN_FLASE                   // false
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
	for i := len(lexer.MirrorTokenStack) - 1; i >= 0; i-- {
		stackInString.WriteString(tokenNameMap[lexer.MirrorTokenStack[i]])
	}
	return stackInString.String()
}

func (lexer *Lexer) isLeftPairToken(token int) bool {
	itIs, hit := leftPairTokens[token]
	return itIs && hit
}

func (lexer *Lexer) isRightPairToken(token int) bool {
	itIs, hit := rightPairTokens[token]
	return itIs && hit
}

func (lexer *Lexer) skipJSONSegment(n int) {
	lexer.JSONSegment = lexer.JSONSegment[n:]
}

func (lexer *Lexer) streamStoppedInAnObject() bool {
	fmt.Printf("[DUMP] streamStoppedInAnObject.MirrorTokenStack: '%+v'\n", lexer.MirrorTokenStack)
	return lexer.getTopTokenOnMirrorStack() == TOKEN_RIGHT_BRACE
}

// check if JSON stream stopped in an object properity's key, like `{"field`
func (lexer *Lexer) streamStoppedInAnObjectKey() bool {
	mirrorStackLen := len(lexer.MirrorTokenStack)
	if mirrorStackLen-1 >= 0 && lexer.MirrorTokenStack[mirrorStackLen-1] == TOKEN_QUOTE {
		if mirrorStackLen-2 >= 0 && lexer.MirrorTokenStack[mirrorStackLen-2] == TOKEN_COLON {
			if mirrorStackLen-3 >= 0 && lexer.MirrorTokenStack[mirrorStackLen-3] == TOKEN_NULL {
				if mirrorStackLen-4 >= 0 && lexer.MirrorTokenStack[mirrorStackLen-4] == TOKEN_RIGHT_BRACE {
					return true
				}
			}
		}
	}
	return false
}

// check if JSON stream stopped in an object properity's value, like `{"field": "value`
func (lexer *Lexer) streamStoppedInAnObjectValue() bool {
	mirrorStackLen := len(lexer.MirrorTokenStack)
	if mirrorStackLen-1 >= 0 && lexer.MirrorTokenStack[mirrorStackLen-1] == TOKEN_QUOTE {
		if mirrorStackLen-2 >= 0 && lexer.MirrorTokenStack[mirrorStackLen-2] == TOKEN_RIGHT_BRACE {
			return true
		}
	}
	return false
}

func (lexer *Lexer) streamStoppedInAnArray() bool {
	return lexer.getTopTokenOnMirrorStack() == TOKEN_RIGHT_BRACKET
}

func (lexer *Lexer) streamStoppedInAString() bool {
	return lexer.getTopTokenOnMirrorStack() == TOKEN_QUOTE
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
		fmt.Printf("[DUMP] AppendString.token: %s\n", tokenNameMap[token])

		switch token {
		case TOKEN_EOF:
			// nothing to do with TOKEN_EOF
		case TOKEN_OTHERS:
			lexer.JSONContent.WriteByte(tokenSymbol)
		case TOKEN_QUOTE:
			fmt.Printf("    case TOKEN_QUOTE:\n")
			fmt.Printf("    lexer.streamStoppedInAnObject():%+v\n", lexer.streamStoppedInAnObject())
			lexer.JSONContent.WriteByte(tokenSymbol)
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObject() {
				fmt.Printf("    lexer.streamStoppedInAnObject()\n")
				// push `null`, `:`, `"` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_NULL)
				lexer.pushMirrorTokenStack(TOKEN_COLON)
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnArray() {
				fmt.Printf("    lexer.streamStoppedInAnArray()\n")

				// push `"` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAString() {
				fmt.Printf("    lexer.streamStoppedInAString()\n")

				// check if stopped in key of object's properity or value of object's properity
				if lexer.streamStoppedInAnObjectKey() {
					fmt.Printf("    lexer.streamStoppedInAnObjectKey()\n")

					// pop `"` from mirror stack
					lexer.popMirrorTokenStack()
				} else if lexer.streamStoppedInAnObjectValue() {
					fmt.Printf("    lexer.streamStoppedInAnObjectValue()\n")

					// pop `"` from mirror stack
					lexer.popMirrorTokenStack()
				} else {
					return fmt.Errorf("invalied quote token in json stream, incompleted object properity")
				}
			} else {
				return fmt.Errorf("invalied quote token in json stream")
			}
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

// complete missing parts for incomplete number, properity of object, null and boolean.
func (lexer *Lexer) completeMissingParts() string {
	// check if "," or ":" symbol on top of lexer.TokenStack
	if lexer.streamStoppedInAnObject() {
		switch lexer.getTopTokenOnStack() {
		case TOKEN_DOT:
			return `0`
		case TOKEN_COLON:
			return `: null`
		}
	}
	return ""
}

func (lexer *Lexer) CompleteJSON() string {
	mirrorTokens := lexer.dumpMirrorTokenStackToString()
	fmt.Printf("[DUMP] mirrorTokens: %s\n", mirrorTokens)
	return lexer.JSONContent.String() + lexer.dumpMirrorTokenStackToString()
}

// {         }
// {"      ":null}
