package test

import (
	"github.com/gin-gonic/gin"

	"homburg/backend/utils"

	"net/http"
)

func GetCountry(c *gin.Context) {
	ipaddr := c.Param("ipaddr")

	country := utils.GetCountry(ipaddr)

	c.JSON(http.StatusOK, gin.H{
		"ip-address": ipaddr,
		"country":    country,
	})
}
