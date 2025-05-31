package service

import (
	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/ahargunyllib/thera-be/pkg/validator"
)

type notificationService struct {
	notificationRepo contracts.NotificationRepository
	validator        validator.CustomValidatorInterface
}

func NewNotificationService(
	notificationRepo contracts.NotificationRepository,
	validator validator.CustomValidatorInterface,
) contracts.NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		validator:        validator,
	}
}
