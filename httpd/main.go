package main

import (
	"go-api/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user", handler.UserGet())

	r.Run()
}
