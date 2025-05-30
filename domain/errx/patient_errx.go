package errx

import "net/http"

var (
	ErrPatientNotFound = NewError(
		http.StatusNotFound,
		"patient_not_found",
		"The specified patient could not be found.",
	)
	ErrIDNumberPatientAlreadyExists = NewError(
		http.StatusConflict,
		"id_number_already_exists",
		"A patient with this ID Number already exists.",
	)
	ErrMedicalRecordNumberPatientAlreadyExists = NewError(
		http.StatusConflict,
		"medical_record_number_already_exists",
		"A patient with this Medical Record Number already exists.",
	)
)
