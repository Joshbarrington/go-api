package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/httpd/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(collection *mongo.Collection) *gin.Engine {
	r := gin.Default()
	r.POST("/user", handler.UserPost(collection))
	return r
}
