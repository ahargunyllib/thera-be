package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
	openai "github.com/ahargunyllib/thera-be/pkg/opeanai"
	"github.com/google/uuid"
)

// CreateMessage implements contracts.ChatbotService.
func (c *chatBotService) CreateMessage(
	ctx context.Context,
	req dto.CreateMessageRequest,
) (dto.CreateMessageResponse, error) {
	valErr := c.validator.Validate(req)
	if valErr != nil {
		return dto.CreateMessageResponse{}, valErr
	}

	isCreatingNewChannel := false
	if req.ChannelID == uuid.Nil {
		channelID, err := c.uuid.NewV7()
		if err != nil {
			return dto.CreateMessageResponse{}, err
		}

		req.ChannelID = channelID

		isCreatingNewChannel = true
	}

	messageID, err := c.uuid.NewV7()
	if err != nil {
		return dto.CreateMessageResponse{}, err
	}

	userMessage := &entity.Message{
		ID:        messageID,
		ChannelID: req.ChannelID,
		Content:   req.Content,
		Role:      enums.MessageRoleUserIdx,
	}

	err = c.chatBotRepo.Begin(ctx)
	if err != nil {
		return dto.CreateMessageResponse{}, err
	}

	messages, err := c.chatBotRepo.GetMessagesByChannelID(ctx, req.ChannelID)
	if err != nil {
		_ = c.chatBotRepo.Rollback()
		return dto.CreateMessageResponse{}, err
	}

	historyMessages := make([]openai.Message, len(messages)+2)
	historyMessages[0] = openai.Message{
		Content: `
			You are an assistant specialized in mental health consultation for doctors and medical professionals.
			Your role is to provide empathetic, professional, and evidence-based advice
				to help them navigate their mental health challenges.
			Please ensure that all responses remain within the context of mental health and avoid discussing unrelated topics.
		`,
		Role: "system",
	}
	for i, message := range messages {
		var role string
		if message.Role == enums.MessageRoleUserIdx {
			role = "user"
		} else {
			role = "assistant"
		}

		historyMessages[i+1] = openai.Message{
			Content: message.Content,
			Role:    role,
		}
	}
	historyMessages[len(messages)+1] = openai.Message{
		Content: req.Content,
		Role:    "user",
	}

	chatRes, err := c.openai.Chat(ctx, historyMessages)
	if err != nil {
		_ = c.chatBotRepo.Rollback()
		return dto.CreateMessageResponse{}, err
	}

	assistantMessageID, err := c.uuid.NewV7()
	if err != nil {
		_ = c.chatBotRepo.Rollback()
		return dto.CreateMessageResponse{}, err
	}

	assistantMessage := &entity.Message{
		ID:        assistantMessageID,
		ChannelID: req.ChannelID,
		Content:   chatRes.Choices[0].Message.Content,
		Role:      enums.MessageRoleAssistantIdx,
	}

	if isCreatingNewChannel {
		// Ask for channel name if creating a new channel
		msgs := make([]openai.Message, 2)
		msgs[0] = openai.Message{
			Content: "Please provide a name for the new channel. Short and concise. For example: 'General Health Discussion'",
			Role:    "user",
		}
		msgs[1] = openai.Message{
			Content: userMessage.Content + " " + chatRes.Choices[0].Message.Content,
			Role:    "assistant",
		}

		chatRes, err = c.openai.Chat(ctx, msgs)
		if err != nil {
			_ = c.chatBotRepo.Rollback()
			return dto.CreateMessageResponse{}, err
		}

		channel := &entity.Channel{
			ID:       req.ChannelID,
			Name:     chatRes.Choices[0].Message.Content,
			DoctorID: req.DoctorID,
		}

		err = c.chatBotRepo.CreateChannel(ctx, channel)
		if err != nil {
			_ = c.chatBotRepo.Rollback()
			return dto.CreateMessageResponse{}, err
		}
	}

	err = c.chatBotRepo.CreateMessage(ctx, userMessage)
	if err != nil {
		_ = c.chatBotRepo.Rollback()
		return dto.CreateMessageResponse{}, err
	}

	err = c.chatBotRepo.CreateMessage(ctx, assistantMessage)
	if err != nil {
		_ = c.chatBotRepo.Rollback()
		return dto.CreateMessageResponse{}, err
	}

	channel, err := c.chatBotRepo.GetChannelByID(ctx, req.ChannelID)
	if err != nil {
		_ = c.chatBotRepo.Rollback()
		return dto.CreateMessageResponse{}, err
	}

	err = c.chatBotRepo.Commit()
	if err != nil {
		_ = c.chatBotRepo.Rollback()
		return dto.CreateMessageResponse{}, err
	}

	res := dto.CreateMessageResponse{
		Message: dto.NewMessageResponse(assistantMessage),
		Channel: dto.NewChannelResponse(channel),
	}

	return res, nil
}
