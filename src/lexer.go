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
	TOKEN_NEGATIVE                    // -
	TOKEN_NULL                        // null
	TOKEN_TRUE                        // true
	TOKEN_FLASE                       // false
	TOKEN_ALPHABET_LOWERCASE_A        // a
	TOKEN_ALPHABET_LOWERCASE_E        // e
	TOKEN_ALPHABET_LOWERCASE_F        // f
	TOKEN_ALPHABET_LOWERCASE_L        // l
	TOKEN_ALPHABET_LOWERCASE_N        // n
	TOKEN_ALPHABET_LOWERCASE_R        // r
	TOKEN_ALPHABET_LOWERCASE_S        // s
	TOKEN_ALPHABET_LOWERCASE_T        // t
	TOKEN_ALPHABET_LOWERCASE_U        // u
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
	TOKEN_NEGATIVE_SYMBOL             = '-'
	TOKEN_ALPHABET_LOWERCASE_A_SYMBOL = 'a'
	TOKEN_ALPHABET_LOWERCASE_E_SYMBOL = 'e'
	TOKEN_ALPHABET_LOWERCASE_F_SYMBOL = 'f'
	TOKEN_ALPHABET_LOWERCASE_L_SYMBOL = 'l'
	TOKEN_ALPHABET_LOWERCASE_N_SYMBOL = 'n'
	TOKEN_ALPHABET_LOWERCASE_R_SYMBOL = 'r'
	TOKEN_ALPHABET_LOWERCASE_S_SYMBOL = 's'
	TOKEN_ALPHABET_LOWERCASE_T_SYMBOL = 't'
	TOKEN_ALPHABET_LOWERCASE_U_SYMBOL = 'u'
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
	TOKEN_NEGATIVE:             "-",
	TOKEN_NULL:                 "null",
	TOKEN_TRUE:                 "true",
	TOKEN_FLASE:                "false",
	TOKEN_ALPHABET_LOWERCASE_A: "a",
	TOKEN_ALPHABET_LOWERCASE_E: "e",
	TOKEN_ALPHABET_LOWERCASE_F: "f",
	TOKEN_ALPHABET_LOWERCASE_L: "l",
	TOKEN_ALPHABET_LOWERCASE_N: "n",
	TOKEN_ALPHABET_LOWERCASE_R: "r",
	TOKEN_ALPHABET_LOWERCASE_S: "s",
	TOKEN_ALPHABET_LOWERCASE_T: "t",
	TOKEN_ALPHABET_LOWERCASE_U: "u",
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

func (lexer *Lexer) streamStoppedInAnObjectStart() bool {
	// `,`, `}` left
	case1 := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_COMMA,
	}
	if matchStack(lexer.MirrorTokenStack, case1) {
		return true
	}

	return false
}

func (lexer *Lexer) streamStoppedInAnObjectEnd() bool {
	// only `}` left
	if lexer.getTopTokenOnMirrorStack() == TOKEN_RIGHT_BRACE {
		return true
	}

	return false
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

func (lexer *Lexer) streamStoppedInANumber() bool {
	return lexer.getTopTokenOnStack() == TOKEN_NUMBER
}

func (lexer *Lexer) streamStoppedInANumberDecimalPart() bool {
	return lexer.getTopTokenOnStack() == TOKEN_DOT
}

func (lexer *Lexer) streamStoppedWithLeadingComma() bool {
	return lexer.getTopTokenOnStack() == TOKEN_COMMA
}

func (lexer *Lexer) streamStoppedWithLeadingEscapeCharacter() bool {
	return lexer.getTopTokenOnStack() == TOKEN_ESCAPE_CHARACTER
}

func (lexer *Lexer) pushCommaIntoJSONContent() {
	lexer.JSONContent.WriteByte(TOKEN_COMMA_SYMBOL)
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

func (lexer *Lexer) appendPaddingContent() {
	lexer.JSONContent.WriteString(lexer.PaddingContent.String())
}

func (lexer *Lexer) cleanPaddingContent() {
	lexer.PaddingContent.Reset()
}

func (lexer *Lexer) matchToken() (int, byte) {
	// finish
	fmt.Printf("[DUMP] len(lexer.JSONSegment): %d\n", len(lexer.JSONSegment))
	fmt.Printf("[DUMP] lexer.JSONSegment: '%s'\n", lexer.JSONSegment)
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
	case TOKEN_NEGATIVE_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_NEGATIVE, tokenSymbol
	case TOKEN_ALPHABET_LOWERCASE_A_SYMBOL:
		lexer.skipJSONSegment(1)
		return TOKEN_ALPHABET_LOWERCASE_A, tokenSymbol
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
		fmt.Printf("\n\n[DUMP] AppendString.token: `%s`\n", tokenNameMap[token])
		fmt.Printf("       lexer.TokenStack: `%+v`\n", lexer.TokenStack)
		fmt.Printf("       lexer.MirrorTokenStack: `%+v`\n", lexer.MirrorTokenStack)

		switch token {
		case TOKEN_EOF:
			// nothing to do with TOKEN_EOF
		case TOKEN_IGNORED:
			lexer.JSONContent.WriteByte(tokenSymbol)
		case TOKEN_OTHERS:
			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
				// pop `,` from  stack
				lexer.popTokenStack()
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
			fmt.Printf("    case TOKEN_LEFT_BRACKET:\n")
			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
			}
			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObjectArrayValueStart() {
				fmt.Printf("    lexer.streamStoppedInAnObjectArrayValueStart()\n")
				// pop `n`, `u`, `l`, `l` from mirror stack
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
			}
			// push `]` into mirror stack
			lexer.pushMirrorTokenStack(TOKEN_RIGHT_BRACKET)
		case TOKEN_RIGHT_BRACKET:
			fmt.Printf("    case TOKEN_RIGHT_BRACKET:\n")

			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			// push `]` into stack
			lexer.pushTokenStack(token)
			// pop `]` from mirror stack
			lexer.popMirrorTokenStack()
		case TOKEN_LEFT_BRACE:
			fmt.Printf("    case TOKEN_LEFT_BRACE:\n")
			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
			}
			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObjectObjectValueStart() {
				fmt.Printf("    lexer.streamStoppedInAnObjectObjectValueStart()\n")
				// pop `n`, `u`, `l`, `l` from mirror stack
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
			}
			// push `}` into mirror stack
			lexer.pushMirrorTokenStack(TOKEN_RIGHT_BRACE)
		case TOKEN_RIGHT_BRACE:
			fmt.Printf("    case TOKEN_RIGHT_BRACE:\n")

			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			// push `}` into stack
			lexer.pushTokenStack(token)
			// pop `}` from mirror stack
			lexer.popMirrorTokenStack()
		case TOKEN_QUOTE:
			fmt.Printf("    case TOKEN_QUOTE:\n")
			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
			}
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
				lexer.JSONContent.WriteByte(tokenSymbol)
				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}
			// start process
			lexer.JSONContent.WriteByte(tokenSymbol)
			lexer.pushTokenStack(token)
			if lexer.streamStoppedInAnObjectStart() {
				// case for new object properity key quote coming
				fmt.Printf("    lexer.streamStoppedInAnObjectStart()\n")

				// push `"`, `:`, `n`, `u`, `l`, `l` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_U)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_N)
				lexer.pushMirrorTokenStack(TOKEN_COLON)
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnArray() {
				fmt.Printf("    lexer.streamStoppedInAnArray()\n")

				// push `"` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnArrayValueEnd() {
				fmt.Printf("    lexer.streamStoppedInAnArrayValueEnd()\n")

				// pop `"` from mirror stack
				lexer.popMirrorTokenStack()
			} else if lexer.streamStoppedInAnObjectKeyStart() {
				// check if stopped in key of object's properity or value of object's properity
				fmt.Printf("    lexer.streamStoppedInAnObjectKeyStart()\n")

				// push `"`, `:`, `n`, `u`, `l`, `l` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_L)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_U)
				lexer.pushMirrorTokenStack(TOKEN_ALPHABET_LOWERCASE_N)
				lexer.pushMirrorTokenStack(TOKEN_COLON)
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnObjectKeyEnd() {
				// check if stopped in key of object's properity or value of object's properity
				fmt.Printf("    lexer.streamStoppedInAnObjectKeyEnd()\n")

				// pop `"` from mirror stack
				lexer.popMirrorTokenStack()
			} else if lexer.streamStoppedInAnObjectStringValueStart() {
				fmt.Printf("    lexer.streamStoppedInAnObjectStringValueStart()\n")
				// pop `n`, `u`, `l`, `l` from mirror stack
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				lexer.popMirrorTokenStack()
				// push `"` into mirror stack
				lexer.pushMirrorTokenStack(TOKEN_QUOTE)
			} else if lexer.streamStoppedInAnObjectValueEnd() {
				fmt.Printf("    lexer.streamStoppedInAnObjectValueEnd()\n")

				// pop `"` from mirror stack
				lexer.popMirrorTokenStack()
			} else {
				return fmt.Errorf("invalied quote token in json stream")
			}
		case TOKEN_COLON:
			fmt.Printf("    case TOKEN_COLON:\n")
			lexer.JSONContent.WriteByte(tokenSymbol)
			if lexer.streamStoppedInAString() {
				continue
			}
			lexer.pushTokenStack(token)
			// pop `:` from mirror stack
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_A:
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_A:\n")

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
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_S,
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_A,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			if !itIsPartOfTokenFalse() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_E:
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_E:\n")

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
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			// check if `t`, `r`, `u` in token stack and `e` in mirror stack
			itIsPartOfTokenTrue := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_T,
					TOKEN_ALPHABET_LOWERCASE_R,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			if !itIsPartOfTokenFalse() && !itIsPartOfTokenTrue() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_F:
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_F:\n")

			// \f escape `\`, `f`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
				lexer.JSONContent.WriteByte(tokenSymbol)
				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
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
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_L:\n")

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
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_S,
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			// check if `n`, `u` in token stack and `l`, `l` in mirror stack
			itIsPartOfTokenNull1 := func() bool {
				fmt.Printf("[]RUN itIsPartOfTokenNull1() !!!!!!!\n")

				left := []int{
					TOKEN_ALPHABET_LOWERCASE_N,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				fmt.Printf("    lexer.TokenStack: %+v\n", lexer.TokenStack)
				if !matchStack(lexer.TokenStack, left) {
					fmt.Printf("left does not match !!!!!!!\n")

					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					fmt.Printf("does not match !!!!!!!\n")
					return false
				}
				fmt.Printf("match !!!!!!!\n")

				return true
			}
			// check if `n`, `u`, `l` in token stack and `l` in mirror stack
			itIsPartOfTokenNull2 := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_N,
					TOKEN_ALPHABET_LOWERCASE_U,
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_L,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			if !itIsPartOfTokenFalse() && !itIsPartOfTokenNull1() && !itIsPartOfTokenNull2() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_N:
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_N:\n")
			// \n escape `\`, `n`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
				lexer.JSONContent.WriteByte(tokenSymbol)
				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
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
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_R:\n")
			// \r escape `\`, `r`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
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
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_U,
					TOKEN_ALPHABET_LOWERCASE_R,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			if !itIsPartOfTokenTrue() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_S:
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_S:\n")

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
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_S,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			if !itIsPartOfTokenFalse() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
		case TOKEN_ALPHABET_LOWERCASE_T:
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_T:\n")

			// \t escape `\`, `t`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
				lexer.JSONContent.WriteByte(tokenSymbol)
				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}

			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
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
			fmt.Printf("    case TOKEN_ALPHABET_LOWERCASE_U:\n")
			// unicode escape `\`, `u`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
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

			// check if `t`, `r` in token stack and, `u`, `e` in mirror stack
			itIsPartOfTokenTrue := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_T,
					TOKEN_ALPHABET_LOWERCASE_R,
				}
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_E,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			// check if `n` in token stack and `u`, `l`, `l` in mirror stack
			itIsPartOfTokenNull := func() bool {
				left := []int{
					TOKEN_ALPHABET_LOWERCASE_N,
				}
				if !matchStack(lexer.TokenStack, left) {
					return false
				}
				right := []int{
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_L,
					TOKEN_ALPHABET_LOWERCASE_U,
				}
				if !matchStack(lexer.MirrorTokenStack, right) {
					return false
				}
				return true
			}
			if !itIsPartOfTokenTrue() && !itIsPartOfTokenNull() {
				continue
			}
			lexer.pushTokenStack(token)
			lexer.popMirrorTokenStack()
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
			fmt.Printf("    case TOKEN_NUMBER:\n")

			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
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
		case TOKEN_ESCAPE_CHARACTER:
			fmt.Printf("    case TOKEN_ESCAPE_CHARACTER:\n")
			// double escape character `\`, `\`
			if lexer.streamStoppedWithLeadingEscapeCharacter() {
				lexer.pushEscapeCharacterIntoJSONContent()
				lexer.JSONContent.WriteByte(tokenSymbol)
				// pop `\` from  stack
				lexer.popTokenStack()
				continue
			}
			// just write escape character into stack and waitting other token trigger escape method.
			lexer.pushTokenStack(token)
		case TOKEN_NEGATIVE:
			fmt.Printf("    case TOKEN_NEGATIVE:\n")

			// in a string, just skip token
			if lexer.streamStoppedInAString() {
				lexer.JSONContent.WriteByte(tokenSymbol)
				continue
			}

			// check if json stream stopped with leading comma
			if lexer.streamStoppedWithLeadingComma() {
				lexer.pushCommaIntoJSONContent()
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
	mirrorTokens := lexer.dumpMirrorTokenStackToString()
	fmt.Printf("[DUMP] mirrorTokens: %s\n", mirrorTokens)
	return lexer.JSONContent.String() + lexer.dumpMirrorTokenStackToString()
}
