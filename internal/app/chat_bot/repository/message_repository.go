package repository

import (
	"context"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

func (c *chatBotRepository) CreateMessage(ctx context.Context, message *entity.Message) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO messages
			(id, channel_id, content, role)
		VALUES
			(:id, :channel_id, :content, :role)
	`)

	var err error
	if c.tx != nil {
		_, err = c.db.NamedExecContext(ctx, qb.String(), message)
	} else {
		_, err = c.tx.NamedExecContext(ctx, qb.String(), message)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *chatBotRepository) GetMessagesByChannelID(ctx context.Context, channelID uuid.UUID) ([]entity.Message, error) {
	var messages []entity.Message

	var qb strings.Builder

	qb.WriteString(`
		SELECT
			id, channel_id, content, role, created_at, updated_at
		FROM
			messages
		WHERE
			channel_id = $1
	`)

	err := c.db.SelectContext(ctx, &messages, qb.String(), channelID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
