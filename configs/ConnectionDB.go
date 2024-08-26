package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionDatabase() *mongo.Database {
	var createdCleint = options.Client().ApplyURI("mongodb://localhost:27017/")
	var client, err = mongo.Connect(context.Background(), createdCleint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conneted database")
	Db := client.Database("newdatabase")
	return Db
}
