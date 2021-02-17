package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-api/internal/db"
	"go-api/internal/model"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbName = "userdb"
const collectionName = "users"

var client, _ = db.MongoDbConnection()
var collection, _ = db.GetMongoDbCollection(dbName, collectionName, client)
var router = SetupRouter(collection)

type PostResponseJSON struct {
	ID string `json:"id"`
}

func testHTTPRequest(collection *mongo.Collection, json []byte, requestType string, endPoint string) (*httptest.ResponseRecorder, *http.Response) {
	writer := httptest.NewRecorder()

	req, _ := http.NewRequest(requestType, endPoint, bytes.NewBuffer(json))
	req.Header.Set("Content-type", "application/json")

	router.ServeHTTP(writer, req)
	resp := writer.Result()

	return writer, resp
}

func TestUserPost(t *testing.T) {

	var respJSON PostResponseJSON
	var user model.User

	testData := map[string]string{
		"username":  "test",
		"firstname": "test",
		"lastname":  "test",
		"password":  "test",
	}
	testJSON, err := json.Marshal(testData)

	if err != nil {
		log.Println(err)
	}

	writer, resp := testHTTPRequest(collection, testJSON, "POST", "/user")
	err = json.NewDecoder(resp.Body).Decode(&respJSON)

	if err != nil {
		log.Println(err)
	}

	log.Printf("[Response] _id: %s", respJSON.ID)
	oid, _ := primitive.ObjectIDFromHex(respJSON.ID)

	err = collection.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: oid}}).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, writer.Code)
	assert.NotEmpty(t, respJSON.ID)
	assert.True(t, db.CheckPasswordHash("test", user.Password))
}
