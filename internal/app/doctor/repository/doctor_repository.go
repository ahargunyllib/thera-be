package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/pgerror"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (d *doctorRepository) GetDoctorByEmail(ctx context.Context, email string) (*entity.Doctor, error) {
	var doctor entity.Doctor

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			doctors.id,
			doctors.full_name,
			doctors.email,
			doctors.phone_number,
			doctors.password,
			doctors.specialty,
			doctors.hospital_id,
			doctors.created_at,
			doctors.updated_at,
			hospitals.id AS "hospital.id",
			hospitals.name AS "hospital.name",
			hospitals.address AS "hospital.address",
			hospitals.phone AS "hospital.phone",
			hospitals.email AS "hospital.email",
			hospitals.website AS "hospital.website",
			hospitals.latitude AS "hospital.latitude",
			hospitals.longitude AS "hospital.longitude",
			hospitals.year_established AS "hospital.year_established"
		FROM doctors
		JOIN hospitals ON doctors.hospital_id = hospitals.id
		WHERE doctors.email = $1
	`)

	err := d.db.GetContext(ctx, &doctor, qb.String(), email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrDoctorNotFound
		}

		return nil, err
	}

	return &doctor, nil
}

func (d *doctorRepository) CountDoctors(ctx context.Context, query *dto.GetDoctorsQuery) (int64, error) {
	var count int64

	var qb strings.Builder
	var args []any

	qb.WriteString(`
		SELECT COUNT(*)
		FROM doctors
		WHERE 1=1
	`)

	err := d.db.GetContext(ctx, &count, qb.String(), args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *doctorRepository) CreateDoctor(ctx context.Context, doctor *entity.Doctor) error {
	var qb strings.Builder
	qb.WriteString(`
		INSERT INTO doctors
			(id, full_name, email, password, phone_number, specialty, hospital_id)
		VALUES (:id, :full_name, :email, :password, :phone_number, :specialty, :hospital_id)
	`)

	_, err := d.db.NamedExecContext(ctx, qb.String(), doctor)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "doctors_email_key",
					Err:            errx.ErrDoctorAlreadyExists,
				},
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "doctors_hospital_id_fkey",
					Err:            errx.ErrHospitalNotFound,
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

func (d *doctorRepository) GetDoctorByID(ctx context.Context, id uuid.UUID) (*entity.Doctor, error) {
	var doctor entity.Doctor

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			doctors.id,
			doctors.full_name,
			doctors.email,
			doctors.phone_number,
			doctors.specialty,
			doctors.hospital_id,
			doctors.created_at,
			doctors.updated_at,
			hospitals.id AS "hospital.id",
			hospitals.name AS "hospital.name",
			hospitals.address AS "hospital.address",
			hospitals.phone AS "hospital.phone",
			hospitals.email AS "hospital.email",
			hospitals.website AS "hospital.website",
			hospitals.latitude AS "hospital.latitude",
			hospitals.longitude AS "hospital.longitude",
			hospitals.year_established AS "hospital.year_established"
		FROM doctors
		JOIN hospitals ON doctors.hospital_id = hospitals.id
		WHERE doctors.id = $1
	`)

	err := d.db.GetContext(ctx, &doctor, qb.String(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrDoctorNotFound
		}

		return nil, err
	}

	return &doctor, nil
}

func (d *doctorRepository) GetDoctors(ctx context.Context, query *dto.GetDoctorsQuery) ([]entity.Doctor, error) {
	var doctors []entity.Doctor

	var qb strings.Builder
	var args []any

	qb.WriteString(`
		SELECT
			doctors.id,
			doctors.full_name,
			doctors.email,
			doctors.phone_number,
			doctors.specialty,
			doctors.hospital_id,
			doctors.created_at,
			doctors.updated_at,
			hospitals.id AS "hospital.id",
			hospitals.name AS "hospital.name",
			hospitals.address AS "hospital.address",
			hospitals.phone AS "hospital.phone",
			hospitals.email AS "hospital.email",
			hospitals.website AS "hospital.website",
			hospitals.latitude AS "hospital.latitude",
			hospitals.longitude AS "hospital.longitude",
			hospitals.year_established AS "hospital.year_established"
		FROM doctors
		JOIN hospitals ON doctors.hospital_id = hospitals.id
		WHERE 1=1
	`)

	normalizeQuery(query)
	appendPagination(&qb, &args, query.SortBy, query.SortOrder, query.Limit, query.Page)

	err := d.db.SelectContext(ctx, &doctors, qb.String(), args...)
	if err != nil {
		return nil, err
	}

	return doctors, nil
}
