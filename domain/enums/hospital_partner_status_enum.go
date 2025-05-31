package enums

import (
	"database/sql/driver"
	"errors"
)

type HospitalPartnerStatusIdx int64
type HospitalPartnerStatusKey string
type HospitalPartnerStatusValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	HospitalPartnerStatusUnknownIdx    HospitalPartnerStatusIdx = 0
	HospitalPartnerStatusPendingIdx    HospitalPartnerStatusIdx = 1
	HospitalPartnerStatusInactiveIdx   HospitalPartnerStatusIdx = 2
	HospitalPartnerStatusActiveIdx     HospitalPartnerStatusIdx = 3
	HospitalPartnerStatusTerminatedIdx HospitalPartnerStatusIdx = 4

	HospitalPartnerStatusUnknownKey    HospitalPartnerStatusKey = "unknown"
	HospitalPartnerStatusPendingKey    HospitalPartnerStatusKey = "pending"
	HospitalPartnerStatusActiveKey     HospitalPartnerStatusKey = "active"
	HospitalPartnerStatusInactiveKey   HospitalPartnerStatusKey = "inactive"
	HospitalPartnerStatusTerminatedKey HospitalPartnerStatusKey = "terminated"
)

var (
	HospitalPartnerStatusUnknownValue = HospitalPartnerStatusValue{
		Idx:        int64(HospitalPartnerStatusUnknownIdx),
		Key:        string(HospitalPartnerStatusUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	HospitalPartnerStatusPendingValue = HospitalPartnerStatusValue{
		Idx:        int64(HospitalPartnerStatusPendingIdx),
		Key:        string(HospitalPartnerStatusPendingKey),
		LongLabel:  map[string]string{"id": "Menunggu", "en": "Pending"},
		ShortLabel: map[string]string{"id": "MNG", "en": "PND"},
	}
	HospitalPartnerStatusActiveValue = HospitalPartnerStatusValue{
		Idx:        int64(HospitalPartnerStatusActiveIdx),
		Key:        string(HospitalPartnerStatusActiveKey),
		LongLabel:  map[string]string{"id": "Aktif", "en": "Active"},
		ShortLabel: map[string]string{"id": "AKT", "en": "ACT"},
	}
	HospitalPartnerStatusInactiveValue = HospitalPartnerStatusValue{
		Idx:        int64(HospitalPartnerStatusInactiveIdx),
		Key:        string(HospitalPartnerStatusInactiveKey),
		LongLabel:  map[string]string{"id": "Tidak Aktif", "en": "Inactive"},
		ShortLabel: map[string]string{"id": "TKT", "en": "INA"},
	}
	HospitalPartnerStatusTerminatedValue = HospitalPartnerStatusValue{
		Idx:        int64(HospitalPartnerStatusTerminatedIdx),
		Key:        string(HospitalPartnerStatusTerminatedKey),
		LongLabel:  map[string]string{"id": "Dihentikan", "en": "Terminated"},
		ShortLabel: map[string]string{"id": "DHT", "en": "TRM"},
	}
)

var (
	HospitalPartnerStatusMapIdx = map[HospitalPartnerStatusIdx]HospitalPartnerStatusValue{
		HospitalPartnerStatusUnknownIdx:    HospitalPartnerStatusUnknownValue,
		HospitalPartnerStatusPendingIdx:    HospitalPartnerStatusPendingValue,
		HospitalPartnerStatusActiveIdx:     HospitalPartnerStatusActiveValue,
		HospitalPartnerStatusInactiveIdx:   HospitalPartnerStatusInactiveValue,
		HospitalPartnerStatusTerminatedIdx: HospitalPartnerStatusTerminatedValue,
	}

	HospitalPartnerStatusMapKey = map[HospitalPartnerStatusKey]HospitalPartnerStatusValue{
		HospitalPartnerStatusUnknownKey:    HospitalPartnerStatusUnknownValue,
		HospitalPartnerStatusActiveKey:     HospitalPartnerStatusActiveValue,
		HospitalPartnerStatusPendingKey:    HospitalPartnerStatusPendingValue,
		HospitalPartnerStatusInactiveKey:   HospitalPartnerStatusInactiveValue,
		HospitalPartnerStatusTerminatedKey: HospitalPartnerStatusTerminatedValue,
	}
)

func (p HospitalPartnerStatusIdx) String() string {
	if gender, ok := HospitalPartnerStatusMapIdx[p]; ok {
		return gender.Key
	}

	return string(HospitalPartnerStatusUnknownKey)
}

func (p *HospitalPartnerStatusIdx) Scan(value interface{}) error {
	if value == nil {
		*p = HospitalPartnerStatusUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*p = HospitalPartnerStatusIdx(v)
		return nil
	case int16:
		*p = HospitalPartnerStatusIdx(v)
		return nil
	case int32:
		*p = HospitalPartnerStatusIdx(v)
		return nil
	case int64:
		*p = HospitalPartnerStatusIdx(v)
		return nil
	case string:
		if val, ok := HospitalPartnerStatusMapKey[HospitalPartnerStatusKey(v)]; ok {
			*p = HospitalPartnerStatusIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := HospitalPartnerStatusMapKey[HospitalPartnerStatusKey(strVal)]; ok {
			*p = HospitalPartnerStatusIdx(val.Idx)
			return nil
		}
	}

	return errors.New("invalid gender value")
}

func (p HospitalPartnerStatusIdx) Value() (int64, error) {
	return int64(p), nil
}

type NullHospitalPartnerStatusIdx struct {
	HospitalPartnerStatusIdx HospitalPartnerStatusIdx
	Valid                    bool // Valid is true if HospitalPartnerStatusIdx is not NULL
}

func (p *NullHospitalPartnerStatusIdx) Scan(value interface{}) error {
	if value == nil {
		p.HospitalPartnerStatusIdx, p.Valid = HospitalPartnerStatusUnknownIdx, false
		return nil
	}

	p.Valid = true
	return p.HospitalPartnerStatusIdx.Scan(value)
}

func (p NullHospitalPartnerStatusIdx) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.HospitalPartnerStatusIdx.Value()
}
