package entity

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID            `db:"id"`
	ChannelID uuid.UUID            `db:"channel_id"`
	Channel   Channel              `db:"channel"`
	Content   string               `db:"content"`
	Role      enums.MessageRoleIdx `db:"role"`
	CreatedAt time.Time            `db:"created_at"`
	UpdatedAt time.Time            `db:"updated_at"`
}
