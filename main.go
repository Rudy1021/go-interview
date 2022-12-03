package main

import (
	"go-interview/api/apis/collections"
	"go-interview/api/config"
	"go-interview/api/database"
)

func main() {
	config.Init()
	database.SelectCollection()
	collections.UserText_r_all()
}
