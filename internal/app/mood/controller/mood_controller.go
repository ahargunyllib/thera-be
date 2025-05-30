package controller

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/gofiber/fiber/v2"
)

func (mc *moodController) createMood(c *fiber.Ctx) error {
	log.Debug(log.CustomLogInfo{
		"message": "Creating mood",
	}, "[moodController.createMood]")

	var req dto.CreateMoodRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	claims, ok := c.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	req.DoctorID = claims.UserID

	err := mc.moodSvc.CreateMood(c.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusCreated, nil)
}

func (mc *moodController) getMyDailyMood(c *fiber.Ctx) error {
	var query dto.GetMyDailyMoodQuery

	claims, ok := c.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	query.DoctorID = claims.UserID

	res, err := mc.moodSvc.GetMyDailyMood(c.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, res)
}

func (mc *moodController) getMyMonthlyOverview(c *fiber.Ctx) error {
	var query dto.GetMyMonthlyOverviewQuery

	claims, ok := c.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	query.DoctorID = claims.UserID

	res, err := mc.moodSvc.GetMyMonthlyOverview(c.Context(), query)
	if err != nil {
		return err
	}

	return response.SendResponse(c, fiber.StatusOK, res)
}
