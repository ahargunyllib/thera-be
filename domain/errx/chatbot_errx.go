package errx

import "net/http"

var (
	ErrChannelNotFound = NewError(
		http.StatusNotFound,
		"channel_not_found",
		"The requested channel was not found.",
	)
)
