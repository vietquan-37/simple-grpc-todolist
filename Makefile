PHONY: run proto buf
run-server:
	@echo "Starting the server..."
	cd cmd/server && go run main.go

run-client:
	@echo "Starting the client..."
	cd cmd/client && go run main.go
buf: 
	@if exist pb\*.go del /Q pb\*.go
	buf generate
