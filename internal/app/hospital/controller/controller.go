package controller

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/gofiber/fiber/v2"
)

type hospitalController struct {
	hospitalSvc contracts.HospitalService
}

func InitHospitalController(
	router fiber.Router,
	hospitalService contracts.HospitalService,
) {
	controller := hospitalController{
		hospitalSvc: hospitalService,
	}

	hospitalRouter := router.Group("/hospitals")
	hospitalRouter.Get("/", controller.getHospitals)
	hospitalRouter.Get("/:id", controller.getHospitalByID)
}
