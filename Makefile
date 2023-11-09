build:
	@go build -C ./cmd/clean-arch -o ../../bin/clean-arch

run: build
	@./bin/clean-arch

test:
	@go test -coverprofile=coverage.out -v ./...

.PHONY: test
