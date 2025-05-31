package errx

import "net/http"

var (
	ErrDoctorAppointmentNotFound = NewError(
		http.StatusNotFound,
		"doctor_appointment_not_found",
		"Please check the appointment ID and try again.",
	)
	ErrDoctorAppointmentAlreadyExists = NewError(
		http.StatusConflict,
		"doctor_appointment_already_exists",
		"An appointment with the same details already exists. Please check the doctor, patient, date, and time.",
	)
)
