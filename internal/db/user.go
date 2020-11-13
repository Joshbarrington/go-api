package db

import (
	"context"

	"go-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(collection *mongo.Collection, user model.User) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), user)
	return res, err
}
