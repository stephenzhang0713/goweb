package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goweb/chapter4/4.3/model"
	"goweb/chapter4/4.3/mongodb"
	"time"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		err        error
		collection *mongo.Collection
		iResult    *mongo.InsertOneResult
		id         primitive.ObjectID
	)

	collection = client.Database("my_db").Collection("my_collection")

	logRecord := model.LogRecord{
		JobName: "job1",
		Command: "echo 1",
		Err:     "",
		Content: "1",
		Tp: model.ExecTime{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10,
		},
	}

	if iResult, err = collection.InsertOne(context.TODO(), logRecord); err != nil {
		fmt.Print(err)
		return
	}

	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())

}
