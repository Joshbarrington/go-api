package handler

import (
	"log"

	"go-api/internal/db"
	"go-api/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserGetAll(collection *mongo.Collection) gin.HandlerFunc {

	return func(c *gin.Context) {

		result, err := db.GetAll(collection)

		if err != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": 500, "message": "unable to retrieve user(s)"}})
		}

		c.JSON(200, gin.H{"result": result})
	}
}

func UserGetByName(collection *mongo.Collection) gin.HandlerFunc {

	return func(c *gin.Context) {

		user_name := c.Param("username")

		result, err := db.GetUserByName(collection, user_name)

		if err != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": 400, "message": "invalid username(s)"}})
			return
		}

		c.JSON(200, gin.H{"result": result})
	}
}

func UserPost(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user model.User

		if err := c.Bind(&user); err != nil {
			log.Print(err)
			c.JSON(400, gin.H{"error": gin.H{"code": 400, "message": "invalid field(s)"}})
			return
		}

		result, err := db.AddUser(collection, user)

		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{"error": gin.H{"code": 500, "message": "user unsuccesfully added"}})
			return
		}

		id := result.InsertedID.(primitive.ObjectID)
		c.JSON(200, gin.H{"id": id})

	}
}
