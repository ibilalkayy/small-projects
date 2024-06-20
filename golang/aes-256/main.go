package main

import (
	"context"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/ibilalkayy/small-projects/aes-256/database"
	"github.com/ibilalkayy/small-projects/aes-256/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	database.Connect()

	// Sample user creation
	username := "example_user"
	email := "user@example.com"
	password := "password123"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	userCollection := database.GetCollection("users")
	res, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	userID := res.InsertedID.(primitive.ObjectID)

	// Encryption
	key := "thisis32bitlongpassphraseimusing"
	plaintext := "This is a secret"
	ciphertext := EncryptAES([]byte(key), plaintext)

	// Store encrypted data
	encryptedData := models.EncryptedData{
		UserID:        userID,
		Plaintext:     plaintext,
		Ciphertext:    ciphertext,
		EncryptionKey: key,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	encryptedDataCollection := database.GetCollection("encrypted_data")
	res, err = encryptedDataCollection.InsertOne(context.Background(), encryptedData)
	if err != nil {
		log.Fatalf("Error storing encrypted data: %v", err)
	}

	encryptedDataID := res.InsertedID.(primitive.ObjectID)

	// Decryption
	DecryptAES([]byte(key), ciphertext, userID, encryptedDataID)
}

func EncryptAES(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string, userID, encryptedDataID primitive.ObjectID) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	decryptedText := string(pt[:])
	fmt.Println("DECRYPTED:", decryptedText)

	// Store decryption log
	decryptionLog := models.DecryptionLog{
		UserID:          userID,
		EncryptedDataID: encryptedDataID,
		DecryptedText:   decryptedText,
		CreatedAt:       time.Now(),
	}

	decryptionLogCollection := database.GetCollection("decryption_logs")
	_, err = decryptionLogCollection.InsertOne(context.Background(), decryptionLog)
	if err != nil {
		log.Fatalf("Error storing decryption log: %v", err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
