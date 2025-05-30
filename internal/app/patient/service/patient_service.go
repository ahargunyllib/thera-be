package service

import (
	"context"
	"database/sql"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
)

func (ps *patientService) CreatePatient(ctx context.Context, req dto.CreatePatientRequest) error {
	valErr := ps.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	id, err := ps.uuid.NewV7()
	if err != nil {
		return err
	}

	patient := entity.Patient{
		ID:       id,
		FullName: req.FullName,
		IDNumber: req.IDNumber,
		PhoneNumber: sql.NullString{
			String: req.PhoneNumber,
			Valid:  req.PhoneNumber != "",
		},
		Address: sql.NullString{
			String: req.Address,
			Valid:  req.Address != "",
		},
		DateOfBirth: sql.NullTime{
			Time:  req.DateOfBirth,
			Valid: !req.DateOfBirth.IsZero(),
		},
		Gender:    1, // TODO
		Height:    req.Height,
		Weight:    req.Weight,
		BloodType: 1, // TODO
		Allergies: sql.NullString{
			String: req.Allergies,
			Valid:  req.Allergies != "",
		},

		MedicalRecordNumber: req.MedicalRecordNumber,
		HospitalID:          req.HospitalID,
	}

	err = ps.patientRepo.CreatePatient(ctx, &patient)
	if err != nil {
		return err
	}

	return nil
}

func (ps *patientService) DeletePatientByID(ctx context.Context, params dto.DeletePatientByIDParams) error {
	valErr := ps.validator.Validate(params)
	if valErr != nil {
		return valErr
	}

	err := ps.patientRepo.DeletePatientByID(ctx, params.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ps *patientService) GetPatientByID(
	ctx context.Context,
	params dto.GetPatientByIDParams,
) (dto.GetPatientResponse, error) {
	valErr := ps.validator.Validate(params)
	if valErr != nil {
		return dto.GetPatientResponse{}, valErr
	}

	patient, err := ps.patientRepo.GetPatientByID(ctx, params.ID)
	if err != nil {
		return dto.GetPatientResponse{}, err
	}

	patientResponse := dto.NewPatientResponse(patient)

	res := dto.GetPatientResponse{
		Patient: patientResponse,
	}

	return res, nil
}

func (ps *patientService) GetPatients(ctx context.Context, query dto.GetPatientsQuery) (
	dto.GetPatientsResponse,
	error,
) {
	valErr := ps.validator.Validate(query)
	if valErr != nil {
		return dto.GetPatientsResponse{}, valErr
	}

	patients, err := ps.patientRepo.GetPatients(ctx, &query)
	if err != nil {
		return dto.GetPatientsResponse{}, err
	}

	count, err := ps.patientRepo.CountPatients(ctx, &query)
	if err != nil {
		return dto.GetPatientsResponse{}, err
	}

	patientResponses := make([]dto.PatientResponse, len(patients))
	for i, patient := range patients {
		patientResponses[i] = dto.NewPatientResponse(&patient)
	}

	paginationResponse := dto.NewPaginationResponse(count, query.Page, query.Limit)

	res := dto.GetPatientsResponse{
		Patients: patientResponses,
	}

	res.Meta.Pagination = paginationResponse

	return res, nil
}

func (ps *patientService) UpdatePatientByID(
	ctx context.Context,
	params dto.UpdatePatientByIDParams,
	req dto.UpdatePatientRequest,
) error {
	valErr := ps.validator.Validate(params)
	if valErr != nil {
		return valErr
	}

	valErr = ps.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	patient, err := ps.patientRepo.GetPatientByID(ctx, params.ID)
	if err != nil {
		return err
	}

	err = parsePatientRequest(patient, &req)
	if err != nil {
		return err
	}

	err = ps.patientRepo.UpdatePatient(ctx, patient)
	if err != nil {
		return err
	}

	return nil
}
