deps.devtools:
	@go get github.com/golang/dep/cmd/dep

deps: deps.devtools 
	@dep ensure -v

install:
	go install -v ./cmd/cli/*.go

build:
	@go build -v -o revive ./cmd/cli/*.go

test.all:
	@go test -v ./tests/...

