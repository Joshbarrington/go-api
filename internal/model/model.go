package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
}

// db.people.save({ username:"Batty", firstname:"Nic", lastname:"Harrison", password:"Rudy"})
// db.people.save({ _id:1, username:"Boy", firstname:"Nic", lastname:"Raboy" , password:"Harrison"})
// db.people.save({ _id:2, username:"Cheese", firstname:"Nic", lastname:"Raboy" , password:"Harrison"})
// db.people.save({ _id:3, username:"hehdh", firstname:"Nic", lastname:"Raboy" , password:"Harrison"})
