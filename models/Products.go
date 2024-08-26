package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Products struct {
	ProductName        string    `json:"productName" bson:"productName"`
	ProductDescription string    `json:"productDescription" bson:"productDescription"`
	CreatedAt          time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt" bson:"updatedAt"`
}

func CreateCollectionProducts(db *mongo.Database) {
	collName := "products"

	// Koleksiyon mevcutsa, bu aşamayı geçin
	if collExists(db, collName) {
		fmt.Println("Collection already exists")
		return
	}

	err := db.CreateCollection(context.Background(), collName)
	if err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	} else {
		fmt.Println("Collection created successfully.")
	}

	// Koleksiyon şemasını oluşturma
	CollectionBaseSchema(db)
}

func collExists(db *mongo.Database, collName string) bool {
	collections, err := db.ListCollectionNames(context.Background(), bson.M{"name": collName})
	if err != nil {
		log.Fatalf("Failed to list collections: %v", err)
	}
	return len(collections) > 0
}

func CollectionBaseSchema(db *mongo.Database) {
	collModCommand := bson.D{
		{Key: "collMod", Value: "products"},
		{Key: "validator", Value: bson.D{
			{Key: "$jsonSchema", Value: bson.D{
				{Key: "bsonType", Value: "object"},
				{Key: "required", Value: bson.A{"productName", "productDescription", "createdAt", "updatedAt"}},
				{Key: "properties", Value: bson.D{
					{Key: "productName", Value: bson.D{
						{Key: "bsonType", Value: "string"},
						{Key: "minLength", Value: 3},
						{Key: "description", Value: "Product name must be a string and at least 3 characters long."},
					}},
					{Key: "productDescription", Value: bson.D{
						{Key: "bsonType", Value: "string"},
						{Key: "description", Value: "Product description must be a string."},
					}},
					{Key: "createdAt", Value: bson.D{
						{Key: "bsonType", Value: "date"},
						{Key: "description", Value: "createdAt must be a valid date."},
					}},
					{Key: "updatedAt", Value: bson.D{
						{Key: "bsonType", Value: "date"},
						{Key: "description", Value: "updatedAt must be a valid date."},
					}},
				}},
			}},
		}},
		{Key: "validationLevel", Value: "strict"},
	}

	result := db.RunCommand(context.Background(), collModCommand)
	if err := result.Err(); err != nil {
		log.Fatalf("Failed to modify collection schema: %v", err)
	} else {
		fmt.Println("Collection schema modified successfully.")
	}
}
