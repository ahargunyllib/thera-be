package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/pgerror"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (mr *moodRepository) CreateMood(ctx context.Context, mood *entity.Mood) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO moods (doctor_id, scale)
		VALUES (:doctor_id, :scale)
	`)

	_, err := mr.db.NamedExecContext(ctx, qb.String(), mood)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "unique_user_per_day",
					Err:            errx.ErrHaveAlreadyCreatedMood,
				},
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "moods_doctor_id_fkey",
					Err: errx.ErrDoctorNotFound.WithDetails(map[string]any{
						"doctor_id": mood.DoctorID,
					}),
				},
			}

			if customPgErr := pgerror.HandlePgError(*pgErr, pgErrors); customPgErr != nil {
				return customPgErr
			}
		}

		return err
	}

	return nil
}

func (mr *moodRepository) GetMyDailyMood(ctx context.Context, doctorID uuid.UUID) ([]entity.Mood, error) {
	var moods []entity.Mood

	var qb strings.Builder

	qb.WriteString(`
		SELECT
			id,
			doctor_id,
			scale,
			created_at
		FROM moods
		WHERE doctor_id = $1
			AND created_at >= (DATE_TRUNC('week', NOW())::date - INTERVAL '1 day')
			AND created_at < ((DATE_TRUNC('week', NOW())::date - INTERVAL '1 day') + INTERVAL '7 day')
		ORDER BY created_at DESC;
	`)

	err := mr.db.SelectContext(ctx, &moods, qb.String(), doctorID)
	if err != nil {
		return nil, err
	}

	return moods, nil
}

func (mr *moodRepository) GetMyMonthlyOverview(
	ctx context.Context,
	doctorID uuid.UUID,
) ([]entity.MonthlyMoodStatistic, error) {
	var monthlyMoodStatistics []entity.MonthlyMoodStatistic

	var qb strings.Builder

	qb.WriteString(`
		SELECT
			SUM(scale) AS scale,
			TO_CHAR(created_at, 'YYYY-MM') AS month
		FROM moods
		WHERE doctor_id = $1
			AND EXTRACT(YEAR FROM created_at) = EXTRACT(YEAR FROM NOW())
		GROUP BY month
		ORDER BY month DESC;
	`)

	err := mr.db.SelectContext(ctx, &monthlyMoodStatistics, qb.String(), doctorID)
	if err != nil {
		return nil, err
	}

	return monthlyMoodStatistics, nil
}
