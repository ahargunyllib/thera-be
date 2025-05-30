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

	hospitalCtrl := router.Group("/hospitals")
	hospitalCtrl.Get("/", controller.getHospitals)
	hospitalCtrl.Get("/:id", controller.getHospitalByID)
}
