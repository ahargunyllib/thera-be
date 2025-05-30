package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/pgerror"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (ar *adminRepository) CreateAdmin(ctx context.Context, admin *entity.Admin) error {
	var qb strings.Builder
	qb.WriteString(`
		INSERT INTO admins
			(id, full_name, email, password, role, hospital_id)
		VALUES
			(:id, :full_name, :email, :password, :role, :hospital_id)
	`)

	_, err := ar.db.NamedExecContext(ctx, qb.String(), admin)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "admins_email_key",
					Err: errx.ErrAdminAlreadyExists.WithDetails(map[string]any{
						"email": admin.Email,
					}).WithLocation("repository.admin.CreateAdmin"),
				},
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "admins_hospital_id_fkey",
					Err: errx.ErrHospitalNotFound.WithDetails(map[string]any{
						"hospital_id": admin.HospitalID,
					}).WithLocation("repository.admin.CreateAdmin"),
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

func (ar *adminRepository) GetAdminByEmail(ctx context.Context, email string) (*entity.Admin, error) {
	var admin entity.Admin

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			id,
			full_name,
			email,
			password,
			role,
			hospital_id,
			created_at,
			updated_at
		FROM admins
		WHERE email = $1
	`)

	err := ar.db.GetContext(ctx, &admin, qb.String(), email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrAdminNotFound.WithDetails(map[string]any{
				"email": email,
			}).WithLocation("repository.admin.GetAdminByEmail")
		}

		return nil, err
	}

	return &admin, nil
}

func (ar *adminRepository) GetAdminByID(ctx context.Context, id uuid.UUID) (*entity.Admin, error) {
	var admin entity.Admin

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			admins.id,
			admins.full_name,
			admins.email,
			admins.password,
			admins.role,
			admins.hospital_id,
			admins.created_at,
			admins.updated_at,
			hospitals.id AS "hospital.id",
			hospitals.name AS "hospital.name",
			hospitals.address AS "hospital.address",
			hospitals.phone AS "hospital.phone",
			hospitals.email AS "hospital.email",
			hospitals.website AS "hospital.website",
			hospitals.latitude AS "hospital.latitude",
			hospitals.longitude AS "hospital.longitude",
			hospitals.year_established AS "hospital.year_established",
			hospitals.created_at AS "hospital.created_at",
			hospitals.updated_at AS "hospital.updated_at"
		FROM admins
		JOIN hospitals ON admins.hospital_id = hospitals.id
		WHERE admins.id = $1
	`)

	err := ar.db.GetContext(ctx, &admin, qb.String(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrAdminNotFound.WithDetails(map[string]any{
				"id": id,
			}).WithLocation("repository.admin.GetAdminByID")
		}

		return nil, err
	}

	return &admin, nil
}
