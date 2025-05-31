package entity

import (
	"database/sql"
	"time"

	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type Notification struct {
	ID         string                    `db:"id"` // ulid
	DoctorID   uuid.NullUUID             `db:"doctor_id"`
	HospitalID sql.NullInt64             `db:"hospital_id"`
	Title      string                    `db:"title"`
	Body       sql.NullString            `db:"body"`
	Type       enums.NotificationTypeIdx `db:"type"`
	ReadAt     sql.NullTime              `db:"read_at"` // Nullable time for read status
	Metadata   map[string]any            `db:"metadata"`
	CreatedAt  time.Time                 `db:"created_at"`
}
