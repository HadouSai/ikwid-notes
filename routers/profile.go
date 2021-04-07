package routers

import (
	"encoding/json"
	"net/http"

	"github.com/HadouSai/ikwid-notes/db"
)

// ViewProfile
func ViewProfile(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(rw, "Need to provide ID", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(id)

	if err != nil {
		http.Error(rw, "An error occurred while searching profile "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(profile)

}
