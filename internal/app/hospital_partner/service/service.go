package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/ulid"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type hospitalPartnerService struct {
	hospitalPartnerRepo contracts.HospitalPartnerRepository
	validator           validator.CustomValidatorInterface
	ulid                ulid.CustomULIDInterface
}

func NewHospitalPartnerService(
	hospitalPartnerRepo contracts.HospitalPartnerRepository,
	validator validator.CustomValidatorInterface,
	ulid ulid.CustomULIDInterface,
) contracts.HospitalPartnerService {
	return &hospitalPartnerService{
		hospitalPartnerRepo: hospitalPartnerRepo,
		validator:           validator,
		ulid:                ulid,
	}
}
