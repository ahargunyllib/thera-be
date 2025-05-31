package errx

import "net/http"

var (
	ErrHospitalPartnerNotFound = NewError(
		http.StatusNotFound,
		"hospital_partner_not_found",
		"The requested hospital partner was not found. Please check the ID and try again.",
	)
)
