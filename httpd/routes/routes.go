package routes

import (
	"go-api/httpd/handler"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(collection *mongo.Collection) *gin.Engine {
	r := gin.Default()
	r.POST("/user", handler.UserPost(collection))
	r.GET("/user/:username", handler.UserGetByName(collection))
	r.GET("/users", handler.UserGetAll(collection))

	return r
}
