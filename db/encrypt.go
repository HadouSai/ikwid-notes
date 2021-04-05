package db

import "golang.org/x/crypto/bcrypt"

/* Encrypt an string */
func Encrypt(data string) (string, error) {
	cost := 7

	bytes, err := bcrypt.GenerateFromPassword([]byte(data), cost)

	return string(bytes), err
}

// []slice => vector with n elements
