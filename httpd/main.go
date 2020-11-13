package main

import (
	"github.com/gin-gonic/gin"
	"go-api/httpd/handler"
	"go-api/internal/db"
)

const dbName = "userdb"
const collectionName = "users"

func main() {

	collection, _ := db.GetMongoDbCollection(dbName, collectionName)

	r := gin.Default()

	r.POST("/user", handler.UserPost(collection))

	r.Run()
}
