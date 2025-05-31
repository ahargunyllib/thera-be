package service

import (
	"context"
	"time"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
)

func (dss *doctorScheduleService) CreateDoctorSchedule(ctx context.Context, req dto.CreateDoctorScheduleRequest) error {
	valErr := dss.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	_, err := time.Parse(req.StartTime, time.TimeOnly)
	if err != nil {
		return errx.ErrInvalidTimeFormat.WithDetails(map[string]any{
			"start_time": req.StartTime,
		})
	}

	_, err = time.Parse(req.EndTime, time.TimeOnly)
	if err != nil {
		return errx.ErrInvalidTimeFormat.WithDetails(map[string]any{
			"end_time": req.EndTime,
		})
	}

	doctorSchedule := &entity.DoctorSchedule{
		DoctorID:  req.DoctorID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	err = dss.doctorScheduleRepo.CreateDoctorSchedule(ctx, doctorSchedule)
	if err != nil {
		return err
	}

	return nil
}

func (dss *doctorScheduleService) DeleteDoctorSchedule(
	ctx context.Context,
	params dto.DeleteDoctorScheduleParams,
) error {
	valErr := dss.validator.Validate(params)
	if valErr != nil {
		return valErr
	}

	err := dss.doctorScheduleRepo.DeleteDoctorSchedule(ctx, params.ID)
	if err != nil {
		return err
	}

	return nil
}

func (dss *doctorScheduleService) GetDoctorSchedules(
	ctx context.Context,
	query dto.GetDoctorSchedulesQuery,
) ([]dto.DoctorScheduleResponse, error) {
	valErr := dss.validator.Validate(query)
	if valErr != nil {
		return nil, valErr
	}

	schedules, err := dss.doctorScheduleRepo.GetDoctorSchedules(ctx, &query)
	if err != nil {
		return nil, err
	}

	schedulesResponse := make([]dto.DoctorScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		schedulesResponse[i] = dto.NewDoctorScheduleResponse(&schedule)
	}

	return schedulesResponse, nil
}

func (dss *doctorScheduleService) UpdateDoctorSchedule(
	ctx context.Context,
	params dto.UpdateDoctorScheduleParams,
	req dto.UpdateDoctorScheduleRequest,
) error {
	valErr := dss.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	doctorSchedule := &entity.DoctorSchedule{
		ID:        params.ID,
		DoctorID:  req.DoctorID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	err := dss.doctorScheduleRepo.UpdateDoctorSchedule(ctx, doctorSchedule)
	if err != nil {
		return err
	}

	return nil
}
