package main

import (
	"github.com/ahargunyllib/thera-be/internal/infra/database"
)

const SeedersFilePath = "data/seeders/"
const SeedersDevPath = SeedersFilePath + "dev/"
const SeedersProdPath = SeedersFilePath + "prod/"

func main() {
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	// var path string
	// if env.AppEnv.AppEnv == "production" {
	// 	path = SeedersProdPath
	// } else {
	// 	path = SeedersDevPath
	// }
}
