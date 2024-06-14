package models

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type EncryptedData struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint `gorm:"not null"`
	Plaintext     string
	Ciphertext    string `gorm:"not null"`
	EncryptionKey string `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DecryptionData struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uint `gorm:"not null"`
	EncryptedDataID uint `gorm:"not null"`
	DecryptedText   string
	CreatedAt       time.Time
}
