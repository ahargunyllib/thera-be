package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type chatBotController struct {
	chatBotSvc contracts.ChatbotService
}

func InitChatBotController(
	router fiber.Router,
	chatBotSvc contracts.ChatbotService,
	middleware *middlewares.Middleware,
) {
	controller := &chatBotController{
		chatBotSvc: chatBotSvc,
	}

	chatBotRouter := router.Group("/chat-bots")
	chatBotRouter.Get("/channels/:id", middleware.RequireAuth(), controller.GetChannelByID)
	chatBotRouter.Post("/channels/messages", middleware.RequireAuth(), controller.createMessage)
	chatBotRouter.Get("/channels/me", middleware.RequireAuth(), controller.getMyChannels)
}
