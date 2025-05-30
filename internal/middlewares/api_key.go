package middlewares

import (
	"strings"

	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/gofiber/fiber/v2"
)

func APIKey() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Path() == "/api/v1/service-transactions/notifications" {
			return ctx.Next()
		}

		apiKey := ctx.Get("x-api-key")
		if apiKey == "" {
			return errx.ErrNoAPIKey
		}

		keySlice := strings.Split(apiKey, " ")
		if len(keySlice) != 2 {
			return errx.ErrInvalidAPIKey
		}

		key := keySlice[1]
		if key != env.AppEnv.APIKey {
			return errx.ErrInvalidAPIKey
		}

		return ctx.Next()
	}
}
