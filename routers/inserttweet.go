package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HadouSai/ikwid-notes/db"
	"github.com/HadouSai/ikwid-notes/models"
)

func InsertTweet(rw http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(rw, "Tweet incorrect message"+err.Error(), http.StatusBadRequest)
	}

	register := models.InsertTweet{
		UserId:  IdUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	if _, status, err := db.InsertTweet(register); err != nil {
		http.Error(rw, "An error occurred while inserting the tweet"+err.Error(), http.StatusBadRequest)
	} else if !status {
		http.Error(rw, "An error occurred while	inserting the tweet, status failed"+err.Error(), http.StatusBadRequest)
	}

	rw.WriteHeader(http.StatusCreated)
}
