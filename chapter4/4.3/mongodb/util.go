package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mgoCli *mongo.Client

func initDb() {
	var err error
	//credential := options.Credential{
	//	AuthMechanism: "PLAIN",
	//	Username:      "root",
	//	Password:      "zah56641189",
	//}
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

func MgoCli() *mongo.Client {
	if mgoCli == nil {
		initDb()
	}
	return mgoCli
}
