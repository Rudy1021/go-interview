package main

import (
	_ "go-interview/api/config"
	"go-interview/api/database"
	"go-interview/api/router"
)

func main() {
	database.SelectCollection()
	router := router.InitRouter()
	router.Run(":5555")
}
