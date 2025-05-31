package errx

import "net/http"

var (
	ErrDoctorScheduleNotFound = NewError(
		http.StatusNotFound,
		"doctor_schedule_not_found",
		"The requested doctor schedule could not be found.",
	)
)
