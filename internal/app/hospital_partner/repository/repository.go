package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type hospitalPartnerRepository struct {
	db *sqlx.DB
}

func NewHospitalPartnerRepository(db *sqlx.DB) contracts.HospitalPartnerRepository {
	return &hospitalPartnerRepository{
		db: db,
	}
}
