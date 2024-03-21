test:
	@go mod download
	@go test -v ./... -coverprofile=coverage.out