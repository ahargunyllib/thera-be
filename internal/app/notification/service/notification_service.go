package service

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/pkg/log"
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
	notificationIDs := make([]string, len(notifications))
	for i, notification := range notifications {
		notificationsResponse[i] = dto.NewNotificationResponse(&notification)
		notificationIDs[i] = notification.ID
	}

	go func() {
		err = n.notificationRepo.ReadNotifications(ctx, notificationIDs)
		if err != nil {
			// Log the error or handle it as needed
			log.Error(log.CustomLogInfo{
				"message": "Failed to mark notifications as read",
				"err":     err,
			}, "[notificationService.GetMyNotifications]")
		}
	}()

	res := dto.GetMyNotificationsResponse{
		Notifications: notificationsResponse,
	}

	return res, nil
}
