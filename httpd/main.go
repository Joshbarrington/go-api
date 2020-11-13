package main

import (
	"fmt"
	"go-api/httpd/handler"
	"go-api/internal/db"

	"github.com/gin-gonic/gin"
)

const dbName = "userdb"
const collectionName = "people"

func main() {
	fmt.Println(dbName + " " + collectionName)
	collection, _ := db.GetMongoDbCollection(dbName, collectionName)

	r := gin.Default()

	r.POST("/user", handler.UserPost(collection))
	r.GET("/user/:id", handler.GetUser(collection))
	r.GET("/users", handler.Users(collection))

	r.Run()
}
