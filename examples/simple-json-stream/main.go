package main

import (
	"fmt"

	streamingjson "github.com/karminski/streaming-json-go"
)

func main() {
	// case A, complete the incomplete JSON object
	jsonSegmentA := `{"a":` // will complete to `{"a":null}`
	lexer := streamingjson.NewLexer()
	lexer.AppendString(jsonSegmentA)
	completedJSON := lexer.CompleteJSON()
	fmt.Printf("completedJSON: %s\n", completedJSON)

	// case B, complete the incomplete JSON array
	jsonSegmentB := `[t` // will complete to `[true]`
	lexer = streamingjson.NewLexer()
	lexer.AppendString(jsonSegmentB)
	completedJSON = lexer.CompleteJSON()
	fmt.Printf("completedJSON: %s\n", completedJSON)
}
