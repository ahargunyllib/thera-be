package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers"
	"github.com/ahargunyllib/thera-be/pkg/helpers/pgerror"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (d *doctorScheduleRepository) CreateDoctorSchedule(ctx context.Context, schedule *entity.DoctorSchedule) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO doctor_schedules (doctor_id, start_time, end_time, day_of_week)
		VALUES (:doctor_id, :start_time, :end_time, :day_of_week)
	`)

	_, err := d.db.NamedExecContext(ctx, qb.String(), schedule)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "doctor_schedules_doctor_id_fkey",
					Err: errx.ErrDoctorNotFound.WithDetails(map[string]any{
						"doctor_id": schedule.DoctorID,
					}).WithLocation("repository.doctor_schedule.CreateDoctorSchedule"),
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

func (d *doctorScheduleRepository) DeleteDoctorSchedule(ctx context.Context, scheduleID int) error {
	var qb strings.Builder

	qb.WriteString(`
		DELETE FROM doctor_schedules
		WHERE id = $1
	`)

	res, err := d.db.ExecContext(ctx, qb.String(), scheduleID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	return helpers.CheckRowsAffected(
		rowsAffected,
		errx.ErrDoctorScheduleNotFound.WithDetails(map[string]any{
			"id": scheduleID,
		}).WithLocation("repository.doctor_schedule.DeleteDoctorSchedule"),
	)
}

func (d *doctorScheduleRepository) GetDoctorSchedules(
	ctx context.Context,
	query *dto.GetDoctorSchedulesQuery,
) ([]entity.DoctorSchedule, error) {
	var schedules []entity.DoctorSchedule

	var qb strings.Builder
	var args []any

	qb.WriteString(`
		SELECT
			doctor_schedules.id,
			doctor_schedules.doctor_id,
			doctor_schedules.start_time,
			doctor_schedules.end_time,
			doctor_schedules.day_of_week,
			doctors.id AS "doctor.id",
			doctors.full_name AS "doctor.full_name",
			doctors.email AS "doctor.email",
			doctors.phone_number AS "doctor.phone_number",
			doctors.specialty AS "doctor.specialty",
			doctors.hospital_id AS "doctor.hospital_id",
			doctors.created_at AS "doctor.created_at",
			doctors.updated_at AS "doctor.updated_at"
		FROM doctor_schedules
		JOIN doctors ON doctor_schedules.doctor_id = doctors.id
		WHERE 1=1
	`)

	if query.DoctorID != uuid.Nil {
		qb.WriteString(fmt.Sprintf(" AND doctor_id = $%d", len(args)+1))
		args = append(args, query.DoctorID)
	}

	if query.DayOfWeek != 0 {
		qb.WriteString(fmt.Sprintf(" AND day_of_week = $%d", len(args)+1))
		args = append(args, query.DayOfWeek)
	}

	err := d.db.SelectContext(ctx, &schedules, qb.String(), args...)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (d *doctorScheduleRepository) GetDoctorScheduleByID(
	ctx context.Context,
	scheduleID int,
) (*entity.DoctorSchedule, error) {
	var schedule entity.DoctorSchedule

	var qb strings.Builder

	qb.WriteString(`
		SELECT
			doctor_schedules.id,
			doctor_schedules.doctor_id,
			doctor_schedules.start_time,
			doctor_schedules.end_time,
			doctor_schedules.day_of_week
		FROM doctor_schedules
		WHERE doctor_schedules.id = $1
	`)

	err := d.db.GetContext(ctx, &schedule, qb.String(), scheduleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrDoctorScheduleNotFound.WithDetails(map[string]any{
				"id": scheduleID,
			}).WithLocation("repository.doctor_schedule.GetDoctorScheduleByID")
		}

		return nil, err
	}

	return &schedule, nil
}

func (d *doctorScheduleRepository) UpdateDoctorSchedule(ctx context.Context, schedule *entity.DoctorSchedule) error {
	var qb strings.Builder

	qb.WriteString(`
		UPDATE doctor_schedules
		SET start_time = :start_time, end_time = :end_time, day_of_week = :day_of_week, doctor_id = :doctor_id
		WHERE id = :id
	`)

	_, err := d.db.NamedExecContext(ctx, qb.String(), schedule)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "doctor_schedules_doctor_id_fkey",
					Err: errx.ErrDoctorNotFound.WithDetails(map[string]any{
						"doctor_id": schedule.DoctorID,
					}).WithLocation("repository.doctor_schedule.UpdateDoctorSchedule"),
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

func (d *doctorScheduleRepository) GetNextScheduleByDoctorID(
	ctx context.Context,
	doctorID uuid.UUID,
) (*entity.DoctorSchedule, error) {
	var schedule entity.DoctorSchedule

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			id,
			doctor_id,
			start_time,
			end_time,
			day_of_week
		FROM doctor_schedules
		WHERE 1=1
		AND doctor_id = $1
		OR day_of_week > EXTRACT(DOW FROM NOW())
		ORDER BY day_of_week ASC, start_time ASC
		LIMIT 1
	`)
	err := d.db.GetContext(ctx, &schedule, qb.String(), doctorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrDoctorScheduleNotFound
		}
		return nil, err
	}

	return &schedule, nil
}
