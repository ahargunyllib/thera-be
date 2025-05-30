package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type MoodResponse struct {
	ID        int       `json:"id"`
	DoctorID  uuid.UUID `json:"doctor_id"`
	Scale     int       `json:"scale"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMoodResponse(moodEntity *entity.Mood) MoodResponse {
	return MoodResponse{
		ID:        moodEntity.ID,
		DoctorID:  moodEntity.DoctorID,
		Scale:     moodEntity.Scale,
		CreatedAt: moodEntity.CreatedAt,
	}
}

type CreateMoodRequest struct {
	DoctorID uuid.UUID `validate:"required"`
	Scale    int       `json:"scale" validate:"required,min=1,max=5"`
}

type GetMyDailyMoodQuery struct {
	DoctorID uuid.UUID `validate:"required"`
}

type GetMyDailyMoodResponse struct {
	Moods []MoodResponse `json:"moods"`
}

type GetMyMonthlyOverviewQuery struct {
	DoctorID uuid.UUID `validate:"required"`
}

type GetMyMonthlyOverviewResponse struct {
	MonthlyMoods map[string]int `json:"monthly_moods"`
}
