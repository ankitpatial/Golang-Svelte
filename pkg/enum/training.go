package enum

import (
	"fmt"
	"io"
	"strconv"
)

type TrainingType string

const (
	TrainingTypeRoofing    TrainingType = "ROOFING"
	TrainingTypeSolar      TrainingType = "SOLAR"
	TrainingTypeSiteSurvey TrainingType = "SITE_SURVEY"
)

var AllTrainingTypes = []TrainingType{
	TrainingTypeRoofing,
	TrainingTypeSolar,
	TrainingTypeSiteSurvey,
}

func (t TrainingType) String() string {
	return string(t)
}

func (t TrainingType) Humanize() string {
	switch t {
	case TrainingTypeRoofing:
		return "Roofing Training"
	case TrainingTypeSolar:
		return "Solar Training"
	case TrainingTypeSiteSurvey:
		return "Site Survey Training"
	default:
		return ""
	}
}

// Values provides list valid values for Enum.
func (TrainingType) Values() (kinds []string) {
	for _, s := range AllTrainingTypes {
		kinds = append(kinds, string(s))
	}
	return
}

func (t TrainingType) IsValid() bool {
	for _, t := range AllTrainingTypes {
		if t == t {
			return true
		}
	}
	return false
}

func (t *TrainingType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = TrainingType(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid TrainingType", str)
	}
	return nil
}

func (t TrainingType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
