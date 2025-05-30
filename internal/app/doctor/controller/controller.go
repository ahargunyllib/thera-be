package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type doctorController struct {
	doctorSvc contracts.DoctorService
}

func InitDoctorController(
	router fiber.Router,
	doctorService contracts.DoctorService,
	middleware *middlewares.Middleware,
) {
	controller := doctorController{
		doctorSvc: doctorService,
	}

	doctorRouter := router.Group("/doctors")
	doctorRouter.Get("/", controller.getDoctors)
	doctorRouter.Post("/login", controller.loginDoctor)
	doctorRouter.Get("/sessions", middleware.RequireAuth(), middleware.RequireRole("doctor"), controller.getDoctorSession)
}
