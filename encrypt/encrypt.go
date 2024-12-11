package encrypt

import (
	"fmt"
	"io/ioutil"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
	"errors"
)

// GetEncryptionKey retrieves the encryption key from the environment variable
func GetEncryptionKey() ([]byte, error) {
	// Retrieve the AES key from the environment variable
	key := os.Getenv("ENCRYPTION_KEY")

	// If the key is not set, return an error
	if key == "" {
		return nil, fmt.Errorf("the ENCRYPTION_KEY environment variable is not set")
	}

	// Return the key as a byte slice
	return []byte(key), nil
}

// EncryptEnvFile encrypts a .env file
func EncryptEnvFile(filename string) error {
	// Read the content of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading the file: %w", err)
	}

	// Get the AES key from the environment variable
	key, err := GetEncryptionKey()
	if err != nil {
		return err
	}

	// Create a new AES cipher block using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error creating AES cipher: %w", err)
	}

	// Create a byte slice to hold the ciphertext (with space for the IV)
	ciphertext := make([]byte, aes.BlockSize+len(content))

	// Generate a random initialization vector (IV)
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return fmt.Errorf("error generating initialization vector: %w", err)
	}

	// Create a new encrypter stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt the content
	stream.XORKeyStream(ciphertext[aes.BlockSize:], content)

	// Save the encrypted file
	outputFile := filename + ".encrypted"
	err = ioutil.WriteFile(outputFile, ciphertext, 0644)
	if err != nil {
		return fmt.Errorf("error writing the encrypted file: %w", err)
	}

	fmt.Printf("The encrypted file has been saved as %s\n", outputFile)
	return nil
}

// DecryptEnvFile decrypts an encrypted .env file
func DecryptEnvFile(filename string) error {
	// Read the encrypted file
	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading the encrypted file: %w", err)
	}

	// Check if the ciphertext is long enough to contain the IV
	if len(ciphertext) < aes.BlockSize {
		return errors.New("the file is too short to be valid")
	}

	// Get the AES key from the environment variable
	key, err := GetEncryptionKey()
	if err != nil {
		return err
	}

	// Create a new AES cipher block using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error creating AES decryption cipher: %w", err)
	}

	// Extract the IV from the beginning of the encrypted file
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Create a new decrypter stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt the content
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	// Save the decrypted file
	outputFile := filename + ".decrypted"
	err = ioutil.WriteFile(outputFile, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("error writing the decrypted file: %w", err)
	}

	fmt.Printf("The decrypted file has been saved as %s\n", outputFile)
	return nil
}
