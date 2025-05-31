package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type DoctorScheduleResponse struct {
	ID        int            `json:"id"`
	Doctor    DoctorResponse `json:"doctor"`
	DayOfWeek string         `json:"day_of_week"`
	StartTime string         `json:"start_time"` // time.TimeOnly
	EndTime   string         `json:"end_time"`   // time.TimeOnly
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func NewDoctorScheduleResponse(
	doctorScheduleEntity *entity.DoctorSchedule,
) DoctorScheduleResponse {
	return DoctorScheduleResponse{
		ID:        doctorScheduleEntity.ID,
		Doctor:    NewDoctorResponse(&doctorScheduleEntity.Doctor),
		DayOfWeek: time.Weekday(doctorScheduleEntity.DayOfWeek).String(),
		StartTime: doctorScheduleEntity.StartTime,
		EndTime:   doctorScheduleEntity.EndTime,
		CreatedAt: doctorScheduleEntity.CreatedAt,
		UpdatedAt: doctorScheduleEntity.UpdatedAt,
	}
}

type GetDoctorSchedulesQuery struct {
	DoctorID  uuid.UUID `validate:"omitempty,uuid"`
	DayOfWeek int       `query:"day_of_week" validate:"omitempty,oneof=0 1 2 3 4 5 6"`
}

type GetDoctorSchedulesResponse struct {
	DoctorSchedules []DoctorScheduleResponse `json:"doctor_schedules"`
}

type CreateDoctorScheduleRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id" validate:"required,uuid"`
	DayOfWeek int       `json:"day_of_week" validate:"required,oneof=0 1 2 3 4 5 6"`
	StartTime string    `json:"start_time" validate:"required"` // time.TimeOnly
	EndTime   string    `json:"end_time" validate:"required"`   // time.TimeOnly
}

type UpdateDoctorScheduleParams struct {
	ID int `param:"id" validate:"required"`
}

type UpdateDoctorScheduleRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id" validate:"required,uuid"`
	DayOfWeek int       `json:"day_of_week" validate:"required,oneof=0 1 2 3 4 5 6"`
	StartTime string    `json:"start_time" validate:"required"` // time.TimeOnly
	EndTime   string    `json:"end_time" validate:"required"`   // time.TimeOnly
}

type DeleteDoctorScheduleParams struct {
	ID int `param:"id" validate:"required"`
}
