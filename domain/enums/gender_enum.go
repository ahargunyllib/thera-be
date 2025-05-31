package enums

import (
	"database/sql/driver"
	"errors"
)

type GenderIdx int64
type GenderKey string
type GenderValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	GenderUnknownIdx GenderIdx = 0
	GenderMaleIdx    GenderIdx = 1
	GenderFemaleIdx  GenderIdx = 2

	GenderUnknownKey GenderKey = "unknown"
	GenderMaleKey    GenderKey = "male"
	GenderFemaleKey  GenderKey = "female"
)

var (
	GenderUnknownValue = GenderValue{
		Idx:        int64(GenderUnknownIdx),
		Key:        string(GenderUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	GenderMaleValue = GenderValue{
		Idx:        int64(GenderMaleIdx),
		Key:        string(GenderMaleKey),
		LongLabel:  map[string]string{"id": "Pria", "en": "Man"},
		ShortLabel: map[string]string{"id": "L", "en": "M"},
	}
	GenderFemaleValue = GenderValue{
		Idx:        int64(GenderFemaleIdx),
		Key:        string(GenderFemaleKey),
		LongLabel:  map[string]string{"id": "Wanita", "en": "Women"},
		ShortLabel: map[string]string{"id": "P", "en": "W"},
	}
)

var (
	GenderMapIdx = map[GenderIdx]GenderValue{
		GenderUnknownIdx: GenderUnknownValue,
		GenderMaleIdx:    GenderMaleValue,
		GenderFemaleIdx:  GenderFemaleValue,
	}

	GenderMapKey = map[GenderKey]GenderValue{
		GenderUnknownKey: GenderUnknownValue,
		GenderMaleKey:    GenderMaleValue,
		GenderFemaleKey:  GenderFemaleValue,
	}
)

func (p GenderIdx) String() string {
	if gender, ok := GenderMapIdx[p]; ok {
		return gender.Key
	}

	return string(GenderUnknownKey)
}

func (p *GenderIdx) Scan(value interface{}) error {
	if value == nil {
		*p = GenderUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*p = GenderIdx(v)
		return nil
	case int16:
		*p = GenderIdx(v)
		return nil
	case int32:
		*p = GenderIdx(v)
		return nil
	case int64:
		*p = GenderIdx(v)
		return nil
	case string:
		if val, ok := GenderMapKey[GenderKey(v)]; ok {
			*p = GenderIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := GenderMapKey[GenderKey(strVal)]; ok {
			*p = GenderIdx(val.Idx)
			return nil
		}
	}

	return errors.New("invalid gender value")
}

func (p GenderIdx) Value() (int64, error) {
	return int64(p), nil
}

type NullGenderIdx struct {
	GenderIdx GenderIdx
	Valid     bool // Valid is true if GenderIdx is not NULL
}

func (p *NullGenderIdx) Scan(value interface{}) error {
	if value == nil {
		p.GenderIdx, p.Valid = GenderUnknownIdx, false
		return nil
	}

	p.Valid = true
	return p.GenderIdx.Scan(value)
}

func (p NullGenderIdx) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.GenderIdx.Value()
}
