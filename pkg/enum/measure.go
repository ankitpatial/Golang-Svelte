package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Measure string

const (
	MeasurePrimaryStructureOnly      Measure = "PrimaryStructureOnly"
	MeasurePrimaryPlusDetachedGarage Measure = "PrimaryPlusDetachedGarage"
	MeasureAllStructuresOnParcel     Measure = "AllStructuresOnParcel"
)

var AllMeasures = []Measure{
	MeasurePrimaryStructureOnly,
	MeasurePrimaryPlusDetachedGarage,
	MeasureAllStructuresOnParcel,
}

func (s Measure) String() string {
	return string(s)
}

func (Measure) Values() (kinds []string) {
	for _, s := range AllMeasures {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Measure) IsValid() bool {
	for _, t := range AllMeasures {
		if s == t {
			return true
		}
	}
	return false
}

func (s *Measure) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Measure(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid Measure", str)
	}
	return nil
}

func (s Measure) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
