package errx

import "net/http"

var (
	ErrDoctorNotFound = NewError(
		http.StatusNotFound,
		"doctor_not_found",
		"The specified doctor could not be found.",
	)
	ErrDoctorAlreadyExists = NewError(
		http.StatusConflict,
		"doctor_already_exists",
		"A doctor with this email already exists.",
	)
	ErrDoctorInvalidCredentials = NewError(
		http.StatusUnauthorized,
		"doctor_invalid_credentials",
		"The provided credentials are invalid.",
	)
)
