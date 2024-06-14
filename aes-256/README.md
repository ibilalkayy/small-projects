To integrate MongoDB with a Go project and use it for storing AES encrypted data, we can use the MongoDB Go Driver instead of GORM, as GORM does not natively support MongoDB. Below is the updated implementation.

### Step 1: Initialize the Project

First, ensure you have Go and MongoDB installed. Then, create a new Go project and initialize it:

```sh
mkdir aes_encryption_project
cd aes_encryption_project
go mod init aes_encryption_project
```

### Step 2: Add Dependencies

Add the necessary dependencies to your `go.mod` file:

```sh
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
```

### Step 3: Create a `.env` File

Create a `.env` file to store your MongoDB connection information:

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=aes_encryption_db
```

### Step 4: Create Database Models

Create a `models` directory and a file `models/models.go` to define your database models:

```go
package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
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
    CreatedAt    time.Time           `bson:"created_at"`
    UpdatedAt    time.Time           `bson:"updated_at"`
}

type DecryptionLog struct {
    ID              primitive.ObjectID `bson:"_id,omitempty"`
    UserID          primitive.ObjectID `bson:"user_id"`
    EncryptedDataID primitive.ObjectID `bson:"encrypted_data_id"`
    DecryptedText   string             `bson:"decrypted_text"`
    CreatedAt       time.Time          `bson:"created_at"`
}
```

### Step 5: Setup Database Connection

Create a file `database/database.go` to handle the database connection:

```go
package database

import (
    "context"
    "log"
    "os"
    "time"

    "aes_encryption_project/models"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Connect() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("DB_NAME")

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatalf("Error creating MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    Client = client
    DB = client.Database(dbName)
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Collection(collectionName)
}
```

### Step 6: Update the Main Function and Encryption/Decryption Logic

Update `main.go` to include database interactions and refactor the encryption/decryption functions:

```go
package main

import (
    "aes_encryption_project/database"
    "aes_encryption_project/models"
    "context"
    "crypto/aes"
    "encoding/hex"
    "fmt"
    "log"
    "time"

    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
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
```

### Explanation

1. **Dependencies**: Use the MongoDB Go Driver for database interactions and `github.com/joho/godotenv` for environment variable management.
2. **Environment Variables**: Store MongoDB connection details in a `.env` file.
3. **Database Models**: Define `User`, `EncryptedData`, and `DecryptionLog` models in `models/models.go`.
4. **Database Connection**: Establish a database connection and set up collections in `database/database.go`.
5. **Main Function**: Connect to the database, create a sample user, encrypt a plaintext string, store the encrypted data, and decrypt it, logging the result in decryption logs.

## Updated README

# AES Encryption/Decryption with Database Integration using MongoDB in Go

This project demonstrates how to perform AES encryption and decryption in Go while integrating with a MongoDB database using the MongoDB Go Driver. The project includes user management, storing encrypted data, and logging decryption operations.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Usage](#usage)
- [Database Models](#database-models)
- [Functions](#functions)
- [Notes](#notes)

## Overview

This project performs AES encryption and decryption, stores encrypted data in a MongoDB database, and logs decryption operations. It utilizes the MongoDB Go Driver for ORM functionalities and dotenv for environment variable management.

## Project Structure

```
aes_encryption_project/
├── database/
│   └── database.go
├── models/
│   └── models.go
├── .env
├── go.mod
├── go.sum
└── main.go
```

## Dependencies

The project uses the following dependencies:

- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) - MongoDB driver for Go
- [MongoDB](https://www.mongodb.com/) - NoSQL database management system
- [godotenv](https://github.com/joho/godotenv) - Package to load environment variables from a `.env` file
## Dependencies (continued)

- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing library

## Setup

### 1. Install Go and MongoDB

Ensure Go and MongoDB are installed on your machine.

### 2. Clone the Repository

```sh
git clone https://github.com/yourusername/aes_encryption_project.git
cd aes_encryption_project
```

### 3. Install Dependencies

```sh
go mod tidy
```

### 4. Create a `.env` File

Create a `.env` file in the project root with your MongoDB connection details:

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=aes_encryption_db
```

### 5. Run the Project

```sh
go run main.go
```

## Usage

When you run the project, it will:
1. Connect to the MongoDB database.
2. Create a sample user.
3. Encrypt a plaintext string.
4. Store the encrypted data in the database.
5. Decrypt the encrypted data and log the operation.

## Database Models

### User Model

Stores user information.

```go
type User struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"`
    Username     string             `bson:"username"`
    Email        string             `bson:"email"`
    PasswordHash string             `bson:"password_hash"`
    CreatedAt    time.Time          `bson:"created_at"`
    UpdatedAt    time.Time          `bson:"updated_at"`
}
```

### Encrypted Data Model

Stores encrypted data for each user.

```go
type EncryptedData struct {
    ID            primitive.ObjectID `bson:"_id,omitempty"`
    UserID        primitive.ObjectID `bson:"user_id"`
    Plaintext     string             `bson:"plaintext"`
    Ciphertext    string             `bson:"ciphertext"`
    EncryptionKey string             `bson:"encryption_key"`
    CreatedAt     time.Time          `bson:"created_at"`
    UpdatedAt     time.Time          `bson:"updated_at"`
}
```

### Decryption Logs Model

Logs decryption operations.

```go
type DecryptionLog struct {
    ID              primitive.ObjectID `bson:"_id,omitempty"`
    UserID          primitive.ObjectID `bson:"user_id"`
    EncryptedDataID primitive.ObjectID `bson:"encrypted_data_id"`
    DecryptedText   string             `bson:"decrypted_text"`
    CreatedAt       time.Time          `bson:"created_at"`
}
```

## Functions

### `main.go`

- **Main Function**: Connects to the database, creates a sample user, encrypts a plaintext string, stores the encrypted data, decrypts it, and logs the operation.

### `database/database.go`

- **Connect**: Establishes a connection to the MongoDB database and sets up collections.
- **GetCollection**: Retrieves a MongoDB collection.

### `models/models.go`

- **User**: Defines the user model.
- **EncryptedData**: Defines the encrypted data model.
- **DecryptionLog**: Defines the decryption log model.

### `EncryptAES` Function

Encrypts plaintext using AES and returns the ciphertext.

```go
func EncryptAES(key []byte, plaintext string) string {
    c, err := aes.NewCipher(key)
    CheckError(err)

    out := make([]byte, len(plaintext))
    c.Encrypt(out, []byte(plaintext))

    return hex.EncodeToString(out)
}
```

### `DecryptAES` Function

Decrypts ciphertext using AES and logs the operation.

```go
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
```

### `CheckError` Function

Utility function to handle errors by panicking.

```go
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
```

## Notes

- Ensure the encryption key length is 16, 24, or 32 bytes as required by AES.
- Handle keys and plaintext securely in a real-world application. Avoid hardcoding sensitive information.
- Customize error handling for production use.
- Adjust the database models and relationships based on specific requirements.