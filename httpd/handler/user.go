package handler

import (
	"log"
	"strconv"

	"go-api/internal/db"
	"go-api/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserPost ... User Post
func UserPost(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user model.User
		log.Print(user.FirstName)
		if err := c.Bind(&user); err != nil {
			log.Print(err)
			c.JSON(400, gin.H{"error": gin.H{"code": 400, "message": "missing required field(s)"}})
			return
		}

		result, err := db.AddUser(collection, user)

		if err != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": 500, "message": "user unsuccesfully added"}})
			return
		}

		id := result.InsertedID.(primitive.ObjectID)
		c.JSON(200, gin.H{"id": id})
	}
}

// GetUser ...
// To bind path variables use c.bind coupled with the id of the path variable.
func GetUser(m *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Print(err)
			c.JSON(400, gin.H{"error": gin.H{"code": 400, "message": "bad request path field(s)"}})		
			return
		}

		result, err := db.GetUser(m, id)

		if err != nil {
			c.JSON(200, gin.H{"user": result})
			return
		}

		c.JSON(200, result)
	}
}

// GetUser ...
// To bind path variables use c.bind coupled with the id of the path variable.
func Users(m *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		result, err := db.AllUsers(m)

		if err != nil {
			c.Bind(err)
		}
		c.JSON(200, result)
	}
}
