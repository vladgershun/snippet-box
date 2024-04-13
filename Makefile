run:
	@go run ./cmd/web -port=":4000"
test:
	@go test -v ./cmd/web
