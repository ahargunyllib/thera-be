package enums

import (
	"database/sql/driver"
	"errors"
)

// DoctorAppointmentType enum
type DoctorAppointmentTypeIdx int64
type DoctorAppointmentTypeKey string
type DoctorAppointmentTypeValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	DoctorAppointmentTypeUnknownIdx      DoctorAppointmentTypeIdx = 0
	DoctorAppointmentTypeConsultationIdx DoctorAppointmentTypeIdx = 1
	DoctorAppointmentTypeCheckupIdx      DoctorAppointmentTypeIdx = 2
	DoctorAppointmentTypeFollowupIdx     DoctorAppointmentTypeIdx = 3
	DoctorAppointmentTypeEmergencyIdx    DoctorAppointmentTypeIdx = 4
	DoctorAppointmentTypeTelemedicineIdx DoctorAppointmentTypeIdx = 5
	DoctorAppointmentTypeSurgeryIdx      DoctorAppointmentTypeIdx = 6

	DoctorAppointmentTypeUnknownKey      DoctorAppointmentTypeKey = "unknown"
	DoctorAppointmentTypeConsultationKey DoctorAppointmentTypeKey = "consultation"
	DoctorAppointmentTypeCheckupKey      DoctorAppointmentTypeKey = "checkup"
	DoctorAppointmentTypeFollowupKey     DoctorAppointmentTypeKey = "followup"
	DoctorAppointmentTypeEmergencyKey    DoctorAppointmentTypeKey = "emergency"
	DoctorAppointmentTypeTelemedicineKey DoctorAppointmentTypeKey = "telemedicine"
	DoctorAppointmentTypeSurgeryKey      DoctorAppointmentTypeKey = "surgery"
)

var (
	DoctorAppointmentTypeUnknownValue = DoctorAppointmentTypeValue{
		Idx:        int64(DoctorAppointmentTypeUnknownIdx),
		Key:        string(DoctorAppointmentTypeUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	DoctorAppointmentTypeConsultationValue = DoctorAppointmentTypeValue{
		Idx:        int64(DoctorAppointmentTypeConsultationIdx),
		Key:        string(DoctorAppointmentTypeConsultationKey),
		LongLabel:  map[string]string{"id": "Konsultasi", "en": "Consultation"},
		ShortLabel: map[string]string{"id": "KNS", "en": "CNS"},
	}
	DoctorAppointmentTypeCheckupValue = DoctorAppointmentTypeValue{
		Idx:        int64(DoctorAppointmentTypeCheckupIdx),
		Key:        string(DoctorAppointmentTypeCheckupKey),
		LongLabel:  map[string]string{"id": "Pemeriksaan", "en": "Check-up"},
		ShortLabel: map[string]string{"id": "PMR", "en": "CHK"},
	}
	DoctorAppointmentTypeFollowupValue = DoctorAppointmentTypeValue{
		Idx:        int64(DoctorAppointmentTypeFollowupIdx),
		Key:        string(DoctorAppointmentTypeFollowupKey),
		LongLabel:  map[string]string{"id": "Tindak Lanjut", "en": "Follow-up"},
		ShortLabel: map[string]string{"id": "TDL", "en": "FLW"},
	}
	DoctorAppointmentTypeEmergencyValue = DoctorAppointmentTypeValue{
		Idx:        int64(DoctorAppointmentTypeEmergencyIdx),
		Key:        string(DoctorAppointmentTypeEmergencyKey),
		LongLabel:  map[string]string{"id": "Darurat", "en": "Emergency"},
		ShortLabel: map[string]string{"id": "DRT", "en": "EMR"},
	}
	DoctorAppointmentTypeTelemedicineValue = DoctorAppointmentTypeValue{
		Idx:        int64(DoctorAppointmentTypeTelemedicineIdx),
		Key:        string(DoctorAppointmentTypeTelemedicineKey),
		LongLabel:  map[string]string{"id": "Telemedicine", "en": "Telemedicine"},
		ShortLabel: map[string]string{"id": "TLM", "en": "TLM"},
	}
	DoctorAppointmentTypeSurgeryValue = DoctorAppointmentTypeValue{
		Idx:        int64(DoctorAppointmentTypeSurgeryIdx),
		Key:        string(DoctorAppointmentTypeSurgeryKey),
		LongLabel:  map[string]string{"id": "Operasi", "en": "Surgery"},
		ShortLabel: map[string]string{"id": "OP", "en": "SRG"},
	}
)

var (
	DoctorAppointmentTypeMapIdx = map[DoctorAppointmentTypeIdx]DoctorAppointmentTypeValue{
		DoctorAppointmentTypeUnknownIdx:      DoctorAppointmentTypeUnknownValue,
		DoctorAppointmentTypeConsultationIdx: DoctorAppointmentTypeConsultationValue,
		DoctorAppointmentTypeCheckupIdx:      DoctorAppointmentTypeCheckupValue,
		DoctorAppointmentTypeFollowupIdx:     DoctorAppointmentTypeFollowupValue,
		DoctorAppointmentTypeEmergencyIdx:    DoctorAppointmentTypeEmergencyValue,
		DoctorAppointmentTypeTelemedicineIdx: DoctorAppointmentTypeTelemedicineValue,
		DoctorAppointmentTypeSurgeryIdx:      DoctorAppointmentTypeSurgeryValue,
	}
	DoctorAppointmentTypeMapKey = map[DoctorAppointmentTypeKey]DoctorAppointmentTypeValue{
		DoctorAppointmentTypeUnknownKey:      DoctorAppointmentTypeUnknownValue,
		DoctorAppointmentTypeConsultationKey: DoctorAppointmentTypeConsultationValue,
		DoctorAppointmentTypeCheckupKey:      DoctorAppointmentTypeCheckupValue,
		DoctorAppointmentTypeFollowupKey:     DoctorAppointmentTypeFollowupValue,
		DoctorAppointmentTypeEmergencyKey:    DoctorAppointmentTypeEmergencyValue,
		DoctorAppointmentTypeTelemedicineKey: DoctorAppointmentTypeTelemedicineValue,
		DoctorAppointmentTypeSurgeryKey:      DoctorAppointmentTypeSurgeryValue,
	}
)

func (p DoctorAppointmentTypeIdx) String() string {
	if appointmentType, ok := DoctorAppointmentTypeMapIdx[p]; ok {
		return appointmentType.Key
	}
	return string(DoctorAppointmentTypeConsultationKey)
}

func (p *DoctorAppointmentTypeIdx) Scan(value interface{}) error {
	if value == nil {
		*p = DoctorAppointmentTypeConsultationIdx
		return nil
	}
	switch v := value.(type) {
	case int8:
		*p = DoctorAppointmentTypeIdx(v)
		return nil
	case int16:
		*p = DoctorAppointmentTypeIdx(v)
		return nil
	case int32:
		*p = DoctorAppointmentTypeIdx(v)
		return nil
	case int64:
		*p = DoctorAppointmentTypeIdx(v)
		return nil
	case string:
		if val, ok := DoctorAppointmentTypeMapKey[DoctorAppointmentTypeKey(v)]; ok {
			*p = DoctorAppointmentTypeIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := DoctorAppointmentTypeMapKey[DoctorAppointmentTypeKey(strVal)]; ok {
			*p = DoctorAppointmentTypeIdx(val.Idx)
			return nil
		}
	}
	return errors.New("invalid doctor appointment type value")
}

func (p DoctorAppointmentTypeIdx) Value() (driver.Value, error) {
	return int64(p), nil
}

type NullDoctorAppointmentTypeIdx struct {
	DoctorAppointmentTypeIdx DoctorAppointmentTypeIdx
	Valid                    bool // Valid is true if DoctorAppointmentTypeIdx is not NULL
}

func (p *NullDoctorAppointmentTypeIdx) Scan(value interface{}) error {
	if value == nil {
		p.DoctorAppointmentTypeIdx, p.Valid = DoctorAppointmentTypeConsultationIdx, false
		return nil
	}
	p.Valid = true
	return p.DoctorAppointmentTypeIdx.Scan(value)
}

func (p NullDoctorAppointmentTypeIdx) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}
	return p.DoctorAppointmentTypeIdx.Value()
}
