package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type moodService struct {
	moodRepo  contracts.MoodRepository
	validator validator.CustomValidatorInterface
}

func NewMoodService(
	moodRepo contracts.MoodRepository,
	validator validator.CustomValidatorInterface,
) contracts.MoodService {
	return &moodService{
		moodRepo:  moodRepo,
		validator: validator,
	}
}
