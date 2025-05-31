package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type DoctorAppointmentResponse struct {
	ID              string          `json:"id"`               // ulid
	Doctor          DoctorResponse  `json:"doctor"`           // Doctor details
	Patient         PatientResponse `json:"patient"`          // Patient details
	AppointmentDate string          `json:"appointment_date"` // date in YYYY-MM-DD format
	StartTime       string          `json:"start_time"`       // time in HH:MM format
	EndTime         string          `json:"end_time"`         // time in HH:MM format
	Status          string          `json:"status"`
	Type            string          `json:"type"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

func NewDoctorAppointmentResponse(doctorAppointmentEntity *entity.DoctorAppointment) DoctorAppointmentResponse {
	return DoctorAppointmentResponse{
		ID:              doctorAppointmentEntity.ID,
		Doctor:          NewDoctorResponse(&doctorAppointmentEntity.Doctor),
		Patient:         NewPatientResponse(&doctorAppointmentEntity.Patient),
		AppointmentDate: doctorAppointmentEntity.AppointmentDate,
		StartTime:       doctorAppointmentEntity.StartTime,
		EndTime:         doctorAppointmentEntity.EndTime,
		Status:          enums.DoctorAppointmentStatusMapIdx[doctorAppointmentEntity.Status].LongLabel["id"],
		Type:            enums.DoctorAppointmentTypeMapIdx[doctorAppointmentEntity.Type].LongLabel["id"],
		CreatedAt:       doctorAppointmentEntity.CreatedAt,
		UpdatedAt:       doctorAppointmentEntity.UpdatedAt,
	}
}

type GetDoctorAppointmentsQuery struct {
	DoctorID  uuid.UUID `valid:"omitempty,uuid"`
	PatientID uuid.UUID `valid:"omitempty,uuid"`
	FromDate  string    `query:"from_date" valid:"omitempty,datetime=2006-01-02 15:04:05"` // YYYY-MM-DD
	ToDate    string    `query:"to_date" valid:"omitempty,datetime=2006-01-02 15:04:05"`   // YYYY-MM-DD
}

type GetDoctorAppointmentsResponse struct {
	DoctorAppointments []DoctorAppointmentResponse `json:"doctor_appointments"` // List of doctor appointments
}

type CreateDoctorAppointmentRequest struct {
	DoctorID        uuid.UUID `json:"doctor_id" valid:"required,uuid"`                       // UUID of the doctor
	PatientID       uuid.UUID `json:"patient_id" valid:"required,uuid"`                      // UUID of the patient
	AppointmentDate string    `json:"appointment_date" valid:"required,datetime=2006-01-02"` // date in YYYY-MM-DD format
	StartTime       string    `json:"start_time" valid:"required,datetime=15:04"`            // time in HH:MM format
	EndTime         string    `json:"end_time" valid:"required,datetime=15:04"`              // time in HH:MM format
	Status          string    `json:"status" valid:"required"`                               // Status of the appointment
	Type            string    `json:"type" valid:"required"`                                 // Type of the appointment
}

type UpdateDoctorAppointmentParams struct {
	ID string `json:"id" valid:"required,ulid"` // ulid of the doctor appointment
}

type UpdateDoctorAppointmentRequest struct {
	DoctorID        uuid.UUID `json:"doctor_id" valid:"required,uuid"`                       // UUID of the doctor
	PatientID       uuid.UUID `json:"patient_id" valid:"required,uuid"`                      // UUID of the patient
	AppointmentDate string    `json:"appointment_date" valid:"required,datetime=2006-01-02"` // date in YYYY-MM-DD format
	StartTime       string    `json:"start_time" valid:"required,datetime=15:04"`            // time in HH:MM format
	EndTime         string    `json:"end_time" valid:"required,datetime=15:04"`              // time in HH:MM format
	Status          string    `json:"status" valid:"required"`                               // Status of the appointment
	Type            string    `json:"type" valid:"required"`                                 // Type of the appointment
}

type DeleteDoctorAppointmentParams struct {
	ID string `json:"id" valid:"required,ulid"` // ulid of the doctor appointment
}
