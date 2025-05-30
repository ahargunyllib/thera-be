package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	openai "github.com/ahargunyllib/thera-be/pkg/opeanai"
	"github.com/bytedance/sonic"
)

func (dss *doctorScheduleService) CreateDoctorSchedule(ctx context.Context, req dto.CreateDoctorScheduleRequest) error {
	valErr := dss.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	_, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		return errx.ErrInvalidTimeFormat.WithDetails(map[string]any{
			"start_time": req.StartTime,
		})
	}

	_, err = time.Parse("15:04", req.EndTime)
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

	doctorSchedule, err := dss.doctorScheduleRepo.GetDoctorScheduleByID(ctx, params.ID)
	if err != nil {
		return err
	}

	// if req.DayOfWeek != 0 && req.DayOfWeek != doctorSchedule.DayOfWeek {
	// 	doctorSchedule.DayOfWeek = req.DayOfWeek
	// }

	if req.StartTime != "" && req.StartTime != doctorSchedule.StartTime {
		_, err := time.Parse("15:04", req.StartTime)
		if err != nil {
			return errx.ErrInvalidTimeFormat.WithDetails(map[string]any{
				"start_time": req.StartTime,
			})
		}
		doctorSchedule.StartTime = req.StartTime
	}

	if req.EndTime != "" && req.EndTime != doctorSchedule.EndTime {
		_, err := time.Parse("15:04", req.EndTime)
		if err != nil {
			return errx.ErrInvalidTimeFormat.WithDetails(map[string]any{
				"end_time": req.EndTime,
			})
		}
		doctorSchedule.EndTime = req.EndTime
	}

	err = dss.doctorScheduleRepo.UpdateDoctorSchedule(ctx, doctorSchedule)
	if err != nil {
		return err
	}

	return nil
}

func (dss *doctorScheduleService) GetPreviewImprovedNextSchedule(
	ctx context.Context,
	query dto.GetPreviewImprovedNextScheduleQuery,
) (dto.GetPreviewImprovedNextScheduleResponse, error) {
	valErr := dss.validator.Validate(query)
	if valErr != nil {
		return dto.GetPreviewImprovedNextScheduleResponse{}, valErr
	}

	nextSchedule, err := dss.doctorScheduleRepo.GetNextScheduleByDoctorID(ctx, query.DoctorID)
	if err != nil {
		return dto.GetPreviewImprovedNextScheduleResponse{}, err
	}

	lastMood, err := dss.doctorScheduleRepo.GetLastMoodByDoctorID(ctx, query.DoctorID)
	if err != nil {
		return dto.GetPreviewImprovedNextScheduleResponse{}, err
	}

	messages := make([]openai.Message, 2)
	messages[0] = openai.Message{
		Role: "system",
		Content: `
			You are an assistant that helps adjust a doctor's schedule based on their mood rating.

			The user will give you the start time, end time, and mood rating (1-5). Your job is to generate a new start_time and end_time that slightly or significantly reduces the work hours based on the mood:
			- If the mood is 1 or 2, shorten the hours by 2-4 hours if possible.
			- If the mood is 3, reduce the hours by 1-2 hours.
			- If the mood is 4 or 5, keep the schedule unchanged.

			Important rules:
			- Always use 24-hour format for times, no AM/PM.
			- Do not include any markdown, triple backticks, or code blocks.
			- Do not include any explanations, comments, or extra text.
			- Only return a valid JSON object with "start_time" and "end_time" as string values.

			Example:
			{
				"start_time": "09:00",
				"end_time": "14:00"
			}

			If no change is possible (e.g., time is too short to adjust), return the original times.
		`,
	}

	messages[1] = openai.Message{
		Role: "user",
		Content: fmt.Sprintf(`
			start_time: %s,
			end_time: %s,
			mood: %d
		`, nextSchedule.StartTime, nextSchedule.EndTime, lastMood.Scale),
	}

	chatRes, err := dss.openai.Chat(ctx, messages)
	if err != nil {
		return dto.GetPreviewImprovedNextScheduleResponse{}, err
	}

	var adjustedSchedule struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	content := chatRes.Choices[0].Message.Content

	err = sonic.Unmarshal([]byte(content), &adjustedSchedule)
	if err != nil {
		return dto.GetPreviewImprovedNextScheduleResponse{}, err
	}

	if adjustedSchedule.StartTime == "" || adjustedSchedule.EndTime == "" {
		return dto.GetPreviewImprovedNextScheduleResponse{}, errx.ErrInvalidResponseFormat.WithDetails(map[string]any{
			"response": content,
		})
	}

	nextScheduleAfterImproved := nextSchedule // * WARN: SHALLOW COPY
	nextScheduleAfterImproved.StartTime = adjustedSchedule.StartTime
	nextScheduleAfterImproved.EndTime = adjustedSchedule.EndTime

	nextScheduleResponse := dto.NewDoctorScheduleResponse(nextSchedule)
	nextScheduleAfterImprovedResponse := dto.NewDoctorScheduleResponse(nextScheduleAfterImproved)

	res := dto.GetPreviewImprovedNextScheduleResponse{
		Schedule:     nextScheduleResponse,
		NextSchedule: nextScheduleAfterImprovedResponse,
	}

	return res, nil
}
