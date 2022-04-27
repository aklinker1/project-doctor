build:
	go build -o bin/doctor main.go

run:
	@echo "──────────"
	@USE_LOCAL_SCHEMA=true go run main.go
	@echo "──────────"

debug:
	@echo "──────────"
	@USE_LOCAL_SCHEMA=true go run main.go --debug
	@echo "──────────"
