lint:
	golangci-lint run ./...
.PHONY: lint

mock:
	mockgen -source ./internal/core/ports/repository.go -package repositories > ./internal/repositories/mock_repository.go
	mockgen -source ./internal/core/ports/events.go -package events > ./internal/events/mock_events.go
.PHONY: mock