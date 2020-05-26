# word-frequency
Go program that finds unique words in the text and counts their corresponding occurrences.

### Usage

``go run . < somefile``

Runs program on somefile.

### Tests

``go test -v -race -bench=. ./...``

Test results (with benchmarked baseline/goroutine versions):
```--- PASS: TestCounter_simple (0.00s)
   === RUN   TestCounter_GetMostCommon
   --- PASS: TestCounter_GetMostCommon (0.00s)
   goos: linux
   goarch: amd64
   BenchmarkCounter_CountBaseline-4               2         634735908 ns/op
   BenchmarkCounter_CountGoroutines-4             2         561476028 ns/op
   PASS
   ok      _/home/student/go-word-frequency/word-frequency 4.734s```
