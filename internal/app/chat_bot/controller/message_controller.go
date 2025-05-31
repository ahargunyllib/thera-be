package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (cbr *chatBotController) createMessage(ctx *fiber.Ctx) error {
	var req dto.CreateMessageRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	req.DoctorID = claims.UserID

	res, err := cbr.chatBotSvc.CreateMessage(ctx.UserContext(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusCreated, res)
}
