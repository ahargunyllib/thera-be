package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/uuid"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type patientService struct {
	patientRepo contracts.PatientRepository
	validator   validator.CustomValidatorInterface
	uuid        uuid.UUIDInterface
}

func NewPatientService(
	patientRepo contracts.PatientRepository,
	validator validator.CustomValidatorInterface,
	uuid uuid.UUIDInterface,
) contracts.PatientService {
	return &patientService{
		patientRepo: patientRepo,
		validator:   validator,
		uuid:        uuid,
	}
}
