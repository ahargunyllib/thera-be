package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type PatientResponse struct {
	ID                  uuid.UUID        `json:"id"`
	FullName            string           `json:"full_name"`
	IDNumber            string           `json:"id_number"`
	PhoneNumber         string           `json:"phone_number,omitempty"`
	Address             string           `json:"address,omitempty"`
	DateOfBirth         time.Time        `json:"date_of_birth,omitempty"`
	Gender              string           `json:"gender"`
	Height              float64          `json:"height"`
	Weight              float64          `json:"weight"`
	BloodType           string           `json:"blood_type"`
	Allergies           string           `json:"allergies,omitempty"`
	MedicalRecordNumber string           `json:"medical_record_number"`
	Hospital            HospitalResponse `json:"hospital"`
	CreatedAt           time.Time        `json:"created_at"`
	UpdatedAt           time.Time        `json:"updated_at"`
}

func NewPatientResponse(patientEntity *entity.Patient) PatientResponse {
	return PatientResponse{
		ID:                  patientEntity.ID,
		FullName:            patientEntity.FullName,
		IDNumber:            patientEntity.IDNumber,
		PhoneNumber:         patientEntity.PhoneNumber.String,
		Address:             patientEntity.Address.String,
		DateOfBirth:         patientEntity.DateOfBirth.Time,
		Gender:              "", // TODO
		Height:              patientEntity.Height,
		Weight:              patientEntity.Weight,
		BloodType:           "", // TODO
		Allergies:           patientEntity.Allergies.String,
		MedicalRecordNumber: patientEntity.MedicalRecordNumber,
		Hospital:            NewHospitalResponse(&patientEntity.Hospital),
		CreatedAt:           patientEntity.CreatedAt,
		UpdatedAt:           patientEntity.UpdatedAt,
	}
}

type GetPatientsQuery struct {
	Limit     int    `json:"limit" validate:"required,gte=1"`
	Page      int    `json:"page" validate:"required,gte=1"`
	SortBy    string `json:"sort_by,omitempty" validate:"omitempty,oneof=id full_name created_at"`
	SortOrder string `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"`
}

type GetPatientsResponse struct {
	Patients []PatientResponse `json:"patients"`
	Meta     struct {
		Pagination PaginationResponse `json:"pagination"`
	} `json:"meta"`
}

type GetPatientByIDParams struct {
	ID uuid.UUID `path:"id" validate:"required,uuid"`
}

type GetPatientResponse struct {
	Patient PatientResponse `json:"patient" validate:"required"`
}

type CreatePatientRequest struct {
	FullName            string    `json:"full_name" validate:"required"`
	IDNumber            string    `json:"id_number" validate:"required"`
	PhoneNumber         string    `json:"phone_number,omitempty" validate:"omitempty,phone"`
	Address             string    `json:"address,omitempty"`
	DateOfBirth         time.Time `json:"date_of_birth,omitempty" validate:"omitempty,date"`
	Gender              string    `json:"gender" validate:"required,oneof=male female"`
	Height              float64   `json:"height" validate:"required,gte=0"`
	Weight              float64   `json:"weight" validate:"required,gte=0"`
	BloodType           string    `json:"blood_type" validate:"required,oneof=a b ab o"`
	Allergies           string    `json:"allergies,omitempty" validate:"omitempty"`
	MedicalRecordNumber string    `json:"medical_record_number" validate:"required"`
	HospitalID          int       `json:"hospital_id" validate:"required"`
}

type UpdatePatientRequest struct {
	FullName            string    `json:"full_name,omitempty" validate:"omitempty"`
	IDNumber            string    `json:"id_number,omitempty" validate:"omitempty"`
	PhoneNumber         string    `json:"phone_number,omitempty" validate:"omitempty,phone"`
	Address             string    `json:"address,omitempty"`
	DateOfBirth         time.Time `json:"date_of_birth,omitempty" validate:"omitempty,date"`
	Gender              string    `json:"gender" validate:"omitempty,oneof=male female"`
	Height              float64   `json:"height,omitempty" validate:"omitempty,gte=0"`
	Weight              float64   `json:"weight,omitempty" validate:"omitempty,gte=0"`
	BloodType           string    `json:"blood_type,omitempty" validate:"omitempty,oneof=a b ab o"`
	Allergies           string    `json:"allergies,omitempty" validate:"omitempty"`
	MedicalRecordNumber string    `json:"medical_record_number,omitempty" validate:"omitempty"`
	HospitalID          int       `json:"hospital_id,omitempty" validate:"omitempty"`
}

type UpdatePatientByIDParams struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}

type DeletePatientByIDParams struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
