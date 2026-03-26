package lib

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateAPIKeys() (string, string, error) {
	pubBytes := make([]byte, 24)
	if _, err := rand.Read(pubBytes); err != nil {
		return "", "", err
	}
	
	privBytes := make([]byte, 32)
	if _, err := rand.Read(privBytes); err != nil {
		return "", "", err
	}

	// Format them safely. e.g pk_test_1234abcd and sk_test_5678efgh
	pubKey := fmt.Sprintf("pk_test_%s", hex.EncodeToString(pubBytes))
	privKey := fmt.Sprintf("sk_test_%s", hex.EncodeToString(privBytes))

	return pubKey, privKey, nil
}


func GenerateLiveApiKeys()(string, string, error){
	pubBytes := make([]byte, 24)
	if _, err := rand.Read(pubBytes); err != nil {
		return "", "", err
	}

	privBytes := make([]byte, 32)
	if _, err := rand.Read(privBytes); err != nil {
		return "", "", err
	}

	// Format them safely. e.g pk_test_1234abcd and sk_test_5678efgh
	pubKey := fmt.Sprintf("pk_live_%s", hex.EncodeToString(pubBytes))
	privKey := fmt.Sprintf("sk_live_%s", hex.EncodeToString(privBytes))

	return pubKey, privKey, nil
}
