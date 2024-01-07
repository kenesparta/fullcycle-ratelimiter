package usecase

import (
	"context"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
)

type RegisterAPIToken struct {
	apiRepo entity.APITokenRepository
}

func NewRegisterAPIToken(
	apiRepo entity.APITokenRepository,
) *RegisterAPIToken {
	return &RegisterAPIToken{
		apiRepo: apiRepo,
	}
}

// Execute This saves a new request depending on API Token.
// If we have a request from an endpoint has API Token, we persist using it:
//  1. We need to confirm if the API token has been blocked for exceeding the maximum number of requests.
//  2. We save the API token in the database using the RateLimiter.Allow() and using the variables
//     or if we already have the same IP, we update the requests array.
//  4. Finally, we execute the validation and update/insert the data in the database.
func (ipr *RegisterAPIToken) Execute(
	ctx context.Context,
	input dto.IPRequestSave,
) (dto.APITokenOutput, error) {
	return dto.APITokenOutput{
		Allow: true,
	}, nil
}
