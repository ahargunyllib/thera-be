package enums

import (
	"database/sql/driver"
	"errors"
)

type HospitalStatusIdx int64
type HospitalStatusKey string
type HospitalStatusValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	HospitalStatusUnknownIdx           HospitalStatusIdx = 0
	HospitalStatusActiveIdx            HospitalStatusIdx = 1
	HospitalStatusTemporaryClosedIdx   HospitalStatusIdx = 2
	HospitalStatusPermanentClosedIdx   HospitalStatusIdx = 3
	HospitalStatusUnderConstructionIdx HospitalStatusIdx = 4
	HospitalStatusPlannedIdx           HospitalStatusIdx = 5

	HospitalStatusUnknownKey           HospitalStatusKey = "unknown"
	HospitalStatusActiveKey            HospitalStatusKey = "active"
	HospitalStatusTemporaryClosedKey   HospitalStatusKey = "temporary_closed"
	HospitalStatusPermanentClosedKey   HospitalStatusKey = "permanent_closed"
	HospitalStatusUnderConstructionKey HospitalStatusKey = "under_construction"
	HospitalStatusPlannedKey           HospitalStatusKey = "planned"
)

var (
	HospitalStatusUnknownValue = HospitalStatusValue{
		Idx:        int64(HospitalStatusUnknownIdx),
		Key:        string(HospitalStatusUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	HospitalStatusActiveValue = HospitalStatusValue{
		Idx:        int64(HospitalStatusActiveIdx),
		Key:        string(HospitalStatusActiveKey),
		LongLabel:  map[string]string{"id": "Beroperasi", "en": "Active"},
		ShortLabel: map[string]string{"id": "Aktif", "en": "Active"},
	}
	HospitalStatusTemporaryClosedValue = HospitalStatusValue{
		Idx:        int64(HospitalStatusTemporaryClosedIdx),
		Key:        string(HospitalStatusTemporaryClosedKey),
		LongLabel:  map[string]string{"id": "Ditutup Sementara", "en": "Temporarily Closed"},
		ShortLabel: map[string]string{"id": "Sementara", "en": "Temp Closed"},
	}
	HospitalStatusPermanentClosedValue = HospitalStatusValue{
		Idx:        int64(HospitalStatusPermanentClosedIdx),
		Key:        string(HospitalStatusPermanentClosedKey),
		LongLabel:  map[string]string{"id": "Ditutup Permanen", "en": "Permanently Closed"},
		ShortLabel: map[string]string{"id": "Permanen", "en": "Perm Closed"},
	}
	HospitalStatusUnderConstructionValue = HospitalStatusValue{
		Idx:        int64(HospitalStatusUnderConstructionIdx),
		Key:        string(HospitalStatusUnderConstructionKey),
		LongLabel:  map[string]string{"id": "Dalam Pembangunan", "en": "Under Construction"},
		ShortLabel: map[string]string{"id": "Pembangunan", "en": "Construction"},
	}
	HospitalStatusPlannedValue = HospitalStatusValue{
		Idx:        int64(HospitalStatusPlannedIdx),
		Key:        string(HospitalStatusPlannedKey),
		LongLabel:  map[string]string{"id": "Dalam Perencanaan", "en": "Planned"},
		ShortLabel: map[string]string{"id": "Rencana", "en": "Planned"},
	}
)

var (
	HospitalStatusMapIdx = map[HospitalStatusIdx]HospitalStatusValue{
		HospitalStatusUnknownIdx:           HospitalStatusUnknownValue,
		HospitalStatusActiveIdx:            HospitalStatusActiveValue,
		HospitalStatusTemporaryClosedIdx:   HospitalStatusTemporaryClosedValue,
		HospitalStatusPermanentClosedIdx:   HospitalStatusPermanentClosedValue,
		HospitalStatusUnderConstructionIdx: HospitalStatusUnderConstructionValue,
		HospitalStatusPlannedIdx:           HospitalStatusPlannedValue,
	}
	HospitalStatusMapKey = map[HospitalStatusKey]HospitalStatusValue{
		HospitalStatusUnknownKey:           HospitalStatusUnknownValue,
		HospitalStatusActiveKey:            HospitalStatusActiveValue,
		HospitalStatusTemporaryClosedKey:   HospitalStatusTemporaryClosedValue,
		HospitalStatusPermanentClosedKey:   HospitalStatusPermanentClosedValue,
		HospitalStatusUnderConstructionKey: HospitalStatusUnderConstructionValue,
		HospitalStatusPlannedKey:           HospitalStatusPlannedValue,
	}
)

func (t HospitalStatusIdx) String() string {
	if v, ok := HospitalStatusMapIdx[t]; ok {
		return v.Key
	}
	return string(HospitalStatusUnknownKey)
}

func (t *HospitalStatusIdx) Scan(value interface{}) error {
	if value == nil {
		*t = HospitalStatusUnknownIdx
		return nil
	}
	switch v := value.(type) {
	case int8:
		*t = HospitalStatusIdx(v)
		return nil
	case int16:
		*t = HospitalStatusIdx(v)
		return nil
	case int32:
		*t = HospitalStatusIdx(v)
		return nil
	case int64:
		*t = HospitalStatusIdx(v)
		return nil
	case string:
		if val, ok := HospitalStatusMapKey[HospitalStatusKey(v)]; ok {
			*t = HospitalStatusIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := HospitalStatusMapKey[HospitalStatusKey(strVal)]; ok {
			*t = HospitalStatusIdx(val.Idx)
			return nil
		}
	}

	return errors.New("failed to scan HospitalStatusIdx")
}

func (t HospitalStatusIdx) Value() (driver.Value, error) {
	return int64(t), nil
}

type NullHospitalStatusIdx struct {
	HospitalStatusIdx HospitalStatusIdx
	Valid             bool // Valid is true if HospitalStatusIdx is not NULL
}

func (n *NullHospitalStatusIdx) Scan(value interface{}) error {
	if value == nil {
		n.HospitalStatusIdx, n.Valid = HospitalStatusUnknownIdx, false
		return nil
	}

	n.Valid = true
	return n.HospitalStatusIdx.Scan(value)
}

func (n NullHospitalStatusIdx) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.HospitalStatusIdx.Value()
}
