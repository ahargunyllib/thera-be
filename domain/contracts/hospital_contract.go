package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
)

type HospitalRepository interface {
	GetHospitals(ctx context.Context, query *dto.GetHospitalsQuery) ([]entity.Hospital, error)
	CountHospitals(ctx context.Context, query *dto.GetHospitalsQuery) (int64, error)
	GetHospitalByID(ctx context.Context, id int) (*entity.Hospital, error)
	CreateHospital(ctx context.Context, hospital *entity.Hospital) error
}

type HospitalService interface {
	GetHospitals(ctx context.Context, query dto.GetHospitalsQuery) (dto.GetHospitalsResponse, error)
	GetHospitalByID(ctx context.Context, params dto.GetHospitalByIDParams) (dto.GetHospitalResponse, error)
}
