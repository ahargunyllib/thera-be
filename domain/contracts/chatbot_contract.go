package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type ChatBotRepository interface {
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error

	GetChannelsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]entity.Channel, error)
	GetChannelByID(ctx context.Context, id uuid.UUID) (*entity.Channel, error)
	CreateChannel(ctx context.Context, channel *entity.Channel) error

	GetMessagesByChannelID(ctx context.Context, channelID uuid.UUID) ([]entity.Message, error)
	CreateMessage(ctx context.Context, message *entity.Message) error
}

type ChatbotService interface {
	GetChannelsByDoctorID(ctx context.Context, params dto.GetChannelsByDoctorIDParams) (dto.GetChannelsResponse, error)
	GetChannelByID(ctx context.Context, params dto.GetChannelByIDParams) (dto.GetChannelByIDResponse, error)
	CreateMessage(ctx context.Context, req dto.CreateMessageRequest) (dto.CreateMessageResponse, error)
}
