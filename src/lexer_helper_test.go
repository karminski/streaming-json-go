package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_matchStack(t *testing.T) {
	stack := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_U,
		TOKEN_ALPHABET_LOWERCASE_N,
		TOKEN_COLON,
	}
	tokens := []int{
		TOKEN_RIGHT_BRACE,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_L,
		TOKEN_ALPHABET_LOWERCASE_U,
		TOKEN_ALPHABET_LOWERCASE_N,
		TOKEN_COLON,
	}

	matchResult := matchStack(stack, tokens)

	assert.Equal(t, true, matchResult, "the tokens should be match")
}

func Test_matchStack2(t *testing.T) {
	stack := []int{
		TOKEN_LEFT_BRACE,
		TOKEN_QUOTE,
		TOKEN_QUOTE,
		TOKEN_COLON,
		TOKEN_ALPHABET_LOWERCASE_N,
		TOKEN_ALPHABET_LOWERCASE_U,
	}
	tokens := []int{
		TOKEN_ALPHABET_LOWERCASE_N,
		TOKEN_ALPHABET_LOWERCASE_U,
	}

	matchResult := matchStack(stack, tokens)

	assert.Equal(t, true, matchResult, "the tokens should be match")
}
