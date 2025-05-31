package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type HospitalPartnerRepository interface {
	GetHospitalPartnersByHospitalID(ctx context.Context, hospitalID uuid.UUID) ([]entity.HospitalPartner, error)
	GetHospitalPartnerByID(ctx context.Context, id string) (*entity.HospitalPartner, error)
	CreateHospitalPartner(ctx context.Context, hospitalPartner *entity.HospitalPartner) error
	UpdateHospitalPartner(ctx context.Context, hospitalPartner *entity.HospitalPartner) error
}

type HospitalPartnerService interface {
	GetHospitalPartnersByHospitalID(ctx context.Context, query dto.GetMyHospitalPartnersQuery) (
		dto.GetHospitalPartnersResponse,
		error,
	)
	CreateHospitalPartner(ctx context.Context, req dto.CreateHospitalPartnerRequest) error
	UpdateHospitalPartner(
		ctx context.Context,
		params dto.UpdateHospitalPartnerParams,
		req dto.UpdateHospitalPartnerRequest,
	) error
}
