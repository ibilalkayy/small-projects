package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

type EncryptedData struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        primitive.ObjectID `bson:"user_id"`
	Plaintext     string             `bson:"plaintext"`
	Ciphertext    string             `bson:"ciphertext"`
	EncryptionKey string             `bson:"encryption_key"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}

type DecryptionLog struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          primitive.ObjectID `bson:"user_id"`
	EncryptedDataID primitive.ObjectID `bson:"encrypted_data_id"`
	DecryptedText   string             `bson:"decrypted_text"`
	CreatedAt       time.Time          `bson:"created_at"`
}
