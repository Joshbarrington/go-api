package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbConnection() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(),
		options.Client().SetDirect(true).ApplyURI("mongodb://db:27017"))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("[MongoDb] Connected to MongoDB!")

	return client, nil
}

func GetMongoDbCollection(DbName string, CollectionName string, MongoClient *mongo.Client) (*mongo.Collection, error) {

	collection := MongoClient.Database(DbName).Collection(CollectionName)

	return collection, nil
}
