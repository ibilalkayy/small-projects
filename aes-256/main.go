package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ibilalkayy/small-projects/aes-256/database"
	"github.com/ibilalkayy/small-projects/aes-256/models"
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

	user := models.User{Username: username, Email: email, PasswordHash: string(hashedPassword)}
	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Fatalf("Error creating user: %v", result.Error)
	}

	// Encryption
	key := "thisis32bitlongpassphraseimusing"
	plaintext := "This is a secret"
	ciphertext := EncryptAES([]byte(key), plaintext)

	// Store encrypted data
	encryptedData := models.EncryptedData{UserID: user.ID, Plaintext: plaintext, Ciphertext: ciphertext, EncryptionKey: key}
	database.DB.Create(&encryptedData)

	// Decryption
	DecryptAES([]byte(key), ciphertext, user.ID)
}

func EncryptAES(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string, userID uint) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	decryptedText := string(pt[:])
	fmt.Println("DECRYPTED:", decryptedText)

	// Store decryption log
	decryptionLog := models.DecryptionData{UserID: userID, EncryptedDataID: 1, DecryptedText: decryptedText}
	database.DB.Create(&decryptionLog)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
