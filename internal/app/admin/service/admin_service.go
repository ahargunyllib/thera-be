package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/errx"
)

func (as *adminService) GetAdminSession(
	ctx context.Context,
	req dto.GetAdminSessionRequest,
) (dto.GetAdminSessionResponse, error) {
	valErr := as.validator.Validate(req)
	if valErr != nil {
		return dto.GetAdminSessionResponse{}, valErr
	}

	admin, err := as.adminRepo.GetAdminByID(ctx, req.AdminID)
	if err != nil {
		return dto.GetAdminSessionResponse{}, err
	}

	adminResponse := dto.NewAdminResponse(admin)

	res := dto.GetAdminSessionResponse{
		Admin: adminResponse,
	}

	return res, nil
}

func (as *adminService) LoginAdmin(ctx context.Context, req dto.LoginAdminRequest) (dto.LoginAdminResponse, error) {
	valErr := as.validator.Validate(req)
	if valErr != nil {
		return dto.LoginAdminResponse{}, valErr
	}

	admin, err := as.adminRepo.GetAdminByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginAdminResponse{}, err
	}

	if !as.bcrypt.Compare(req.Password, admin.Password) {
		return dto.LoginAdminResponse{}, errx.ErrAdminInvalidCredentials
	}

	token, err := as.jwt.Create(admin.ID, "admin")
	if err != nil {
		return dto.LoginAdminResponse{}, err
	}

	res := dto.LoginAdminResponse{
		AccessToken: token,
	}

	return res, nil
}
