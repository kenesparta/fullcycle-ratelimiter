package usecase

import (
	"context"
	"log"

	"github.com/kenesparta/fullcycle-ratelimiter/config"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
)

type RegisterIPRequest struct {
	config config.IConfig
	ipRepo entity.IPRepository
}

func NewRegisterRequest(
	ipRepo entity.IPRepository,
	config config.IConfig,
) *RegisterIPRequest {
	return &RegisterIPRequest{
		ipRepo: ipRepo,
		config: config,
	}
}

// Execute This saves a new request depending on IP or APIToken.
// If we have a request from an endpoint that does not have API Token, we persist the request using the IP
//  1. We need to confirm if the IP has been blocked for exceeding the maximum number of requests.
//  2. We save the IP in the database using the RateLimiter.Allow() and using the environmental variables
//     or if we already have the same IP, we update the requests array.
//  4. Finally, we execute the validation and update/insert the data in the database.
func (ipr *RegisterIPRequest) Execute(ctx context.Context, input dto.RequestSave) error {
	ipValue, blockedErr := ipr.ipRepo.GetBlockedDuration(ctx, input.IP)
	if blockedErr != nil {
		return blockedErr
	}

	if ipValue != "" {
		log.Println("ip is blocked due to exceeding the maximum number of requests")
		return entity.ErrIPExceededAmountRequest
	}

	return nil
}
