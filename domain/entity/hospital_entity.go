package entity

import (
	"database/sql"
	"time"

	"github.com/ahargunyllib/thera-be/domain/enums"
)

type Hospital struct {
	ID              int                     `db:"id"`
	Name            string                  `db:"name"`
	Address         string                  `db:"address"`
	Phone           sql.NullString          `db:"phone"`
	Email           sql.NullString          `db:"email"`
	Website         sql.NullString          `db:"website"`
	Type            enums.HospitalTypeIdx   `db:"type"`
	Status          enums.HospitalStatusIdx `db:"status"`
	Latitude        float64                 `db:"latitude"`
	Longitude       float64                 `db:"longitude"`
	YearEstablished int                     `db:"year_established"`
	CreatedAt       time.Time               `db:"created_at"`
	UpdatedAt       time.Time               `db:"updated_at"`
}
