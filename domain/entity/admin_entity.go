package entity

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID         uuid.UUID `db:"id"`
	Email      string    `db:"email"`
	FullName   string    `db:"full_name"`
	Password   string    `db:"password"`
	Role       int       `db:"role"`
	HospitalID int       `db:"hospital_id"`
	Hospital   Hospital  `db:"hospital"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
