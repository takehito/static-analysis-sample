goimports:
	goimports -w *.go
build: goimports
	go build
