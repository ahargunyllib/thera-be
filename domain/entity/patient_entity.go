package entity

import (
	"database/sql"
	"time"

	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type Patient struct {
	ID                  uuid.UUID          `db:"id"`
	FullName            string             `db:"full_name"`
	IDNumber            string             `db:"id_number"`
	PhoneNumber         sql.NullString     `db:"phone_number"`
	Address             string             `db:"address"`
	DateOfBirth         time.Time          `db:"date_of_birth"`
	Gender              enums.GenderIdx    `db:"gender"`
	Height              float64            `db:"height"`
	Weight              float64            `db:"weight"`
	BloodType           enums.BloodTypeIdx `db:"blood_type"`
	Allergies           sql.NullString     `db:"allergies"`
	MedicalRecordNumber string             `db:"medical_record_number"`
	HospitalID          int                `db:"hospital_id"`
	Hospital            Hospital           `db:"hospital"`
	CreatedAt           time.Time          `db:"created_at"`
	UpdatedAt           time.Time          `db:"updated_at"`
}
