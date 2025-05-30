package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type hospitalService struct {
	hospitalRepo contracts.HospitalRepository
	validator    validator.CustomValidatorInterface
}

func NewHospitalService(
	hospitalRepo contracts.HospitalRepository,
	validator validator.CustomValidatorInterface,
) contracts.HospitalService {
	return &hospitalService{
		hospitalRepo: hospitalRepo,
		validator:    validator,
	}
}
