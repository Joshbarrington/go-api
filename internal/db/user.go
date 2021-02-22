package db

import (
	"context"
	"log"

	"go-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetAll(collection *mongo.Collection) ([]bson.M, error) {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var users []bson.M

	if err = cursor.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}

	return users, err
}

func GetUserByName(collection *mongo.Collection, user_name string) (bson.M, error) {

	var user bson.M

	err := collection.FindOne(context.Background(), bson.D{primitive.E{Key: "username", Value: user_name}}).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, err
}
