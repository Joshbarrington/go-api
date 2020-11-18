package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-api/httpd/routes"
	"go-api/internal/db"
)

const dbName = "userdb"
const collectionName = "users"

var collection, _ = db.GetMongoDbCollection(dbName, collectionName)
var router = routes.SetupRouter(collection)

func TestUserPost(t *testing.T) {

	testJson, err := json.Marshal(map[string]string{
		"username":  "test",
		"firstname": "test",
		"lastname":  "test",
		"password":  "test",
	})

	if err != nil {
		log.Println(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(testJson))
	req.Header.Set("Content-type", "application/json")

	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	assert.Equal(t, 200, w.Code)
}
