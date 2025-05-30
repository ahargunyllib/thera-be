package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type DoctorResponse struct {
	ID          uuid.UUID        `json:"id"`
	FullName    string           `json:"full_name"`
	Email       string           `json:"email"`
	PhoneNumber string           `json:"phone_number,omitempty"`
	Specialty   string           `json:"specialty"`
	Hospital    HospitalResponse `json:"hospital"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

func NewDoctorResponse(doctorEntity *entity.Doctor) DoctorResponse {
	return DoctorResponse{
		ID:          doctorEntity.ID,
		FullName:    doctorEntity.FullName,
		Email:       doctorEntity.Email,
		PhoneNumber: doctorEntity.PhoneNumber.String,
		Specialty:   "",
		Hospital:    NewHospitalResponse(&doctorEntity.Hospital),
		CreatedAt:   doctorEntity.CreatedAt,
		UpdatedAt:   doctorEntity.UpdatedAt,
	}
}

type LoginDoctorRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type LoginDoctorResponse struct {
	AccessToken string `json:"access_token"`
}

type GetDoctorSessionRequest struct {
	DoctorID uuid.UUID `validate:"required,uuid"`
}

type GetDoctorSessionResponse struct {
	Doctor DoctorResponse `json:"doctor" validate:"required"`
}

// type DoctorForgotPasswordRequest struct {
// 	Email string `json:"email" validate:"required,email"`
// }

// type DoctorForgotPasswordResponse struct {
// 	OTP string `json:"otp"`
// }

type GetDoctorsQuery struct {
	Page      int    `query:"page" validate:"omitempty,min=1"`
	Limit     int    `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    string `query:"sort_by" validate:"omitempty,oneof=full_name email specialty id created_at"`
	SortOrder string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

type GetDoctorsResponse struct {
	Doctors []DoctorResponse `json:"doctors"`
	Meta    struct {
		Pagination PaginationResponse `json:"pagination"`
	} `json:"meta"`
}
