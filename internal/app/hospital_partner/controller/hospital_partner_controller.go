package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (c *hospitalPartnerController) getMyHospitalPartners(ctx *fiber.Ctx) error {
	var query dto.GetMyHospitalPartnersQuery
	if err := ctx.QueryParser(&query); err != nil {
		return err
	}

	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	query.HospitalID = claims.HospitalID

	res, err := c.hospitalPartnerSvc.GetHospitalPartnersByHospitalID(ctx.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (c *hospitalPartnerController) createHospitalPartner(ctx *fiber.Ctx) error {
	var req dto.CreateHospitalPartnerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	req.FromHospitalID = claims.HospitalID

	err := c.hospitalPartnerSvc.CreateHospitalPartner(ctx.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusCreated, nil)
}

func (c *hospitalPartnerController) updateHospitalPartner(ctx *fiber.Ctx) error {
	var params dto.UpdateHospitalPartnerParams
	if err := ctx.ParamsParser(&params); err != nil {
		return err
	}

	var req dto.UpdateHospitalPartnerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := c.hospitalPartnerSvc.UpdateHospitalPartner(ctx.Context(), params, req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, nil)
}
