package enum

import (
	"fmt"
	"io"
	"strconv"
)

// **
// AccountStatus
// **

type AccountStatus string

const (
	AccountStatusPending  AccountStatus = "PENDING"
	AccountStatusActive   AccountStatus = "ACTIVE"
	AccountStatusDisabled AccountStatus = "DISABLED"
)

var AllAccountStatus = []AccountStatus{
	AccountStatusPending,  // 0
	AccountStatusActive,   // 1
	AccountStatusDisabled, // 2
}

func (e AccountStatus) ToUInt8() uint8 {
	for idx, i := range AllAccountStatus {
		if e == i {
			return uint8(idx)
		}
	}
	return 0
}

func (e *AccountStatus) FromUInt8(v uint8) {
	for idx, i := range AllAccountStatus {
		if uint8(idx) == v {
			*e = i
			break
		}
	}
}

func (s AccountStatus) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (AccountStatus) Values() (kinds []string) {
	for _, s := range AllAccountStatus {
		kinds = append(kinds, string(s))
	}
	return
}

func (s AccountStatus) IsValid() bool {
	for _, t := range AllAccountStatus {
		if s == t {
			return true
		}
	}
	return false
}

func (s *AccountStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = AccountStatus(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid AccountStatus", str)
	}
	return nil
}

func (s AccountStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

// **
// ROLE
// **

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RoleSubAdmin  Role = "SUB_ADMIN"
	RoleEstimator Role = "ESTIMATOR"
	RoleReviewer  Role = "REVIEWER"
	RoleSiteUser  Role = "SITE_USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleSubAdmin,
	RoleEstimator,
	RoleReviewer,
	RoleSiteUser,
}

func (s Role) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (Role) Values() (kinds []string) {
	for _, s := range AllRole {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Role) IsValid() bool {
	for _, t := range AllRole {
		if s == t {
			return true
		}
	}
	return false
}

func (s *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Role(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (s Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
