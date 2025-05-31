package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
)

func (h *hospitalPartnerService) CreateHospitalPartner(ctx context.Context, req dto.CreateHospitalPartnerRequest) error {
	valErr := h.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	id, err := h.ulid.New()
	if err != nil {
		return err
	}

	hospitalPartner := &entity.HospitalPartner{
		ID:             id.String(),
		FromHospitalID: req.FromHospitalID,
		ToHospitalID:   req.ToHospitalID,
		PartnerType:    enums.HospitalPartnerTypeIdx(enums.HospitalPartnerTypeMapKey[enums.HospitalPartnerTypeKey(req.PartnerType)].Idx),
		Status:         enums.HospitalPartnerStatusPendingIdx,
	}

	err = h.hospitalPartnerRepo.CreateHospitalPartner(ctx, hospitalPartner)
	if err != nil {
		return err
	}

	return nil
}

func (h *hospitalPartnerService) GetHospitalPartnersByHospitalID(ctx context.Context, query dto.GetMyHospitalPartnersQuery) (dto.GetHospitalPartnersResponse, error) {
	valErr := h.validator.Validate(query)
	if valErr != nil {
		return dto.GetHospitalPartnersResponse{}, valErr
	}

	hospitalPartners, err := h.hospitalPartnerRepo.GetHospitalPartnersByHospitalID(ctx, query.HospitalID)
	if err != nil {
		return dto.GetHospitalPartnersResponse{}, err
	}

	hospitalPartnersResponse := make([]dto.HospitalPartnerResponse, len(hospitalPartners))
	for i, partner := range hospitalPartners {
		hospitalPartnersResponse[i] = dto.NewHospitalPartnerResponse(&partner)
	}

	res := dto.GetHospitalPartnersResponse{
		HospitalPartners: hospitalPartnersResponse,
	}

	return res, nil
}

func (h *hospitalPartnerService) UpdateHospitalPartner(ctx context.Context, params dto.UpdateHospitalPartnerParams, req dto.UpdateHospitalPartnerRequest) error {
	valErr := h.validator.Validate(params)
	if valErr != nil {
		return valErr
	}

	hospitalPartner, err := h.hospitalPartnerRepo.GetHospitalPartnerByID(ctx, params.PartnerID)
	if err != nil {
		return err
	}

	hospitalPartner.Status = enums.HospitalPartnerStatusIdx(enums.HospitalPartnerStatusMapKey[enums.HospitalPartnerStatusKey(req.PartnerStatus)].Idx)

	err = h.hospitalPartnerRepo.UpdateHospitalPartner(ctx, hospitalPartner)
	if err != nil {
		return err
	}

	return nil
}
