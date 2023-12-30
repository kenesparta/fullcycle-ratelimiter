package entity

import (
	"crypto/rand"
	"encoding/hex"
)

// APIToken Thi token is related with the token API that we can generate for each client
type APIToken struct {
	value string

	// BlockedDuration is the number of SECONDS that it blocks the token if it reaches the RateLimiter.MaxRequests each
	// RateLimiter.TimeWindowSec amount of seconds.
	BlockedDuration int64
	RateLimiter     RateLimiter
}

func (at *APIToken) GenerateValue() error {
	bytes, err := generateRandomBytes(32)
	if err != nil {
		return err
	}
	at.value = hex.EncodeToString(bytes)
	return nil
}

func (at *APIToken) Value() string {
	return at.value
}

func generateRandomBytes(length int) ([]byte, error) {
	byteSlice := make([]byte, length)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	return byteSlice, nil
}
