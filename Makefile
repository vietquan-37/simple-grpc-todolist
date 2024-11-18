PHONY: run proto buf
run:
	@cd cmd && go run main.go
buf: 
	@buf generate

