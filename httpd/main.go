package main

import (
	"go-api/httpd/routes"
	"go-api/internal/db"
	"log"
)

const dbName = "userdb"
const collectionName = "people"

func main() {
	log.Println(dbName + " " + collectionName)
	collection, _ := db.GetMongoDbCollection(dbName, collectionName)

	r := routes.SetupRouter(collection)

	r.Run()
}
