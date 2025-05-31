package errx

import "net/http"

var (
	ErrDoctorScheduleNotFound = NewError(
		http.StatusNotFound,
		"doctor_schedule_not_found",
		"The requested doctor schedule could not be found.",
	)
	ErrInvalidResponseFormat = NewError(
		http.StatusBadRequest,
		"invalid_response_format",
		"The response format is invalid or does not match the expected structure.",
	)
)
