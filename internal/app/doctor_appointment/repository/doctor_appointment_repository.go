package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/errx"
	"github.com/ahargunyllib/thera-be/pkg/helpers"
	"github.com/ahargunyllib/thera-be/pkg/helpers/pgerror"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (d *doctorAppointmentRepository) CreateDoctorAppointment(
	ctx context.Context,
	doctorAppointment *entity.DoctorAppointment,
) error {
	var qb strings.Builder

	qb.WriteString(`
		INSERT INTO doctor_appointments (
			id, doctor_id, patient_id, appointment_date, start_time, end_time, status, type
		) VALUES (
			:id, :doctor_id, :patient_id, :appointment_date, :start_time, :end_time, :status, :type
		)
	`)

	_, err := d.db.NamedExecContext(ctx, qb.String(), doctorAppointment)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "doctor_appointments_doctor_id_fkey",
					Err: errx.ErrDoctorNotFound.WithDetails(map[string]any{
						"doctor_id": doctorAppointment.DoctorID,
					}).WithLocation("repository.doctor_appointment.CreateDoctorAppointment"),
				},
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "doctor_appointments_patient_id_fkey",
					Err: errx.ErrPatientNotFound.WithDetails(map[string]any{
						"patient_id": doctorAppointment.PatientID,
					}).WithLocation("repository.doctor_appointment.CreateDoctorAppointment"),
				},
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "idx_doctor_appointments_unique",
					Err: errx.ErrDoctorAppointmentAlreadyExists.WithDetails(map[string]any{
						"doctor_id":        doctorAppointment.DoctorID,
						"patient_id":       doctorAppointment.PatientID,
						"appointment_date": doctorAppointment.AppointmentDate,
						"start_time":       doctorAppointment.StartTime,
						"end_time":         doctorAppointment.EndTime,
					}).WithLocation("repository.doctor_appointment.CreateDoctorAppointment"),
				},
			}

			if customPgErr := pgerror.HandlePgError(*pgErr, pgErrors); customPgErr != nil {
				return customPgErr
			}
		}

		return err
	}

	return nil
}

func (d *doctorAppointmentRepository) DeleteDoctorAppointment(ctx context.Context, id string) error {
	var qb strings.Builder

	qb.WriteString(`
		DELETE FROM doctor_appointments
		WHERE id = $1
	`)

	res, err := d.db.ExecContext(ctx, qb.String(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	return helpers.CheckRowsAffected(rowsAffected, errx.ErrDoctorAppointmentNotFound.WithDetails(map[string]any{
		"id": id,
	}).WithLocation("repository.doctor_appointment.DeleteDoctorAppointment"))
}

func (d *doctorAppointmentRepository) GetDoctorAppointments(
	ctx context.Context,
	query *dto.GetDoctorAppointmentsQuery,
) ([]entity.DoctorAppointment, error) {
	var doctorAppointments []entity.DoctorAppointment

	var qb strings.Builder
	var args []any

	qb.WriteString(`
		SELECT
			doctor_appointments.id,
			doctor_appointments.doctor_id,
			doctor_appointments.patient_id,
			doctor_appointments.appointment_date,
			doctor_appointments.start_time,
			doctor_appointments.end_time,
			doctor_appointments.status,
			doctor_appointments.type,
			doctor_appointments.created_at,
			doctor_appointments.updated_at,
			doctors.id AS "doctor.id",
			doctors.full_name AS "doctor.full_name",
			doctors.email AS "doctor.email",
			doctors.phone_number AS "doctor.phone_number",
			doctors.specialty AS "doctor.specialty",
			doctors.hospital_id AS "doctor.hospital_id",
			doctors.created_at AS "doctor.created_at",
			doctors.updated_at AS "doctor.updated_at",
			patients.id AS "patient.id",
			patients.full_name AS "patient.full_name",
			patients.id_number AS "patient.id_number",
			patients.phone_number AS "patient.phone_number",
			patients.address AS "patient.address",
			patients.date_of_birth AS "patient.date_of_birth",
			patients.gender AS "patient.gender",
			patients.height AS "patient.height",
			patients.weight AS "patient.weight",
			patients.blood_type AS "patient.blood_type",
			patients.allergies AS "patient.allergies",
			patients.medical_record_number AS "patient.medical_record_number",
			patients.hospital_id AS "patient.hospital_id",
			patients.created_at AS "patient.created_at",
			patients.updated_at AS "patient.updated_at"
		FROM doctor_appointments
		JOIN doctors ON doctor_appointments.doctor_id = doctors.id
		JOIN patients ON doctor_appointments.patient_id = patients.id
		WHERE 1=1
	`)

	if query.DoctorID != uuid.Nil {
		qb.WriteString(fmt.Sprintf(`
			AND doctor_appointments.doctor_id = $%d
		`, len(args)+1))
		args = append(args, query.DoctorID)
	}

	if query.PatientID != uuid.Nil {
		qb.WriteString(fmt.Sprintf(`
			AND doctor_appointments.patient_id = $%d
		`, len(args)+1))
		args = append(args, query.PatientID)
	}

	if query.FromDate != "" {
		qb.WriteString(fmt.Sprintf(`
			AND doctor_appointments.appointment_date >= $%d
		`, len(args)+1))
		args = append(args, query.FromDate)
	}

	if query.ToDate != "" {
		qb.WriteString(fmt.Sprintf(`
			AND doctor_appointments.appointment_date <= $%d
		`, len(args)+1))
		args = append(args, query.ToDate)
	}

	err := d.db.SelectContext(ctx, &doctorAppointments, qb.String(), args...)
	if err != nil {
		return nil, err
	}

	return doctorAppointments, nil
}

func (d *doctorAppointmentRepository) UpdateDoctorAppointment(
	ctx context.Context,
	doctorAppointment *entity.DoctorAppointment,
) error {
	var qb strings.Builder

	qb.WriteString(`
		UPDATE doctor_appointments
		SET
			doctor_id = :doctor_id,
			patient_id = :patient_id,
			appointment_date = :appointment_date,
			start_time = :start_time,
			end_time = :end_time,
			status = :status,
			type = :type
		WHERE id = :id
	`)

	res, err := d.db.NamedExecContext(ctx, qb.String(), doctorAppointment)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "doctor_appointments_doctor_id_fkey",
					Err: errx.ErrDoctorNotFound.WithDetails(map[string]any{
						"doctor_id": doctorAppointment.DoctorID,
					}).WithLocation("repository.doctor_appointment.UpdateDoctorAppointment"),
				},
				{
					Code:           pgerror.ForeignKey,
					ConstraintName: "doctor_appointments_patient_id_fkey",
					Err: errx.ErrPatientNotFound.WithDetails(map[string]any{
						"patient_id": doctorAppointment.PatientID,
					}).WithLocation("repository.doctor_appointment.UpdateDoctorAppointment"),
				},
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "idx_doctor_appointments_unique",
					Err: errx.ErrDoctorAppointmentAlreadyExists.WithDetails(map[string]any{
						"doctor_id":        doctorAppointment.DoctorID,
						"patient_id":       doctorAppointment.PatientID,
						"appointment_date": doctorAppointment.AppointmentDate,
						"start_time":       doctorAppointment.StartTime,
						"end_time":         doctorAppointment.EndTime,
					}).WithLocation("repository.doctor_appointment.UpdateDoctorAppointment"),
				},
			}

			if customPgErr := pgerror.HandlePgError(*pgErr, pgErrors); customPgErr != nil {
				return customPgErr
			}
		}

		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	return helpers.CheckRowsAffected(rowsAffected, errx.ErrDoctorAppointmentNotFound.WithDetails(map[string]any{
		"id": doctorAppointment.ID,
	}).WithLocation("repository.doctor_appointment.UpdateDoctorAppointment"))
}
