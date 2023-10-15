package enum

import (
	"fmt"
	"io"
	"strconv"
)

type EstimateStatus string

const (
	EstimateStatusNew      EstimateStatus = "New"
	EstimateStatusPending  EstimateStatus = "Pending"
	EstimateStatusApproved EstimateStatus = "Approved"
	EstimateStatusDenied   EstimateStatus = "Denied"
	EstimateStatusOnHold   EstimateStatus = "OnHold"
	EstimateStatusFailed   EstimateStatus = "Failed"
)

var AllEstimateStatuses = []EstimateStatus{
	EstimateStatusNew,
	EstimateStatusPending,
	EstimateStatusApproved,
	EstimateStatusDenied,
	EstimateStatusOnHold,
	EstimateStatusFailed,
}

func (s EstimateStatus) String() string {
	return string(s)
}

func (EstimateStatus) Values() (kinds []string) {
	for _, s := range AllEstimateStatuses {
		kinds = append(kinds, string(s))
	}
	return
}

func (s EstimateStatus) IsValid() bool {
	for _, t := range AllEstimateStatuses {
		if s == t {
			return true
		}
	}
	return false
}

func (s *EstimateStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = EstimateStatus(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid EstimateStatus", str)
	}
	return nil
}

func (s EstimateStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
