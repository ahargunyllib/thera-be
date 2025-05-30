package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type patientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(
	db *sqlx.DB,
) contracts.PatientRepository {
	return &patientRepository{
		db: db,
	}
}
