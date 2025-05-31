package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type DoctorScheduleRepository interface {
	GetDoctorSchedules(ctx context.Context, query *dto.GetDoctorSchedulesQuery) ([]entity.DoctorSchedule, error)
	GetDoctorScheduleByID(ctx context.Context, scheduleID int) (*entity.DoctorSchedule, error)
	CreateDoctorSchedule(ctx context.Context, schedule *entity.DoctorSchedule) error
	UpdateDoctorSchedule(ctx context.Context, schedule *entity.DoctorSchedule) error
	DeleteDoctorSchedule(ctx context.Context, scheduleID int) error

	GetNextScheduleByDoctorID(
		ctx context.Context,
		doctorID uuid.UUID,
	) (*entity.DoctorSchedule, error)

	GetLastMoodByDoctorID(
		ctx context.Context,
		doctorID uuid.UUID,
	) (*entity.Mood, error)
}

type DoctorScheduleService interface {
	GetDoctorSchedules(ctx context.Context, query dto.GetDoctorSchedulesQuery) ([]dto.DoctorScheduleResponse, error)
	CreateDoctorSchedule(ctx context.Context, req dto.CreateDoctorScheduleRequest) error
	UpdateDoctorSchedule(
		ctx context.Context,
		params dto.UpdateDoctorScheduleParams,
		req dto.UpdateDoctorScheduleRequest,
	) error
	DeleteDoctorSchedule(ctx context.Context, params dto.DeleteDoctorScheduleParams) error

	GetPreviewImprovedNextSchedule(
		ctx context.Context,
		query dto.GetPreviewImprovedNextScheduleQuery,
	) (dto.GetPreviewImprovedNextScheduleResponse, error)
}
