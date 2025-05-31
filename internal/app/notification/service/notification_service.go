package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

func (n *notificationService) GetMyNotifications(
	ctx context.Context,
	query dto.GetMyNotificationsQuery,
) (dto.GetMyNotificationsResponse, error) {
	valErr := n.validator.Validate(query)
	if valErr != nil {
		return dto.GetMyNotificationsResponse{}, valErr
	}

	var notifications []entity.Notification
	var err error

	if query.HospitalID == 0 && query.DoctorID == uuid.Nil {
		return dto.GetMyNotificationsResponse{}, nil
	}

	if query.HospitalID != 0 {
		notifications, err = n.notificationRepo.GetNotificationsByHospitalID(ctx, query.HospitalID)
	}

	if query.DoctorID != uuid.Nil {
		notifications, err = n.notificationRepo.GetNotificationsByDoctorID(ctx, query.DoctorID)
	}

	if err != nil {
		return dto.GetMyNotificationsResponse{}, err
	}

	notificationsResponse := make([]dto.NotificationResponse, len(notifications))
	for i, notification := range notifications {
		notificationsResponse[i] = dto.NewNotificationResponse(&notification)
	}

	res := dto.GetMyNotificationsResponse{
		Notifications: notificationsResponse,
	}

	return res, nil
}
