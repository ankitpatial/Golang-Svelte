package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Product string

const (
	ProductHVAC      Product = "HVAC"
	ProductSmartHome Product = "SMART_HOME"
)

var AllPackages = []Product{
	ProductHVAC,
	ProductSmartHome,
}

func (s Product) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (Product) Values() (kinds []string) {
	for _, s := range AllPackages {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Product) IsValid() bool {
	for _, t := range AllPackages {
		if s == t {
			return true
		}
	}
	return false
}

func (s *Product) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Product(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s Product) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
