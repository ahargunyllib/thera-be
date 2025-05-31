package enums

import (
	"database/sql/driver"
	"errors"
)

type HospitalTypeIdx int64
type HospitalTypeKey string
type HospitalTypeValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	HospitalTypeUnknownIdx        HospitalTypeIdx = 0
	HospitalTypeGeneralIdx        HospitalTypeIdx = 1
	HospitalTypeSpecialistIdx     HospitalTypeIdx = 2
	HospitalTypeClinicIdx         HospitalTypeIdx = 3
	HospitalTypeTeachingIdx       HospitalTypeIdx = 4
	HospitalTypePsychiatricIdx    HospitalTypeIdx = 5
	HospitalTypeRehabilitationIdx HospitalTypeIdx = 6

	HospitalTypeUnknownKey        HospitalTypeKey = "unknown"
	HospitalTypeGeneralKey        HospitalTypeKey = "general"
	HospitalTypeSpecialistKey     HospitalTypeKey = "specialist"
	HospitalTypeClinicKey         HospitalTypeKey = "clinic"
	HospitalTypeTeachingKey       HospitalTypeKey = "teaching"
	HospitalTypePsychiatricKey    HospitalTypeKey = "psychiatric"
	HospitalTypeRehabilitationKey HospitalTypeKey = "rehabilitation"
)

var (
	HospitalTypeUnknownValue = HospitalTypeValue{
		Idx:        int64(HospitalTypeUnknownIdx),
		Key:        string(HospitalTypeUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	HospitalTypeGeneralValue = HospitalTypeValue{
		Idx:        int64(HospitalTypeGeneralIdx),
		Key:        string(HospitalTypeGeneralKey),
		LongLabel:  map[string]string{"id": "Rumah Sakit Umum", "en": "General Hospital"},
		ShortLabel: map[string]string{"id": "Umum", "en": "General"},
	}
	HospitalTypeSpecialistValue = HospitalTypeValue{
		Idx:        int64(HospitalTypeSpecialistIdx),
		Key:        string(HospitalTypeSpecialistKey),
		LongLabel:  map[string]string{"id": "Rumah Sakit Khusus", "en": "Specialist Hospital"},
		ShortLabel: map[string]string{"id": "Khusus", "en": "Specialist"},
	}
	HospitalTypeClinicValue = HospitalTypeValue{
		Idx:        int64(HospitalTypeClinicIdx),
		Key:        string(HospitalTypeClinicKey),
		LongLabel:  map[string]string{"id": "Klinik", "en": "Clinic"},
		ShortLabel: map[string]string{"id": "Klinik", "en": "Clinic"},
	}
	HospitalTypeTeachingValue = HospitalTypeValue{
		Idx:        int64(HospitalTypeTeachingIdx),
		Key:        string(HospitalTypeTeachingKey),
		LongLabel:  map[string]string{"id": "Rumah Sakit Pendidikan", "en": "Teaching Hospital"},
		ShortLabel: map[string]string{"id": "Pendidikan", "en": "Teaching"},
	}
	HospitalTypePsychiatricValue = HospitalTypeValue{
		Idx:        int64(HospitalTypePsychiatricIdx),
		Key:        string(HospitalTypePsychiatricKey),
		LongLabel:  map[string]string{"id": "Rumah Sakit Jiwa", "en": "Psychiatric Hospital"},
		ShortLabel: map[string]string{"id": "Jiwa", "en": "Psych"},
	}
	HospitalTypeRehabilitationValue = HospitalTypeValue{
		Idx:        int64(HospitalTypeRehabilitationIdx),
		Key:        string(HospitalTypeRehabilitationKey),
		LongLabel:  map[string]string{"id": "Rumah Sakit Rehabilitasi", "en": "Rehabilitation Hospital"},
		ShortLabel: map[string]string{"id": "Rehab", "en": "Rehab"},
	}
)

var (
	HospitalTypeMapIdx = map[HospitalTypeIdx]HospitalTypeValue{
		HospitalTypeUnknownIdx:        HospitalTypeUnknownValue,
		HospitalTypeGeneralIdx:        HospitalTypeGeneralValue,
		HospitalTypeSpecialistIdx:     HospitalTypeSpecialistValue,
		HospitalTypeClinicIdx:         HospitalTypeClinicValue,
		HospitalTypeTeachingIdx:       HospitalTypeTeachingValue,
		HospitalTypePsychiatricIdx:    HospitalTypePsychiatricValue,
		HospitalTypeRehabilitationIdx: HospitalTypeRehabilitationValue,
	}
	HospitalTypeMapKey = map[HospitalTypeKey]HospitalTypeValue{
		HospitalTypeUnknownKey:        HospitalTypeUnknownValue,
		HospitalTypeGeneralKey:        HospitalTypeGeneralValue,
		HospitalTypeSpecialistKey:     HospitalTypeSpecialistValue,
		HospitalTypeClinicKey:         HospitalTypeClinicValue,
		HospitalTypeTeachingKey:       HospitalTypeTeachingValue,
		HospitalTypePsychiatricKey:    HospitalTypePsychiatricValue,
		HospitalTypeRehabilitationKey: HospitalTypeRehabilitationValue,
	}
)

func (t HospitalTypeIdx) String() string {
	if v, ok := HospitalTypeMapIdx[t]; ok {
		return v.Key
	}
	return string(HospitalTypeUnknownKey)
}

func (t *HospitalTypeIdx) Scan(value interface{}) error {
	if value == nil {
		*t = HospitalTypeUnknownIdx
		return nil
	}
	switch v := value.(type) {
	case int8:
		*t = HospitalTypeIdx(v)
		return nil
	case int16:
		*t = HospitalTypeIdx(v)
		return nil
	case int32:
		*t = HospitalTypeIdx(v)
		return nil
	case int64:
		*t = HospitalTypeIdx(v)
		return nil
	case string:
		if val, ok := HospitalTypeMapKey[HospitalTypeKey(v)]; ok {
			*t = HospitalTypeIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := HospitalTypeMapKey[HospitalTypeKey(strVal)]; ok {
			*t = HospitalTypeIdx(val.Idx)
			return nil
		}
	}

	return errors.New("failed to scan HospitalTypeIdx")
}

func (t HospitalTypeIdx) Value() (driver.Value, error) {
	return int64(t), nil
}

type NullHospitalTypeIdx struct {
	HospitalTypeIdx HospitalTypeIdx
	Valid           bool
}

func (n *NullHospitalTypeIdx) Scan(value interface{}) error {
	if value == nil {
		n.HospitalTypeIdx, n.Valid = HospitalTypeUnknownIdx, false
		return nil
	}

	n.Valid = true
	return n.HospitalTypeIdx.Scan(value)
}

func (n NullHospitalTypeIdx) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.HospitalTypeIdx.Value()
}
