package repository

import (
	"context"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

func (d *doctorScheduleRepository) GetLastMoodByDoctorID(
	ctx context.Context,
	doctorID uuid.UUID,
) (*entity.Mood, error) {
	var mood entity.Mood

	var qb strings.Builder

	qb.WriteString(`
		SELECT scale
		FROM moods
		WHERE doctor_id = $1
		ORDER BY created_at DESC
		LIMIT 1
	`)

	err := d.db.GetContext(ctx, &mood, qb.String(), doctorID)
	if err != nil {
		return nil, err
	}

	return &mood, nil
}
