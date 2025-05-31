package repository

import (
	"context"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
)

func (mr *moodRepository) CreateMoodNotification(
	ctx context.Context,
	notification *entity.Notification,
) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO notifications (id, doctor_id, title, body, type)
		VALUES (:id, :doctor_id, :title, :body, :type)
	`)

	_, err := mr.db.NamedExecContext(ctx, qb.String(), notification)
	if err != nil {
		return err
	}

	return nil
}
