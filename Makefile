binary:
	$(info Making Linux x64 CLI binary...)
	$(info Binary saving to ./cmd/accelerator)
	GOOS=linux GOARCH=amd64 go build -o cmd/accelerator cmd/main.go
