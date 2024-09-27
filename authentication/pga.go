package authentication

import (
	"homburg/backend/models"

	"net/http"
	// "time"
)

//
// VerifyPGA()
//  1. check if the headers are empty
//  2. check if the headers are valid
//
// Public-Key header should have a valid/registered public key.
// Signed-Message header should have the message "YY-MM-DD HH" signed by the
//  private key.
//

func VerifyPGA(publicKey, signedMessage string) (models.User, int, string) {
	var user models.User

	//currentTime := time.Now().UTC()

	// 1. check if the headers are empty
	if publicKey == "" || signedMessage == "" {
		return user, http.StatusBadRequest, "invalid headers"
	}

	// 2. check if the headers are valid

	return user, http.StatusOK, ""
}
