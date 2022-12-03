package database

import (
	"context"
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

func SelectCollection() *mongo.Collection {
	var (
		client     = GetMgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)

	db = client.Database(config.Getstr("mongo.database"))
	collection = db.Collection(config.Getstr("mongo.collection"))
	collection = collection
	return collection
}
