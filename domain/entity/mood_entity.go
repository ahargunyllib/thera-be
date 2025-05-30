package entity

import (
	"time"

	"github.com/google/uuid"
)

type Mood struct {
	ID        int       `db:"id"`
	DoctorID  uuid.UUID `db:"doctor_id"`
	Scale     int       `db:"scale"`
	CreatedAt time.Time `db:"created_at"`
}

type MonthlyMoodStatistic struct {
	Scale int    `db:"scale"`
	Month string `db:"month"`
}
