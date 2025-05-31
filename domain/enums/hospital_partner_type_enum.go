package enums

import (
	"database/sql/driver"
	"errors"
)

type HospitalPartnerTypeIdx int64
type HospitalPartnerTypeKey string
type HospitalPartnerTypeValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	HospitalPartnerTypeUnknownIdx       HospitalPartnerTypeIdx = 0
	HospitalPartnerTypeCollaborationIdx HospitalPartnerTypeIdx = 1

	HospitalPartnerTypeUnknownKey       HospitalPartnerTypeKey = "unknown"
	HospitalPartnerTypeCollaborationKey HospitalPartnerTypeKey = "collaboration"
)

var (
	HospitalPartnerTypeUnknownValue = HospitalPartnerTypeValue{
		Idx:        int64(HospitalPartnerTypeUnknownIdx),
		Key:        string(HospitalPartnerTypeUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	HospitalPartnerTypeCollaborationValue = HospitalPartnerTypeValue{
		Idx:        int64(HospitalPartnerTypeCollaborationIdx),
		Key:        string(HospitalPartnerTypeCollaborationKey),
		LongLabel:  map[string]string{"id": "Kolaborasi", "en": "Collaboration"},
		ShortLabel: map[string]string{"id": "KOL", "en": "COL"},
	}
)

var (
	HospitalPartnerTypeMapIdx = map[HospitalPartnerTypeIdx]HospitalPartnerTypeValue{
		HospitalPartnerTypeUnknownIdx:       HospitalPartnerTypeUnknownValue,
		HospitalPartnerTypeCollaborationIdx: HospitalPartnerTypeCollaborationValue,
	}

	HospitalPartnerTypeMapKey = map[HospitalPartnerTypeKey]HospitalPartnerTypeValue{
		HospitalPartnerTypeUnknownKey:       HospitalPartnerTypeUnknownValue,
		HospitalPartnerTypeCollaborationKey: HospitalPartnerTypeCollaborationValue,
	}
)

func (p HospitalPartnerTypeIdx) String() string {
	if gender, ok := HospitalPartnerTypeMapIdx[p]; ok {
		return gender.Key
	}

	return string(HospitalPartnerTypeUnknownKey)
}

func (p *HospitalPartnerTypeIdx) Scan(value interface{}) error {
	if value == nil {
		*p = HospitalPartnerTypeUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*p = HospitalPartnerTypeIdx(v)
		return nil
	case int16:
		*p = HospitalPartnerTypeIdx(v)
		return nil
	case int32:
		*p = HospitalPartnerTypeIdx(v)
		return nil
	case int64:
		*p = HospitalPartnerTypeIdx(v)
		return nil
	case string:
		if val, ok := HospitalPartnerTypeMapKey[HospitalPartnerTypeKey(v)]; ok {
			*p = HospitalPartnerTypeIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := HospitalPartnerTypeMapKey[HospitalPartnerTypeKey(strVal)]; ok {
			*p = HospitalPartnerTypeIdx(val.Idx)
			return nil
		}
	}

	return errors.New("invalid gender value")
}

func (p HospitalPartnerTypeIdx) Value() (int64, error) {
	return int64(p), nil
}

type NullHospitalPartnerTypeIdx struct {
	HospitalPartnerTypeIdx HospitalPartnerTypeIdx
	Valid                  bool // Valid is true if HospitalPartnerTypeIdx is not NULL
}

func (p *NullHospitalPartnerTypeIdx) Scan(value interface{}) error {
	if value == nil {
		p.HospitalPartnerTypeIdx, p.Valid = HospitalPartnerTypeUnknownIdx, false
		return nil
	}

	p.Valid = true
	return p.HospitalPartnerTypeIdx.Scan(value)
}

func (p NullHospitalPartnerTypeIdx) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.HospitalPartnerTypeIdx.Value()
}
