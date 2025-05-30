package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/pkg/log"
)

func (hs *hospitalService) GetHospitalByID(
	ctx context.Context,
	params dto.GetHospitalByIDParams,
) (dto.GetHospitalResponse, error) {
	valErr := hs.validator.Validate(params)
	if valErr != nil {
		return dto.GetHospitalResponse{}, valErr
	}

	hospital, err := hs.hospitalRepo.GetHospitalByID(ctx, params.ID)
	if err != nil {
		return dto.GetHospitalResponse{}, err
	}

	hospitalResponse := dto.NewHospitalResponse(hospital)

	res := dto.GetHospitalResponse{
		Hospital: hospitalResponse,
	}

	return res, nil
}

func (hs *hospitalService) GetHospitals(
	ctx context.Context,
	query dto.GetHospitalsQuery,
) (dto.GetHospitalsResponse, error) {
	log.Debug(log.CustomLogInfo{
		"query": query,
	}, "[HospitalService][GetHospitals] query")

	valErr := hs.validator.Validate(query)
	if valErr != nil {
		return dto.GetHospitalsResponse{}, valErr
	}

	hospitals, err := hs.hospitalRepo.GetHospitals(ctx, &query)
	if err != nil {
		return dto.GetHospitalsResponse{}, err
	}

	hospitalsResponse := make([]dto.HospitalResponse, len(hospitals))
	for i, hospital := range hospitals {
		hospitalsResponse[i] = dto.NewHospitalResponse(&hospital)
	}

	count, err := hs.hospitalRepo.CountHospitals(ctx, &query)
	if err != nil {
		return dto.GetHospitalsResponse{}, err
	}

	paginationResponse := dto.NewPaginationResponse(count, query.Page, query.Limit)

	res := dto.GetHospitalsResponse{
		Hospitals: hospitalsResponse,
	}

	res.Meta.Pagination = paginationResponse

	return res, nil
}
