package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/keygen-test/controllers"
	"github.com/keygen-test/models"
	"strconv"
)

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for new key")

	resp := models.Response{}
	result, err := controllers.GetKeyController()

	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Result = result
	}

	log.Println("Reponse for new key:", resp)
	json.NewEncoder(w).Encode(resp)
}

func SubmitKeyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for submit key")

	req := models.Request{}
	resp := models.Response{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		log.Println("Reponse for submit key:", resp)
		return
	}

	err = controllers.SubmitKeyController(req.Key)

	if err != nil {
		resp.Error = err.Error()
	}

	log.Println("Reponse for submit key:", resp)
	json.NewEncoder(w).Encode(resp)
}

func CheckKeyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for check key")

	req := models.Request{}
	resp := models.Response{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		log.Println("Reponse for check key:", resp)
		return
	}

	result, err := controllers.CheckKeyController(req.Key)

	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Result = result
	}

	log.Println("Reponse for check key:", resp)
	json.NewEncoder(w).Encode(resp)
}

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for key pool info")

	resp := models.Response{}
	result, err := controllers.GetInfoController()

	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Result = strconv.Itoa(result)
	}

	log.Println("Reponse for key pool info:", resp)
	json.NewEncoder(w).Encode(resp)
}