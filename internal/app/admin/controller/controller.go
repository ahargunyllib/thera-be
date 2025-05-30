package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type adminController struct {
	adminSvc contracts.AdminService
}

func InitAdminController(router fiber.Router, adminSvc contracts.AdminService, middleware *middlewares.Middleware) {
	controller := &adminController{
		adminSvc: adminSvc,
	}

	adminRouter := router.Group("/admins")
	adminRouter.Post("/login", controller.loginAdmin)
	adminRouter.Get("/sessions", middleware.RequireAuth(), middleware.RequireRole("admin"), controller.getAdminSession)
}
