package cryptography

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyHash(encrypted string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain))
	if err != nil {
		return false
	}

	return true
}
