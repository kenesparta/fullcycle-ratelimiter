package usecase

import (
	"context"
	"log"

	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
)

type CreateAPIKeyUseCase struct {
	apiKeyRepo entity.APIKeyRepository
}

func NewCreateAPIKeyUseCase(apiKeyRepo entity.APIKeyRepository) *CreateAPIKeyUseCase {
	return &CreateAPIKeyUseCase{apiKeyRepo: apiKeyRepo}
}

// Execute Creates a new API Key with its own value and persists it
func (cr *CreateAPIKeyUseCase) Execute(ctx context.Context, input dto.APIKeyInput) (dto.APIKeyCreateOutput, error) {
	apiKey := entity.APIKey{
		BlockedDuration: input.BlockedDuration,
		RateLimiter: entity.RateLimiter{
			TimeWindowSec: input.TimeWindowSec,
			MaxRequests:   input.MaxRequests,
		},
	}

	if err := apiKey.GenerateValue(); err != nil {
		log.Printf("Error on CreateAPIKeyUseCase generating key value: %s\n", err.Error())
		return dto.APIKeyCreateOutput{}, err
	}

	keyValue, saveErr := cr.apiKeyRepo.Save(ctx, &apiKey)
	if saveErr != nil {
		log.Printf("Error on CreateAPIKeyUseCase saving data: %s\n", saveErr.Error())
		return dto.APIKeyCreateOutput{}, saveErr
	}

	log.Println("Saved with success")
	return dto.APIKeyCreateOutput{
		KeyValue: keyValue,
	}, nil
}
