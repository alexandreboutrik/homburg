package test

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func GetHeaders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Public-Key":     c.GetHeader("Public-Key"),
		"Signed-Message": c.GetHeader("Signed-Message"),
		"User-Agent":     c.GetHeader("User-Agent"),
		"Authorization":  c.GetHeader("Authorization"),
	})
}
