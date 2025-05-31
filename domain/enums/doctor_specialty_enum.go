package enums

import (
	"database/sql/driver"
	"errors"
)

type DoctorSpecialtyIdx int64
type DoctorSpecialtyKey string
type DoctorSpecialtyValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	DoctorSpecialtyUnknownIdx             DoctorSpecialtyIdx = 0
	DoctorSpecialtyGeneralPractitionerIdx DoctorSpecialtyIdx = 1
	DoctorSpecialtyPediatricsIdx          DoctorSpecialtyIdx = 2
	DoctorSpecialtyInternalMedicineIdx    DoctorSpecialtyIdx = 3
	DoctorSpecialtyObgynIdx               DoctorSpecialtyIdx = 4
	DoctorSpecialtyGeneralSurgeryIdx      DoctorSpecialtyIdx = 5
	DoctorSpecialtyCardiologyIdx          DoctorSpecialtyIdx = 6
	DoctorSpecialtyNeurologyIdx           DoctorSpecialtyIdx = 7
	DoctorSpecialtyEntIdx                 DoctorSpecialtyIdx = 8
	DoctorSpecialtyOphthalmologyIdx       DoctorSpecialtyIdx = 9
	DoctorSpecialtyDermatologyIdx         DoctorSpecialtyIdx = 10
	DoctorSpecialtyPulmonologyIdx         DoctorSpecialtyIdx = 11
	DoctorSpecialtyPsychiatryIdx          DoctorSpecialtyIdx = 12
	DoctorSpecialtyOrthopedicsIdx         DoctorSpecialtyIdx = 13
	DoctorSpecialtyAnesthesiologyIdx      DoctorSpecialtyIdx = 14
	DoctorSpecialtyRadiologyIdx           DoctorSpecialtyIdx = 15
)

const (
	DoctorSpecialtyUnknownKey             DoctorSpecialtyKey = "unknown"
	DoctorSpecialtyGeneralPractitionerKey DoctorSpecialtyKey = "general_practitioner"
	DoctorSpecialtyPediatricsKey          DoctorSpecialtyKey = "pediatrics"
	DoctorSpecialtyInternalMedicineKey    DoctorSpecialtyKey = "internal_medicine"
	DoctorSpecialtyObgynKey               DoctorSpecialtyKey = "obgyn"
	DoctorSpecialtyGeneralSurgeryKey      DoctorSpecialtyKey = "general_surgery"
	DoctorSpecialtyCardiologyKey          DoctorSpecialtyKey = "cardiology"
	DoctorSpecialtyNeurologyKey           DoctorSpecialtyKey = "neurology"
	DoctorSpecialtyEntKey                 DoctorSpecialtyKey = "ent"
	DoctorSpecialtyOphthalmologyKey       DoctorSpecialtyKey = "ophthalmology"
	DoctorSpecialtyDermatologyKey         DoctorSpecialtyKey = "dermatology"
	DoctorSpecialtyPulmonologyKey         DoctorSpecialtyKey = "pulmonology"
	DoctorSpecialtyPsychiatryKey          DoctorSpecialtyKey = "psychiatry"
	DoctorSpecialtyOrthopedicsKey         DoctorSpecialtyKey = "orthopedics"
	DoctorSpecialtyAnesthesiologyKey      DoctorSpecialtyKey = "anesthesiology"
	DoctorSpecialtyRadiologyKey           DoctorSpecialtyKey = "radiology"
)

var (
	DoctorSpecialtyUnknownValue = DoctorSpecialtyValue{
		Idx:        int64(DoctorSpecialtyUnknownIdx),
		Key:        string(DoctorSpecialtyUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	DoctorSpecialtyGeneralPractitionerValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyGeneralPractitionerIdx),
		Key: string(DoctorSpecialtyGeneralPractitionerKey),
		LongLabel: map[string]string{
			"id": "Dokter Umum",
			"en": "General Practitioner",
		},
		ShortLabel: map[string]string{
			"id": "Umum",
			"en": "GP",
		},
	}
	DoctorSpecialtyPediatricsValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyPediatricsIdx),
		Key: string(DoctorSpecialtyPediatricsKey),
		LongLabel: map[string]string{
			"id": "Spesialis Anak (Pediatri)",
			"en": "Pediatrics",
		},
		ShortLabel: map[string]string{
			"id": "Anak",
			"en": "Peds",
		},
	}
	DoctorSpecialtyInternalMedicineValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyInternalMedicineIdx),
		Key: string(DoctorSpecialtyInternalMedicineKey),
		LongLabel: map[string]string{
			"id": "Spesialis Penyakit Dalam",
			"en": "Internal Medicine",
		},
		ShortLabel: map[string]string{
			"id": "Penyakit Dalam",
			"en": "IM",
		},
	}
	DoctorSpecialtyObgynValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyObgynIdx),
		Key: string(DoctorSpecialtyObgynKey),
		LongLabel: map[string]string{
			"id": "Spesialis Kebidanan dan Kandungan",
			"en": "Obstetrics and Gynecology",
		},
		ShortLabel: map[string]string{
			"id": "Obgyn",
			"en": "Obgyn",
		},
	}
	DoctorSpecialtyGeneralSurgeryValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyGeneralSurgeryIdx),
		Key: string(DoctorSpecialtyGeneralSurgeryKey),
		LongLabel: map[string]string{
			"id": "Spesialis Bedah Umum",
			"en": "General Surgery",
		},
		ShortLabel: map[string]string{
			"id": "Bedah",
			"en": "Surg",
		},
	}
	DoctorSpecialtyCardiologyValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyCardiologyIdx),
		Key: string(DoctorSpecialtyCardiologyKey),
		LongLabel: map[string]string{
			"id": "Spesialis Kardiologi",
			"en": "Cardiology",
		},
		ShortLabel: map[string]string{
			"id": "Kardiologi",
			"en": "Cardio",
		},
	}
	DoctorSpecialtyNeurologyValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyNeurologyIdx),
		Key: string(DoctorSpecialtyNeurologyKey),
		LongLabel: map[string]string{
			"id": "Spesialis Neurologi",
			"en": "Neurology",
		},
		ShortLabel: map[string]string{
			"id": "Neurologi",
			"en": "Neuro",
		},
	}
	DoctorSpecialtyEntValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyEntIdx),
		Key: string(DoctorSpecialtyEntKey),
		LongLabel: map[string]string{
			"id": "Spesialis THT (Telinga, Hidung, Tenggorokan)",
			"en": "ENT (Ear, Nose, Throat)",
		},
		ShortLabel: map[string]string{
			"id": "THT",
			"en": "ENT",
		},
	}
	DoctorSpecialtyOphthalmologyValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyOphthalmologyIdx),
		Key: string(DoctorSpecialtyOphthalmologyKey),
		LongLabel: map[string]string{
			"id": "Spesialis Mata (Oftalmologi)",
			"en": "Ophthalmology",
		},
		ShortLabel: map[string]string{
			"id": "Mata",
			"en": "Ophth",
		},
	}
	DoctorSpecialtyDermatologyValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyDermatologyIdx),
		Key: string(DoctorSpecialtyDermatologyKey),
		LongLabel: map[string]string{
			"id": "Spesialis Kulit dan Kelamin",
			"en": "Dermatology and Venereology",
		},
		ShortLabel: map[string]string{
			"id": "Kulit",
			"en": "Derm",
		},
	}
	DoctorSpecialtyPulmonologyValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyPulmonologyIdx),
		Key: string(DoctorSpecialtyPulmonologyKey),
		LongLabel: map[string]string{
			"id": "Spesialis Paru",
			"en": "Pulmonology",
		},
		ShortLabel: map[string]string{
			"id": "Paru",
			"en": "Pulmo",
		},
	}
	DoctorSpecialtyPsychiatryValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyPsychiatryIdx),
		Key: string(DoctorSpecialtyPsychiatryKey),
		LongLabel: map[string]string{
			"id": "Spesialis Psikiatri",
			"en": "Psychiatry",
		},
		ShortLabel: map[string]string{
			"id": "Psikiatri",
			"en": "Psych",
		},
	}
	DoctorSpecialtyOrthopedicsValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyOrthopedicsIdx),
		Key: string(DoctorSpecialtyOrthopedicsKey),
		LongLabel: map[string]string{
			"id": "Spesialis Orthopedi dan Traumatologi",
			"en": "Orthopedics and Traumatology",
		},
		ShortLabel: map[string]string{
			"id": "Ortho",
			"en": "Ortho",
		},
	}
	DoctorSpecialtyAnesthesiologyValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyAnesthesiologyIdx),
		Key: string(DoctorSpecialtyAnesthesiologyKey),
		LongLabel: map[string]string{
			"id": "Spesialis Anestesiologi dan Terapi Intensif",
			"en": "Anesthesiology and Intensive Therapy",
		},
		ShortLabel: map[string]string{
			"id": "Anestesi",
			"en": "Anesth",
		},
	}
	DoctorSpecialtyRadiologyValue = DoctorSpecialtyValue{
		Idx: int64(DoctorSpecialtyRadiologyIdx),
		Key: string(DoctorSpecialtyRadiologyKey),
		LongLabel: map[string]string{
			"id": "Spesialis Radiologi",
			"en": "Radiology",
		},
		ShortLabel: map[string]string{
			"id": "Radiologi",
			"en": "Radio",
		},
	}
)

var (
	DoctorSpecialtyMapIdx = map[DoctorSpecialtyIdx]DoctorSpecialtyValue{
		DoctorSpecialtyUnknownIdx:             DoctorSpecialtyUnknownValue,
		DoctorSpecialtyGeneralPractitionerIdx: DoctorSpecialtyGeneralPractitionerValue,
		DoctorSpecialtyPediatricsIdx:          DoctorSpecialtyPediatricsValue,
		DoctorSpecialtyInternalMedicineIdx:    DoctorSpecialtyInternalMedicineValue,
		DoctorSpecialtyObgynIdx:               DoctorSpecialtyObgynValue,
		DoctorSpecialtyGeneralSurgeryIdx:      DoctorSpecialtyGeneralSurgeryValue,
		DoctorSpecialtyCardiologyIdx:          DoctorSpecialtyCardiologyValue,
		DoctorSpecialtyNeurologyIdx:           DoctorSpecialtyNeurologyValue,
		DoctorSpecialtyEntIdx:                 DoctorSpecialtyEntValue,
		DoctorSpecialtyOphthalmologyIdx:       DoctorSpecialtyOphthalmologyValue,
		DoctorSpecialtyDermatologyIdx:         DoctorSpecialtyDermatologyValue,
		DoctorSpecialtyPulmonologyIdx:         DoctorSpecialtyPulmonologyValue,
		DoctorSpecialtyPsychiatryIdx:          DoctorSpecialtyPsychiatryValue,
		DoctorSpecialtyOrthopedicsIdx:         DoctorSpecialtyOrthopedicsValue,
		DoctorSpecialtyAnesthesiologyIdx:      DoctorSpecialtyAnesthesiologyValue,
		DoctorSpecialtyRadiologyIdx:           DoctorSpecialtyRadiologyValue,
	}

	DoctorSpecialtyMapKey = map[DoctorSpecialtyKey]DoctorSpecialtyValue{
		DoctorSpecialtyUnknownKey:             DoctorSpecialtyUnknownValue,
		DoctorSpecialtyGeneralPractitionerKey: DoctorSpecialtyGeneralPractitionerValue,
		DoctorSpecialtyPediatricsKey:          DoctorSpecialtyPediatricsValue,
		DoctorSpecialtyInternalMedicineKey:    DoctorSpecialtyInternalMedicineValue,
		DoctorSpecialtyObgynKey:               DoctorSpecialtyObgynValue,
		DoctorSpecialtyGeneralSurgeryKey:      DoctorSpecialtyGeneralSurgeryValue,
		DoctorSpecialtyCardiologyKey:          DoctorSpecialtyCardiologyValue,
		DoctorSpecialtyNeurologyKey:           DoctorSpecialtyNeurologyValue,
		DoctorSpecialtyEntKey:                 DoctorSpecialtyEntValue,
		DoctorSpecialtyOphthalmologyKey:       DoctorSpecialtyOphthalmologyValue,
		DoctorSpecialtyDermatologyKey:         DoctorSpecialtyDermatologyValue,
		DoctorSpecialtyPulmonologyKey:         DoctorSpecialtyPulmonologyValue,
		DoctorSpecialtyPsychiatryKey:          DoctorSpecialtyPsychiatryValue,
		DoctorSpecialtyOrthopedicsKey:         DoctorSpecialtyOrthopedicsValue,
		DoctorSpecialtyAnesthesiologyKey:      DoctorSpecialtyAnesthesiologyValue,
		DoctorSpecialtyRadiologyKey:           DoctorSpecialtyRadiologyValue,
	}
)

func (d DoctorSpecialtyIdx) String() string {
	if val, ok := DoctorSpecialtyMapIdx[d]; ok {
		return val.Key
	}
	return string(DoctorSpecialtyUnknownKey)
}

func (d DoctorSpecialtyIdx) Value() (driver.Value, error) {
	return int64(d), nil
}

func (d *DoctorSpecialtyIdx) Scan(value interface{}) error {
	if value == nil {
		*d = DoctorSpecialtyUnknownIdx
		return nil
	}

	switch v := value.(type) {
	case int8:
		*d = DoctorSpecialtyIdx(v)
		return nil
	case int16:
		*d = DoctorSpecialtyIdx(v)
		return nil
	case int32:
		*d = DoctorSpecialtyIdx(v)
		return nil
	case int64:
		*d = DoctorSpecialtyIdx(v)
		return nil
	case string:
		if val, ok := DoctorSpecialtyMapKey[DoctorSpecialtyKey(v)]; ok {
			*d = DoctorSpecialtyIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := DoctorSpecialtyMapKey[DoctorSpecialtyKey(strVal)]; ok {
			*d = DoctorSpecialtyIdx(val.Idx)
			return nil
		}
	}

	return errors.New("invalid doctor specialty value")
}

type NullDoctorSpecialtyIdx struct {
	DoctorSpecialtyIdx DoctorSpecialtyIdx
	Valid              bool // Valid is true if DoctorSpecialtyIdx is not NULL
}

func (n *NullDoctorSpecialtyIdx) Scan(value interface{}) error {
	if value == nil {
		n.DoctorSpecialtyIdx, n.Valid = DoctorSpecialtyUnknownIdx, false
		return nil
	}

	n.Valid = true
	return n.DoctorSpecialtyIdx.Scan(value)
}

func (n NullDoctorSpecialtyIdx) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.DoctorSpecialtyIdx.Value()
}
