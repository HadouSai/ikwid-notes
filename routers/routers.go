package routers

import (
	"encoding/json"
	"net/http"

	"github.com/HadouSai/ikwid-notes/db"
	"github.com/HadouSai/ikwid-notes/models"
)

/* Register route create a register in the database */
func Register(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(rw, "Error malformed data "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(rw, "Email not provided", 404)
		return
	}

	if len(user.Password) < 6 {
		http.Error(rw, "Password need to be at least 6 characters", 404)
		return
	}

	if _, emailExist, _ := db.CheckUserExist(user.Email); emailExist == true {
		http.Error(rw, "Email already exist", 404)
		return
	}

	if _, status, err := db.InsertUser(user); err != nil {
		http.Error(rw, "An error occurred inserting user "+err.Error(), 404)
		return
	} else if status == false {
		http.Error(rw, "An error occurred inserting user, failed in status "+err.Error(), 404)
		return
	}

	rw.WriteHeader(http.StatusCreated)

	// body solo se puede leer una vez luego se destruye
}
