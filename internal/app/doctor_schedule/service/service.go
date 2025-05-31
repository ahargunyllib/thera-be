package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type doctorScheduleService struct {
	doctorScheduleRepo contracts.DoctorScheduleRepository
	validator          validator.CustomValidatorInterface
}

func NewDoctorScheduleService(
	doctorScheduleRepo contracts.DoctorScheduleRepository,
	validator validator.CustomValidatorInterface,
) contracts.DoctorScheduleService {
	return &doctorScheduleService{
		doctorScheduleRepo: doctorScheduleRepo,
		validator:          validator,
	}
}
