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
	"github.com/jackc/pgx/v5/pgconn"
)

func (hr *hospitalRepository) GetHospitals(
	ctx context.Context,
	query *dto.GetHospitalsQuery,
) ([]entity.Hospital, error) {
	var hospitals []entity.Hospital

	var qb strings.Builder
	var args []any

	qb.WriteString(`
		SELECT
			id,
			name,
			address,
			phone,
			email,
			website,
			latitude,
			longitude,
			year_established,
			created_at,
			updated_at
		FROM hospitals WHERE 1=1
	`)
	normalizeQuery(query)
	appendPagination(&qb, &args, query.SortBy, query.SortOrder, query.Limit, query.Page)

	err := hr.db.SelectContext(ctx, &hospitals, qb.String(), args...)
	if err != nil {
		return nil, err
	}

	return hospitals, nil
}

func (hr *hospitalRepository) CountHospitals(
	ctx context.Context,
	query *dto.GetHospitalsQuery,
) (int64, error) {
	var count int64

	var qb strings.Builder
	qb.WriteString(`
		SELECT COUNT(*) FROM hospitals
		WHERE 1=1
	`)

	err := hr.db.GetContext(ctx, &count, qb.String())
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (hr *hospitalRepository) GetHospitalByID(ctx context.Context, id int) (*entity.Hospital, error) {
	var hospital entity.Hospital

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			id,
			name,
			address,
			phone,
			email,
			website,
			latitude,
			longitude,
			year_established,
			created_at,
			updated_at
		FROM hospitals
		WHERE id = $1
	`)

	err := hr.db.GetContext(ctx, &hospital, qb.String(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrHospitalNotFound.
				WithDetails(map[string]any{
					"id": id,
				}).
				WithLocation("repository.hospital.GetHospitalByID")
		}

		return nil, err
	}

	return &hospital, nil
}

func (hr *hospitalRepository) CreateHospital(ctx context.Context, hospital *entity.Hospital) error {
	var qb strings.Builder
	qb.WriteString(`
		INSERT INTO hospitals
			(name, address, phone, email, website, latitude, longitude, year_established)
		VALUES
			(:name, :address, :phone, :email, :website, :latitude, :longitude, :year_established)
	`)

	_, err := hr.db.NamedExecContext(ctx, qb.String(), hospital)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{}

			if customPgErr := pgerror.HandlePgError(*pgErr, pgErrors); customPgErr != nil {
				return customPgErr
			}
		}
		return err
	}

	return nil
}
