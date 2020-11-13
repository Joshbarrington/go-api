package db

import (
	"context"
	"fmt"
	"log"

	"go-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddUser ...
func AddUser(collection *mongo.Collection, user model.User) {
	res, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

// GetUser ...
func GetUser(collection *mongo.Collection, id int) {
	filter := bson.D{{Key: "ID", Value: id}}

	var result model.User

	fmt.Println(filter)

	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
