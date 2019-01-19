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

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/get", a.GetKeyHandler).Methods("GET")
	a.Router.HandleFunc("/submit", a.SubmitKeyHandler).Methods("POST")
	a.Router.HandleFunc("/check", a.CheckKeyHandler).Methods("POST")
	a.Router.HandleFunc("/info", a.GetInfoHandler).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
