package main

func matchStack(stack []int, tokens []int) bool {
	pointer := len(stack)
	tokensLeft := len(tokens)
	for {
		tokensLeft--
		pointer--
		if tokensLeft < 0 {
			break
		}
		if pointer < 0 {
			return false
		}
		if stack[pointer] != tokens[tokensLeft] {
			return false
		}
	}
	return true
}
