package enum

import (
	"fmt"
	"io"
	"strconv"
)

type InstallationStatus string

const (
	InstallationStatusPending   InstallationStatus = "PENDING"
	InstallationStatusNew       InstallationStatus = "NEW"
	InstallationStatusScheduled InstallationStatus = "SCHEDULED"
	InstallationStatusInstalled InstallationStatus = "INSTALLED"
)

var AllInstallationStatuses = []InstallationStatus{
	InstallationStatusPending,
	InstallationStatusNew,
	InstallationStatusScheduled,
	InstallationStatusInstalled,
}

func (s InstallationStatus) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (InstallationStatus) Values() (kinds []string) {
	for _, s := range AllInstallationStatuses {
		kinds = append(kinds, string(s))
	}
	return
}

func (s InstallationStatus) IsValid() bool {
	for _, t := range AllInstallationStatuses {
		if s == t {
			return true
		}
	}
	return false
}

func (s *InstallationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = InstallationStatus(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s InstallationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
