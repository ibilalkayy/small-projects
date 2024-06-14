# AES Encryption/Decryption with Database Integration using GORM in Go

This project demonstrates how to perform AES encryption and decryption in Go while integrating with a PostgreSQL database using GORM. The project includes user management, storing encrypted data, and logging decryption operations.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Usage](#usage)
- [Functions](#functions)
- [Notes](#notes)

## Overview

This project performs AES encryption and decryption, stores encrypted data in a PostgreSQL database, and logs decryption operations. It utilizes GORM for ORM functionalities and dotenv for environment variable management.

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

- [GORM](https://gorm.io/) - ORM library for Go
- [PostgreSQL](https://www.postgresql.org/) - Relational database management system
- [godotenv](https://github.com/joho/godotenv) - Package to load environment variables from a `.env` file
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing library

## Setup

### 1. Install Go and PostgreSQL

Ensure Go and PostgreSQL are installed on your machine.

### 2. Clone the Repository

```sh
git clone https://github.com/ibilalkayy/small-projects.git
cd aes-256
```

### 3. Install Dependencies

```sh
go mod tidy
```

### 4. Create a `.env` File

Create a `.env` file in the project root with your database connection details:

```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=5432
```

### 5. Run the Project

```sh
go run main.go
```

## Usage

When you run the project, it will:
1. Connect to the PostgreSQL database.
2. Create a sample user.
3. Encrypt a plaintext string.
4. Store the encrypted data in the database.
5. Decrypt the encrypted data and log the operation.

## Functions

### `main.go`

- **Main Function**: Connects to the database, creates a sample user, encrypts a plaintext string, stores the encrypted data, decrypts it, and logs the operation.

### `database/database.go`

- **Connect**: Establishes a connection to the PostgreSQL database using GORM and performs migrations.

### `models/models.go`

- **User**: Defines the user model.
- **EncryptedData**: Defines the encrypted data model.
- **DecryptionLog**: Defines the decryption log model.

### `EncryptAES` Function

Encrypts plaintext using AES and returns the ciphertext.

### `DecryptAES` Function

Decrypts ciphertext using AES and logs the operation.

### `CheckError` Function

Utility function to handle errors by panicking.

## Notes

- Ensure the encryption key length is 16, 24, or 32 bytes as required by AES.
- Handle keys and plaintext securely in a real-world application. Avoid hardcoding sensitive information.
- Customize error handling for production use.
- Adjust the database models and relationships based on specific requirements.