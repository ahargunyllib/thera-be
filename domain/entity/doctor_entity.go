package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Doctor struct {
	ID          uuid.UUID      `db:"id"`
	FullName    string         `db:"full_name"`
	Email       string         `db:"email"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Specialty   int            `db:"specialty"`
	HospitalID  int            `db:"hospital_id"`
	Hospital    Hospital       `db:"hospital"`
	Password    string         `db:"password"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}
