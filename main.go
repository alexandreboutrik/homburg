package main

import (
	"github.com/gin-gonic/gin"

	"homburg/backend/endpoints/test"
	"homburg/backend/endpoints/user"
	"homburg/backend/initializers"
	"homburg/backend/utils"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Cross Origin Resource Share
	router.Use(utils.CORSMiddleware())

	// user endpoints
	router.POST("/api/register", endpoints.UserRegister)

	// testing endpoints, for dev purposes only
	if initializers.TESTING_EP != false {
		router.GET("/test/get-country/:ipaddr", test.GetCountry)
		router.GET("/test/get-headers", test.GetHeaders)
	}

	// logging endpoints, required for legal purposes
	if initializers.LOGGING_EP != false {
	}

	router.Run("localhost:8000")
}
