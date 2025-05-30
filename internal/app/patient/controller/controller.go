package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type patientController struct {
	patientSvc contracts.PatientService
}

func InitPatientController(
	router fiber.Router,
	patientSvc contracts.PatientService,
	middleware *middlewares.Middleware,
) {
	controller := &patientController{
		patientSvc: patientSvc,
	}

	patientRouter := router.Group("/patients")
	patientRouter.Get("/", middleware.RequireAuth(), middleware.RequireRole("admin"), controller.getPatients)
	patientRouter.Get("/:id", middleware.RequireAuth(), middleware.RequireRole("admin"), controller.getPatientByID)
	patientRouter.Post("/", middleware.RequireAuth(), middleware.RequireRole("admin"), controller.createPatient)
	patientRouter.Patch("/:id", middleware.RequireAuth(), middleware.RequireRole("admin"), controller.updatePatientByID)
	patientRouter.Delete("/:id", middleware.RequireAuth(), middleware.RequireRole("admin"), controller.deletePatientByID)
}
