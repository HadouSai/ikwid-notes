package routers

import (
	"encoding/json"
	"net/http"

	"github.com/HadouSai/ikwid-notes/db"
	"github.com/HadouSai/ikwid-notes/models"
)

// EditProfile
func EditProfile(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(rw, "Incorrect data"+err.Error(), http.StatusBadRequest)
	}

	if status, err := db.EditProfile(user, IdUser); err != nil {
		http.Error(rw, "An error occurred while editing profile "+err.Error(), http.StatusBadRequest)
		return
	} else if !status {
		http.Error(rw, "The profile was not updated "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
