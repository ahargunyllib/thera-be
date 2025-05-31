package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
)

func (c *chatBotService) GetChannelByID(
	ctx context.Context,
	params dto.GetChannelByIDParams,
) (dto.GetChannelByIDResponse, error) {
	valErr := c.validator.Validate(params)
	if valErr != nil {
		return dto.GetChannelByIDResponse{}, valErr
	}

	channel, err := c.chatBotRepo.GetChannelByID(ctx, params.ChannelID)
	if err != nil {
		return dto.GetChannelByIDResponse{}, err
	}

	messages, err := c.chatBotRepo.GetMessagesByChannelID(ctx, params.ChannelID)
	if err != nil {
		return dto.GetChannelByIDResponse{}, err
	}

	messagesResponse := make([]dto.MessageResponse, len(messages))
	for i, message := range messages {
		messagesResponse[i] = dto.NewMessageResponse(&message)
	}

	channelResponse := dto.NewChannelResponse(channel)

	res := dto.GetChannelByIDResponse{
		Channel:  channelResponse,
		Messages: messagesResponse,
	}

	return res, nil
}

// GetChannels implements contracts.ChatbotService.
func (c *chatBotService) GetChannelsByDoctorID(
	ctx context.Context,
	params dto.GetChannelsByDoctorIDParams,
) (dto.GetChannelsResponse, error) {
	valErr := c.validator.Validate(params)
	if valErr != nil {
		return dto.GetChannelsResponse{}, valErr
	}

	channels, err := c.chatBotRepo.GetChannelsByDoctorID(ctx, params.DoctorID)
	if err != nil {
		return dto.GetChannelsResponse{}, err
	}

	channelsResponse := make([]dto.ChannelResponse, len(channels))
	for i, channel := range channels {
		channelsResponse[i] = dto.NewChannelResponse(&channel)
	}

	res := dto.GetChannelsResponse{
		Channels: channelsResponse,
	}

	return res, nil
}
