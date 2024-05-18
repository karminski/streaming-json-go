package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCompleteJSON(t *testing.T) {
	json_p0001 := `{"name"`
	lexer := NewLexer()
	errInAppendString := lexer.AppendString(json_p0001)

	ret := lexer.CompleteJSON()

	assert.Nil(t, errInAppendString)

	assert.Equal(t, `{"name"}`, ret, "the token should be equal")
}
