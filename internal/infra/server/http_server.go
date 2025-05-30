package server

import (
	adminController "github.com/ahargunyllib/thera-be/internal/app/admin/controller"
	adminRepo "github.com/ahargunyllib/thera-be/internal/app/admin/repository"
	adminSvc "github.com/ahargunyllib/thera-be/internal/app/admin/service"
	doctorController "github.com/ahargunyllib/thera-be/internal/app/doctor/controller"
	doctorRepo "github.com/ahargunyllib/thera-be/internal/app/doctor/repository"
	doctorSvc "github.com/ahargunyllib/thera-be/internal/app/doctor/service"
	hospitalController "github.com/ahargunyllib/thera-be/internal/app/hospital/controller"
	hospitalRepo "github.com/ahargunyllib/thera-be/internal/app/hospital/repository"
	hospitalSvc "github.com/ahargunyllib/thera-be/internal/app/hospital/service"
	moodController "github.com/ahargunyllib/thera-be/internal/app/mood/controller"
	moodRepo "github.com/ahargunyllib/thera-be/internal/app/mood/repository"
	moodSvc "github.com/ahargunyllib/thera-be/internal/app/mood/service"
	patientController "github.com/ahargunyllib/thera-be/internal/app/patient/controller"
	patientRepo "github.com/ahargunyllib/thera-be/internal/app/patient/repository"
	patientSvc "github.com/ahargunyllib/thera-be/internal/app/patient/service"
	"github.com/ahargunyllib/thera-be/internal/middlewares"
	"github.com/ahargunyllib/thera-be/pkg/bcrypt"
	errorhandler "github.com/ahargunyllib/thera-be/pkg/helpers/http/error_handler"
	"github.com/ahargunyllib/thera-be/pkg/helpers/http/response"
	"github.com/ahargunyllib/thera-be/pkg/jwt"
	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/ahargunyllib/thera-be/pkg/uuid"
	"github.com/ahargunyllib/thera-be/pkg/validator"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type HTTPServer interface {
	Start(part string)
	MountMiddlewares()
	MountRoutes(db *sqlx.DB, redis *redis.Client)
	GetApp() *fiber.App
}

type httpServer struct {
	app *fiber.App
}

func NewHTTPServer() HTTPServer {
	config := fiber.Config{
		CaseSensitive: true,
		AppName:       "Thera BE",
		ServerHeader:  "Thera BE",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		ErrorHandler:  errorhandler.ErrorHandler,
	}

	app := fiber.New(config)

	return &httpServer{
		app: app,
	}
}

func (s *httpServer) GetApp() *fiber.App {
	return s.app
}

func (s *httpServer) Start(port string) {
	if port[0] != ':' {
		port = ":" + port
	}

	err := s.app.Listen(port)

	if err != nil {
		log.Fatal(log.CustomLogInfo{
			"error": err.Error(),
		}, "[SERVER][Start] failed to start server")
	}
}

func (s *httpServer) MountMiddlewares() {
	s.app.Use(middlewares.LoggerConfig())
	s.app.Use(middlewares.Helmet())
	s.app.Use(middlewares.Compress())
	s.app.Use(middlewares.Cors())
	s.app.Use(middlewares.RecoverConfig())
}

func (s *httpServer) MountRoutes(db *sqlx.DB, redis *redis.Client) {
	validator := validator.Validator
	bcrypt := bcrypt.Bcrypt
	jwt := jwt.Jwt
	uuid := uuid.UUID

	s.app.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "Thera BE is running")
	})

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "Thera BE is running")
	})

	hospitalRepository := hospitalRepo.NewHospitalRepository(db)
	adminRepository := adminRepo.NewAdminRepository(db)
	doctorRepository := doctorRepo.NewDoctorRepository(db)
	patientRepository := patientRepo.NewPatientRepository(db)
	moodRepository := moodRepo.NewMoodRepository(db)

	hospitalService := hospitalSvc.NewHospitalService(hospitalRepository, validator)
	adminService := adminSvc.NewAdminService(adminRepository, validator, bcrypt, jwt)
	doctorService := doctorSvc.NewDoctorService(doctorRepository, validator, bcrypt, jwt)
	patientService := patientSvc.NewPatientService(patientRepository, validator, uuid)
	moodService := moodSvc.NewMoodService(moodRepository, validator)

	middleware := middlewares.NewMiddleware(jwt)

	hospitalController.InitHospitalController(v1, hospitalService)
	adminController.InitAdminController(v1, adminService, middleware)
	doctorController.InitDoctorController(v1, doctorService, middleware)
	patientController.InitPatientController(v1, patientService, middleware)
	moodController.InitMoodController(v1, moodService, middleware)

	s.app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("./web/not-found.html")
	})
}
