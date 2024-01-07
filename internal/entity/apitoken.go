package entity

import (
	"crypto/rand"
	"encoding/hex"
)

const (
	APIKeyPrefixRateKey     = "rate:api-key"
	APIKeyPrefixDurationKey = "duration:api-key"
	APIKeyPrefixValueKey    = "value:api-key"
	StatusAPIKeyBlocked     = "APIKeyBlocked"
)

// APIKey Thi token is related with the token API that we can generate for each client
type APIKey struct {
	value string

	// BlockedDuration is the number of SECONDS that it blocks the token if it reaches the RateLimiter.MaxRequests each
	// RateLimiter.TimeWindowSec amount of seconds.
	BlockedDuration int64
	RateLimiter     RateLimiter
}

func (at *APIKey) GenerateValue() error {
	bytes, err := generateRandomBytes(32)
	if err != nil {
		return err
	}
	at.value = hex.EncodeToString(bytes)
	return nil
}

func (at *APIKey) SetValue(value string) {
	at.value = value
}

func (at *APIKey) Value() string {
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
