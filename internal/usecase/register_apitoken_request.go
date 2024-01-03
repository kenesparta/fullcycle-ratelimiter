package usecase

import (
	"context"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
)

type RegisterAPITokenRequest struct {
	apiRepo entity.APITokenRepository
}

func NewRegisterAPITokenRequest(
	apiRepo entity.APITokenRepository,
) *RegisterAPITokenRequest {
	return &RegisterAPITokenRequest{
		apiRepo: apiRepo,
	}
}

// Execute This saves a new request depending on IP or APIToken.
// If we have a request from an endpoint that does not have API Token, we persist the request using the IP
//  1. We need to confirm if the IP has been blocked for exceeding the maximum number of requests.
//  2. We save the IP in the database using the RateLimiter.Allow() and using the environmental variables
//     or if we already have the same IP, we update the requests array.
//  4. Finally, we execute the validation and update/insert the data in the database.
func (ipr *RegisterAPITokenRequest) Execute(
	ctx context.Context,
	input dto.IPRequestSave,
) (dto.IPRequestResult, error) {
	return dto.IPRequestResult{
		Allow: true,
	}, nil
}
