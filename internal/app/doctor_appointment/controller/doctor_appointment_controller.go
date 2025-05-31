package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (dac *doctorAppointmentController) getDoctorAppointments(c *fiber.Ctx) error {
	var query dto.GetDoctorAppointmentsQuery
	if err := c.QueryParser(&query); err != nil {
		return err
	}

	doctorIDStr := c.Query("doctor_id")
	if doctorIDStr != "" {
		doctorID, err := uuid.Parse(doctorIDStr)
		if err != nil {
			return err
		}

		query.DoctorID = doctorID
	}

	patientIDStr := c.Query("patient_id")
	if patientIDStr != "" {
		patientID, err := uuid.Parse(patientIDStr)
		if err != nil {
			return err
		}

		query.PatientID = patientID
	}

	res, err := dac.doctorAppointmentSvc.GetDoctorAppointments(c.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, res)
}

func (dac *doctorAppointmentController) createDoctorAppointment(c *fiber.Ctx) error {
	var req dto.CreateDoctorAppointmentRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := dac.doctorAppointmentSvc.CreateDoctorAppointment(c.Context(), req); err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusCreated, nil)
}

func (dac *doctorAppointmentController) updateDoctorAppointment(c *fiber.Ctx) error {
	var params dto.UpdateDoctorAppointmentParams
	if err := c.ParamsParser(&params); err != nil {
		return err
	}

	var req dto.UpdateDoctorAppointmentRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := dac.doctorAppointmentSvc.UpdateDoctorAppointment(c.Context(), params, req); err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, nil)
}

func (dac *doctorAppointmentController) deleteDoctorAppointment(c *fiber.Ctx) error {
	var params dto.DeleteDoctorAppointmentParams
	if err := c.ParamsParser(&params); err != nil {
		return err
	}

	if err := dac.doctorAppointmentSvc.DeleteDoctorAppointment(c.Context(), params); err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusNoContent, nil)
}
