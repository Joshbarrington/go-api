package main

import (
	"log"

	"go-api/httpd/routes"
	"go-api/internal/db"
)

const dbName = "userdb"
const collectionName = "users"

func main() {

	collection, _ := db.GetMongoDbCollection(dbName, collectionName)

	log.Printf("[MongoDb] Database: %s", dbName)
	log.Printf("[MongoDb] Collection: %s", collectionName)

	r := routes.SetupRouter(collection)
	r.Run()
}
