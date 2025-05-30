package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type AdminRepository interface {
	GetAdminByEmail(ctx context.Context, email string) (*entity.Admin, error)
	GetAdminByID(ctx context.Context, id uuid.UUID) (*entity.Admin, error)
	CreateAdmin(ctx context.Context, admin *entity.Admin) error
}

type AdminService interface {
	LoginAdmin(ctx context.Context, req dto.LoginAdminRequest) (dto.LoginAdminResponse, error)
	GetAdminSession(ctx context.Context, req dto.GetAdminSessionRequest) (dto.GetAdminSessionResponse, error)
}
