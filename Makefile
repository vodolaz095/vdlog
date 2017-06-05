check: test


test:
		go get -v .
		gofmt -s -w .
		go vet .
		go test -v -coverprofile=c.out
		go tool cover -func=c.out

all: test
