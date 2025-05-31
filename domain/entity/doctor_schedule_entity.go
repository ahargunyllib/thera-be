package entity

import (
	"time"

	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ID        int       `db:"id"`
	DoctorID  uuid.UUID `db:"doctor_id"`
	Doctor    Doctor    `db:"doctor"`
	DayOfWeek int       `db:"day_of_week"`
	StartTime string    `db:"start_time"` // time.TimeOnly
	EndTime   string    `db:"end_time"`   // time.TimeOnly
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
