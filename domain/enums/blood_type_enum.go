package enums

import (
	"database/sql/driver"
	"errors"
)

type BloodTypeIdx int64
type BloodTypeKey string
type BloodTypeValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	BloodTypeUnknownIdx BloodTypeIdx = 0
	BloodTypeAIdx       BloodTypeIdx = 1
	BloodTypeBIdx       BloodTypeIdx = 2
	BloodTypeABIdx      BloodTypeIdx = 3
	BloodTypeOIdx       BloodTypeIdx = 4

	BloodTypeUnknownKey BloodTypeKey = "unknown"
	BloodTypeAKey       BloodTypeKey = "a"
	BloodTypeBKey       BloodTypeKey = "b"
	BloodTypeABKey      BloodTypeKey = "ab"
	BloodTypeOKey       BloodTypeKey = "o"
)

var (
	BloodTypeUnknownValue = BloodTypeValue{
		Idx:        int64(BloodTypeUnknownIdx),
		Key:        string(BloodTypeUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	BloodTypeAValue = BloodTypeValue{
		Idx:        int64(BloodTypeAIdx),
		Key:        string(BloodTypeAKey),
		LongLabel:  map[string]string{"id": "Golongan Darah A", "en": "Blood Type A"},
		ShortLabel: map[string]string{"id": "A", "en": "A"},
	}
	BloodTypeBValue = BloodTypeValue{
		Idx:        int64(BloodTypeBIdx),
		Key:        string(BloodTypeBKey),
		LongLabel:  map[string]string{"id": "Golongan Darah B", "en": "Blood Type B"},
		ShortLabel: map[string]string{"id": "B", "en": "B"},
	}
	BloodTypeABValue = BloodTypeValue{
		Idx:        int64(BloodTypeABIdx),
		Key:        string(BloodTypeABKey),
		LongLabel:  map[string]string{"id": "Golongan Darah AB", "en": "Blood Type AB"},
		ShortLabel: map[string]string{"id": "AB", "en": "AB"},
	}
	BloodTypeOValue = BloodTypeValue{
		Idx:        int64(BloodTypeOIdx),
		Key:        string(BloodTypeOKey),
		LongLabel:  map[string]string{"id": "Golongan Darah O", "en": "Blood Type O"},
		ShortLabel: map[string]string{"id": "O", "en": "O"},
	}
)

var (
	BloodTypeMapIdx = map[BloodTypeIdx]BloodTypeValue{
		BloodTypeUnknownIdx: BloodTypeUnknownValue,
		BloodTypeAIdx:       BloodTypeAValue,
		BloodTypeBIdx:       BloodTypeBValue,
		BloodTypeABIdx:      BloodTypeABValue,
		BloodTypeOIdx:       BloodTypeOValue,
	}

	BloodTypeMapKey = map[BloodTypeKey]BloodTypeValue{
		BloodTypeUnknownKey: BloodTypeUnknownValue,
		BloodTypeAKey:       BloodTypeAValue,
		BloodTypeBKey:       BloodTypeBValue,
		BloodTypeABKey:      BloodTypeABValue,
		BloodTypeOKey:       BloodTypeOValue,
	}
)

func (b BloodTypeIdx) String() string {
	if bt, ok := BloodTypeMapIdx[b]; ok {
		return bt.Key
	}
	return string(BloodTypeUnknownKey)
}

func (b *BloodTypeIdx) Scan(value interface{}) error {
	if value == nil {
		*b = BloodTypeUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*b = BloodTypeIdx(v)
		return nil
	case int16:
		*b = BloodTypeIdx(v)
		return nil
	case int32:
		*b = BloodTypeIdx(v)
		return nil
	case int64:
		*b = BloodTypeIdx(v)
		return nil
	case string:
		if val, ok := BloodTypeMapKey[BloodTypeKey(v)]; ok {
			*b = BloodTypeIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := BloodTypeMapKey[BloodTypeKey(strVal)]; ok {
			*b = BloodTypeIdx(val.Idx)
			return nil
		}
	}

	return errors.New("invalid blood type value")
}

func (b BloodTypeIdx) Value() (int64, error) {
	return int64(b), nil
}

type NullBloodTypeIdx struct {
	BloodTypeIdx BloodTypeIdx
	Valid        bool // Valid is true if BloodTypeIdx is not NULL
}

func (b *NullBloodTypeIdx) Scan(value interface{}) error {
	if value == nil {
		b.BloodTypeIdx, b.Valid = BloodTypeUnknownIdx, false
		return nil
	}

	b.Valid = true
	return b.BloodTypeIdx.Scan(value)
}

func (b NullBloodTypeIdx) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}

	return b.BloodTypeIdx.Value()
}
