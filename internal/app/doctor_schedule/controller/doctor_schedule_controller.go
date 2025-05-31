package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (dsc *doctorScheduleController) createDoctorSchedule(c *fiber.Ctx) error {
	var req dto.CreateDoctorScheduleRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := dsc.doctorScheduleSvc.CreateDoctorSchedule(c.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusCreated, nil)
}

func (dsc *doctorScheduleController) getDoctorSchedules(c *fiber.Ctx) error {
	var query dto.GetDoctorSchedulesQuery
	if err := c.QueryParser(&query); err != nil {
		return err
	}

	doctorIDStr := c.Query("doctor_id")
	if doctorIDStr != "" {
		doctorID, err := uuid.Parse(doctorIDStr)
		if err != nil {
			return errx.ErrCannotParseUUID.WithDetails(map[string]any{
				"doctor_id": doctorIDStr,
			})
		}
		query.DoctorID = doctorID
	}

	res, err := dsc.doctorScheduleSvc.GetDoctorSchedules(c.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, res)
}

func (dsc *doctorScheduleController) updateDoctorSchedule(c *fiber.Ctx) error {
	var params dto.UpdateDoctorScheduleParams
	if err := c.ParamsParser(&params); err != nil {
		return err
	}

	var req dto.UpdateDoctorScheduleRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := dsc.doctorScheduleSvc.UpdateDoctorSchedule(c.Context(), params, req)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, nil)
}

func (dsc *doctorScheduleController) deleteDoctorSchedule(c *fiber.Ctx) error {
	var params dto.DeleteDoctorScheduleParams
	if err := c.ParamsParser(&params); err != nil {
		return err
	}

	err := dsc.doctorScheduleSvc.DeleteDoctorSchedule(c.Context(), params)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusNoContent, nil)
}

func (dsc *doctorScheduleController) getImprovedNextSchedulePreview(c *fiber.Ctx) error {
	var query dto.GetPreviewImprovedNextScheduleQuery
	if err := c.QueryParser(&query); err != nil {
		return err
	}

	res, err := dsc.doctorScheduleSvc.GetPreviewImprovedNextSchedule(c.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, res)
}
