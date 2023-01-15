# go-testing-course

Project to learn unit testing in Golang

- go get github.com/stretchr/testify
- go get github.com/jarcoal/httpmock

- go test ./controller -coverprofile="coverage.out"

- go test ./util -run=Parser -bench= > "bench.out"
- go test ./util -run=Parser -bench=.

## How to read the benchmark output

https://blog.logrocket.com/benchmarking-golang-improve-function-performance/

### BDD

Ginkgo
