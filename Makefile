.PHONY: fmt test vet install clean

install:
	@go get -t -v ./...

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test -cover -v -race ./...

clean:
	@rm -rf bin
