package usecase

import (
	"context"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
)

type CreateAPITokenUseCase struct {
	apitokenRepo entity.APITokenRepository
}

func NewCreateAPITokenUseCase(apitokenRepo entity.APITokenRepository) *CreateAPITokenUseCase {
	return &CreateAPITokenUseCase{apitokenRepo: apitokenRepo}
}

func (cr *CreateAPITokenUseCase) Execute(ctx context.Context, input dto.APITokenInput) error {
	return nil
}
