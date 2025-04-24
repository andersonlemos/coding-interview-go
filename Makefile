.PHONY: test

clean-cache:
	@go clean
test: clean-cache
	@go test -v ./...