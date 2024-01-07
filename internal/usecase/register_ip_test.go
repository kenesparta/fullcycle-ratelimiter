package usecase

import (
	"context"
	"github.com/kenesparta/fullcycle-ratelimiter/config"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestRegisterIPRequest_Execute(t *testing.T) {
	startTime := time.Date(2024, time.January, 1, 12, 34, 56, 0, time.UTC)
	cfg := &config.Config{
		RateLimiter: config.RateLimiter{
			ByIP: config.LimitValues{
				MaxRequests:     4,
				TimeWindow:      1,
				BlockedDuration: 600,
			},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIPRepo := mock.NewMockIPRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name       string
		input      dto.IPRequestSave
		getRequest *entity.RateLimiter
		expected   dto.IPRequestResult
	}{
		{
			name: "with api token and result in getRequest",
			input: dto.IPRequestSave{
				IP:      "127.0.0.1",
				TimeAdd: startTime,
			},
			getRequest: func() *entity.RateLimiter {
				rt := &entity.RateLimiter{}
				rt.AddRequests(startTime.Add(-30 * time.Millisecond))
				rt.AddRequests(startTime.Add(-20 * time.Millisecond))
				rt.AddRequests(startTime.Add(-10 * time.Millisecond))
				return rt
			}(),
			expected: dto.IPRequestResult{
				Allow: true,
			},
		},
		{
			name: "with api token and result in getRequest",
			input: dto.IPRequestSave{
				IP:      "192.168.0.166",
				TimeAdd: startTime,
			},
			getRequest: func() *entity.RateLimiter {
				rt := &entity.RateLimiter{}
				rt.AddRequests(startTime.Add(-40 * time.Millisecond))
				rt.AddRequests(startTime.Add(-30 * time.Millisecond))
				rt.AddRequests(startTime.Add(-20 * time.Millisecond))
				rt.AddRequests(startTime.Add(-10 * time.Millisecond))
				return rt
			}(),
			expected: dto.IPRequestResult{
				Allow: false,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockIPRepo.EXPECT().GetBlockedDuration(ctx, tc.input.IP).Return("", nil).AnyTimes()
			mockIPRepo.EXPECT().GetRequest(ctx, tc.input.IP).Return(tc.getRequest, nil).AnyTimes()
			mockIPRepo.EXPECT().UpsertRequest(ctx, tc.input.IP, tc.getRequest).Return(nil).AnyTimes()
			mockIPRepo.EXPECT().
				SaveBlockedDuration(ctx, tc.input.IP, cfg.RateLimiter.ByIP.BlockedDuration).
				Return(nil).
				AnyTimes()
			rrq := NewRegisterIPUseCase(mockIPRepo, cfg)
			exRes, err := rrq.Execute(ctx, tc.input)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, exRes)
		})
	}
}
