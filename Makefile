PHONY: run proto buf
run:
	@cd cmd && go run main.go
buf: 
	@if exist pb\*.go del /Q pb\*.go
	buf generate
