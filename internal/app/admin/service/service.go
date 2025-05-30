package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/bcrypt"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type adminService struct {
	adminRepo contracts.AdminRepository
	validator validator.CustomValidatorInterface
	bcrypt    bcrypt.BcryptInterface
	jwt       jwt.CustomJwtInterface
}

func NewAdminService(
	adminRepo contracts.AdminRepository,
	validator validator.CustomValidatorInterface,
	bcrypt bcrypt.BcryptInterface,
	jwt jwt.CustomJwtInterface,
) contracts.AdminService {
	return &adminService{
		adminRepo: adminRepo,
		validator: validator,
		bcrypt:    bcrypt,
		jwt:       jwt,
	}
}
