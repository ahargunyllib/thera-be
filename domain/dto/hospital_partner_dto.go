package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
)

type HospitalPartnerResponse struct {
	ID           string           `json:"id"`
	FromHospital HospitalResponse `json:"from_hospital"`
	ToHospital   HospitalResponse `json:"to_hospital"`
	PartnerType  string           `json:"partner_type"`
	Status       string           `json:"status"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

func NewHospitalPartnerResponse(hospitalPartnerEntity *entity.HospitalPartner) HospitalPartnerResponse {
	return HospitalPartnerResponse{
		ID:           hospitalPartnerEntity.ID,
		FromHospital: NewHospitalResponse(&hospitalPartnerEntity.FromHospital),
		ToHospital:   NewHospitalResponse(&hospitalPartnerEntity.ToHospital),
		PartnerType:  enums.HospitalPartnerTypeMapIdx[hospitalPartnerEntity.PartnerType].LongLabel["id"],
		Status:       enums.HospitalPartnerStatusMapIdx[hospitalPartnerEntity.Status].LongLabel["id"],
		CreatedAt:    hospitalPartnerEntity.CreatedAt,
		UpdatedAt:    hospitalPartnerEntity.UpdatedAt,
	}
}

type GetMyHospitalPartnersQuery struct {
	HospitalID int `validate:"required"`
}

type GetHospitalPartnersResponse struct {
	HospitalPartners []HospitalPartnerResponse `json:"hospital_partners"`
}

type CreateHospitalPartnerRequest struct {
	FromHospitalID int    `validate:"required"`
	ToHospitalID   int    `json:"to_hospital_id" validate:"required"`
	PartnerType    string `json:"partner_type" validate:"required,oneof=collaboration"`
}

type UpdateHospitalPartnerParams struct {
	PartnerID string `param:"partner_id" validate:"required"`
}

type UpdateHospitalPartnerRequest struct {
	PartnerStatus string `json:"partner_status" validate:"required,oneof=confirmed canceled"`
}
