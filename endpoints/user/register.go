package endpoints

import (
	"github.com/gin-gonic/gin"

	"homburg/backend/initializers"
	"homburg/backend/models"

	"net/http"
)

//
// UserRegister (POST :registerRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the user is already registered
//  3. create a new user record
//

type user_registerRequest struct {
	PublicKey string `json:"public-key"`
}

func UserRegister(c *gin.Context) {
	var body user_registerRequest
	var user models.User

	// WARNING: NOT IMPLEMENTED YET
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
	return

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if body.PublicKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the user is already registered
	result := initializers.DB.Model(&models.User{}).Where("public-key = ?", body.PublicKey).First(&user)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is already registered"})
		return
	}

	// 3. create a new user record
	user = models.User{
		PublicKey: body.PublicKey,
	}
	result = initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed creating the record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
