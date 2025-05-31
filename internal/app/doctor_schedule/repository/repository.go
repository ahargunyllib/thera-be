package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type doctorScheduleRepository struct {
	db *sqlx.DB
}

func NewDoctorScheduleRepository(db *sqlx.DB) contracts.DoctorScheduleRepository {
	return &doctorScheduleRepository{
		db: db,
	}
}
