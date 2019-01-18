package controllers

import (
	"github.com/keygen-test/db"
	"github.com/keygen-test/models"
	"time"
	"fmt"
)

func GetKeyController() (string, error) {

	var k *models.SecretKey
	var err error
	k, err = db.DB.GetNextKey()

	if err != nil {
		k, err = db.DB.GenerateNewKey()
		if err != nil {
			return "", err
		}
	}

	tNow := time.Now()
	k.SentAt = &tNow
	err = db.DB.UpdateKey(k)
	if err != nil {
		return "", err
	}

	return k.Key, nil
}

func SubmitKeyController(key string) error {

	k, err := db.DB.GetKey(key)
	if err != nil{
		return err
	}

	if k.SentAt == nil {
		return fmt.Errorf("key was not sent")
	} else if k.SubmitAt != nil {
		return fmt.Errorf("key already used")
	} else {
		tNow := time.Now()
		k.SubmitAt = &tNow
		return db.DB.UpdateKey(k)
	}
}

func CheckKeyController(key string) (string, error) {
	k, err := db.DB.GetKey(key)
	if err != nil{
		return "", err
	}

	if k.SentAt == nil {
		return "NOT SENT", nil
	} else if k.SubmitAt == nil {
		return "NOT USED", nil
	} else {
		return "USED", nil
	}
}

func GetInfoController() (int, error) {
	return db.DB.FreeKeyCount()
}