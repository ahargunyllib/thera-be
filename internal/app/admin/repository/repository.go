package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type adminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) contracts.AdminRepository {
	return &adminRepository{
		db: db,
	}
}
