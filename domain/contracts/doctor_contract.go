package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type DoctorRepository interface {
	GetDoctors(ctx context.Context, query *dto.GetDoctorsQuery) ([]entity.Doctor, error)
	CountDoctors(ctx context.Context, query *dto.GetDoctorsQuery) (int64, error)
	CreateDoctor(ctx context.Context, doctor *entity.Doctor) error
	GetDoctorByID(ctx context.Context, id uuid.UUID) (*entity.Doctor, error)
	GetDoctorByEmail(ctx context.Context, email string) (*entity.Doctor, error)
}

type DoctorService interface {
	GetDoctors(ctx context.Context, query dto.GetDoctorsQuery) (dto.GetDoctorsResponse, error)
	LoginDoctor(ctx context.Context, req dto.LoginDoctorRequest) (dto.LoginDoctorResponse, error)
	GetDoctorSession(ctx context.Context, req dto.GetDoctorSessionRequest) (dto.GetDoctorSessionResponse, error)
}
