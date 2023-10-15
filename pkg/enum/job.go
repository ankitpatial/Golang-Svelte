package enum

import (
	"fmt"
	"io"
	"strconv"
)

type JobSource string

const (
	JobSourceManual   JobSource = "Manual"
	JobSourceExternal JobSource = "External"
)

var AllJobSource = []JobSource{
	JobSourceManual,
	JobSourceExternal,
}

func (s JobSource) String() string {
	return string(s)
}

func (JobSource) Values() (kinds []string) {
	for _, s := range AllJobSource {
		kinds = append(kinds, string(s))
	}
	return
}

func (s JobSource) IsValid() bool {
	for _, t := range AllJobSource {
		if s == t {
			return true
		}
	}
	return false
}

func (s *JobSource) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = JobSource(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid JobSource", str)
	}
	return nil
}

func (s JobSource) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

// **
// JobProgress
// **

type JobProgress string

const (
	JobProgressNew                 JobProgress = "New"
	JobProgressCustomerContacted   JobProgress = "CustomerContacted"
	JobProgressJobDetailsConfirmed JobProgress = "JobDetailsConfirmed"
	JobProgressPermitting          JobProgress = "Permitting"
	JobProgressDelayed             JobProgress = "Delayed"
	JobProgressScheduled           JobProgress = "Scheduled"
	JobProgressInProgress          JobProgress = "InProgress"
	JobProgressInstalled           JobProgress = "Installed"
	JobProgressInvoiced            JobProgress = "Invoiced"
	JobProgressInvoiceApproved     JobProgress = "InvoiceApproved"
	JobProgressInvoicePaid         JobProgress = "InvoicePaid"
)

var AllJobProgress = []JobProgress{
	JobProgressNew,
	JobProgressCustomerContacted,
	JobProgressJobDetailsConfirmed,
	JobProgressPermitting,
	JobProgressDelayed,
	JobProgressScheduled,
	JobProgressInProgress,
	JobProgressInstalled,
	JobProgressInvoiced,
	JobProgressInvoiceApproved,
	JobProgressInvoicePaid,
}

func (s JobProgress) String() string {
	return string(s)
}

func (JobProgress) Values() (kinds []string) {
	for _, s := range AllJobProgress {
		kinds = append(kinds, string(s))
	}
	return
}

func (s JobProgress) IsValid() bool {
	for _, t := range AllJobProgress {
		if s == t {
			return true
		}
	}
	return false
}

func (s *JobProgress) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = JobProgress(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid JobProgress", str)
	}
	return nil
}

func (s JobProgress) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

//
// JobAssignmentStatus
//

type JobAssignmentStatus string

const (
	JobAssignmentStatusAssigned JobAssignmentStatus = "Assigned"
	JobAssignmentStatusAccepted JobAssignmentStatus = "Accepted"
	JobAssignmentStatusRejected JobAssignmentStatus = "Rejected"
	JobAssignmentStatusExpired  JobAssignmentStatus = "Expired"
)

var AllJobAssignmentStatus = []JobAssignmentStatus{
	JobAssignmentStatusAssigned,
	JobAssignmentStatusAccepted,
	JobAssignmentStatusRejected,
	JobAssignmentStatusExpired,
}

func (s JobAssignmentStatus) String() string {
	return string(s)
}

func (JobAssignmentStatus) Values() (kinds []string) {
	for _, s := range AllJobAssignmentStatus {
		kinds = append(kinds, string(s))
	}
	return
}

func (s JobAssignmentStatus) IsValid() bool {
	for _, t := range AllJobAssignmentStatus {
		if s == t {
			return true
		}
	}
	return false
}

func (s *JobAssignmentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = JobAssignmentStatus(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid JobAssignmentStatus", str)
	}
	return nil
}

func (s JobAssignmentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

// **
// JobMeasurement type of parcel
// **

type JobMeasurement string

const (
	Undefined                JobMeasurement = "Undefined"
	PrimaryAndDetachedGarage JobMeasurement = "PrimaryPlusDetachedGarage"
	PrimaryStructure         JobMeasurement = "PrimaryStructureOnly"
	AllStructures            JobMeasurement = "AllStructuresOnParcel"
)

var AllJobMeasurement = []JobMeasurement{
	Undefined,                // 0
	PrimaryAndDetachedGarage, // 1
	PrimaryStructure,         // 2
	AllStructures,            // 3
}

func (e JobMeasurement) ToUInt8() uint8 {
	for idx, i := range AllJobMeasurement {
		if e == i {
			return uint8(idx)
		}
	}
	return 0
}

func (e *JobMeasurement) FromUInt8(v uint8) {
	for idx, i := range AllJobMeasurement {
		if uint8(idx) == v {
			*e = i
			break
		}
	}
}

func (s JobMeasurement) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (JobMeasurement) Values() (kinds []string) {
	for _, s := range AllJobMeasurement {
		kinds = append(kinds, string(s))
	}
	return
}

func (s JobMeasurement) IsValid() bool {
	for _, t := range AllJobMeasurement {
		if s == t {
			return true
		}
	}
	return false
}

func (s *JobMeasurement) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = JobMeasurement(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid JobMeasurement", str)
	}
	return nil
}

func (s JobMeasurement) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
