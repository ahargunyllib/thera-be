package service

import (
	"github.com/ahargunyllib/thera-be/domain/dto"
	"github.com/ahargunyllib/thera-be/domain/entity"
)

func parsePatientRequest(data *entity.Patient, req *dto.UpdatePatientRequest) error {
	if req.FullName != "" && data.FullName != req.FullName {
		data.FullName = req.FullName
	}

	if req.IDNumber != "" && data.IDNumber != req.IDNumber {
		data.IDNumber = req.IDNumber
	}

	if req.PhoneNumber != "" && data.PhoneNumber.String != req.PhoneNumber {
		data.PhoneNumber.String = req.PhoneNumber
		data.PhoneNumber.Valid = true
	}

	if req.Address != "" && data.Address != req.Address {
		data.Address = req.Address
	}

	if !req.DateOfBirth.IsZero() && data.DateOfBirth != req.DateOfBirth {
		data.DateOfBirth = req.DateOfBirth
	}

	// if req.Gender != "" && data.Gender != req.Gender {
	// 	data.Gender = req.Gender
	// }

	if req.Height != 0 && data.Height != req.Height {
		data.Height = req.Height
	}

	if req.Weight != 0 && data.Weight != req.Weight {
		data.Weight = req.Weight
	}

	// if req.BloodType != "" && data.BloodType != req.BloodType {
	// 	data.BloodType = req.BloodType
	// }

	if req.Allergies != "" && data.Allergies.String != req.Allergies {
		data.Allergies.String = req.Allergies
		data.Allergies.Valid = true
	}

	if req.MedicalRecordNumber != "" && data.MedicalRecordNumber != req.MedicalRecordNumber {
		data.MedicalRecordNumber = req.MedicalRecordNumber
	}

	if req.HospitalID != 0 && data.HospitalID != req.HospitalID {
		data.HospitalID = req.HospitalID
	}

	return nil
}
