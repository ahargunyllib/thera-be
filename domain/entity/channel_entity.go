package entity

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	DoctorID  uuid.UUID `db:"doctor_id"`
	Doctor    Doctor    `db:"doctor"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
