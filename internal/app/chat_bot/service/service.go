package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	openai "github.com/ahargunyllib/thera-be/pkg/opeanai"
	"github.com/ahargunyllib/thera-be/pkg/uuid"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type chatBotService struct {
	chatBotRepo contracts.ChatBotRepository
	validator   validator.CustomValidatorInterface
	uuid        uuid.UUIDInterface
	openai      openai.CustomOpenAIInterface
}

func NewChatBotService(
	chatBotRepo contracts.ChatBotRepository,
	validator validator.CustomValidatorInterface,
	uuid uuid.UUIDInterface,
	openai openai.CustomOpenAIInterface,
) contracts.ChatbotService {
	return &chatBotService{
		chatBotRepo: chatBotRepo,
		validator:   validator,
		uuid:        uuid,
		openai:      openai,
	}
}
