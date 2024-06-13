# AES Encryption/Decryption in Go

Here is an explanation of a simple AES encryption and decryption example in Go. The code demonstrates how to encrypt a plaintext string using a specified key and decrypt it back to its original form.

## Overview

The code consists of three main parts:
1. **Main function**: Initializes the key and plaintext, performs encryption, prints the results, and calls the decryption function.
2. **EncryptAES function**: Handles the encryption process using the AES algorithm.
3. **DecryptAES function**: Handles the decryption process to retrieve the original plaintext.
4. **CheckError function**: Utility function to handle errors.

### Main Function

- **Cipher Key**: A 32-byte key used for the AES encryption. This must be either 16, 24, or 32 bytes long, as per AES requirements.
- **Plaintext**: The string to be encrypted.
- **EncryptAES**: Calls the `EncryptAES` function to encrypt the plaintext using the provided key.
- **Print Statements**: Displays the original plaintext and the resulting ciphertext.
- **DecryptAES**: Calls the `DecryptAES` function to decrypt the ciphertext back to plaintext.

### EncryptAES Function

This function takes a key and plaintext string as inputs and returns the encrypted text as a hex-encoded string.

1. **New Cipher**: Creates a new AES cipher using the provided key.
2. **Output Byte Slice**: Initializes a byte slice to hold the encrypted data.
3. **Encrypt**: Encrypts the plaintext into the byte slice.
4. **Hex Encoding**: Converts the encrypted byte slice into a hex-encoded string for easy display and transmission.

### DecryptAES Function

This function takes a key and a ciphertext string as inputs and prints the decrypted plaintext.

1. **Hex Decoding**: Converts the hex-encoded ciphertext back to a byte slice.
2. **New Cipher**: Creates a new AES cipher using the provided key.
3. **Output Byte Slice**: Initializes a byte slice to hold the decrypted data.
4. **Decrypt**: Decrypts the ciphertext into the byte slice.
5. **Print Decrypted Text**: Converts the decrypted byte slice to a string and prints it.

### CheckError Function

This utility function checks for errors and panics if any are found. It is used throughout the code to ensure that functions are executing correctly.

```go
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
```

## How to Run

1. **Set Up Go Environment**: Ensure you have Go installed and properly set up on your machine.
2. **Create a New File**: Copy the code into a new Go file, for example, `main.go`.
3. **Run the Code**: Execute the code using `go run main.go`.

## Example Output

When you run the code, it will display the following output:
- The original plaintext.
- The encrypted ciphertext in hex format.
- The decrypted plaintext.

This verifies that the encryption and decryption processes are working correctly.