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

	"go-api/httpd/routes"
	"go-api/internal/db"

	"github.com/stretchr/testify/assert"
)

const dbName = "userdb"
const collectionName = "people"

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

func TestUserGet(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))

	assert.Equal(t, 200, w.Code)
}

func TestUsers(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))

	assert.Equal(t, 200, w.Code)
}