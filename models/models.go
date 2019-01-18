package models

import (
	"time"
)

type SecretKey struct {
	Key        string     `gorm:"primary_key"`
	SentAt     *time.Time
	SubmitAt   *time.Time
}
