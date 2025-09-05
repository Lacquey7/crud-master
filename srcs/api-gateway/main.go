package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

const tokenLength = 32 // 256 bits

func generateSecureToken() (string, error) {
	bytes := make([]byte, tokenLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func main() {
	targetToken := "8ad493e7e180e481efa4d282f412ad4a67ebddc516646219b40ff26341176535" // token cible

	attempt := 0
	start := time.Now()

	for {
		attempt++
		token, err := generateSecureToken()
		if err != nil {
			fmt.Println("Erreur génération token:", err)
			break
		}

		if token == targetToken {
			fmt.Printf("Token trouvé après %d tentatives en %v !\n", attempt, time.Since(start))
			break
		}

		// Affiche une mise à jour toutes les 100 000 tentatives
		if attempt%100000000 == 0 {
			fmt.Printf("Tentatives: %d - Temps écoulé: %v\n", attempt, time.Since(start))
		}

		// Limite max de tentatives pour ne pas tourner indéfiniment
		if attempt >= 1000000000 {
			fmt.Println("Limite de 100 million de tentatives atteinte, arrêt du programme")
			break
		}
	}
}
