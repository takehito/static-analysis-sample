goimports:
	goimports -w *.go
gofmt:
	go fmt
build: goimports gofmt
	go build
test: goimports gofmt
	go test
