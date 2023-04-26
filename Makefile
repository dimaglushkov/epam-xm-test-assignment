lint:
	golangci-lint run ./...
.PHONY: lint

up:
	docker-compose up
.PHONY: up