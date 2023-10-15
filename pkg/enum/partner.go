package enum

import (
	"fmt"
	"io"
	"strconv"
)

// Partner  ==>
type Partner string

const (
	PartnerNone        Partner = "NONE"
	PartnerRFX         Partner = "RFX"
	PartnerRoofing     Partner = "ROOFING"
	PartnerSolar       Partner = "SOLAR"
	PartnerEpc         Partner = "EPC"
	PartnerIntegration Partner = "INTEGRATION"
	PartnerLender      Partner = "LENDER"
)

var AllPartnerType = []Partner{
	PartnerNone,
	PartnerRFX,
	PartnerRoofing,
	PartnerSolar,
	PartnerEpc,
	PartnerIntegration,
	PartnerLender,
}

func (s Partner) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (Partner) Values() (kinds []string) {
	for _, s := range AllPartnerType {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Partner) IsValid() bool {
	for _, t := range AllPartnerType {
		if s == t {
			return true
		}
	}
	return false
}

func (s *Partner) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Partner(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid Partner", str)
	}
	return nil
}

func (s Partner) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

// **
// Partner Contact
// **

type PartnerContact string

const (
	PartnerContactPrimary         PartnerContact = "PRIMARY"
	PartnerContactOperations      PartnerContact = "OPERATIONS"
	PartnerContactApproval        PartnerContact = "APPROVAL"
	PartnerContactAccounting      PartnerContact = "ACCOUNTING"
	PartnerContactInvoicing       PartnerContact = "INVOICING"
	PartnerContactCustomerService PartnerContact = "CUSTOMER_SERVICE"
	PartnerContactCustom          PartnerContact = "CUSTOM"
)

var AllPartnerContact = []PartnerContact{
	PartnerContactPrimary,
	PartnerContactOperations,
	PartnerContactApproval,
	PartnerContactAccounting,
	PartnerContactInvoicing,
	PartnerContactCustomerService,
	PartnerContactCustom,
}

func (s PartnerContact) String() string {
	return string(s)
}

func (s PartnerContact) Humanize() string {
	switch s {
	case PartnerContactPrimary:
		return "Primary"
	case PartnerContactOperations:
		return "Operations"
	case PartnerContactAccounting:
		return "Accounting"
	case PartnerContactCustomerService:
		return "Customer Service"
	case PartnerContactInvoicing:
		return "Invoicing"
	default:
		return ""
	}
}

// Values provides list valid values for Enum.
func (PartnerContact) Values() (kinds []string) {
	for _, s := range AllPartnerContact {
		kinds = append(kinds, string(s))
	}
	return
}

func (s PartnerContact) IsValid() bool {
	for _, t := range AllPartnerContact {
		if s == t {
			return true
		}
	}
	return false
}

func (s *PartnerContact) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = PartnerContact(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid PartnerContact", str)
	}
	return nil
}

func (s PartnerContact) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

// **
// PartnerStatus
// **

type PartnerStatus string

const (
	PartnerStatusActive         PartnerStatus = "Active"
	PartnerStatusInActive       PartnerStatus = "InActive"
	PartnerStatusOnboarding     PartnerStatus = "Onboarding"
	PartnerStatusOnboardingDone PartnerStatus = "OnboardingDone"
)

var AllPartnerStatus = []PartnerStatus{
	PartnerStatusActive,
	PartnerStatusInActive,
	PartnerStatusOnboarding,
	PartnerStatusOnboardingDone,
}

func (s PartnerStatus) String() string {
	return string(s)
}

// Values provides list valid values for Enum.
func (PartnerStatus) Values() (kinds []string) {
	for _, s := range AllPartnerStatus {
		kinds = append(kinds, string(s))
	}
	return
}

func (s PartnerStatus) IsValid() bool {
	for _, t := range AllPartnerStatus {
		if s == t {
			return true
		}
	}
	return false
}

func (s *PartnerStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = PartnerStatus(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid PartnerStatus", str)
	}
	return nil
}

func (s PartnerStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}
