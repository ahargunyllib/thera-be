package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/gofiber/fiber/v2"
)

func (pc *patientController) getPatients(ctx *fiber.Ctx) error {
	var query dto.GetPatientsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return err
	}

	res, err := pc.patientSvc.GetPatients(ctx.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (pc *patientController) getPatientByID(ctx *fiber.Ctx) error {
	var params dto.GetPatientByIDParams
	if err := ctx.ParamsParser(&params); err != nil {
		return err
	}

	res, err := pc.patientSvc.GetPatientByID(ctx.Context(), params)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (pc *patientController) createPatient(ctx *fiber.Ctx) error {
	var req dto.CreatePatientRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := pc.patientSvc.CreatePatient(ctx.Context(), req); err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusCreated, nil)
}

func (pc *patientController) updatePatientByID(ctx *fiber.Ctx) error {
	var params dto.UpdatePatientByIDParams
	if err := ctx.ParamsParser(&params); err != nil {
		return err
	}

	var req dto.UpdatePatientRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := pc.patientSvc.UpdatePatientByID(ctx.Context(), params, req); err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, nil)
}

func (pc *patientController) deletePatientByID(ctx *fiber.Ctx) error {
	var params dto.DeletePatientByIDParams
	if err := ctx.ParamsParser(&params); err != nil {
		return err
	}

	if err := pc.patientSvc.DeletePatientByID(ctx.Context(), params); err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusNoContent, nil)
}
