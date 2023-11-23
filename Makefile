build:
	@echo "Building application..."
	@env GOOS=linux CGO_ENABLED=0 go build -o ./build/bin/clean-arch ./cmd/clean-arch
	@echo "Done!"

test:
	@echo "Starting tests..."
	@go test -coverprofile=coverage.out -v ./...
	@echo "Done!"

prepare:
	@echo "Generating .env file..."
	@cp -r ./examples/.env.example ./deploy/.env
	@echo "Done!"

grpc:
	@echo "Generating gRPC files..."
	@protoc --go_out=. --go-grpc_out=. ./internal/main/grpc/protofile/*.proto
	@echo "Done!"

swagger:
	@echo "Generating swagger files..."
	swag init -g cmd/clean-arch/main.go
	@mv ./docs/swagger* ./api/
	@echo "Done!"

up:
	@echo "Starting docker-compose..."
	@cd ./deploy && docker-compose up --remove-orphans -d
	@echo "Done!"

up_build: build
	@echo "Stopping docker-compose..."
	cd ./deploy && docker-compose down
	@echo "Starting docker-compose..."
	cd ./deploy && docker-compose up --remove-orphans --build -d
	@echo "Done!"

down:
	@echo "Stopping docker-compose..."
	cd ./deploy && docker-compose down
	@echo "Done!"

.PHONY: test, build
