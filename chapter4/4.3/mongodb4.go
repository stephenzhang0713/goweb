package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goweb/chapter4/4.3/model"
	"goweb/chapter4/4.3/mongodb"
	"log"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)

	collection = client.Database("my_db").Collection("test")
	cond := model.FindByJobName{JobName: "job multi1"}

	if cursor, err = collection.Find(context.TODO(), cond, options.Find().SetSkip(0), options.Find().SetLimit(2)); err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	for cursor.Next(context.TODO()) {
		var lr model.LogRecord
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		fmt.Println(lr)
	}

	var results []model.LogRecord
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
