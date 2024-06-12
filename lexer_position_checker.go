package streamingjsongo

// check if JSON stream stopped at an object properity's key start, like `{"`
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

// check if JSON stream stopped in an object properity's negative number value, like `:-`
func (lexer *Lexer) streamStoppedInAnObjectNegativeNumberValueStart() bool {
	// `:`, `-` in stack
	case1 := []int{
		TOKEN_COLON,
		TOKEN_NEGATIVE,
	}
	return matchStack(lexer.TokenStack, case1)
}

// check if JSON stream stopped in an object properity's negative number value, like `-`
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

// check if JSON stream stopped in an array
func (lexer *Lexer) streamStoppedInAnArray() bool {
	return lexer.getTopTokenOnMirrorStack() == TOKEN_RIGHT_BRACKET
}

// check if JSON stream stopped in an array's string value end, like `["value"`
func (lexer *Lexer) streamStoppedInAnArrayStringValueEnd() bool {
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

// check if JSON stream stopped in a string, like `""`
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

// check if JSON stream stopped in a number, like `[0-9]`
func (lexer *Lexer) streamStoppedInANumber() bool {
	return lexer.getTopTokenOnStack() == TOKEN_NUMBER
}

// check if JSON stream stopped in a number first decimal place, like `.?`
func (lexer *Lexer) streamStoppedInANumberDecimalPart() bool {
	return lexer.getTopTokenOnStack() == TOKEN_DOT
}

// check if JSON stream stopped in a number other decimal place (except first place), like `.[0-9]?`
func (lexer *Lexer) streamStoppedInANumberDecimalPartMiddle() bool {
	// `.`, TOKEN_NUMBER in stack
	case1 := []int{
		TOKEN_DOT,
		TOKEN_NUMBER,
	}
	return matchStack(lexer.TokenStack, case1)
}

// check if JSON stream stopped in escape character, like `\`
func (lexer *Lexer) streamStoppedWithLeadingEscapeCharacter() bool {
	return lexer.getTopTokenOnStack() == TOKEN_ESCAPE_CHARACTER
}
