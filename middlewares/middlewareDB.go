package middlewaredb

import (
	"net/http"

	"github.com/HadouSai/ikwid-notes/db"
)

/* CheckDB is the middleware that check the status in database */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(rw, "Connection lost with database", 500)
			return
		}

		next.ServeHTTP(rw, r)
	}
}
