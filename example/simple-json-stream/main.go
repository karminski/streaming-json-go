package main

import (
	"fmt"

	streamingjson "github.com/karminski/streaming-json-go"
)

func main() {
	jsonSegmentA := `{"a":` // will complete to `{"a":null}`
	lexer := streamingjson.NewLexer()
	lexer.AppendString(jsonSegmentA)
	completedJSON := lexer.CompleteJSON()
	fmt.Printf("completedJSON: %s\n", completedJSON)
}
