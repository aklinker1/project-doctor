build:
	@echo
	@time -f "%C\nin %E" go build -o bin/doctor main.go
	@echo

run:
	@echo "──────────"
	@go run main.go
	@echo "──────────"
