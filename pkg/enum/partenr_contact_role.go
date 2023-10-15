package enum

import (
	"fmt"
	"io"
	"strconv"
)

type PartnerContactRole string

const (
	PartnerContactRoleNone     PartnerContactRole = "NONE"
	PartnerContactRoleAdmin    PartnerContactRole = "ADMIN"
	PartnerContactRoleSalesRep PartnerContactRole = "SALES_REP"
)

var AllPartnerUserRoles = []PartnerContactRole{
	PartnerContactRoleNone,
	PartnerContactRoleAdmin,
	PartnerContactRoleSalesRep,
}

func (s PartnerContactRole) String() string {
	return string(s)
}

func (s PartnerContactRole) Humanize() string {
	switch s {
	case PartnerContactRoleAdmin:
		return "Admin"
	case PartnerContactRoleSalesRep:
		return "Sales Rep"
	default:
		return "Unknown"
	}
}

// Values provides list valid values for Enum.
func (PartnerContactRole) Values() (kinds []string) {
	for _, s := range AllPartnerUserRoles {
		kinds = append(kinds, string(s))
	}
	return
}

func (s PartnerContactRole) IsValid() bool {
	for _, t := range AllPartnerUserRoles {
		if s == t {
			return true
		}
	}
	return false
}

func (s *PartnerContactRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = PartnerContactRole(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid PartnerStatus", str)
	}
	return nil
}

func (s PartnerContactRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
