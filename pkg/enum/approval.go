package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Approval string

const (
	ApprovalPending  Approval = "PENDING"
	ApprovalApproved Approval = "APPROVED"
	ApprovalDenied   Approval = "DENIED"
)

var AllApprovals = []Approval{
	ApprovalPending,
	ApprovalApproved,
	ApprovalDenied,
}

func (s Approval) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (Approval) Values() (kinds []string) {
	for _, s := range AllApprovals {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Approval) IsValid() bool {
	for _, t := range AllApprovals {
		if s == t {
			return true
		}
	}
	return false
}

func (s *Approval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Approval(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s Approval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
