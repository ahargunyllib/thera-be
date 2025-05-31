package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type ChannelResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Doctor    DoctorResponse `json:"doctor"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func NewChannelResponse(channelEntity *entity.Channel) ChannelResponse {
	return ChannelResponse{
		ID:        channelEntity.ID,
		Name:      channelEntity.Name,
		Doctor:    NewDoctorResponse(&channelEntity.Doctor),
		CreatedAt: channelEntity.CreatedAt,
		UpdatedAt: channelEntity.UpdatedAt,
	}
}

type GetChannelsByDoctorIDParams struct {
	DoctorID uuid.UUID `validate:"required,uuid"`
}

type GetChannelsResponse struct {
	Channels []ChannelResponse `json:"channels"`
}

type GetChannelByIDParams struct {
	ChannelID uuid.UUID `params:"id" validate:"required,uuid"`
}

type GetChannelByIDResponse struct {
	Channel  ChannelResponse   `json:"channel"`
	Messages []MessageResponse `json:"messages"`
}

type MessageResponse struct {
	ID        uuid.UUID       `json:"id"`
	Channel   ChannelResponse `json:"channel"`
	Content   string          `json:"content"`
	Role      string          `json:"role"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func NewMessageResponse(messageEntity *entity.Message) MessageResponse {
	return MessageResponse{
		ID: messageEntity.ID,
		// Channel:   NewChannelResponse(&messageEntity.Channel),
		Content:   messageEntity.Content,
		Role:      enums.MessageRoleMapIdx[messageEntity.Role].LongLabel["id"],
		CreatedAt: messageEntity.CreatedAt,
		UpdatedAt: messageEntity.UpdatedAt,
	}
}

type CreateMessageRequest struct {
	ChannelID uuid.UUID `json:"channel_id,omitempty" validate:"omitempty,uuid"`
	DoctorID  uuid.UUID `validate:"uuid"`
	Content   string    `json:"content" validate:"required"`
}

type CreateMessageResponse struct {
	Message MessageResponse `json:"message"`
	Channel ChannelResponse `json:"channel"`
}
