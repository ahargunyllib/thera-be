package contracts

import (
	"context"

	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/google/uuid"
)

type NotificationRepository interface {
	GetNotificationsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]entity.Notification, error)
	GetNotificationsByHospitalID(ctx context.Context, hospitalID int) ([]entity.Notification, error)
	ReadNotifications(ctx context.Context, notificationIDs []string) error
}

type NotificationService interface {
	GetMyNotifications(ctx context.Context, query dto.GetMyNotificationsQuery) (dto.GetMyNotificationsResponse, error)
}
