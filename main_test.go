package main

import (
	"os"
	"testing"
	"net/http"
	"net/http/httptest"
	"log"
	"encoding/json"
	"github.com/keygen-test/models"
	"bytes"
	"strconv"
)

var a App
var testKeys = []string{"1111", "2222", "3333", "4444", "5555", "aaaa", "bbbb"}
var testKey string

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("127.0.0.1", "testuser", "mypass", "mydb_test")
	code := m.Run()
	a.DB.Exec("DROP TABLE secret_keys")
	os.Exit(code)
}


func TestEmptyTable(t *testing.T) {
	req, _ := http.NewRequest("GET", "/get", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	resp := models.Response{}
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Error != "service not available" {
		t.Errorf("Expected error 'service not available'. Got |%s|", resp.Error)
	}
}


func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	log.Println(a)
	log.Println(a.Router)
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}


func TestTable(t *testing.T) {
	for _, key := range testKeys {
		k := models.SecretKey{Key: key}
		err := a.DB.Save(&k).Error
		if err != nil {
			t.Errorf("Expected no error. Got |%s|", err)
		}
	}
}

func TestGetKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "/get", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	resp := models.Response{}
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Error != "" {
		t.Errorf("Expected no errors. Got |%s|", resp.Error)
	}

	if len(resp.Result) != 4 {
		t.Errorf("Expected 4 symbol key. Got |%s|", resp.Result)
	}

	// keep key to reuse
	testKey = resp.Result
}

func TestCheckReceivedKey(t *testing.T) {
	payload := []byte(`{"key":"`+testKey+`"}`)
	req, _ := http.NewRequest("POST", "/check", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	resp := models.Response{}
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Error != "" {
		t.Errorf("Expected no errors. Got |%s|", resp.Error)
	}

	if resp.Result != "NOT USED" {
		t.Errorf("Expected 'NOT USED' state. Got |%s|", resp.Result)
	}
}

func TestSubmitKey(t *testing.T) {
	payload := []byte(`{"key":"`+testKey+`"}`)
	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	resp := models.Response{}
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Error != "" {
		t.Errorf("Expected no errors. Got |%s|", resp.Error)
	}

	if resp.Result != "" {
		t.Errorf("Expected empty result. Got |%s|", resp.Result)
	}
}

func TestDoubleSubmitKey(t *testing.T) {
	payload := []byte(`{"key":"`+testKey+`"}`)
	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	resp := models.Response{}
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Error != "key already used" {
		t.Errorf("Expected error 'key already used'. Got |%s|", resp.Error)
	}
}

func TestCheckUsedKey(t *testing.T) {
	payload := []byte(`{"key":"`+testKey+`"}`)
	req, _ := http.NewRequest("POST", "/check", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	resp := models.Response{}
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Error != "" {
		t.Errorf("Expected no errors. Got |%s|", resp.Error)
	}

	if resp.Result != "USED" {
		t.Errorf("Expected 'USED' state. Got |%s|", resp.Result)
	}
}

func TestFreeKeys(t *testing.T) {
	req, _ := http.NewRequest("GET", "/info", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	resp := models.Response{}
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Error != "" {
		t.Errorf("Expected no errors. Got |%s|", resp.Error)
	}

	leftKeys := strconv.Itoa(len(testKeys) - 1)
	if resp.Result != leftKeys {
		t.Errorf("Expected %s free keys. Got |%s|", leftKeys, resp.Result)
	}
}