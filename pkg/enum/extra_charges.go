package enum

import (
	"fmt"
	"io"
	"strconv"
)

type ExtraCharge string

const (
	ExtraChargeNone        ExtraCharge = "NONE"
	ExtraChargeAmount      ExtraCharge = "AMOUNT"
	ExtraChargePercent     ExtraCharge = "PERCENT"
	ExtraChargePerSQ       ExtraCharge = "PER_SQ"
	ExtraChargeConditional ExtraCharge = "CONDITIONAL"
)

var AllExtraCharges = []ExtraCharge{
	ExtraChargeNone,
	ExtraChargeAmount,
	ExtraChargePercent,
	ExtraChargePerSQ,
	ExtraChargeConditional,
}

func (t ExtraCharge) String() string {
	return string(t)
}

// Values provides list valid values for Enum.
func (ExtraCharge) Values() (kinds []string) {
	for _, s := range AllExtraCharges {
		kinds = append(kinds, string(s))
	}
	return
}

func (t ExtraCharge) IsValid() bool {
	for _, t := range AllExtraCharges {
		if t == t {
			return true
		}
	}
	return false
}

func (t *ExtraCharge) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = ExtraCharge(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid ExtraCharge", str)
	}
	return nil
}

func (t ExtraCharge) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
