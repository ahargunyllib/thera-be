package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type doctorScheduleController struct {
	doctorScheduleSvc contracts.DoctorScheduleService
}

func InitDoctorScheduleController(
	router fiber.Router,
	doctorScheduleSvc contracts.DoctorScheduleService,
	middleware *middlewares.Middleware,
) {
	controller := &doctorScheduleController{
		doctorScheduleSvc: doctorScheduleSvc,
	}

	doctorScheduleRouter := router.Group("/doctor-schedules")
	doctorScheduleRouter.Post(
		"/",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.createDoctorSchedule,
	)
	doctorScheduleRouter.Get(
		"/",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.getDoctorSchedules,
	)
	doctorScheduleRouter.Get(
		"/improved-next-schedule-preview",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.getImprovedNextSchedulePreview,
	)
	doctorScheduleRouter.Patch(
		"/:id",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.updateDoctorSchedule,
	)
	doctorScheduleRouter.Delete(
		"/:id",
		middleware.RequireAuth(),
		middleware.RequireRole("admin"),
		controller.deleteDoctorSchedule,
	)
}
