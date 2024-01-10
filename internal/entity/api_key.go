package entity

import (
	"crypto/rand"
	"encoding/hex"
)

const (
	APIKeyPrefixRateKey            = "rate:api-key"
	APIKeyPrefixBlockedDurationKey = "blocked:api-key"
	StatusAPIKeyBlocked            = "APIKeyBlocked"
	APIKeyHeaderName               = "API_KEY"
)

// APIKey This key is related with the API Key that we can generate for each client
type APIKey struct {
	value string

	// BlockedDuration is the number of SECONDS that it blocks the API Key if it reaches the RateLimiter.MaxRequests
	BlockedDuration int64

	// RateLimiter.TimeWindowSec amount of seconds.
	RateLimiter RateLimiter
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

func (at *APIKey) Validate() error {
	if at.BlockedDuration == 0 {
		return ErrBlockedTimeDuration
	}

	if err := at.RateLimiter.Validate(); err != nil {
		return err
	}

	return nil
}

func generateRandomBytes(length int) ([]byte, error) {
	byteSlice := make([]byte, length)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	return byteSlice, nil
}
