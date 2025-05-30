package errx

import "net/http"

var (
	ErrAdminNotFound = NewError(
		http.StatusNotFound,
		"admin_not_found",
		"The specified admin could not be found.",
	)
	ErrAdminAlreadyExists = NewError(
		http.StatusConflict,
		"admin_already_exists",
		"An admin with this email already exists.",
	)
	ErrAdminInvalidCredentials = NewError(
		http.StatusUnauthorized,
		"admin_invalid_credentials",
		"The provided credentials are invalid.",
	)
)
