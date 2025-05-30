package uuid

import (
	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/google/uuid"
)

type UUIDInterface interface {
	NewV7() (uuid.UUID, error)
}

type UUIDStruct struct{}

var UUID = getUUID()

func getUUID() UUIDInterface {
	return &UUIDStruct{}
}

func (u *UUIDStruct) NewV7() (uuid.UUID, error) {
	uuid, err := uuid.NewV7()

	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[UUID][New] failed to create uuid v7")

		return uuid, err
	}

	return uuid, err
}
