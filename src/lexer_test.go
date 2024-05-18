package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCompleteJSON(t *testing.T) {
	json_p0000 := `{"name`
	lexer := NewLexer()
	errInAppendString := lexer.AppendString(json_p0000)

	ret := lexer.CompleteJSON()

	assert.Nil(t, errInAppendString)

	assert.Equal(t, `{"name"}`, ret, "the token should be equal")
}

func TestCompleteJSON(t *testing.T) {
	streamingJSONCase := map[string]string{
		`{`:  `{}`,
		`[`:  `[]`,
		`["`: `[""]`,
	}
	for testCase, expect := range streamingJSONCase {
		lexer := NewLexer()
		errInAppendString := lexer.AppendString(testCase)
		ret := lexer.CompleteJSON()
		assert.Nil(t, errInAppendString)
		assert.Equal(t, expect, ret, "unexpected JSON")
	}
}
