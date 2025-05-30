package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (ac *adminController) loginAdmin(ctx *fiber.Ctx) error {
	var req dto.LoginAdminRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := ac.adminSvc.LoginAdmin(ctx.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (ac *adminController) getAdminSession(ctx *fiber.Ctx) error {
	var req dto.GetAdminSessionRequest

	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	req.AdminID = claims.UserID

	res, err := ac.adminSvc.GetAdminSession(ctx.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}
