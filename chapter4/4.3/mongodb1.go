package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goweb/chapter4/4.3/mongodb"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)

	db = client.Database("my_db")

	collection = db.Collection("my_collection")
	collection = collection
}
