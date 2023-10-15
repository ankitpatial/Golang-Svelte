package enum

import (
	"fmt"
	"io"
	"strconv"
)

type InstallationType string

const (
	InstallationTypeSmartHome InstallationType = "SMART_HOME"
	InstallationTypeHVAC      InstallationType = "HVAC"
)

var AllInstallationTypes = []InstallationType{
	InstallationTypeSmartHome,
	InstallationTypeHVAC,
}

func (s InstallationType) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (InstallationType) Values() (kinds []string) {
	for _, s := range AllInstallationTypes {
		kinds = append(kinds, string(s))
	}
	return
}

func (s InstallationType) IsValid() bool {
	for _, t := range AllInstallationTypes {
		if s == t {
			return true
		}
	}
	return false
}

func (s *InstallationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = InstallationType(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s InstallationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
