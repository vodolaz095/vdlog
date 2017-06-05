check: test


test:
		gofmt -s -w .
		go vet .
		go test -v -coverprofile=c.out
		go tool cover -func=c.out