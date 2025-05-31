package repository

import (
	"context"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

func (mr *moodRepository) GetDoctorByID(
	ctx context.Context,
	doctorID uuid.UUID,
) (*entity.Doctor, error) {
	var doctor entity.Doctor

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			full_name,
			hospital_id
		FROM doctors
		WHERE id = $1
	`)

	err := mr.db.GetContext(ctx, &doctor, qb.String(), doctorID)
	if err != nil {
		return nil, err
	}

	return &doctor, nil
}
