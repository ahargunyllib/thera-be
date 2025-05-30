package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers"
	"github.com/ahargunyllib/thera-be/pkg/helpers/pgerror"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (p *patientRepository) CountPatients(
	ctx context.Context,
	query *dto.GetPatientsQuery,
) (int64, error) {
	var count int64

	var qb strings.Builder
	var args []any

	qb.WriteString(`
		SELECT COUNT(*)
		FROM patients
		WHERE 1=1
	`)

	err := p.db.GetContext(ctx, &count, qb.String(), args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *patientRepository) CreatePatient(ctx context.Context, patient *entity.Patient) error {
	var qb strings.Builder
	qb.WriteString(`
		INSERT INTO patients
			(id, full_name, id_number, phone_number, address,
				date_of_birth, gender, height, weight, blood_type, allergies, medical_record_number, hospital_id)
		VALUES
			(:id, :full_name, :id_number, :phone_number, :address,
				:date_of_birth, :gender, :height, :weight, :blood_type, :allergies, :medical_record_number, :hospital_id)
	`)

	_, err := p.db.NamedExecContext(ctx, qb.String(), patient)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "patients_id_number_key",
					Err:            errx.ErrIDNumberPatientAlreadyExists,
				},
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "patients_medical_record_number_key",
					Err:            errx.ErrMedicalRecordNumberPatientAlreadyExists,
				},
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "patients_hospital_id_fkey",
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

func (p *patientRepository) DeletePatientByID(ctx context.Context, id uuid.UUID) error {
	var qb strings.Builder
	qb.WriteString(`
		DELETE FROM patients
		WHERE id = $1
	`)

	res, err := p.db.ExecContext(ctx, qb.String(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	return helpers.CheckRowsAffected(rowsAffected, errx.ErrPatientNotFound)
}

func (p *patientRepository) GetPatientByID(ctx context.Context, id uuid.UUID) (*entity.Patient, error) {
	var patient entity.Patient

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			patients.id,
			patients.full_name,
			patients.id_number,
			patients.phone_number,
			patients.address,
			patients.date_of_birth,
			patients.gender,
			patients.height,
			patients.weight,
			patients.blood_type,
			patients.allergies,
			patients.medical_record_number,
			patients.hospital_id,
			patients.created_at,
			patients.updated_at,
			hospital.id AS "hospital.id",
			hospital.name AS "hospital.name",
			hospital.address AS "hospital.address",
			hospital.phone AS "hospital.phone",
			hospital.email AS "hospital.email",
			hospital.website AS "hospital.website",
			hospital.latitude AS "hospital.latitude",
			hospital.longitude AS "hospital.longitude",
			hospital.year_established AS "hospital.year_established"
		FROM patients
		JOIN hospitals AS hospital ON patients.hospital_id = hospital.id
		WHERE patients.id = $1
	`)

	err := p.db.GetContext(ctx, &patient, qb.String(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrPatientNotFound
		}
		return nil, err
	}

	return &patient, nil
}

func (p *patientRepository) GetPatients(ctx context.Context, query *dto.GetPatientsQuery) ([]entity.Patient, error) {
	var patients []entity.Patient

	var qb strings.Builder
	var args []any

	qb.WriteString(`
		SELECT
			patients.id,
			patients.full_name,
			patients.id_number,
			patients.phone_number,
			patients.address,
			patients.date_of_birth,
			patients.gender,
			patients.height,
			patients.weight,
			patients.blood_type,
			patients.allergies,
			patients.medical_record_number,
			patients.hospital_id,
			patients.created_at,
			patients.updated_at,
			hospital.id AS "hospital.id",
			hospital.name AS "hospital.name",
			hospital.address AS "hospital.address",
			hospital.phone AS "hospital.phone",
			hospital.email AS "hospital.email",
			hospital.website AS "hospital.website",
			hospital.latitude AS "hospital.latitude",
			hospital.longitude AS "hospital.longitude",
			hospital.year_established AS "hospital.year_established"
		FROM patients
		JOIN hospitals AS hospital ON patients.hospital_id = hospital.id
		WHERE 1=1
	`)

	normalizeQuery(query)
	appendPagination(&qb, &args, query.SortBy, query.SortOrder, query.Limit, query.Page)

	err := p.db.SelectContext(ctx, &patients, qb.String(), args...)
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (p *patientRepository) UpdatePatient(ctx context.Context, patient *entity.Patient) error {
	var qb strings.Builder
	qb.WriteString(`
		UPDATE patients
		SET
			full_name = :full_name,
			id_number = :id_number,
			phone_number = :phone_number,
			address = :address,
			date_of_birth = :date_of_birth,
			gender = :gender,
			height = :height,
			weight = :weight,
			blood_type = :blood_type,
			allergies = :allergies,
			medical_record_number = :medical_record_number,
			hospital_id = :hospital_id
		WHERE id = :id
	`)

	res, err := p.db.NamedExecContext(ctx, qb.String(), patient)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	return helpers.CheckRowsAffected(rowsAffected, errx.ErrPatientNotFound)
}
