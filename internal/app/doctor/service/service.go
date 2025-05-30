package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/bcrypt"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type doctorService struct {
	doctorRepo contracts.DoctorRepository
	validator  validator.CustomValidatorInterface
	bcrypt     bcrypt.BcryptInterface
	jwt        jwt.CustomJwtInterface
}

func NewDoctorService(
	doctorRepo contracts.DoctorRepository,
	validator validator.CustomValidatorInterface,
	bcrypt bcrypt.BcryptInterface,
	jwt jwt.CustomJwtInterface,
) contracts.DoctorService {
	return &doctorService{
		doctorRepo: doctorRepo,
		validator:  validator,
		bcrypt:     bcrypt,
		jwt:        jwt,
	}
}
