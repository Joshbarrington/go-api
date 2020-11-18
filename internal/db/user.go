package db

import (
	"context"
<<<<<<< HEAD
=======
	"fmt"
	"log"
>>>>>>> e5865f8... Get method added and code tidied up.

	"go-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

<<<<<<< HEAD
func AddUser(collection *mongo.Collection, user model.User) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), user)
	return res, err
=======
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
>>>>>>> e5865f8... Get method added and code tidied up.
}
