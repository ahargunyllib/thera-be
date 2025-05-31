package repository

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

func (d *doctorAppointmentRepository) CheckDoctorScheduleAvailability(
	ctx context.Context, doctorID uuid.UUID, startTime string, endTime string,
) (bool, error) {
	var qb strings.Builder

	qb.WriteString(`
		SELECT COUNT(*) > 0
		FROM doctor_schedules
		WHERE doctor_id = $1
		AND start_time <= $2
		AND end_time >= $3
	`)
	var exists bool

	err := d.db.GetContext(ctx, &exists, qb.String(), doctorID, startTime, endTime)
	if err != nil {
		return false, err
	}

	return exists, nil
}
