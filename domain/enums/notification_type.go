package enums

import (
	"database/sql/driver"
	"errors"
)

type NotificationTypeIdx int64
type NotificationTypeKey string
type NotificationTypeValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	LongLabel  map[string]string `json:"long_label"`
	ShortLabel map[string]string `json:"short_label"`
}

const (
	NotificationTypeUnknownIdx NotificationTypeIdx = 0
	NotificationTypeBasicIdx   NotificationTypeIdx = 1

	NotificationTypeUnknownKey NotificationTypeKey = "unknown"
	NotificationTypeBasicKey   NotificationTypeKey = "basic"
)

var (
	NotificationTypeUnknownValue = NotificationTypeValue{
		Idx:        int64(NotificationTypeUnknownIdx),
		Key:        string(NotificationTypeUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}

	NotificationTypeBasicValue = NotificationTypeValue{
		Idx:        int64(NotificationTypeBasicIdx),
		Key:        string(NotificationTypeBasicKey),
		LongLabel:  map[string]string{"id": "Dasar", "en": "Basic"},
		ShortLabel: map[string]string{"id": "Dasar", "en": "Basic"},
	}
)

var (
	NotificationTypeMapIdx = map[NotificationTypeIdx]NotificationTypeValue{
		NotificationTypeUnknownIdx: NotificationTypeUnknownValue,
		NotificationTypeBasicIdx:   NotificationTypeBasicValue,
	}

	NotificationTypeMapKey = map[NotificationTypeKey]NotificationTypeValue{
		NotificationTypeUnknownKey: NotificationTypeUnknownValue,
		NotificationTypeBasicKey:   NotificationTypeBasicValue,
	}
)

func (m NotificationTypeIdx) String() string {
	if role, ok := NotificationTypeMapIdx[m]; ok {
		return role.Key
	}
	return string(NotificationTypeUnknownKey)
}

func (m *NotificationTypeIdx) Scan(value interface{}) error {
	if value == nil {
		*m = NotificationTypeUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*m = NotificationTypeIdx(v)
		return nil
	case int16:
		*m = NotificationTypeIdx(v)
		return nil
	case int32:
		*m = NotificationTypeIdx(v)
		return nil
	case int64:
		*m = NotificationTypeIdx(v)
		return nil
	case int:
		*m = NotificationTypeIdx(v)
		return nil
	case string:
		if val, ok := NotificationTypeMapKey[NotificationTypeKey(v)]; ok {
			*m = NotificationTypeIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := NotificationTypeMapKey[NotificationTypeKey(strVal)]; ok {
			*m = NotificationTypeIdx(val.Idx)
			return nil
		}
	}

	return errors.New("failed to scan NotificationTypeIdx")
}

func (m NotificationTypeIdx) Value() (driver.Value, error) {
	return int64(m), nil
}
