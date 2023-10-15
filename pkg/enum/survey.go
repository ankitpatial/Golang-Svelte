package enum

import (
	"fmt"
	"io"
	"strconv"
)

// SurveyStatus defines the type for the "status" enum field.
type SurveyStatus string

// SurveyStatus values.
const (
	SurveyStatusRequesting SurveyStatus = "REQUESTING"
	SurveyStatusRequested  SurveyStatus = "REQUESTED"
)

var AllSurveyStatus = []SurveyStatus{
	SurveyStatusRequesting,
	SurveyStatusRequested,
}

func (s SurveyStatus) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (SurveyStatus) Values() (kinds []string) {
	for _, s := range AllSurveyStatus {
		kinds = append(kinds, string(s))
	}
	return
}

func (s SurveyStatus) IsValid() bool {
	for _, t := range AllSurveyStatus {
		if s == t {
			return true
		}
	}
	return false
}

func (s *SurveyStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = SurveyStatus(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s SurveyStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

// **
// SurveyProgress
// **

type SurveyProgress string

const (
	SurveyProgressScheduled  SurveyProgress = "SCHEDULED"
	SurveyStatusEnRoute      SurveyProgress = "EN_ROUTE"
	SurveyStatusOnSite       SurveyProgress = "ON_SITE"
	SurveyStatusCompleted    SurveyProgress = "COMPLETED"
	SurveyStatusDocsUploaded SurveyProgress = "DOCS_UPLOADED"
)

var AllSurveyProgress = []SurveyProgress{
	SurveyProgressScheduled,
	SurveyStatusEnRoute,
	SurveyStatusOnSite,
	SurveyStatusCompleted,
	SurveyStatusDocsUploaded,
}

func (s SurveyProgress) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (SurveyProgress) Values() (kinds []string) {
	for _, s := range AllSurveyProgress {
		kinds = append(kinds, string(s))
	}
	return
}

func (s SurveyProgress) IsValid() bool {
	for _, t := range AllSurveyProgress {
		if s == t {
			return true
		}
	}
	return false
}

func (s *SurveyProgress) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = SurveyProgress(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (s SurveyProgress) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
