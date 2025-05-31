package entity

import (
	"time"

	"github.com/ahargunyllib/thera-be/domain/enums"
)

type HospitalPartner struct {
	ID             string                         `db:"id"`
	FromHospitalID int                            `db:"from_hospital_id"`
	FromHospital   Hospital                       `db:"from_hospital"`
	ToHospitalID   int                            `db:"to_hospital_id"`
	ToHospital     Hospital                       `db:"to_hospital"`
	PartnerType    enums.HospitalPartnerTypeIdx   `db:"partner_type"` // 0: referral, 1: transfer
	Status         enums.HospitalPartnerStatusIdx `db:"status"`       // 0: pending, 1: accepted, 2: rejected
	CreatedAt      time.Time                      `db:"created_at"`
	UpdatedAt      time.Time                      `db:"updated_at"`
}
