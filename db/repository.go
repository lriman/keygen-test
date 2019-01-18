package db

import (
	"github.com/keygen-test/models"
	"math/rand"
)

func randString(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (db *DatabaseConnection) GetNextKey() (*models.SecretKey, error) {
	k := new(models.SecretKey)

	err := db.conn.Where("sent_at is NULL").First(&k).Error

	if err != nil {
		db.checkConnection(err)
		return nil, err
	}

	return k, err
}

func (db *DatabaseConnection) GenerateNewKey() (*models.SecretKey, error) {
	k := new(models.SecretKey)
	k.Key = randString(4)

	err := db.conn.Save(k).Error
	if err != nil {
		db.checkConnection(err)
	}

	return k, err
}

func (db *DatabaseConnection) UpdateKey(k *models.SecretKey) error {
	err := db.conn.Save(k).Error
	if err != nil {
		db.checkConnection(err)
	}
	return err
}

func (db *DatabaseConnection) GetKey(key string) (*models.SecretKey, error) {
	k := new(models.SecretKey)

	err := db.conn.Where("Key = ?", key).First(&k).Error
	if err != nil {
		db.checkConnection(err)
	}

	return k, err
}

func (db *DatabaseConnection) FreeKeyCount() (int, error) {
	var count int
	err := db.conn.Model(models.SecretKey{}).Where("sent_at is NULL").Count(&count).Error
	if err != nil {
		db.checkConnection(err)
	}
	return count, err
}
