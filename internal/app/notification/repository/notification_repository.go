package repository

import (
	"context"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

func (n *notificationRepository) GetNotificationsByDoctorID(
	ctx context.Context,
	doctorID uuid.UUID,
) ([]entity.Notification, error) {
	var notifications []entity.Notification
	var qb strings.Builder

	qb.WriteString(`
		SELECT id, title, body, type, read_at, created_at
		FROM notifications
		WHERE doctor_id = $1
	`)

	err := n.db.SelectContext(ctx, &notifications, qb.String(), doctorID)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (n *notificationRepository) GetNotificationsByHospitalID(
	ctx context.Context,
	hospitalID int,
) ([]entity.Notification, error) {
	var notifications []entity.Notification
	var qb strings.Builder

	qb.WriteString(`
		SELECT id, title, body, type, read_at, created_at
		FROM notifications
		WHERE hospital_id = $1
	`)

	err := n.db.SelectContext(ctx, &notifications, qb.String(), hospitalID)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
