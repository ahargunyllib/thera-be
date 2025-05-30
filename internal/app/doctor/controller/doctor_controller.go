package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (dc *doctorController) getDoctors(c *fiber.Ctx) error {
	var query dto.GetDoctorsQuery
	if err := c.QueryParser(&query); err != nil {
		return err
	}

	doctors, err := dc.doctorSvc.GetDoctors(c.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, doctors)
}

func (dc *doctorController) loginDoctor(c *fiber.Ctx) error {
	var req dto.LoginDoctorRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	res, err := dc.doctorSvc.LoginDoctor(c.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, res)
}

func (dc *doctorController) getDoctorSession(c *fiber.Ctx) error {
	var req dto.GetDoctorSessionRequest

	claims, ok := c.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	req.DoctorID = claims.UserID

	res, err := dc.doctorSvc.GetDoctorSession(c.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, res)
}
