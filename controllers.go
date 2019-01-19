package main

import (
	"github.com/keygen-test/repository"
	"time"
	"fmt"
	"log"
)

func (a *App) GetKeyController() (string, error) {
	k, err := repository.GetRandomFreeKey(a.DB)

	tNow := time.Now()
	k.SentAt = &tNow
	err = repository.UpdateKey(a.DB, k)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("service not available")
	}

	return k.Key, nil
}

func (a *App) SubmitKeyController(key string) error {
	k, err := repository.GetKey(a.DB, key)
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
		return repository.UpdateKey(a.DB, k)
	}
}

func (a *App) CheckKeyController(key string) (string, error) {
	k, err := repository.GetKey(a.DB, key)
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

func (a *App) GetInfoController() (int, error) {
	return repository.FreeKeyCount(a.DB)
}