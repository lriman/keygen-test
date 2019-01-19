package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/keygen-test/models"
)

func GetRandomFreeKey(db *gorm.DB) (*models.SecretKey, error) {
	k := new(models.SecretKey)
	err := db.Where("sent_at is NULL").Order(gorm.Expr("random()")).First(&k).Error
	return k, err
}

func GetLastKey(db *gorm.DB) (*models.SecretKey, error) {
	k := new(models.SecretKey)
	err := db.Last(&k).Error
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
