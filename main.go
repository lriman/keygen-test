package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/keygen-test/db"
	"github.com/keygen-test/handlers"
)

func main() {
	db.InitDatabaseConnection("127.0.0.1", "testuser", "mypass", "mydb")

	router := mux.NewRouter()

	router.HandleFunc("/get", handlers.GetKeyHandler).Methods("GET")
	router.HandleFunc("/submit", handlers.SubmitKeyHandler).Methods("POST")
	router.HandleFunc("/check", handlers.CheckKeyHandler).Methods("POST")
	router.HandleFunc("/info", handlers.GetInfoHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", router))
}

