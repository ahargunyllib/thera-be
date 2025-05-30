package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/domain/entity"
	adminRepository "github.com/ahargunyllib/thera-be/internal/app/admin/repository"
	hospitalRepository "github.com/ahargunyllib/thera-be/internal/app/hospital/repository"
	"github.com/ahargunyllib/thera-be/internal/infra/database"
	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/ahargunyllib/thera-be/pkg/bcrypt"
	"github.com/ahargunyllib/thera-be/pkg/helpers/flag"
	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/ahargunyllib/thera-be/pkg/uuid"
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
	adminRepo := adminRepository.NewAdminRepository(psqlDB)

	bcrypt := bcrypt.Bcrypt
	uuid := uuid.UUID

	switch flag.FlagVars.SeederEntity {
	case "hospitals":
		seedHospitals(path, hospitalRepo)
	case "admins":
		seedAdmins(path, adminRepo, bcrypt, uuid)
	default:
		log.Error(log.CustomLogInfo{
			"seeder_entity": flag.FlagVars.SeederEntity,
		}, "[seed][main] Invalid seeder entity specified")
		return
	}
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

func seedAdmins(path string, adminRepo contracts.AdminRepository, bcrypt bcrypt.BcryptInterface, uuid uuid.UUIDInterface) {
	path += "admins.csv"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(log.CustomLogInfo{
			"error": err,
		}, "[seed][seedAdmins] Error opening file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err,
		}, "[seed][seedAdmins] Error reading file")
		return
	}

	ctx := context.Background()
	for idx, record := range records {
		if idx == 0 { // skip header
			continue
		}

		email := record[0]
		fullName := record[1]
		password := record[2]

		roleInt, err := strconv.Atoi(record[3])
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
			}, "[seed][seedAdmins] Error parsing role")
			continue
		}

		hospitalID, err := strconv.Atoi(record[4])
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
			}, "[seed][seedAdmins] Error parsing hospital ID")
			continue
		}

		hashedPassword, err := bcrypt.Hash(password)
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
			}, "[seed][seedAdmins] Error hashing password")
			continue
		}

		id, err := uuid.NewV7()
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
			}, "[seed][seedAdmins] Error generating UUID")
			continue
		}

		admin := &entity.Admin{
			ID:         id,
			FullName:   fullName,
			Email:      email,
			Password:   hashedPassword,
			Role:       roleInt,
			HospitalID: hospitalID,
		}

		err = adminRepo.CreateAdmin(ctx, admin)
		if err != nil {
			log.Error(log.CustomLogInfo{
				"error": err,
				"admin": admin,
			}, "[seed][seedAdmins] Error creating admin")
			continue
		}
	}
}
