package entity

import (
	"database/sql"
	"time"

	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type Doctor struct {
	ID          uuid.UUID                `db:"id"`
	FullName    string                   `db:"full_name"`
	Email       string                   `db:"email"`
	PhoneNumber sql.NullString           `db:"phone_number"`
	Specialty   enums.DoctorSpecialtyIdx `db:"specialty"`
	HospitalID  int                      `db:"hospital_id"`
	Hospital    Hospital                 `db:"hospital"`
	Password    string                   `db:"password"`
	CreatedAt   time.Time                `db:"created_at"`
	UpdatedAt   time.Time                `db:"updated_at"`
}
