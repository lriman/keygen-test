package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/keygen-test/models"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(host, user, pwd, db string) {
	var err error

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pwd, db)
	a.DB, err = gorm.Open("postgres", conn)

	if err != nil {
		log.Fatal("Could not connect to Database:", err)
	}

	a.DB.AutoMigrate(models.SecretKey{})
}

func (a *App) Run(addr string) {
	router := mux.NewRouter()

	router.HandleFunc("/get", a.GetKeyHandler).Methods("GET")
	router.HandleFunc("/submit", a.SubmitKeyHandler).Methods("POST")
	router.HandleFunc("/check", a.CheckKeyHandler).Methods("POST")
	router.HandleFunc("/info", a.GetInfoHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(addr, router))
}
