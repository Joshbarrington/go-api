package db

import (
	"context"

	"go-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddUser(collection *mongo.Collection, user model.User) (*mongo.InsertOneResult, error) {
	hash, err := HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hash
	res, err := collection.InsertOne(context.Background(), user)

	return res, err
}
