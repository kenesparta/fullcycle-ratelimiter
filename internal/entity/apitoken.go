package entity

import (
	"crypto/rand"
	"encoding/hex"
)

// APIToken Thi token is related with the token API that we can generate for each client
type APIToken struct {
	Value           string
	BlockedDuration int64
	RateLimit       RateLimiter
}

func (at *APIToken) GenerateValue() (string, error) {
	bytes, err := generateRandomBytes(32)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func generateRandomBytes(length int) ([]byte, error) {
	byteSlice := make([]byte, length)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	return byteSlice, nil
}
