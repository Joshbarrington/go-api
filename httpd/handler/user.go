package handler

import (
	"fmt"
	"log"
	"strconv"

	"go-api/internal/db"
	"go-api/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserPost ... User Post
func UserPost(m *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		c.Bind(&user)
		fmt.Println(user)
	}
}

// GetUser ...
// To bind path variables use c.bind coupled with the id of the path variable.
func GetUser(m *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id)
		db.GetUser(m, id)
	}
}
