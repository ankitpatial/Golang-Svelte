package enum

import (
	"fmt"
	"io"
	"strconv"
)

type ShingleColor string

const (
	ShingleColorBrown     ShingleColor = "Brown"
	ShingleColorSandstorm ShingleColor = "Sandstorm"
	ShingleColorBlue      ShingleColor = "Blue"
)

var AllShingleColor = []ShingleColor{
	ShingleColorBrown,
	ShingleColorSandstorm,
	ShingleColorBlue,
}

func (s ShingleColor) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (ShingleColor) Values() (kinds []string) {
	for _, s := range AllShingleColor {
		kinds = append(kinds, string(s))
	}
	return
}

func (s ShingleColor) IsValid() bool {
	for _, t := range AllShingleColor {
		if s == t {
			return true
		}
	}
	return false
}

func (s *ShingleColor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = ShingleColor(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid ShingleColor", str)
	}
	return nil
}

func (s ShingleColor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
