# streaming-json-go

[![codecov](https://codecov.io/github/karminski/streaming-json-go/graph/badge.svg?token=FOIHJ1NA7U)](https://codecov.io/github/karminski/streaming-json-go)


```go
import streamingjson "github.com/karminski/streaming-json-go"
```

Welcome to **streaming-json-go**, a groundbreaking library designed to revolutionize the way we handle stream JSON parsing.  

In an era dominated by LLMs (Large Language Models), the ability to efficiently parse JSON streams is more critical than ever. Traditionally, JSON parsing libraries have fallen short, requiring JSON data to be fully generated before any parsing can begin. streaming-json-go challenges this limitation head-on.

### Key Features

- **Real-Time JSON Parsing**: With streaming-json-go, you no longer need to wait for the entire JSON data to be generated. This library allows for the parsing of JSON as it is being streamed (this means JSON stream can stops at any position), significantly cutting down the time-to-first-token.
- **Seamless Integration**: Designed to complement existing JSON parsing libraries, streaming-json-go preprocesses incomplete JSON strings, transforming them into valid, parseable JSON. This means you can continue using your preferred JSON library with our tool seamlessly.
- **Enhanced User Experience**: By enabling real-time data processing, our library drastically reduces the wait time for end-users. Display JSON structures to users without the delay typically associated with complete JSON generation.

### Example Usage

Here’s a quick example to get you started:

```go
// init
lexer := streamingjson.NewLexer()

// append your JSON segment
targetJSONSegmentA := `{"a":` 
lexer.AppendString(targetJSONSegmentA)

// complete the JSON
completedJSONA := lexer.CompleteJSON()
fmt.Printf("%s\n", completedJSONA) // will print `{"a":null}`

// append more JSON segment
targetJSONSegmentB := `[tr`
lexer.AppendString(targetJSONSegmentB)

// complete the JSON again
completedJSONB := lexer.CompleteJSON()
fmt.Printf("%s\n", completedJSONB) // will print `{"a":[true]}`
```


For more examples please see: [examples](./examples/)

### Benchmarks

Using Go 1.21.1, single thread on Intel(R) Xeon(R) Platinum 8252C CPU @ 3.80GHz.

```
$ GOMAXPROCS=1 go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/karminski/streaming-json-go
cpu: Intel(R) Xeon(R) Platinum 8252C CPU @ 3.80GHz
BenchmarkParse/streaming-json-go-append-json-segment                142856              8148 ns/op          47.25 MB/s        3544 B/op         50 allocs/op
BenchmarkParse/streaming-json-go-append-and-complete-json-segment   142857              8131 ns/op          47.35 MB/s        3544 B/op         50 allocs/op
PASS
ok      github.com/karminski/streaming-json-go  2.539s
```

Using Go 1.22.4, single thread on Apple M2 Ultra.
```
karminski@kurumi streaming-json-go % GOMAXPROCS=1 go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/karminski/streaming-json-go
BenchmarkParse/streaming-json-go-append-json-segment                199710              5468 ns/op          70.41 MB/s        3544 B/op         50 allocs/op
BenchmarkParse/streaming-json-go-append-and-complete-json-segment   219049              5471 ns/op          70.37 MB/s        3544 B/op         50 allocs/op
PASS
ok  	github.com/karminski/streaming-json-go	2.548s
```

### About Golang Version

This library itself does not use any third-party Golang libraries, so it can run on version 1.14 and above. 

However, since the testify library is used in the tests, and the testify library requires at least Golang 1.17, this library is limited to requiring at least Golang 1.17 and above. 

If you need to run it on lower versions, you can consider copying the source code of this library directly into your project.


### License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
