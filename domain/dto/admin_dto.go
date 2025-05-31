package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type AdminResponse struct {
	ID        uuid.UUID        `json:"id"`
	Email     string           `json:"email"`
	FullName  string           `json:"full_name"`
	Role      string           `json:"role"`
	Hospital  HospitalResponse `json:"hospital"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

func NewAdminResponse(adminEntity *entity.Admin) AdminResponse {
	return AdminResponse{
		ID:        adminEntity.ID,
		Email:     adminEntity.Email,
		FullName:  adminEntity.FullName,
		Role:      enums.AdminRoleMapIdx[adminEntity.Role].LongLabel["id"],
		Hospital:  NewHospitalResponse(&adminEntity.Hospital),
		CreatedAt: adminEntity.CreatedAt,
		UpdatedAt: adminEntity.UpdatedAt,
	}
}

type LoginAdminRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type LoginAdminResponse struct {
	AccessToken string `json:"access_token"`
}

type GetAdminSessionRequest struct {
	AdminID uuid.UUID `validate:"required,uuid"`
}

type GetAdminSessionResponse struct {
	Admin AdminResponse `json:"admin"`
}

// type AdminForgotPasswordRequest struct {
// 	Email string `json:"email" validate:"required,email"`
// }

// type AdminForgotPasswordResponse struct {
// 	OTP string `json:"otp"`
// }
