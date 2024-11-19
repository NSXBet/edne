.PHONY: all
all: test

.PHONY: clean
clean:

.PHONY: test
test:
	@go test -v ./...

.PHONY: test-cover
test-cover:
	@go test -v -cover ./...