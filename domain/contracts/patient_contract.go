package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type PatientRepository interface {
	GetPatients(ctx context.Context, query *dto.GetPatientsQuery) ([]entity.Patient, error)
	CountPatients(ctx context.Context, query *dto.GetPatientsQuery) (int64, error)
	GetPatientByID(ctx context.Context, id uuid.UUID) (*entity.Patient, error)
	CreatePatient(ctx context.Context, patient *entity.Patient) error
	UpdatePatient(ctx context.Context, patient *entity.Patient) error
	DeletePatientByID(ctx context.Context, id uuid.UUID) error
}

type PatientService interface {
	GetPatients(ctx context.Context, query dto.GetPatientsQuery) (dto.GetPatientsResponse, error)
	GetPatientByID(ctx context.Context, params dto.GetPatientByIDParams) (dto.GetPatientResponse, error)
	CreatePatient(ctx context.Context, req dto.CreatePatientRequest) error
	UpdatePatientByID(ctx context.Context, params dto.UpdatePatientByIDParams, req dto.UpdatePatientRequest) error
	DeletePatientByID(ctx context.Context, params dto.DeletePatientByIDParams) error
}
