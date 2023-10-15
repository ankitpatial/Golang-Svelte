package enum

import (
	"fmt"
	"io"
	"strconv"
)

type EPCStatus string

const (
	EPCStatusNone EPCStatus = "None"
	// EPCStatusDealer is a sales only
	EPCStatusDealer EPCStatus = "Dealer"
	// EPCStatusMultipleDealers installs Solar for multiple dealers, but they might also sell solar
	EPCStatusMultipleDealers EPCStatus = "MultipleDealers"
	// EPCStatusVerticallyIntegrated 1 company that sells and installs the solar
	EPCStatusVerticallyIntegrated EPCStatus = "VerticallyIntegrated"
)

var AllEPCStatuses = []EPCStatus{
	EPCStatusNone,
	EPCStatusDealer,
	EPCStatusMultipleDealers,
	EPCStatusVerticallyIntegrated,
}

func (s EPCStatus) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (EPCStatus) Values() (kinds []string) {
	for _, s := range AllEPCStatuses {
		kinds = append(kinds, string(s))
	}
	return
}

func (s EPCStatus) IsValid() bool {
	for _, t := range AllEPCStatuses {
		if s == t {
			return true
		}
	}
	return false
}

func (s *EPCStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = EPCStatus(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid EPCStatus", str)
	}
	return nil
}

func (s EPCStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
