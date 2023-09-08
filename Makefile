test:
	@echo "Running unit tests"
	go test -v ./...
	@echo "Unit tests finished successfully"

run:
	GO111MODULE=on go run -mod=vendor ./cmd/main.go