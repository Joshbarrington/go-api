package main

import (
	"context"
	"log"

	"go-api/httpd/routes"
	"go-api/internal/db"
)

const dbName = "userdb"
const collectionName = "users"

func main() {

	client, _ := db.MongoDbConnection()
	collection, _ := db.GetMongoDbCollection(dbName, collectionName, client)

	log.Printf("[MongoDb] Database: %s", dbName)
	log.Printf("[MongoDb] Collection: %s", collectionName)

	r := routes.SetupRouter(collection)
	r.Run()

	err := client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[MongoDb] Connection closed")
}
