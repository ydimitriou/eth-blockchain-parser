test:
	@echo "Running unit tests"
	go test -v ./...
	@echo "Unit tests finished successfully"

run:
	GO111MODULE=on go run -mod=vendor ./cmd/main.go

lint:
	@echo "Checking lint"
	golangci-lint run --enable errcheck,gofmt,goimports --exclude-use-default=false --modules-download-mode=vendor --build-tags integration --timeout=3m
	@echo "Lint success"