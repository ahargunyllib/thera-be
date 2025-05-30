package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/gofiber/fiber/v2"
)

func (hc *hospitalController) getHospitals(ctx *fiber.Ctx) error {
	var query dto.GetHospitalsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return err
	}

	res, err := hc.hospitalSvc.GetHospitals(ctx.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (hc *hospitalController) getHospitalByID(ctx *fiber.Ctx) error {
	var params dto.GetHospitalByIDParams
	if err := ctx.ParamsParser(&params); err != nil {
		return err
	}

	res, err := hc.hospitalSvc.GetHospitalByID(ctx.Context(), params)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}
