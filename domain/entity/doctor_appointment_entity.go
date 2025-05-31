package entity

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type DoctorAppointment struct {
	ID              string                           `db:"id"`        // ulid
	DoctorID        uuid.UUID                        `db:"doctor_id"` // uuid
	Doctor          Doctor                           `db:"doctor"`
	PatientID       uuid.UUID                        `db:"patient_id"` // uuid
	Patient         Patient                          `db:"patient"`
	AppointmentDate string                           `db:"appointment_date"` // date in YYYY-MM-DD format
	StartTime       string                           `db:"start_time"`       // time in HH:MM format
	EndTime         string                           `db:"end_time"`         // time in HH:MM format
	Status          enums.DoctorAppointmentStatusIdx `db:"status"`
	Type            enums.DoctorAppointmentTypeIdx   `db:"type"`
	CreatedAt       time.Time                        `db:"created_at"`
	UpdatedAt       time.Time                        `db:"updated_at"`
}
