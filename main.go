package main

import (
	"log"

	"github.com/HadouSai/ikwid-notes/db"
	"github.com/HadouSai/ikwid-notes/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Coudln't connect to database")
		return
	}

	handlers.Handlers()

}
