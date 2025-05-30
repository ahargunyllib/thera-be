package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type moodController struct {
	moodSvc contracts.MoodService
}

func InitMoodController(
	router fiber.Router,
	moodSvc contracts.MoodService,
	middleware *middlewares.Middleware,
) {
	controller := &moodController{
		moodSvc: moodSvc,
	}

	moodRouter := router.Group("/moods")
	moodRouter.Post("/", middleware.RequireAuth(), middleware.RequireRole("doctor"), controller.createMood)
	moodRouter.Get("/me/weekly", middleware.RequireAuth(), middleware.RequireRole("doctor"), controller.getMyDailyMood)
	moodRouter.Get(
		"/me/monthly",
		middleware.RequireAuth(),
		middleware.RequireRole("doctor"),
		controller.getMyMonthlyOverview,
	)
}
