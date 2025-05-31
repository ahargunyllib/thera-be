package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type notificationController struct {
	notificationSvc contracts.NotificationService
}

func InitNotificationController(
	router fiber.Router,
	notificationSvc contracts.NotificationService,
	middleware *middlewares.Middleware,
) {
	controller := notificationController{
		notificationSvc: notificationSvc,
	}

	notificationRouter := router.Group("/notifications")
	notificationRouter.Get("/me", middleware.RequireAuth(), controller.getMyNotifications)
}

func (c *notificationController) getMyNotifications(ctx *fiber.Ctx) error {
	var query dto.GetMyNotificationsQuery

	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	if claims.Role == "admin" {
		query.HospitalID = claims.HospitalID
	} else if claims.Role == "doctor" {
		query.DoctorID = claims.UserID
	}

	res, err := c.notificationSvc.GetMyNotifications(ctx.Context(), query)
	if err != nil {
		return err
	}

return response.SendResponse(ctx, fiber.StatusOK, res)
}
