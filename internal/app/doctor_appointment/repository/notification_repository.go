package repository

import (
	"context"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
)

func (d *doctorAppointmentRepository) CreateDoctorAppointmentNotification(
	ctx context.Context,
	notification *entity.Notification,
) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO notifications (id, doctor_id, title, body)
		VALUES (:id, :doctor_id, :title, :body)
	`)

	_, err := d.db.NamedExecContext(ctx, qb.String(), notification)
	if err != nil {
		return err
	}

	return nil
}
