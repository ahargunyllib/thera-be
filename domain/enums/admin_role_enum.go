package enums

import "errors"

type AdminRoleIdx int64
type AdminRoleKey string
type AdminRoleValue struct {
	Idx        int64             `json:"idx"`
	Key        string            `json:"key"`
	ShortLabel map[string]string `json:"short_label"`
	LongLabel  map[string]string `json:"long_label"`
}

const (
	AdminRoleUnknownIdx    AdminRoleIdx = 0
	AdminRoleSuperAdminIdx AdminRoleIdx = 1
	AdminRoleAdminIdx      AdminRoleIdx = 2

	AdminRoleUnknownKey    AdminRoleKey = "unknown"
	AdminRoleSuperAdminKey AdminRoleKey = "super_admin"
	AdminRoleAdminKey      AdminRoleKey = "admin"
)

var (
	AdminRoleUnknownValue = AdminRoleValue{
		Idx:        int64(AdminRoleUnknownIdx),
		Key:        string(AdminRoleUnknownKey),
		LongLabel:  map[string]string{"id": "", "en": ""},
		ShortLabel: map[string]string{"id": "", "en": ""},
	}
	AdminRoleSuperAdminValue = AdminRoleValue{
		Idx:        int64(AdminRoleSuperAdminIdx),
		Key:        string(AdminRoleSuperAdminKey),
		LongLabel:  map[string]string{"id": "Super Admin", "en": "Super Admin"},
		ShortLabel: map[string]string{"id": "SA", "en": "SA"},
	}
	AdminRoleAdminValue = AdminRoleValue{
		Idx:        int64(AdminRoleAdminIdx),
		Key:        string(AdminRoleAdminKey),
		LongLabel:  map[string]string{"id": "Admin", "en": "Admin"},
		ShortLabel: map[string]string{"id": "AD", "en": "AD"},
	}
)

var (
	AdminRoleMapIdx = map[AdminRoleIdx]AdminRoleValue{
		AdminRoleUnknownIdx:    AdminRoleUnknownValue,
		AdminRoleSuperAdminIdx: AdminRoleSuperAdminValue,
		AdminRoleAdminIdx:      AdminRoleAdminValue,
	}
	AdminRoleMapKey = map[AdminRoleKey]AdminRoleValue{
		AdminRoleUnknownKey:    AdminRoleUnknownValue,
		AdminRoleSuperAdminKey: AdminRoleSuperAdminValue,
		AdminRoleAdminKey:      AdminRoleAdminValue,
	}
)

func (a AdminRoleIdx) String() string {
	if role, ok := AdminRoleMapIdx[a]; ok {
		return role.Key
	}
	return string(AdminRoleUnknownKey)
}
func (a *AdminRoleIdx) Scan(value interface{}) error {
	if value == nil {
		*a = AdminRoleUnknownIdx
		return nil
	}
	switch v := value.(type) {
	case int8:
		*a = AdminRoleIdx(v)
		return nil
	case int16:
		*a = AdminRoleIdx(v)
		return nil
	case int32:
		*a = AdminRoleIdx(v)
		return nil
	case int64:
		*a = AdminRoleIdx(v)
		return nil
	case int:
		*a = AdminRoleIdx(v)
		return nil
	case string:
		if val, ok := AdminRoleMapKey[AdminRoleKey(v)]; ok {
			*a = AdminRoleIdx(val.Idx)
			return nil
		}
	case []byte:
		strVal := string(v)
		if val, ok := AdminRoleMapKey[AdminRoleKey(strVal)]; ok {
			*a = AdminRoleIdx(val.Idx)
			return nil
		}
	}

	return errors.New("failed to scan AdminRoleIdx")
}

func (a AdminRoleIdx) Value() (int64, error) {
	return int64(a), nil
}
