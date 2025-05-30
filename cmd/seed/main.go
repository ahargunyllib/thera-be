package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/domain/entity"
	hospitalRepository "github.com/ahargunyllib/thera-be/internal/app/hospital/repository"
	"github.com/ahargunyllib/thera-be/internal/infra/database"
	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/ahargunyllib/thera-be/pkg/log"
)

const SeedersFilePath = "data/seeders/"
const SeedersDevPath = SeedersFilePath + "dev/"
const SeedersProdPath = SeedersFilePath + "prod/"

func main() {
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	var path string
	if env.AppEnv.AppEnv == "production" {
		path = SeedersProdPath
	} else {
		path = SeedersDevPath
	}

	hospitalRepo := hospitalRepository.NewHospitalRepository(psqlDB)

	seedHospitals(path, hospitalRepo)
}

func seedHospitals(path string, hospitalRepo contracts.HospitalRepository) {
	path += "hospitals.csv"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(log.CustomLogInfo{
			"error": err,
		}, "[seed][seedHospitals] Error opening file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err,
		}, "[seed][seedHospitals] Error reading file")
		return
	}

	ctx := context.Background()

	for idx, record := range records {
		if idx == 0 { // skip header
			continue
		}

		latitude, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
			}, "[seed][seedHospitals] Error parsing latitude")
		}

		longitude, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
			}, "[seed][seedHospitals] Error parsing longitude")
		}

		yearEstablished, err := strconv.Atoi(record[7])
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
			}, "[seed][seedHospitals] Error parsing year established")
		}

		hospital := entity.Hospital{
			Name:    record[0],
			Address: record[1],
			Phone: sql.NullString{
				String: record[2],
				Valid:  true,
			},
			Email: sql.NullString{
				String: record[3],
				Valid:  true,
			},
			Website: sql.NullString{
				String: record[4],
				Valid:  true,
			},
			Latitude:        latitude,
			Longitude:       longitude,
			YearEstablished: yearEstablished,
		}

		err = hospitalRepo.CreateHospital(ctx, &hospital)
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error":    err,
				"hospital": hospital,
			}, "[seed][seedHospitals] Error creating hospital")
		}
	}
}
