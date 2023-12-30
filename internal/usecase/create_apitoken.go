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
func (cr *CreateAPITokenUseCase) Execute(ctx context.Context, input dto.APITokenInput) error {
	apiToken := entity.APIToken{
		BlockedDuration: input.BlockedDuration,
		RateLimiter: entity.RateLimiter{
			TimeWindowSec: input.TimeWindowSec,
			MaxRequests:   input.MaxRequests,
		},
	}

	if err := apiToken.GenerateValue(); err != nil {
		log.Printf("Error on CreateAPITokenUseCase generating token value: %s\n", err.Error())
		return err
	}

	if err := cr.apitokenRepo.Save(ctx, &apiToken); err != nil {
		log.Printf("Error on CreateAPITokenUseCase saving data: %s\n", err.Error())
		return err
	}

	log.Println("Saving with success")
	return nil
}
