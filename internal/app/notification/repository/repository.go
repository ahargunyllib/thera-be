package repository

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type notificationRepository struct {
	db *sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) contracts.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}
