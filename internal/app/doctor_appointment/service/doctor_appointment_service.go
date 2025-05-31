package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/ahargunyllib/thera-be/domain/errx"
)

func (d *doctorAppointmentService) CreateDoctorAppointment(
	ctx context.Context,
	req dto.CreateDoctorAppointmentRequest,
) error {
	valErr := d.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	id, err := d.ulid.New()
	if err != nil {
		return err
	}

	isAvailable, err := d.doctorAppointmentRepo.CheckDoctorScheduleAvailability(
		ctx,
		req.DoctorID,
		req.StartTime,
		req.EndTime,
	)
	if err != nil {
		return err
	}
	if !isAvailable {
		return errx.ErrDoctorAppointmentNotAvailable
	}

	doctorAppointment := entity.DoctorAppointment{
		ID:              id.String(),
		DoctorID:        req.DoctorID,
		PatientID:       req.PatientID,
		AppointmentDate: req.AppointmentDate,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		Status:          enums.DoctorAppointmentStatusPendingIdx,
		Type:            enums.DoctorAppointmentTypeConsultationIdx,
	}

	err = d.doctorAppointmentRepo.CreateDoctorAppointment(ctx, &doctorAppointment)
	if err != nil {
		return err
	}

	return nil
}

func (d *doctorAppointmentService) DeleteDoctorAppointment(
	ctx context.Context,
	params dto.DeleteDoctorAppointmentParams,
) error {
	valErr := d.validator.Validate(params)
	if valErr != nil {
		return valErr
	}

	err := d.doctorAppointmentRepo.DeleteDoctorAppointment(ctx, params.ID)
	if err != nil {
		return err
	}

	return nil
}

func (d *doctorAppointmentService) GetDoctorAppointments(
	ctx context.Context,
	query dto.GetDoctorAppointmentsQuery,
) (dto.GetDoctorAppointmentsResponse, error) {
	valErr := d.validator.Validate(query)
	if valErr != nil {
		return dto.GetDoctorAppointmentsResponse{}, valErr
	}

	doctorAppointments, err := d.doctorAppointmentRepo.GetDoctorAppointments(ctx, &query)
	if err != nil {
		return dto.GetDoctorAppointmentsResponse{}, err
	}

	doctorAppointmentsResponse := make([]dto.DoctorAppointmentResponse, len(doctorAppointments))
	for i, doctorAppointment := range doctorAppointments {
		doctorAppointmentsResponse[i] = dto.NewDoctorAppointmentResponse(&doctorAppointment)
	}

	res := dto.GetDoctorAppointmentsResponse{
		DoctorAppointments: doctorAppointmentsResponse,
	}

	return res, nil
}

func (d *doctorAppointmentService) UpdateDoctorAppointment(
	ctx context.Context,
	params dto.UpdateDoctorAppointmentParams,
	req dto.UpdateDoctorAppointmentRequest,
) error {
	valErr := d.validator.Validate(params)
	if valErr != nil {
		return valErr
	}

	valErr = d.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	doctorAppointment := entity.DoctorAppointment{
		ID:              params.ID,
		DoctorID:        req.DoctorID,
		PatientID:       req.PatientID,
		AppointmentDate: req.AppointmentDate,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		Status:          enums.DoctorAppointmentStatusPendingIdx,
		Type:            enums.DoctorAppointmentTypeConsultationIdx,
	}

	err := d.doctorAppointmentRepo.UpdateDoctorAppointment(ctx, &doctorAppointment)
	if err != nil {
		return err
	}

	return nil
}
