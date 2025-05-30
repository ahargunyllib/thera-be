package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type moodRepository struct {
	db *sqlx.DB
}

func NewMoodRepository(db *sqlx.DB) contracts.MoodRepository {
	return &moodRepository{
		db: db,
	}
}
