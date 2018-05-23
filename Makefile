setup:
	./scripts/setup_database
test:
	go test ./cmd/... ./internal/...
