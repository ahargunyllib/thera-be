package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type MoodRepository interface {
	CreateMood(ctx context.Context, mood *entity.Mood) error
	GetMyDailyMood(ctx context.Context, doctorID uuid.UUID) ([]entity.Mood, error)
	GetMyMonthlyOverview(ctx context.Context, doctorID uuid.UUID) ([]entity.MonthlyMoodStatistic, error)
}

type MoodService interface {
	CreateMood(ctx context.Context, req dto.CreateMoodRequest) error
	GetMyDailyMood(ctx context.Context, query dto.GetMyDailyMoodQuery) (dto.GetMyDailyMoodResponse, error)
	GetMyMonthlyOverview(
		ctx context.Context,
		query dto.GetMyMonthlyOverviewQuery,
	) (dto.GetMyMonthlyOverviewResponse, error)
}
