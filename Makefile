build:
	@echo "Building application..."
	@env GOOS=linux CGO_ENABLED=0 go build -o ./bin/clean-arch ./cmd/clean-arch
	@echo "Done!"

run: build
	@./bin/clean-arch

test:
	@go test -coverprofile=coverage.out -v ./...

prepare:
	@echo "Generating .env file..."
	@cp -r ./examples/.env.example ./config/.env
	@echo "Done!"

.PHONY: test, build
