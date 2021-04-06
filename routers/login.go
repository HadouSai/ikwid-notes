package routers

import (
	"encoding/json"
	"net/http"

	"github.com/HadouSai/ikwid-notes/db"
	"github.com/HadouSai/ikwid-notes/jwt"
	"github.com/HadouSai/ikwid-notes/models"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(rw, "User or password incorrect"+err.Error(), http.StatusUnauthorized)
	}

	if len(user.Email) == 0 {
		http.Error(rw, "Email is required", http.StatusUnauthorized)
		return
	}

	user, userExist := db.SignIn(user.Email, user.Password)
	if !userExist {
		http.Error(rw, "User or password incorrect", http.StatusUnauthorized)
		return
	}

	jwtUser, err := jwt.GenerateJWT(&user)
	if err != nil {
		http.Error(rw, "An error occurred in token generation"+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.LoginResponse{
		Token: jwtUser,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

}
