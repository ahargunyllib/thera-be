package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID                  uuid.UUID      `db:"id"`
	FullName            string         `db:"full_name"`
	IDNumber            string         `db:"id_number"`
	PhoneNumber         sql.NullString `db:"phone_number"`
	Address             sql.NullString `db:"address"`
	DateOfBirth         sql.NullTime   `db:"date_of_birth"`
	Gender              int            `db:"gender"`
	Height              float64        `db:"height"`
	Weight              float64        `db:"weight"`
	BloodType           int            `db:"blood_type"`
	Allergies           sql.NullString `db:"allergies"`
	MedicalRecordNumber string         `db:"medical_record_number"`
	HospitalID          int            `db:"hospital_id"`
	Hospital            Hospital       `db:"hospital"`
	CreatedAt           time.Time      `db:"created_at"`
	UpdatedAt           time.Time      `db:"updated_at"`
}
