build:
	@echo
	@time -f "%C\nin %E" go build -o bin/doctor main.go
	@echo

run:
	@echo "──────────"
	@go run main.go
	@echo "──────────"

debug:
	@echo "──────────"
	@go run main.go --debug
	@echo "──────────"
