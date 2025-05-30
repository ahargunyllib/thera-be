package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
)

func (ms *moodService) CreateMood(ctx context.Context, req dto.CreateMoodRequest) error {
	valErr := ms.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	mood := entity.Mood{
		DoctorID: req.DoctorID,
		Scale:    req.Scale,
	}

	err := ms.moodRepo.CreateMood(ctx, &mood)
	if err != nil {
		return err
	}

	return nil
}

func (ms *moodService) GetMyDailyMood(
	ctx context.Context,
	query dto.GetMyDailyMoodQuery,
) (dto.GetMyDailyMoodResponse, error) {
	valErr := ms.validator.Validate(query)
	if valErr != nil {
		return dto.GetMyDailyMoodResponse{}, valErr
	}

	moods, err := ms.moodRepo.GetMyDailyMood(ctx, query.DoctorID)
	if err != nil {
		return dto.GetMyDailyMoodResponse{}, err
	}

	moodsResponse := make([]dto.MoodResponse, len(moods))
	for i, mood := range moods {
		moodsResponse[i] = dto.NewMoodResponse(&mood)
	}

	res := dto.GetMyDailyMoodResponse{
		Moods: moodsResponse,
	}

	return res, nil
}

func (ms *moodService) GetMyMonthlyOverview(
	ctx context.Context,
	query dto.GetMyMonthlyOverviewQuery,
) (dto.GetMyMonthlyOverviewResponse, error) {
	valErr := ms.validator.Validate(query)
	if valErr != nil {
		return dto.GetMyMonthlyOverviewResponse{}, valErr
	}

	monthlyMoodStatistics, err := ms.moodRepo.GetMyMonthlyOverview(ctx, query.DoctorID)
	if err != nil {
		return dto.GetMyMonthlyOverviewResponse{}, err
	}

	monthlyMoods := map[string]int{}

	for _, stat := range monthlyMoodStatistics {
		monthlyMoods[stat.Month] = stat.Scale
	}

	res := dto.GetMyMonthlyOverviewResponse{
		MonthlyMoods: monthlyMoods,
	}

	return res, nil
}
