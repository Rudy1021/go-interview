package database

import (
	"fmt"

	"github.com/jinzhu/gorm" //使用 gorm 作為 orm 工具
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB
var err error

func init() {
	username := "pmsadmin"
	password := "pms5566"
	dbName := "pms"
	dbHost := "database-1.c1doybqwvpca.ap-northeast-1.rds.amazonaws.com"
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	Db, err = gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	Db.SingularTable(true)

}
