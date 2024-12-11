package main

import (
	"fmt"
	"os"
	"log"
	"envbuddy/encrypt"
	"envbuddy/utils"
)

func main() {
	// Charger la variable d'environnement depuis le fichier .env
	err := utils.LoadEnvVariable(".env", "ENCRYPTION_KEY")
	if err != nil {
		log.Fatalf("Error loading .env variable: %v\n", err)
	}

	// Exemple de traitement : chiffrer un fichier .env
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./envbuddy encrypt|decrypt <filename>")
		return
	}

	action := os.Args[1]
	filename := os.Args[2]

	switch action {
	case "encrypt":
		err := encrypt.EncryptEnvFile(filename)
		if err != nil {
			log.Fatalf("Error encrypting file: %v\n", err)
		}
	case "decrypt":
		err := encrypt.DecryptEnvFile(filename)
		if err != nil {
			log.Fatalf("Error decrypting file: %v\n", err)
		}
	default:
		fmt.Println("Invalid action. Use 'encrypt' or 'decrypt'")
	}
}
