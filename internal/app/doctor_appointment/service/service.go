package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/ulid"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type doctorAppointmentService struct {
	doctorAppointmentRepo contracts.DoctorAppointmentRepository
	validator             validator.CustomValidatorInterface
	ulid                  ulid.CustomULIDInterface
}

func NewDoctorAppointmentService(
	doctorAppointmentRepo contracts.DoctorAppointmentRepository,
	validator validator.CustomValidatorInterface,
	ulid ulid.CustomULIDInterface,
) contracts.DoctorAppointmentService {
	return &doctorAppointmentService{
		doctorAppointmentRepo: doctorAppointmentRepo,
		validator:             validator,
		ulid:                  ulid,
	}
}
