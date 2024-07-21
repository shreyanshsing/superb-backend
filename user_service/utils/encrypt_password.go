package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}