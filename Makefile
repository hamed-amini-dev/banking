dependencies:
	go mod download


run:
	@clear
	@echo 'Running Server...'
	@go run main.go s


test:
	@go test -tags testing ./...	


build-linux:
	@GOOS=linux
	@go build -o account.bin main.go


build-windows:
	@GOOS=windows
	@go build -o account.exe main.go 

build-mac:
	@GOOS=darwin
	@go build -o account main.go     	