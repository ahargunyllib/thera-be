package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/google/uuid"
)

func (c *chatBotRepository) CreateChannel(ctx context.Context, channel *entity.Channel) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO channels
			(id, name, doctor_id)
		VALUES
			(:id, :name, :doctor_id)
	`)

	var err error
	if c.tx != nil {
		_, err = c.db.NamedExecContext(ctx, qb.String(), channel)
	} else {
		_, err = c.tx.NamedExecContext(ctx, qb.String(), channel)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *chatBotRepository) GetChannelByID(ctx context.Context, id uuid.UUID) (*entity.Channel, error) {
	var channel entity.Channel

	var qb strings.Builder

	qb.WriteString(`
		SELECT
			id, name, doctor_id, created_at, updated_at
		FROM
			channels
		WHERE
			id = $1
	`)

	err := c.db.GetContext(ctx, &channel, qb.String(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errx.ErrChannelNotFound
		}

		return nil, err
	}

	return &channel, nil
}

func (c *chatBotRepository) GetChannelsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]entity.Channel, error) {
	var channels []entity.Channel

	var qb strings.Builder

	qb.WriteString(`
		SELECT
			id, name, doctor_id, created_at, updated_at
		FROM
			channels
		WHERE
			doctor_id = $1
	`)

	err := c.db.SelectContext(ctx, &channels, qb.String(), doctorID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
