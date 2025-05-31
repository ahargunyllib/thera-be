package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
)

func (ds *doctorService) GetDoctorSession(
	ctx context.Context,
	req dto.GetDoctorSessionRequest,
) (dto.GetDoctorSessionResponse, error) {
	valErr := ds.validator.Validate(req)
	if valErr != nil {
		return dto.GetDoctorSessionResponse{}, valErr
	}

	doctor, err := ds.doctorRepo.GetDoctorByID(ctx, req.DoctorID)
	if err != nil {
		return dto.GetDoctorSessionResponse{}, err
	}

	doctorResponse := dto.NewDoctorResponse(doctor)

	res := dto.GetDoctorSessionResponse{
		Doctor: doctorResponse,
	}

	return res, nil
}

func (ds *doctorService) GetDoctors(ctx context.Context, query dto.GetDoctorsQuery) (dto.GetDoctorsResponse, error) {
	valErr := ds.validator.Validate(query)
	if valErr != nil {
		return dto.GetDoctorsResponse{}, valErr
	}

	doctors, err := ds.doctorRepo.GetDoctors(ctx, &query)
	if err != nil {
		return dto.GetDoctorsResponse{}, err
	}

	doctorsResponse := make([]dto.DoctorResponse, len(doctors))
	for i, doctor := range doctors {
		doctorsResponse[i] = dto.NewDoctorResponse(&doctor)
	}

	count, err := ds.doctorRepo.CountDoctors(ctx, &query)
	if err != nil {
		return dto.GetDoctorsResponse{}, err
	}

	paginationResponse := dto.NewPaginationResponse(count, query.Page, query.Limit)

	res := dto.GetDoctorsResponse{
		Doctors: doctorsResponse,
	}

	res.Meta.Pagination = paginationResponse

	return res, nil
}

func (ds *doctorService) LoginDoctor(ctx context.Context, req dto.LoginDoctorRequest) (dto.LoginDoctorResponse, error) {
	valErr := ds.validator.Validate(req)
	if valErr != nil {
		return dto.LoginDoctorResponse{}, valErr
	}

	docter, err := ds.doctorRepo.GetDoctorByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginDoctorResponse{}, err
	}

	if !ds.bcrypt.Compare(req.Password, docter.Password) {
		return dto.LoginDoctorResponse{}, errx.ErrDoctorInvalidCredentials
	}

	token, err := ds.jwt.Create(docter.ID, "doctor", docter.HospitalID)
	if err != nil {
		return dto.LoginDoctorResponse{}, err
	}

	res := dto.LoginDoctorResponse{
		AccessToken: token,
	}

	return res, nil
}
