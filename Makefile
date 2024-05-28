build:
	@go build  -o bin/main /home/khero/proxies/cmd/main.go
run: build
	@./bin/main
