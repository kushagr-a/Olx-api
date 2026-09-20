.PHONY: build run clean 

build:
	@go build -o bin/api ./cmd/api/main.go

run: build
	@./bin/api

clean:
	@rm bin/api

migrate-up: # used to apply migration
	@go run ./cmd/migrate up

migrate-down: # used to undo migration
	@go run ./cmd/migrate down