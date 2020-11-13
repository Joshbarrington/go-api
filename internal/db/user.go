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
func GetUser(collection *mongo.Collection, id int) (bson.M, error) {
	var result bson.M

	err := collection.FindOne(context.Background(),
		bson.D{{}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("found document %v", result)
			return result, err
		}
		log.Fatal(err)
	}

	return result, nil
}

// AllUsers ...
func AllUsers(collection *mongo.Collection) ([]bson.M, error) {

	var users []bson.M

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No users found: %v", users)
			return users, nil
		}
		log.Fatal(err)
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}

	return users, nil
}
