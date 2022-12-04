package main

import (
	"go-interview/api/apis/collections"
	_ "go-interview/api/config"
	"go-interview/api/database"
	"go-interview/api/router"
)

func main() {
	database.SelectCollection()
	collections.UserText_r_all()
	router := router.InitRouter()
	router.Run(":5555")
}
