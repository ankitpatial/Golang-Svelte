package enum

import (
	"fmt"
	"io"
	"strconv"
)

type JobDocUrlType string

const (
	JobDocUrlInspectionDocs JobDocUrlType = "InspectionDocs"
	JobDocUrlProductionPics JobDocUrlType = "ProductionPics"
)

var AllJobDocUrlTypes = []JobDocUrlType{
	JobDocUrlInspectionDocs,
	JobDocUrlProductionPics,
}

func (t JobDocUrlType) String() string {
	return string(t)
}

// Values provides list valid values for Enum.
func (JobDocUrlType) Values() (kinds []string) {
	for _, s := range AllJobDocUrlTypes {
		kinds = append(kinds, string(s))
	}
	return
}

func (t JobDocUrlType) IsValid() bool {
	for _, t := range AllJobDocUrlTypes {
		if t == t {
			return true
		}
	}
	return false
}

func (t *JobDocUrlType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = JobDocUrlType(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid DocSection", str)
	}
	return nil
}

func (t JobDocUrlType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
