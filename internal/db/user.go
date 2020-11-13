package db

import (
	"context"
	"fmt"

	"go-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(collection *mongo.Collection, user model.User) {
	res, err := collection.InsertOne(context.Background(), user)
}
