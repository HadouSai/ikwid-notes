package db

import (
	"github.com/HadouSai/ikwid-notes/models"
	"golang.org/x/crypto/bcrypt"
)

/* SignIn function to authenticate users */
func SignIn(email string, password string) (models.User, bool) {
	userModel, userExist, _ := CheckUserExist(email)

	if userExist == false {
		return userModel, false
	}

	passwordBytes := []byte(password)
	passwordDb := []byte(userModel.Password)

	if err := bcrypt.CompareHashAndPassword(passwordDb, passwordBytes); err != nil {
		return userModel, false
	}

	return userModel, true
}
