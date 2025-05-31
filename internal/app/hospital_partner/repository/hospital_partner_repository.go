package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
)

func (h *hospitalPartnerRepository) GetHospitalPartnerByID(ctx context.Context, id string) (
	*entity.HospitalPartner,
	error,
) {
	var hospitalPartner entity.HospitalPartner

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			hospital_partners.id,
			hospital_partners.from_hospital_id,
			hospital_partners.to_hospital_id,
			hospital_partners.partner_type,
			hospital_partners.status,
			hospital_partners.created_at,
			hospital_partners.updated_at,
			from_hospital.id AS "from_hospital.id",
			from_hospital.name AS "from_hospital.name",
			to_hospital.id AS "to_hospital.id",
			to_hospital.name AS "to_hospital.name"
		FROM
			hospital_partners
		JOIN hospitals AS from_hospital ON from_hospital.id = hospital_partners.from_hospital_id
		JOIN hospitals AS to_hospital ON to_hospital.id = hospital_partners.to_hospital_id
		WHERE
			id = $1
	`)

	err := h.db.GetContext(ctx, &hospitalPartner, qb.String(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrHospitalPartnerNotFound
		}

		return nil, err
	}

	return &hospitalPartner, nil
}

func (h *hospitalPartnerRepository) CreateHospitalPartner(
	ctx context.Context,
	hospitalPartner *entity.HospitalPartner,
) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO hospital_partners
			(id, from_hospital_id, to_hospital_id, partner_type, status)
		VALUES
			(:id, :from_hospital_id, :to_hospital_id, :partner_type, :status)
	`)
	_, err := h.db.NamedExecContext(ctx, qb.String(), hospitalPartner)
	if err != nil {
		return err
	}

	return nil
}

func (h *hospitalPartnerRepository) GetHospitalPartnersByHospitalID(
	ctx context.Context,
	hospitalID int,
) ([]entity.HospitalPartner, error) {
	var hospitalPartners []entity.HospitalPartner

	var qb strings.Builder
	qb.WriteString(`
		SELECT
			hospital_partners.id,
			hospital_partners.from_hospital_id,
			hospital_partners.to_hospital_id,
			hospital_partners.partner_type,
			hospital_partners.status,
			hospital_partners.created_at,
			hospital_partners.updated_at,
			from_hospital.id AS "from_hospital.id",
			from_hospital.name AS "from_hospital.name",
			to_hospital.id AS "to_hospital.id",
			to_hospital.name AS "to_hospital.name"
		FROM
			hospital_partners
		JOIN hospitals AS from_hospital ON from_hospital.id = hospital_partners.from_hospital_id
		JOIN hospitals AS to_hospital ON to_hospital.id = hospital_partners.to_hospital_id
		WHERE
			from_hospital_id = $1 OR to_hospital_id = $1
	`)

	err := h.db.SelectContext(ctx, &hospitalPartners, qb.String(), hospitalID)
	if err != nil {
		return nil, err
	}

	return hospitalPartners, nil
}

func (h *hospitalPartnerRepository) UpdateHospitalPartner(
	ctx context.Context,
	hospitalPartner *entity.HospitalPartner,
) error {
	var qb strings.Builder

	qb.WriteString(`
		UPDATE hospital_partners
		SET
			from_hospital_id = :from_hospital_id,
			to_hospital_id = :to_hospital_id,
			type = :type,
			status = :status,
			updated_at = NOW()
		WHERE
			id = :id
	`)

	_, err := h.db.NamedExecContext(ctx, qb.String(), hospitalPartner)
	if err != nil {
		return err
	}

	return nil
}
