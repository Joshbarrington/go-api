package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"user": "test",
		})
	}
}
