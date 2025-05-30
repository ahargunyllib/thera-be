package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type doctorRepository struct {
	db *sqlx.DB
}

func NewDoctorRepository(db *sqlx.DB) contracts.DoctorRepository {
	return &doctorRepository{
		db: db,
	}
}
