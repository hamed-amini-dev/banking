dependencies:
	go mod download


run:
	@clear
	@echo 'Running Server...'
	@go run main.go s


test:
	go test -tags testing ./...	