package main

import (
	"go-api/httpd/handler"
	"go-api/internal/db"

	"github.com/gin-gonic/gin"
)

const dbName = "userdb"
const collectionName = "users"

func main() {

	collection, _ := db.GetMongoDbCollection(dbName, collectionName)

	r := gin.Default()

	r.POST("/user", handler.UserPost(collection))
	r.GET("/user/:id", handler.GetUser(collection))

	r.Run()
}
