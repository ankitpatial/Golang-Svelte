package enum

import (
	"fmt"
	"io"
	"strconv"
)

type OptionList string

const (
	OptionListFinance OptionList = "FINANCE"
	OptionListEPC     OptionList = "EPC"
)

var AllOptionList = []OptionList{
	OptionListFinance,
	OptionListEPC,
}

func (s OptionList) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (OptionList) Values() (kinds []string) {
	for _, s := range AllOptionList {
		kinds = append(kinds, string(s))
	}
	return
}

func (s OptionList) IsValid() bool {
	for _, t := range AllOptionList {
		if s == t {
			return true
		}
	}
	return false
}

func (s *OptionList) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = OptionList(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s OptionList) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
