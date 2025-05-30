package errx

import "net/http"

var (
	ErrHaveAlreadyCreatedMood = NewError(
		http.StatusConflict,
		"have_already_created_mood",
		"You have already created a mood for today.",
	)
)
