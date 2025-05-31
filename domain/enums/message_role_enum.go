package enums

import (
	"database/sql/driver"
	"errors"
)

type MessageRoleIdx int64
type MessageRoleKey string
type MessageRoleValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	MessageRoleUnknownIdx   MessageRoleIdx = 0
	MessageRoleUserIdx      MessageRoleIdx = 1
	MessageRoleAssistantIdx MessageRoleIdx = 2

	MessageRoleUnknownKey   MessageRoleKey = "unknown"
	MessageRoleUserKey      MessageRoleKey = "user"
	MessageRoleAssistantKey MessageRoleKey = "assistant"
)

var (
	MessageRoleUnknownValue = MessageRoleValue{
		Idx:        int64(MessageRoleUnknownIdx),
		Key:        string(MessageRoleUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	MessageRoleUserValue = MessageRoleValue{
		Idx:        int64(MessageRoleUserIdx),
		Key:        string(MessageRoleUserKey),
		LongLabel:  map[string]string{"id": "Pengguna", "en": "User"},
		ShortLabel: map[string]string{"id": "User", "en": "User"},
	}
	MessageRoleAssistantValue = MessageRoleValue{
		Idx:        int64(MessageRoleAssistantIdx),
		Key:        string(MessageRoleAssistantKey),
		LongLabel:  map[string]string{"id": "Asisten", "en": "Assistant"},
		ShortLabel: map[string]string{"id": "Asisten", "en": "Assistant"},
	}
)

var (
	MessageRoleMapIdx = map[MessageRoleIdx]MessageRoleValue{
		MessageRoleUnknownIdx:   MessageRoleUnknownValue,
		MessageRoleUserIdx:      MessageRoleUserValue,
		MessageRoleAssistantIdx: MessageRoleAssistantValue,
	}

	MessageRoleMapKey = map[MessageRoleKey]MessageRoleValue{
		MessageRoleUnknownKey:   MessageRoleUnknownValue,
		MessageRoleUserKey:      MessageRoleUserValue,
		MessageRoleAssistantKey: MessageRoleAssistantValue,
	}
)

func (m MessageRoleIdx) String() string {
	if role, ok := MessageRoleMapIdx[m]; ok {
		return role.Key
	}
	return string(MessageRoleUnknownKey)
}

func (m *MessageRoleIdx) Scan(value interface{}) error {
	if value == nil {
		*m = MessageRoleUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*m = MessageRoleIdx(v)
		return nil
	case int16:
		*m = MessageRoleIdx(v)
		return nil
	case int32:
		*m = MessageRoleIdx(v)
		return nil
	case int64:
		*m = MessageRoleIdx(v)
		return nil
	case int:
		*m = MessageRoleIdx(v)
		return nil
	case string:
		if val, ok := MessageRoleMapKey[MessageRoleKey(v)]; ok {
			*m = MessageRoleIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := MessageRoleMapKey[MessageRoleKey(strVal)]; ok {
			*m = MessageRoleIdx(val.Idx)
			return nil
		}
	}

	return errors.New("failed to scan MessageRoleIdx")
}

func (m MessageRoleIdx) Value() (driver.Value, error) {
	return int64(m), nil
}

type NullMessageRoleIdx struct {
	MessageRoleIdx MessageRoleIdx
	Valid          bool // Valid is true if MessageRoleIdx is not NULL
}

func (n *NullMessageRoleIdx) Scan(value interface{}) error {
	if value == nil {
		n.MessageRoleIdx, n.Valid = MessageRoleUnknownIdx, false
		return nil
	}

	n.Valid = true
	return n.MessageRoleIdx.Scan(value)
}

func (n NullMessageRoleIdx) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.MessageRoleIdx.Value()
}
