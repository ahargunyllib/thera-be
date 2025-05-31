package dto

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/entity"
	"github.com/ahargunyllib/thera-be/domain/enums"
	"github.com/google/uuid"
)

type NotificationResponse struct {
	ID        string    `json:"id"` // ulid
	Title     string    `json:"title"`
	Body      string    `json:"body,omitempty"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNotificationResponse(notificationEntity *entity.Notification) NotificationResponse {
	return NotificationResponse{
		ID:        notificationEntity.ID,
		Title:     notificationEntity.Title,
		Body:      notificationEntity.Body.String,
		Type:      enums.NotificationTypeMapIdx[notificationEntity.Type].LongLabel["id"],
		IsRead:    notificationEntity.ReadAt.Valid,
		CreatedAt: notificationEntity.CreatedAt,
	}
}

type GetMyNotificationsQuery struct {
	DoctorID   uuid.UUID `validate:"uuid"` // UUID of the doctor
	HospitalID int       `validate:"uuid"` // UUID of the hospital
}

type GetMyNotificationsResponse struct {
	Notifications []NotificationResponse `json:"notifications"`
}
