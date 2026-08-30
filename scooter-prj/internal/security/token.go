package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSecureToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	rawToken := hex.EncodeToString(b)
	hashedToken := HashToken(rawToken)
	return hashedToken
}

func HashToken(rawToken string) string {
	hash := sha256.Sum256([]byte(rawToken))
	hashedToken := hex.EncodeToString(hash[:])
	return hashedToken
}
