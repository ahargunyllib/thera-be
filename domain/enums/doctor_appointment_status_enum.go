package enums

import (
	"database/sql/driver"
	"errors"
)

// DoctorAppointmentStatus enum
type DoctorAppointmentStatusIdx int64
type DoctorAppointmentStatusKey string
type DoctorAppointmentStatusValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	DoctorAppointmentStatusUnknownIdx     DoctorAppointmentStatusIdx = 0
	DoctorAppointmentStatusPendingIdx     DoctorAppointmentStatusIdx = 1
	DoctorAppointmentStatusConfirmedIdx   DoctorAppointmentStatusIdx = 2
	DoctorAppointmentStatusCompletedIdx   DoctorAppointmentStatusIdx = 3
	DoctorAppointmentStatusCancelledIdx   DoctorAppointmentStatusIdx = 4
	DoctorAppointmentStatusRescheduledIdx DoctorAppointmentStatusIdx = 5

	DoctorAppointmentStatusUnknownKey     DoctorAppointmentStatusKey = "unknown"
	DoctorAppointmentStatusPendingKey     DoctorAppointmentStatusKey = "pending"
	DoctorAppointmentStatusConfirmedKey   DoctorAppointmentStatusKey = "confirmed"
	DoctorAppointmentStatusCompletedKey   DoctorAppointmentStatusKey = "completed"
	DoctorAppointmentStatusCancelledKey   DoctorAppointmentStatusKey = "cancelled"
	DoctorAppointmentStatusRescheduledKey DoctorAppointmentStatusKey = "rescheduled"
)

var (
	DoctorAppointmentStatusUnknownValue = DoctorAppointmentStatusValue{
		Idx:        int64(DoctorAppointmentStatusUnknownIdx),
		Key:        string(DoctorAppointmentStatusUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	DoctorAppointmentStatusPendingValue = DoctorAppointmentStatusValue{
		Idx:        int64(DoctorAppointmentStatusPendingIdx),
		Key:        string(DoctorAppointmentStatusPendingKey),
		LongLabel:  map[string]string{"id": "Menunggu", "en": "Pending"},
		ShortLabel: map[string]string{"id": "MTG", "en": "PND"},
	}
	DoctorAppointmentStatusConfirmedValue = DoctorAppointmentStatusValue{
		Idx:        int64(DoctorAppointmentStatusConfirmedIdx),
		Key:        string(DoctorAppointmentStatusConfirmedKey),
		LongLabel:  map[string]string{"id": "Dikonfirmasi", "en": "Confirmed"},
		ShortLabel: map[string]string{"id": "KFM", "en": "CFM"},
	}
	DoctorAppointmentStatusCompletedValue = DoctorAppointmentStatusValue{
		Idx:        int64(DoctorAppointmentStatusCompletedIdx),
		Key:        string(DoctorAppointmentStatusCompletedKey),
		LongLabel:  map[string]string{"id": "Selesai", "en": "Completed"},
		ShortLabel: map[string]string{"id": "SLS", "en": "CMP"},
	}
	DoctorAppointmentStatusCancelledValue = DoctorAppointmentStatusValue{
		Idx:        int64(DoctorAppointmentStatusCancelledIdx),
		Key:        string(DoctorAppointmentStatusCancelledKey),
		LongLabel:  map[string]string{"id": "Dibatalkan", "en": "Cancelled"},
		ShortLabel: map[string]string{"id": "BTL", "en": "CNC"},
	}
	DoctorAppointmentStatusRescheduledValue = DoctorAppointmentStatusValue{
		Idx:        int64(DoctorAppointmentStatusRescheduledIdx),
		Key:        string(DoctorAppointmentStatusRescheduledKey),
		LongLabel:  map[string]string{"id": "Dijadwal Ulang", "en": "Rescheduled"},
		ShortLabel: map[string]string{"id": "JDU", "en": "RSC"},
	}
)

var (
	DoctorAppointmentStatusMapIdx = map[DoctorAppointmentStatusIdx]DoctorAppointmentStatusValue{
		DoctorAppointmentStatusUnknownIdx:     DoctorAppointmentStatusUnknownValue,
		DoctorAppointmentStatusPendingIdx:     DoctorAppointmentStatusPendingValue,
		DoctorAppointmentStatusConfirmedIdx:   DoctorAppointmentStatusConfirmedValue,
		DoctorAppointmentStatusCompletedIdx:   DoctorAppointmentStatusCompletedValue,
		DoctorAppointmentStatusCancelledIdx:   DoctorAppointmentStatusCancelledValue,
		DoctorAppointmentStatusRescheduledIdx: DoctorAppointmentStatusRescheduledValue,
	}
	DoctorAppointmentStatusMapKey = map[DoctorAppointmentStatusKey]DoctorAppointmentStatusValue{
		DoctorAppointmentStatusUnknownKey:     DoctorAppointmentStatusUnknownValue,
		DoctorAppointmentStatusPendingKey:     DoctorAppointmentStatusPendingValue,
		DoctorAppointmentStatusConfirmedKey:   DoctorAppointmentStatusConfirmedValue,
		DoctorAppointmentStatusCompletedKey:   DoctorAppointmentStatusCompletedValue,
		DoctorAppointmentStatusCancelledKey:   DoctorAppointmentStatusCancelledValue,
		DoctorAppointmentStatusRescheduledKey: DoctorAppointmentStatusRescheduledValue,
	}
)

func (p DoctorAppointmentStatusIdx) String() string {
	if status, ok := DoctorAppointmentStatusMapIdx[p]; ok {
		return status.Key
	}
	return string(DoctorAppointmentStatusPendingKey)
}

func (p *DoctorAppointmentStatusIdx) Scan(value interface{}) error {
	if value == nil {
		*p = DoctorAppointmentStatusPendingIdx
		return nil
	}
	switch v := value.(type) {
	case int8:
		*p = DoctorAppointmentStatusIdx(v)
		return nil
	case int16:
		*p = DoctorAppointmentStatusIdx(v)
		return nil
	case int32:
		*p = DoctorAppointmentStatusIdx(v)
		return nil
	case int64:
		*p = DoctorAppointmentStatusIdx(v)
		return nil
	case string:
		if val, ok := DoctorAppointmentStatusMapKey[DoctorAppointmentStatusKey(v)]; ok {
			*p = DoctorAppointmentStatusIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := DoctorAppointmentStatusMapKey[DoctorAppointmentStatusKey(strVal)]; ok {
			*p = DoctorAppointmentStatusIdx(val.Idx)
			return nil
		}
	}
	return errors.New("invalid doctor appointment status value")
}

func (p DoctorAppointmentStatusIdx) Value() (driver.Value, error) {
	return int64(p), nil
}

type NullDoctorAppointmentStatusIdx struct {
	DoctorAppointmentStatusIdx DoctorAppointmentStatusIdx
	Valid                      bool // Valid is true if DoctorAppointmentStatusIdx is not NULL
}

func (p *NullDoctorAppointmentStatusIdx) Scan(value interface{}) error {
	if value == nil {
		p.DoctorAppointmentStatusIdx, p.Valid = DoctorAppointmentStatusPendingIdx, false
		return nil
	}
	p.Valid = true
	return p.DoctorAppointmentStatusIdx.Scan(value)
}

func (p NullDoctorAppointmentStatusIdx) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}
	return p.DoctorAppointmentStatusIdx.Value()
}
