package main

import (
	"github.com/ahargunyllib/thera-be/internal/infra/database"
	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/ahargunyllib/thera-be/internal/infra/redis"
	"github.com/ahargunyllib/thera-be/internal/infra/server"
)

func main() {
	server := server.NewHTTPServer()
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()
	redis := redis.NewRedisConn()
	defer redis.Close()

	server.MountMiddlewares()
	server.MountRoutes(psqlDB, redis)
	server.Start(env.AppEnv.AppPort)
}
