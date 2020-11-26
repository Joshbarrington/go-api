package routes

import (
	"bytes"
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-api/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbName = "userdb"
const collectionName = "users"

var collection, _ = db.GetMongoDbCollection(dbName, collectionName)
var router = SetupRouter(collection)

type PostResponseJSON struct {
	ID string `json:"id"`
}

func testHttpRequest(collection *mongo.Collection, json []byte) (*httptest.ResponseRecorder, *http.Response) {
	writer := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(json))
	req.Header.Set("Content-type", "application/json")

	router.ServeHTTP(writer, req)
	resp := writer.Result()

	return writer, resp
}

func TestUserPost(t *testing.T) {

	var respJSON PostResponseJSON

	testJSON, err := json.Marshal(map[string]string{
		"username":  "test",
		"firstname": "test",
		"lastname":  "test",
		"password":  "test",
	})

	if err != nil {
		log.Println(err)
	}

	writer, resp := testHttpRequest(collection, testJSON)
	err = json.NewDecoder(resp.Body).Decode(&respJSON)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("[Response] status code: %d", writer.Code)
	log.Printf("[Response] contents: %s", string(body))
	log.Printf("[Response] respJSON struct contents: %s", respJSON)

	assert.Equal(t, 200, writer.Code)
	assert.NotEmpty(t, respJSON.ID)
}
