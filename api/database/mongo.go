package database

import (
	"context"
	"fmt"
	"go-interview/api/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

func RunDatabase() {
	var (
		client     = GetMgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)
	//2.选择数据库 my_db
	db = client.Database(config.Getstr("mongo.database"))

	//选择表 my_collection
	collection = db.Collection(config.Getstr("mongo.collection"))
	collection = collection
	fmt.Println(collection)
}
