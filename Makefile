PROJECT = $(shell pwd)
GO = $(shell which go)

setup:
	docker pull rudy1021/mongo:4.4-test-env
	docker run --name mongo4 -v $(PROJECT)/dbEngine:/dbEngine/db -d -p 27017:27017 --rm rudy1021/mongo:4.4-test-env
	go mod tidy
	go get -u github.com/spf13/cobra@latest
	go get -d github.com/spf13/viper
	go get -d go.mongodb.org/mongo-driver/mongo

start:
	go run main.go

ngrok:
	ngrok http 5555