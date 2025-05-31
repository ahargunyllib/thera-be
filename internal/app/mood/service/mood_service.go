package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/ahargunyllib/thera-be/pkg/log"
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

	go func() {
		if req.Scale <= 3 {
			doctor, getDoctorError := ms.moodRepo.GetDoctorByID(ctx, req.DoctorID)
			if getDoctorError != nil {
				log.Error(log.CustomLogInfo{
					"message": "Failed to get doctor by ID for mood notification",
					"err":     getDoctorError,
				}, "[moodService.CreateMood]")
				return
			}

			id, getDoctorError := ms.ulid.New()
			if getDoctorError != nil {
				log.Error(log.CustomLogInfo{
					"message": "Failed to generate ULID for mood notification",
					"err":     getDoctorError,
				}, "[moodService.CreateMood]")
				return
			}

			notification := entity.Notification{
				ID: id.String(),
				HospitalID: sql.NullInt64{
					Int64: int64(doctor.HospitalID),
					Valid: true,
				},
				Title: "Low Mood Alert",
				Body: sql.NullString{
					String: fmt.Sprintf(
						"Doctor %s has reported a low mood with a scale of %d. Please check in.",
						doctor.FullName,
						req.Scale,
					),
					Valid: true,
				},
				Type: enums.NotificationTypeImproveNextDoctorScheduleIdx,
			}

			getDoctorError = ms.moodRepo.CreateMoodNotification(ctx, &notification)
			if getDoctorError != nil {
				// Log the error or handle it as needed
				log.Error(log.CustomLogInfo{
					"message": "Failed to create mood notification",
					"err":     getDoctorError,
				}, "[moodService.CreateMood]")
			}
		}
	}()

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
