package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type hospitalRepository struct {
	db *sqlx.DB
}

func NewHospitalRepository(db *sqlx.DB) contracts.HospitalRepository {
	return &hospitalRepository{
		db: db,
	}
}
