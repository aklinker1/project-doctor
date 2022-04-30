build:
	go build -o bin/doctor main.go

run:
	@echo "──────────"
	@go run main.go
	@echo "──────────"

debug:
	@echo "──────────"
	@go run main.go --debug
	@echo "──────────"

test:
	@echo "──────────"
	@go test ./cli/...
	@echo "──────────"
