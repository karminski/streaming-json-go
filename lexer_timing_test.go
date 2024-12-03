package streamingjsongo

import (
	"fmt"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	testCaseA := `{"string": "这是一个字符串", "integer": 42, "float": 3.14159, "boolean_true": true, "boolean_false": false, "null": null, "object": {"empty_object": {}, "non_empty_object": {"key": "value"}, "nested_object": {"nested_key": {"sub_nested_key": "sub_nested_value"}}}, "array":["string in array", 123, 45.67, true, false, null, {"object_in_array": "object_value"},["nested_array"]]}`
	b.Run("streaming-json-go-append-json-segment", func(b *testing.B) {
		benchmarkAppendString(b, testCaseA)
	})
	b.Run("streaming-json-go-append-and-complete-json-segment", func(b *testing.B) {
		benchmarkAppendAndCompleteJSON(b, testCaseA)
	})

}

func benchmarkAppendString(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lexer := NewLexer()
			if err := lexer.AppendString(s); err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
		}
	})
}

func benchmarkAppendAndCompleteJSON(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	var completedJSON string
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lexer := NewLexer()
			if err := lexer.AppendString(s); err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			completedJSON = lexer.CompleteJSON()
			if len(completedJSON) == 0 {
				panic(fmt.Errorf("invalid completed JSON length"))
			}
		}
	})
}
