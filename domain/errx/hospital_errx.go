package errx

import "net/http"

var (
	ErrHospitalNotFound = NewError(
		http.StatusNotFound,
		"hospital_not_found",
		"The specified hospital could not be found.",
	)
)
