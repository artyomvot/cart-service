run-all:
	go run ./cmd/server/server.go

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out
	
complexity-check:
	gocyclo -over 15 ./internal/
	gocognit -over 15 ./internal/