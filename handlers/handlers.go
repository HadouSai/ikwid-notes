package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	middlewaredb "github.com/HadouSai/ikwid-notes/middlewares"
	"github.com/HadouSai/ikwid-notes/routers"
)

// Handlers :  Sett the port and ListenAndServe with mux
func Handlers() {
	router := mux.NewRouter() // capture http and hangle w/r, if has information to send a response

	router.HandleFunc("/signup", middlewaredb.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/signin", middlewaredb.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // cors take the control

	//log.Fatal(http.ListenAndServe(":"+PORT, handler))

	if err := http.ListenAndServe(":"+PORT, handler); err != nil {
		log.New(os.Stdout, "[Error]", 0).Fatalf("Error: No se pudo iniciar el servidor: 127.0.0.1:%s. \n Message: %v", PORT, err)
	}

}
