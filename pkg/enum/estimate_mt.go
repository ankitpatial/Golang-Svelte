package enum

import (
	"fmt"
	"io"
	"strconv"
)

type EstimateType string

const (
	EstimateTypePrimaryPlusDetachedGarage EstimateType = "PrimaryPlusDetachedGarage"
	EstimateTypePrimaryStructureOnly      EstimateType = "PrimaryStructureOnly"
	EstimateTypeAllStructuresOnParcel     EstimateType = "AllStructuresOnParcel"
)

var AllEstimateTypes = []EstimateType{
	EstimateTypePrimaryStructureOnly,
	EstimateTypePrimaryPlusDetachedGarage,
	EstimateTypeAllStructuresOnParcel,
}

func (s EstimateType) String() string {
	return string(s)
}

func (EstimateType) Values() (kinds []string) {
	for _, s := range AllEstimateTypes {
		kinds = append(kinds, string(s))
	}
	return
}

func (s EstimateType) IsValid() bool {
	for _, t := range AllEstimateTypes {
		if s == t {
			return true
		}
	}
	return false
}

func (s *EstimateType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = EstimateType(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid EstimateType", str)
	}
	return nil
}

func (s EstimateType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
