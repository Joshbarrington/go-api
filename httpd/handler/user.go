package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserPost(m *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		c.Bind(&user)
		fmt.Println(user)
	}
}
