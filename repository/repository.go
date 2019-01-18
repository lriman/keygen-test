package repository

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"github.com/keygen-test/models"
)

func RandString(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetNextKey(db *gorm.DB) (*models.SecretKey, error) {
	k := new(models.SecretKey)
	err := db.Where("sent_at is NULL").First(&k).Error
	return k, err
}

func GenerateNewKey(db *gorm.DB) (*models.SecretKey, error) {
	k := new(models.SecretKey)
	k.Key = RandString(4)
	err := db.Save(k).Error
	return k, err
}

func UpdateKey(db *gorm.DB, k *models.SecretKey) error {
	return db.Save(k).Error
}

func GetKey(db *gorm.DB, key string) (*models.SecretKey, error) {
	k := new(models.SecretKey)
	err := db.Where("Key = ?", key).First(k).Error
	return k, err
}

func FreeKeyCount(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(models.SecretKey{}).Where("sent_at is NULL").Count(&count).Error
	return count, err
}
