package usecase

import (
	"context"
	"log"

	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
)

type CreateAPITokenUseCase struct {
	apitokenRepo entity.APITokenRepository
}

func NewCreateAPITokenUseCase(apitokenRepo entity.APITokenRepository) *CreateAPITokenUseCase {
	return &CreateAPITokenUseCase{apitokenRepo: apitokenRepo}
}

// Execute Creates a new API token with its own value and persists it
func (cr *CreateAPITokenUseCase) Execute(ctx context.Context, input dto.APITokenInput) (dto.APITokenCreateOutput, error) {
	apiToken := entity.APIToken{
		BlockedDuration: input.BlockedDuration,
		RateLimiter: entity.RateLimiter{
			TimeWindowSec: input.TimeWindowSec,
			MaxRequests:   input.MaxRequests,
		},
	}

	if err := apiToken.GenerateValue(); err != nil {
		log.Printf("Error on CreateAPITokenUseCase generating token value: %s\n", err.Error())
		return dto.APITokenCreateOutput{}, err
	}

	tokenValue, saveErr := cr.apitokenRepo.Save(ctx, &apiToken)
	if saveErr != nil {
		log.Printf("Error on CreateAPITokenUseCase saving data: %s\n", saveErr.Error())
		return dto.APITokenCreateOutput{}, saveErr
	}

	log.Println("Saved with success")
	return dto.APITokenCreateOutput{
		TokenValue: tokenValue,
	}, nil
}
