package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" binding:"required" bson:"username"`
	FirstName string             `json:"firstname" binding:"required" bson:"firstname"`
	LastName  string             `json:"lastname" binding:"required" bson:"lastname"`
	Password  string             `json:"password" binding:"required" bson:"password"`
}
