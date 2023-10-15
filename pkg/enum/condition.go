package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Condition string

const (
	ConditionAddIfStateIn      Condition = "AddIfStateIn"
	ConditionSubtractIfStateIn Condition = "SubtractIfStateIn"
)

var AllConditions = []Condition{
	ConditionAddIfStateIn,
	ConditionSubtractIfStateIn,
}

func (t Condition) String() string {
	return string(t)
}

// Values provides list valid values for Enum.
func (Condition) Values() (kinds []string) {
	for _, s := range AllConditions {
		kinds = append(kinds, string(s))
	}
	return
}

func (t Condition) IsValid() bool {
	for _, t := range AllConditions {
		if t == t {
			return true
		}
	}
	return false
}

func (t *Condition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = Condition(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid Condition", str)
	}
	return nil
}

func (t Condition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
