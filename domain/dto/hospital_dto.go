package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
)

type HospitalResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Address         string    `json:"address"`
	Phone           string    `json:"phone,omitempty"`
	Email           string    `json:"email,omitempty"`
	Website         string    `json:"website,omitempty"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	YearEstablished int       `json:"year_established"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewHospitalResponse(hospitalEntity *entity.Hospital) HospitalResponse {
	return HospitalResponse{
		ID:              hospitalEntity.ID,
		Name:            hospitalEntity.Name,
		Address:         hospitalEntity.Address,
		Phone:           hospitalEntity.Phone.String,
		Email:           hospitalEntity.Email.String,
		Website:         hospitalEntity.Website.String,
		Type:            enums.HospitalTypeMapIdx[hospitalEntity.Type].LongLabel["id"],
		Status:          enums.HospitalStatusMapIdx[hospitalEntity.Status].LongLabel["id"],
		Latitude:        hospitalEntity.Latitude,
		Longitude:       hospitalEntity.Longitude,
		YearEstablished: hospitalEntity.YearEstablished,
		CreatedAt:       hospitalEntity.CreatedAt,
		UpdatedAt:       hospitalEntity.UpdatedAt,
	}
}

type GetHospitalsQuery struct {
	Page      int    `query:"page" validate:"omitempty,min=1"`
	Limit     int    `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    string `query:"sort_by" validate:"omitempty,oneof=id name created_at"`
	SortOrder string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

type GetHospitalsResponse struct {
	Hospitals []HospitalResponse `json:"hospitals"`
	Meta      struct {
		Pagination PaginationResponse `json:"pagination"`
	} `json:"meta"`
}

type GetHospitalByIDParams struct {
	ID int `path:"id" validate:"required,min=1"`
}

type GetHospitalResponse struct {
	Hospital HospitalResponse `json:"hospital"`
}
