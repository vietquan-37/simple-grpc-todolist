.PHONY: run proto
run:
	@cd cmd go run main.go
proto:
	@if exist pb\*.go del /Q pb\*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto