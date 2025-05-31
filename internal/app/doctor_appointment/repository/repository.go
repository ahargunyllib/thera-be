package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type doctorAppointmentRepository struct {
	db *sqlx.DB
}

func NewDoctorAppointmentRepository(db *sqlx.DB) contracts.DoctorAppointmentRepository {
	return &doctorAppointmentRepository{
		db: db,
	}
}
