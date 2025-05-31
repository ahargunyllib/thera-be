package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type hospitalPartnerController struct {
	hospitalPartnerSvc contracts.HospitalPartnerService
}

func InitHospitalPartnerController(
	router fiber.Router,
	hospitalPartnerSvc contracts.HospitalPartnerService,
	middleware *middlewares.Middleware,
) {
	controller := &hospitalPartnerController{
		hospitalPartnerSvc: hospitalPartnerSvc,
	}

	hospitalPartnerRouter := router.Group("/hospital-partners")
	hospitalPartnerRouter.Get(
		"/me",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.getMyHospitalPartners,
	)
	hospitalPartnerRouter.Post(
		"/",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.createHospitalPartner,
	)
	hospitalPartnerRouter.Patch(
		"/:id",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.updateHospitalPartner,
	)
}
