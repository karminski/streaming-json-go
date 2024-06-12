package main

import (
	"fmt"
	"strings"
)

// token const
const (
	TOKEN_EOF                  = iota // end-of-file
	TOKEN_IGNORED                     // \t', '\n', '\v', '\f', '\r', ' '
	TOKEN_LEFT_BRACKET                // [
	TOKEN_RIGHT_BRACKET               // ]
	TOKEN_LEFT_BRACE                  // {
	TOKEN_RIGHT_BRACE                 // }
	TOKEN_COLON                       // :
	TOKEN_DOT                         // .
	TOKEN_COMMA                       // ,
	TOKEN_QUOTE                       // "
	TOKEN_ESCAPE_CHARACTER            // \
	TOKEN_SLASH                       // /
	TOKEN_NEGATIVE                    // -
	TOKEN_NULL                        // null
	TOKEN_TRUE                        // true
	TOKEN_FLASE                       // false
	TOKEN_ALPHABET_LOWERCASE_A        // a
	TOKEN_ALPHABET_LOWERCASE_B        // b
	TOKEN_ALPHABET_LOWERCASE_C        // c
	TOKEN_ALPHABET_LOWERCASE_D        // d
	TOKEN_ALPHABET_LOWERCASE_E        // e
	TOKEN_ALPHABET_LOWERCASE_F        // f
	TOKEN_ALPHABET_LOWERCASE_L        // l
	TOKEN_ALPHABET_LOWERCASE_N        // n
	TOKEN_ALPHABET_LOWERCASE_R        // r
	TOKEN_ALPHABET_LOWERCASE_S        // s
	TOKEN_ALPHABET_LOWERCASE_T        // t
	TOKEN_ALPHABET_LOWERCASE_U        // u
	TOKEN_ALPHABET_UPPERCASE_A        // A
	TOKEN_ALPHABET_UPPERCASE_B        // B
	TOKEN_ALPHABET_UPPERCASE_C        // C
	TOKEN_ALPHABET_UPPERCASE_D        // D
	TOKEN_ALPHABET_UPPERCASE_E        // E
	TOKEN_ALPHABET_UPPERCASE_F        // F
	TOKEN_NUMBER                      // number
	TOKEN_NUMBER_0                    // 0
	TOKEN_NUMBER_1                    // 1
	TOKEN_NUMBER_2                    // 2
	TOKEN_NUMBER_3                    // 3
	TOKEN_NUMBER_4                    // 4
	TOKEN_NUMBER_5                    // 5
	TOKEN_NUMBER_6                    // 6
	TOKEN_NUMBER_7                    // 7
	TOKEN_NUMBER_8                    // 8
	TOKEN_NUMBER_9                    // 9
	TOKEN_OTHERS                      // anything else in json
)

// token symbol const
const (
	TOKEN_LEFT_BRACKET_SYMBOL         = '['
	TOKEN_RIGHT_BRACKET_SYMBOL        = ']'
	TOKEN_LEFT_BRACE_SYMBOL           = '{'
	TOKEN_RIGHT_BRACE_SYMBOL          = '}'
	TOKEN_COLON_SYMBOL                = ':'
	TOKEN_DOT_SYMBOL                  = '.'
	TOKEN_COMMA_SYMBOL                = ','
	TOKEN_QUOTE_SYMBOL                = '"'
	TOKEN_ESCAPE_CHARACTER_SYMBOL     = '\\'
	TOKEN_SLASH_SYMBOL                = '/'
	TOKEN_NEGATIVE_SYMBOL             = '-'
	TOKEN_ALPHABET_LOWERCASE_A_SYMBOL = 'a'
	TOKEN_ALPHABET_LOWERCASE_B_SYMBOL = 'b'
	TOKEN_ALPHABET_LOWERCASE_C_SYMBOL = 'c'
	TOKEN_ALPHABET_LOWERCASE_D_SYMBOL = 'd'
	TOKEN_ALPHABET_LOWERCASE_E_SYMBOL = 'e'
	TOKEN_ALPHABET_LOWERCASE_F_SYMBOL = 'f'
	TOKEN_ALPHABET_LOWERCASE_L_SYMBOL = 'l'
	TOKEN_ALPHABET_LOWERCASE_N_SYMBOL = 'n'
	TOKEN_ALPHABET_LOWERCASE_R_SYMBOL = 'r'
	TOKEN_ALPHABET_LOWERCASE_S_SYMBOL = 's'
	TOKEN_ALPHABET_LOWERCASE_T_SYMBOL = 't'
	TOKEN_ALPHABET_LOWERCASE_U_SYMBOL = 'u'
	TOKEN_ALPHABET_UPPERCASE_A_SYMBOL = 'A'
	TOKEN_ALPHABET_UPPERCASE_B_SYMBOL = 'B'
	TOKEN_ALPHABET_UPPERCASE_C_SYMBOL = 'C'
	TOKEN_ALPHABET_UPPERCASE_D_SYMBOL = 'D'
	TOKEN_ALPHABET_UPPERCASE_E_SYMBOL = 'E'
	TOKEN_ALPHABET_UPPERCASE_F_SYMBOL = 'F'
	TOKEN_NUMBER_0_SYMBOL             = '0'
	TOKEN_NUMBER_1_SYMBOL             = '1'
	TOKEN_NUMBER_2_SYMBOL             = '2'
	TOKEN_NUMBER_3_SYMBOL             = '3'
	TOKEN_NUMBER_4_SYMBOL             = '4'
	TOKEN_NUMBER_5_SYMBOL             = '5'
	TOKEN_NUMBER_6_SYMBOL             = '6'
	TOKEN_NUMBER_7_SYMBOL             = '7'
	TOKEN_NUMBER_8_SYMBOL             = '8'
	TOKEN_NUMBER_9_SYMBOL             = '9'
)

var tokenNameMap = map[int]string{
	TOKEN_EOF:                  "EOF",
	TOKEN_LEFT_BRACKET:         "[",
	TOKEN_RIGHT_BRACKET:        "]",
	TOKEN_LEFT_BRACE:           "{",
	TOKEN_RIGHT_BRACE:          "}",
	TOKEN_COLON:                ":",
	TOKEN_DOT:                  ".",
	TOKEN_COMMA:                ",",
	TOKEN_QUOTE:                "\"",
	TOKEN_ESCAPE_CHARACTER:     "\\",
	TOKEN_SLASH:                "/",
	TOKEN_NEGATIVE:             "-",
	TOKEN_NULL:                 "null",
	TOKEN_TRUE:                 "true",
	TOKEN_FLASE:                "false",
	TOKEN_ALPHABET_LOWERCASE_A: "a",
	TOKEN_ALPHABET_LOWERCASE_B: "b",
	TOKEN_ALPHABET_LOWERCASE_C: "c",
	TOKEN_ALPHABET_LOWERCASE_D: "d",
	TOKEN_ALPHABET_LOWERCASE_E: "e",
	TOKEN_ALPHABET_LOWERCASE_F: "f",
	TOKEN_ALPHABET_LOWERCASE_L: "l",
	TOKEN_ALPHABET_LOWERCASE_N: "n",
	TOKEN_ALPHABET_LOWERCASE_R: "r",
	TOKEN_ALPHABET_LOWERCASE_S: "s",
	TOKEN_ALPHABET_LOWERCASE_T: "t",
	TOKEN_ALPHABET_LOWERCASE_U: "u",
	TOKEN_ALPHABET_UPPERCASE_A: "A",
	TOKEN_ALPHABET_UPPERCASE_B: "B",
	TOKEN_ALPHABET_UPPERCASE_C: "C",
	TOKEN_ALPHABET_UPPERCASE_D: "D",
	TOKEN_ALPHABET_UPPERCASE_E: "E",
	TOKEN_ALPHABET_UPPERCASE_F: "F",
	TOKEN_NUMBER_0:             "0",
	TOKEN_NUMBER_1:             "1",
	TOKEN_NUMBER_2:             "2",
	TOKEN_NUMBER_3:             "3",
	TOKEN_NUMBER_4:             "4",
	TOKEN_NUMBER_5:             "5",
	TOKEN_NUMBER_6:             "6",
	TOKEN_NUMBER_7:             "7",
	TOKEN_NUMBER_8:             "8",
	TOKEN_NUMBER_9:             "9",
}

type Lexer struct {
	JSONContent      strings.Builder
	PaddingContent   strings.Builder
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
	return lexer.MirrorTokenStack[mirrotTokenStackLen-1]
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

func (lexer *Lexer) skipJSONSegment(n int) {
	lexer.JSONSegment = lexer.JSONSegment[n:]
}

func (lexer *Lexer) streamStoppedInAnObjectStart() bool {
	// `,`, `}` left
	case1 := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_COMMA,
	}
	return matchStack(lexer.MirrorTokenStack, case1)
}

func (lexer *Lexer) streamStoppedInAnObjectKeyStart() bool {
	// `{`, `"` in stack, or `,`, `"` in stack
	case1 := []int{
		TOKEN_LEFT_BRACE,
		TOKEN_QUOTE,
	}
	case2 := []int{
		TOKEN_COMMA,
		TOKEN_QUOTE,
	}
	// `}` in mirror stack
	case3 := []int{
		TOKEN_RIGHT_BRACE,
	}
	return (matchStack(lexer.TokenStack, case1) || matchStack(lexer.TokenStack, case2)) && matchStack(lexer.MirrorTokenStack, case3)
}

// check if JSON stream stopped in an object properity's key, like `{"field`
func (lexer *Lexer) streamStoppedInAnObjectKeyEnd() bool {
	// `{`, `"`, `"` in stack, or `,`, `"`, `"` in stack
	case1 := []int{
		TOKEN_LEFT_BRACE,
		TOKEN_QUOTE,
		TOKEN_QUOTE,
	}
	case2 := []int{
		TOKEN_COMMA,
		TOKEN_QUOTE,
		TOKEN_QUOTE,
	}
	// `"`, `:`, `n`, `u`, `l`, `l`, `}` in mirror stack
	case3 := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_U,
		TOKEN_ALPHABET_LOWERCASE_N,
		TOKEN_COLON,
		TOKEN_QUOTE,
	}
	return (matchStack(lexer.TokenStack, case1) || matchStack(lexer.TokenStack, case2)) && matchStack(lexer.MirrorTokenStack, case3)
}

// check if JSON stream stopped in an object properity's value start, like `{"field": "`
func (lexer *Lexer) streamStoppedInAnObjectStringValueStart() bool {
	// `:`, `"` in stack
	case1 := []int{
		TOKEN_COLON,
		TOKEN_QUOTE,
	}
	// `n`, `u`, `l`, `l`, `}` in mirror stack
	case2 := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_U,
		TOKEN_ALPHABET_LOWERCASE_N,
	}
	return matchStack(lexer.TokenStack, case1) && matchStack(lexer.MirrorTokenStack, case2)
}

// check if JSON stream stopped in an object properity's value finish, like `{"field": "value"`
func (lexer *Lexer) streamStoppedInAnObjectValueEnd() bool {
	// `"`, `}` left
	tokens := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_QUOTE,
	}
	return matchStack(lexer.MirrorTokenStack, tokens)
}

// check if JSON stream stopped in an object properity's value start by array, like `{"field":[`
func (lexer *Lexer) streamStoppedInAnObjectArrayValueStart() bool {
	// `:`, `[` in stack
	case1 := []int{
		TOKEN_COLON,
		TOKEN_LEFT_BRACKET,
	}
	// `n`, `u`, `l`, `l`, `}` in mirror stack
	case2 := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_U,
		TOKEN_ALPHABET_LOWERCASE_N,
	}
	return matchStack(lexer.TokenStack, case1) && matchStack(lexer.MirrorTokenStack, case2)
}

// check if JSON stream stopped in an object properity's value start by array, like `{"field":{`
func (lexer *Lexer) streamStoppedInAnObjectObjectValueStart() bool {
	// `:`, `{` in stack
	case1 := []int{
		TOKEN_COLON,
		TOKEN_LEFT_BRACE,
	}
	// `n`, `u`, `l`, `l`, `}` in mirror stack
	case2 := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_U,
		TOKEN_ALPHABET_LOWERCASE_N,
	}
	return matchStack(lexer.TokenStack, case1) && matchStack(lexer.MirrorTokenStack, case2)
}

func (lexer *Lexer) streamStoppedInAnObjectNegativeNumberValueStart() bool {
	// `:`, `-` in stack
	case1 := []int{
		TOKEN_COLON,
		TOKEN_NEGATIVE,
	}

	return matchStack(lexer.TokenStack, case1)
}

func (lexer *Lexer) streamStoppedInANegativeNumberValueStart() bool {
	// `-` in stack
	case1 := []int{
		TOKEN_NEGATIVE,
	}
	// `0`in mirror stack
	case2 := []int{
		TOKEN_NUMBER_0,
	}
	return matchStack(lexer.TokenStack, case1) && matchStack(lexer.MirrorTokenStack, case2)
}

func (lexer *Lexer) streamStoppedInAnArray() bool {
	return lexer.getTopTokenOnMirrorStack() == TOKEN_RIGHT_BRACKET
}

func (lexer *Lexer) streamStoppedInAnArrayValueEnd() bool {
	// `"`, `"` in stack
	case1 := []int{
		TOKEN_QUOTE,
		TOKEN_QUOTE,
	}
	// `"`, `]` in mirror stack
	case2 := []int{
		TOKEN_RIGHT_BRACKET,
		TOKEN_QUOTE,
	}
	return matchStack(lexer.TokenStack, case1) && matchStack(lexer.MirrorTokenStack, case2)
}

// check if JSON stream stopped in an object properity's value start by array, like `{"field":{`
func (lexer *Lexer) streamStoppedInAnObjectNullValuePlaceholderStart() bool {
	// `n`, `u`, `l`, `l`, `}` in mirror stack
	case1 := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_U,
		TOKEN_ALPHABET_LOWERCASE_N,
	}
	return matchStack(lexer.MirrorTokenStack, case1)
}

func (lexer *Lexer) streamStoppedInAString() bool {
	return lexer.getTopTokenOnStack() == TOKEN_QUOTE && lexer.getTopTokenOnMirrorStack() == TOKEN_QUOTE
}

// check if JSON stream stopped in a string's unicode escape, like `\u????`
func (lexer *Lexer) streamStoppedInAnStringUnicodeEscape() bool {
	// `\`, `u` in stack
	case1 := []int{
		TOKEN_ESCAPE_CHARACTER,
		TOKEN_ALPHABET_LOWERCASE_U,
	}
	// `"` in mirror stack
	case2 := []int{
		TOKEN_QUOTE,
	}
	return matchStack(lexer.TokenStack, case1) && matchStack(lexer.MirrorTokenStack, case2)
}

func (lexer *Lexer) streamStoppedInANumber() bool {
	return lexer.getTopTokenOnStack() == TOKEN_NUMBER
}

func (lexer *Lexer) streamStoppedInANumberDecimalPart() bool {
	return lexer.getTopTokenOnStack() == TOKEN_DOT
}

func (lexer *Lexer) streamStoppedInANumberDecimalPartMiddle() bool {
	// `.`, TOKEN_NUMBER in stack
	case1 := []int{
		TOKEN_DOT,
		TOKEN_NUMBER,
	}
	return matchStack(lexer.TokenStack, case1)
}

func (lexer *Lexer) streamStoppedWithLeadingEscapeCharacter() bool {
	return lexer.getTopTokenOnStack() == TOKEN_ESCAPE_CHARACTER
}

func (lexer *Lexer) pushEscapeCharacterIntoJSONContent() {
	lexer.JSONContent.WriteByte(TOKEN_ESCAPE_CHARACTER_SYMBOL)
}

func (lexer *Lexer) pushNegativeIntoJSONContent() {
	lexer.JSONContent.WriteByte(TOKEN_NEGATIVE_SYMBOL)
}

func (lexer *Lexer) pushByteIntoPaddingContent(b byte) {
	lexer.PaddingContent.WriteByte(b)
}

func (lexer *Lexer) appendPaddingContentToJSONContent() {
	lexer.JSONContent.WriteString(lexer.PaddingContent.String())
}

func (lexer *Lexer) havePaddingContent() bool {
	return lexer.PaddingContent.Len() > 0
}

func (lexer *Lexer) cleanPaddingContent() {
	lexer.PaddingContent.Reset()
}

func (lexer *Lexer) matchToken() (int, byte) {
	// finish

	if len(lexer.JSONSegment) == 0 {
		return TOKEN_EOF, byte(0)
	}
	tokenSymbol := lexer.JSONSegment[0]

	// check if ignored token
	if isIgnoreToken(tokenSymbol) {
		lexer.skipJSONSegment(1)
		return TOKEN_IGNORED, tokenSymbol
	}

	// check token
	switch tokenSymbol {
	case TOKEN_LEFT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACKET, tokenSymbol
	case TOKEN_RIGHT_BRACKET_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACKET, tokenSymbol
	case TOKEN_LEFT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_LEFT_BRACE, tokenSymbol
	case TOKEN_RIGHT_BRACE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_RIGHT_BRACE, tokenSymbol
	case TOKEN_COLON_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COLON, tokenSymbol
	case TOKEN_DOT_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_DOT, tokenSymbol
	case TOKEN_COMMA_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_COMMA, tokenSymbol
	case TOKEN_QUOTE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_QUOTE, tokenSymbol
	case TOKEN_ESCAPE_CHARACTER_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ESCAPE_CHARACTER, tokenSymbol
	case TOKEN_SLASH_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_SLASH, tokenSymbol
	case TOKEN_NEGATIVE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NEGATIVE, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_A_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_A, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_B_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_B, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_C_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_C, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_D_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_D, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_E_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_E, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_F_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_F, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_L_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_L, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_N_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_N, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_R_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_R, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_S_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_S, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_T_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_T, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_U_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_U, tokenSymbol
	case TOKEN_ALPHABET_UPPERCASE_A_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_UPPERCASE_A, tokenSymbol
	case TOKEN_ALPHABET_UPPERCASE_B_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_UPPERCASE_B, tokenSymbol
	case TOKEN_ALPHABET_UPPERCASE_C_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_UPPERCASE_C, tokenSymbol
	case TOKEN_ALPHABET_UPPERCASE_D_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_UPPERCASE_D, tokenSymbol
	case TOKEN_ALPHABET_UPPERCASE_E_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_UPPERCASE_E, tokenSymbol
	case TOKEN_ALPHABET_UPPERCASE_F_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_UPPERCASE_F, tokenSymbol
	case TOKEN_NUMBER_0_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_0, tokenSymbol
	case TOKEN_NUMBER_1_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_1, tokenSymbol
	case TOKEN_NUMBER_2_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_2, tokenSymbol
	case TOKEN_NUMBER_3_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_3, tokenSymbol
	case TOKEN_NUMBER_4_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_4, tokenSymbol
	case TOKEN_NUMBER_5_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_5, tokenSymbol
	case TOKEN_NUMBER_6_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_6, tokenSymbol
	case TOKEN_NUMBER_7_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_7, tokenSymbol
	case TOKEN_NUMBER_8_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_8, tokenSymbol
	case TOKEN_NUMBER_9_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NUMBER_9, tokenSymbol
	default:
		lexer.skipJSONSegment(1)
		return TOKEN_OTHERS, tokenSymbol
	}
}

func (lexer *Lexer) AppendString(str string) error {
	lexer.JSONSegment = str
	for {
		token, tokenSymbol := lexer.matchToken()

		switch token {
		case TOKEN_EOF:
			// nothing to do with TOKEN_EOF
		case TOKEN_IGNORED:
			if lexer.streamStoppedInAString() {
				lexer.JSONContent.WriteByte(tokenSymbol)
				continue
			}
			lexer.pushByteIntoPaddingContent(tokenSymbol)
		case TOKEN_OTHERS:
			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			// double escape character `\`, `\`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
				lexer.JSONContent.WriteByte(tokenSymbol)
				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}
			lexer.JSONContent.WriteByte(tokenSymbol)
		case TOKEN_LEFT_BRACKET:

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}
			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObjectArrayValueStart() {

				// pop `n`, `u`, `l`, `l` from mirror stack
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
			}
			// push `]` into mirror stack
			lexer.pushMirrorTokenStack(TOKEN_RIGHT_BRACKET)
		case TOKEN_RIGHT_BRACKET:

			if lexer.streamStoppedInAString() {
				lexer.JSONContent.WriteByte(tokenSymbol)
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			lexer.JSONContent.WriteByte(tokenSymbol)

			// push `]` into stack
			lexer.pushTokenStack(token)
			// pop `]` from mirror stack
			lexer.popMirrorTokenStack()
		case TOKEN_LEFT_BRACE:

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObjectObjectValueStart() {

				// pop `n`, `u`, `l`, `l` from mirror stack
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
			}
			// push `}` into mirror stack
			lexer.pushMirrorTokenStack(TOKEN_RIGHT_BRACE)
		case TOKEN_RIGHT_BRACE:

			if lexer.streamStoppedInAString() {
				lexer.JSONContent.WriteByte(tokenSymbol)
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}
			lexer.JSONContent.WriteByte(tokenSymbol)

			// push `}` into stack
			lexer.pushTokenStack(token)
			// pop `}` from mirror stack
			lexer.popMirrorTokenStack()
		case TOKEN_QUOTE:

			// check if escape quote `\"`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			// start process
			lexer.JSONContent.WriteByte(tokenSymbol)
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObjectStart() {
				// case for new object properity key quote coming

				// push `"`, `:`, `n`, `u`, `l`, `l` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_U)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_N)
				lexer.pushMirrorTokenStack(TOKEN_COLON)
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnArray() {

				// push `"` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnArrayValueEnd() {

				// pop `"` from mirror stack
				lexer.popMirrorTokenStack()
			} else if lexer.streamStoppedInAnObjectKeyStart() {
				// check if stopped in key of object's properity or value of object's properity

				// push `"`, `:`, `n`, `u`, `l`, `l` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_U)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_N)
				lexer.pushMirrorTokenStack(TOKEN_COLON)
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnObjectKeyEnd() {
				// check if stopped in key of object's properity or value of object's properity

				// pop `"` from mirror stack
				lexer.popMirrorTokenStack()
			} else if lexer.streamStoppedInAnObjectStringValueStart() {

				// pop `n`, `u`, `l`, `l` from mirror stack
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				// push `"` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnObjectValueEnd() {

				// pop `"` from mirror stack
				lexer.popMirrorTokenStack()
			} else {
				return fmt.Errorf("invalied quote token in json stream")
			}
		case TOKEN_COLON:

			if lexer.streamStoppedInAString() {
				lexer.JSONContent.WriteByte(tokenSymbol)
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}
			lexer.JSONContent.WriteByte(tokenSymbol)

			lexer.pushTokenStack(token)
			// pop `:` from mirror stack
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_A:

			// as hex in unicode
			if lexer.streamStoppedInAnStringUnicodeEscape() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				// check if unicode escape is full length
				if lexer.PaddingContent.Len() == 6 {
					lexer.appendPaddingContentToJSONContent()
					lexer.cleanPaddingContent()
					// pop `\`, `u` from stack
					lexer.popTokenStack()
					lexer.popTokenStack()
				}
				continue
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}
			// check if `f` in token stack and `a`, `l`, `s`, `e in mirror stack
			itIsPartOfTokenFalse := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_F,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_S,
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_A,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			if !itIsPartOfTokenFalse() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_B:

			// as hex in unicode
			if lexer.streamStoppedInAnStringUnicodeEscape() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				// check if unicode escape is full length
				if lexer.PaddingContent.Len() == 6 {
					lexer.appendPaddingContentToJSONContent()
					lexer.cleanPaddingContent()
					// pop `\`, `u` from stack
					lexer.popTokenStack()
					lexer.popTokenStack()
				}
				continue
			}

			// \b escape `\`, `b`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}
		case TOKEN_ALPHABET_LOWERCASE_E:

			// as hex in unicode
			if lexer.streamStoppedInAnStringUnicodeEscape() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				// check if unicode escape is full length
				if lexer.PaddingContent.Len() == 6 {
					lexer.appendPaddingContentToJSONContent()
					lexer.cleanPaddingContent()
					// pop `\`, `u` from stack
					lexer.popTokenStack()
					lexer.popTokenStack()
				}
				continue
			}

			// check if in a number, as `e` (exponent) in scientific notation
			if lexer.streamStoppedInANumberDecimalPartMiddle() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				continue
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}

			// check if `f`, `a`, `l`, `s` in token stack and `e` in mirror stack
			itIsPartOfTokenFalse := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_F,
					TOKEN_ALPHABET_LOWERCASE_A,
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_S,
				}

				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			// check if `t`, `r`, `u` in token stack and `e` in mirror stack
			itIsPartOfTokenTrue := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_T,
					TOKEN_ALPHABET_LOWERCASE_R,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			if !itIsPartOfTokenFalse() && !itIsPartOfTokenTrue() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_F:

			// as hex in unicode
			if lexer.streamStoppedInAnStringUnicodeEscape() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				// check if unicode escape is full length
				if lexer.PaddingContent.Len() == 6 {
					lexer.appendPaddingContentToJSONContent()
					lexer.cleanPaddingContent()
					// pop `\`, `u` from stack
					lexer.popTokenStack()
					lexer.popTokenStack()
				}
				continue
			}

			// \f escape `\`, `f`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}

			// push `f` into stack
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnArray() {
				// in array
				// push `a`, `l`, `s`, `e`
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_E)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_S)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_A)
			} else {
				// in object
				// pop `n`, `u`, `l`, `l`
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				// push `a`, `l`, `s`, `e`
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_E)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_S)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_A)
			}
		case TOKEN_ALPHABET_LOWERCASE_L:

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}
			// check if `f`, `a` in token stack and, `l`, `s`, `e` in mirror stack
			itIsPartOfTokenFalse := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_F,
					TOKEN_ALPHABET_LOWERCASE_A,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_S,
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			// check if `n`, `u` in token stack and `l`, `l` in mirror stack
			itIsPartOfTokenNull1 := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_N,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			// check if `n`, `u`, `l` in token stack and `l` in mirror stack
			itIsPartOfTokenNull2 := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_N,
					TOKEN_ALPHABET_LOWERCASE_U,
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			if !itIsPartOfTokenFalse() && !itIsPartOfTokenNull1() && !itIsPartOfTokenNull2() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_N:

			// \n escape `\`, `n`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}

			// push `n`
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnArray() {
				// in array, push `u`, `l`, `l`
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_U)
			} else {
				// in object, pop `n`
				lexer.popMirrorTokenStack()
			}
		case TOKEN_ALPHABET_LOWERCASE_R:

			// \r escape `\`, `r`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}
			// check if `t` in token stack and `r`, `u`, `e in mirror stack
			itIsPartOfTokenTrue := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_T,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_U,
					TOKEN_ALPHABET_LOWERCASE_R,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			if !itIsPartOfTokenTrue() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_S:

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}
			// check if `f`, `a`, `l` in token stack and `s`, `e in mirror stack
			itIsPartOfTokenFalse := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_F,
					TOKEN_ALPHABET_LOWERCASE_A,
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_S,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			if !itIsPartOfTokenFalse() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_T:

			// \t escape `\`, `t`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}

			// push `t` to stack
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnArray() {
				// in array
				// push `r`, `u`, `e`
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_E)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_U)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_R)
			} else {
				// in object
				// pop `n`, `u`, `l`, `l`
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				// push `r`, `u`, `e`
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_E)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_U)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_R)
			}
		case TOKEN_ALPHABET_LOWERCASE_U:

			// unicode escape `\`, `u`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushTokenStack(token)
				lexer.PaddingContent.WriteByte(tokenSymbol)
				continue
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}

			// check if `t`, `r` in token stack and, `u`, `e` in mirror stack
			itIsPartOfTokenTrue := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_T,
					TOKEN_ALPHABET_LOWERCASE_R,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			// check if `n` in token stack and `u`, `l`, `l` in mirror stack
			itIsPartOfTokenNull := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_N,
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				return matchStack(lexer.TokenStack, left) && matchStack(lexer.MirrorTokenStack, right)
			}
			if !itIsPartOfTokenTrue() && !itIsPartOfTokenNull() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_UPPERCASE_A:
			fallthrough
		case TOKEN_ALPHABET_UPPERCASE_B:
			fallthrough
		case TOKEN_ALPHABET_UPPERCASE_C:
			fallthrough
		case TOKEN_ALPHABET_UPPERCASE_D:
			fallthrough
		case TOKEN_ALPHABET_LOWERCASE_C:
			fallthrough
		case TOKEN_ALPHABET_LOWERCASE_D:
			fallthrough
		case TOKEN_ALPHABET_UPPERCASE_F:

			// as hex in unicode
			if lexer.streamStoppedInAnStringUnicodeEscape() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				// check if unicode escape is full length
				if lexer.PaddingContent.Len() == 6 {
					lexer.appendPaddingContentToJSONContent()
					lexer.cleanPaddingContent()
					// pop `\`, `u` from stack
					lexer.popTokenStack()
					lexer.popTokenStack()
				}
				continue
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}
		case TOKEN_ALPHABET_UPPERCASE_E:

			// as hex in unicode
			if lexer.streamStoppedInAnStringUnicodeEscape() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				// check if unicode escape is full length
				if lexer.PaddingContent.Len() == 6 {
					lexer.appendPaddingContentToJSONContent()
					lexer.cleanPaddingContent()
					// pop `\`, `u` from stack
					lexer.popTokenStack()
					lexer.popTokenStack()
				}
				continue
			}

			// check if in a number, as `E` (exponent) in scientific notation
			if lexer.streamStoppedInANumberDecimalPartMiddle() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				continue
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				continue
			}
		case TOKEN_NUMBER_0:
			fallthrough
		case TOKEN_NUMBER_1:
			fallthrough
		case TOKEN_NUMBER_2:
			fallthrough
		case TOKEN_NUMBER_3:
			fallthrough
		case TOKEN_NUMBER_4:
			fallthrough
		case TOKEN_NUMBER_5:
			fallthrough
		case TOKEN_NUMBER_6:
			fallthrough
		case TOKEN_NUMBER_7:
			fallthrough
		case TOKEN_NUMBER_8:
			fallthrough
		case TOKEN_NUMBER_9:

			if lexer.streamStoppedInAnStringUnicodeEscape() {
				lexer.pushByteIntoPaddingContent(tokenSymbol)
				// check if unicode escape is full length
				if lexer.PaddingContent.Len() == 6 {
					lexer.appendPaddingContentToJSONContent()
					lexer.cleanPaddingContent()
					// pop `\`, `u` from stack
					lexer.popTokenStack()
					lexer.popTokenStack()
				}
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			// in negative part of a number
			if lexer.streamStoppedInANegativeNumberValueStart() {
				lexer.pushNegativeIntoJSONContent()
				// pop `0` from mirror stack
				lexer.popMirrorTokenStack()
			}

			lexer.JSONContent.WriteByte(tokenSymbol)
			// in a string or a number, just skip token
			if lexer.streamStoppedInAString() || lexer.streamStoppedInANumber() {
				continue
			}

			// in decimal part of a number
			if lexer.streamStoppedInANumberDecimalPart() {
				lexer.pushTokenStack(TOKEN_NUMBER)
				// pop placeholder `0` in decimal part
				lexer.popMirrorTokenStack()
				continue
			}

			// first number token, push token into stack
			lexer.pushTokenStack(TOKEN_NUMBER)
			if lexer.streamStoppedInAnArray() {
				continue
			} else if lexer.streamStoppedInAnObjectNullValuePlaceholderStart() {
				// pop `n`, `u`, `l`, `l`
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
			}

		case TOKEN_COMMA:
			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				lexer.JSONContent.WriteByte(tokenSymbol)
				continue
			}
			// in a object or a array, keep the comma in stack but not write it into JSONContent, until next token arrival
			// the comma must following with token: quote, null, true, false, number
			lexer.pushByteIntoPaddingContent(tokenSymbol)
			lexer.pushTokenStack(token)
		case TOKEN_DOT:
			// in a string, just skip token
			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			// use 0 for decimal part place holder
			lexer.pushTokenStack(token)
			lexer.pushMirrorTokenStack(TOKEN_NUMBER_0)
		case TOKEN_SLASH:

			// escape character `\`, `/`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}
			// error, invalied json
		case TOKEN_ESCAPE_CHARACTER:

			// double escape character `\`, `\`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				// push padding escape character `\` into JSON content
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()

				// write current character into JSON content
				lexer.JSONContent.WriteByte(tokenSymbol)

				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}
			// just write escape character into stack and waitting other token trigger escape method.
			lexer.pushTokenStack(token)
			lexer.pushByteIntoPaddingContent(TOKEN_ESCAPE_CHARACTER_SYMBOL)
		case TOKEN_NEGATIVE:

			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				lexer.JSONContent.WriteByte(tokenSymbol)
				continue
			}

			// check if json stream stopped with padding content
			if lexer.havePaddingContent() {
				lexer.appendPaddingContentToJSONContent()
				lexer.cleanPaddingContent()
			}

			// just write negative character into stack and waitting other token trigger it.
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObjectNegativeNumberValueStart() {
				// pop `n`, `u`, `l`, `l` from mirror stack
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
			}
			// push `0` into mirror stack for placeholder
			lexer.pushMirrorTokenStack(TOKEN_NUMBER_0)
		default:
			return fmt.Errorf("unexpected token: `%d`, token symbol: `%c`", token, tokenSymbol)
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
