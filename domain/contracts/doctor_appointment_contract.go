package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type DoctorAppointmentRepository interface {
	GetDoctorAppointments(ctx context.Context, query *dto.GetDoctorAppointmentsQuery) ([]entity.DoctorAppointment, error)
	CreateDoctorAppointment(ctx context.Context, doctorAppointment *entity.DoctorAppointment) error
	UpdateDoctorAppointment(ctx context.Context, doctorAppointment *entity.DoctorAppointment) error
	DeleteDoctorAppointment(ctx context.Context, id string) error

	CheckDoctorScheduleAvailability(
		ctx context.Context,
		doctorID uuid.UUID,
		startTime string,
		endTime string,
	) (bool, error)
}

type DoctorAppointmentService interface {
	GetDoctorAppointments(ctx context.Context, query dto.GetDoctorAppointmentsQuery) (
		dto.GetDoctorAppointmentsResponse,
		error,
	)
	CreateDoctorAppointment(ctx context.Context, req dto.CreateDoctorAppointmentRequest) error
	UpdateDoctorAppointment(
		ctx context.Context,
		params dto.UpdateDoctorAppointmentParams,
		req dto.UpdateDoctorAppointmentRequest,
	) error
	DeleteDoctorAppointment(ctx context.Context, params dto.DeleteDoctorAppointmentParams) error
}
