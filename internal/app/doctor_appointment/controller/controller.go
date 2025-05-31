package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type doctorAppointmentController struct {
	doctorAppointmentSvc contracts.DoctorAppointmentService
}

func InitDoctorAppointmentController(
	router fiber.Router,
	doctorAppointmentSvc contracts.DoctorAppointmentService,
	middleware *middlewares.Middleware,
) {
	controller := &doctorAppointmentController{
		doctorAppointmentSvc: doctorAppointmentSvc,
	}

	doctorAppointmentRouter := router.Group("/doctor-appointments")
	doctorAppointmentRouter.Get("/", middleware.RequireAuth(), controller.getDoctorAppointments)
	doctorAppointmentRouter.Post("/", middleware.RequireAuth(), controller.createDoctorAppointment)
	doctorAppointmentRouter.Put("/:id", middleware.RequireAuth(), controller.updateDoctorAppointment)
	doctorAppointmentRouter.Delete("/:id", middleware.RequireAuth(), controller.deleteDoctorAppointment)
}
