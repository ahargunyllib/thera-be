package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	openai "github.com/ahargunyllib/thera-be/pkg/opeanai"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type doctorScheduleService struct {
	doctorScheduleRepo contracts.DoctorScheduleRepository
	validator          validator.CustomValidatorInterface
	openai             openai.CustomOpenAIInterface
}

func NewDoctorScheduleService(
	doctorScheduleRepo contracts.DoctorScheduleRepository,
	validator validator.CustomValidatorInterface,
	openai openai.CustomOpenAIInterface,
) contracts.DoctorScheduleService {
	return &doctorScheduleService{
		doctorScheduleRepo: doctorScheduleRepo,
		validator:          validator,
		openai:             openai,
	}
}
