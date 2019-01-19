package models

import (
	"time"
)

type SecretKey struct {
	Key        string     `gorm:"type:varchar(4);primary_key"`
	SentAt     *time.Time `gorm:"index"`
	SubmitAt   *time.Time
}

func (SecretKey) TableName() string {
	return "secret_keys"
}