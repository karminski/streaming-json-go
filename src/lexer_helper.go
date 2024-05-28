package main

func isIgnoreToken(c byte) bool {
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	}
	return false
}

func matchStack(stack []int, tokens []int) bool {
	pointer := len(stack)
	tokensLeft := len(tokens)
	// fmt.Printf("current pointer: %+v, current tokensLeft: %+v\n", pointer, tokensLeft)

	for {
		tokensLeft--
		pointer--
		if tokensLeft < 0 {
			break
		}
		if pointer < 0 {
			return false
		}
		// fmt.Printf("current stack: %+v, current token: %+v\n", stack[pointer], tokens[tokensLeft])
		if stack[pointer] != tokens[tokensLeft] {
			return false
		}
	}
	return true
}
