package main

import (
	"fmt"
)

func main() {
	defer orm.Db.Close()
	router := router.InitRouter()
	router.Run(":5555")
}
