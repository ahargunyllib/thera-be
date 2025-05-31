package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (cbr *chatBotController) GetChannelByID(ctx *fiber.Ctx) error {
	var params dto.GetChannelByIDParams
	if err := ctx.ParamsParser(&params); err != nil {
		return err
	}

	res, err := cbr.chatBotSvc.GetChannelByID(ctx.UserContext(), params)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (cbr *chatBotController) getMyChannels(ctx *fiber.Ctx) error {
	var params dto.GetChannelsByDoctorIDParams

	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	params.DoctorID = claims.UserID

	res, err := cbr.chatBotSvc.GetChannelsByDoctorID(ctx.UserContext(), params)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}
