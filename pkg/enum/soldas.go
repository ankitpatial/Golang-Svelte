package enum

import (
	"fmt"
	"io"
	"strconv"
)

type SoldAs string

const (
	SoldAsPackage        SoldAs = "PACKAGE"
	SoldAsIndividualItem SoldAs = "INDIVIDUAL_ITEM"
)

var AllSoldAs = []SoldAs{
	SoldAsPackage,
	SoldAsIndividualItem,
}

func (s SoldAs) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (SoldAs) Values() (kinds []string) {
	for _, s := range AllSoldAs {
		kinds = append(kinds, string(s))
	}
	return
}

func (s SoldAs) IsValid() bool {
	for _, t := range AllSoldAs {
		if s == t {
			return true
		}
	}
	return false
}

func (s *SoldAs) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = SoldAs(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s SoldAs) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
