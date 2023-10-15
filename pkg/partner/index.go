package partner

import (
	"roofix/pkg/enum"
	"time"
)

const (
	EcostellaID = "01H3CYX8K04WJGMMC755TVZVCS"
	ResideoID   = "01H3CYX8K0M7780GDHV91STK1B"
)

type Contact struct {
	ID          string              `json:"id"`
	Email       string              `json:"email"`
	Name        string              `json:"name"`
	Phone       string              `json:"phone"`
	Type        enum.PartnerContact `json:"typeID"`
	PartnerID   string              `json:"partnerID"`
	PartnerName string              `json:"partnerName"`
}

type Invite struct {
	ID          string       `json:"id"`
	Type        enum.Partner `json:"type"`
	CompanyName string       `json:"companyName"`
	ContactID   *string      `json:"contactID"`
	UserID      string       `json:"userID"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	CreatedAt   *time.Time   `json:"createdAt"`
}

type JobStats struct {
	// ID is partnerID
	ID string `json:"id"`
	// Name is partner name
	Name                   string             `json:"name"`
	Status                 enum.PartnerStatus `json:"status"`
	NewCount               int                `json:"newCount"`
	NewCountFlagged        int                `json:"newCountFlagged"`
	ContactedCount         int                `json:"contactedCount"`
	ContactedCountFlagged  int                `json:"contactedCountFlagged"`
	ConfirmedCount         int                `json:"confirmedCount"`
	ConfirmedCountFlagged  int                `json:"confirmedCountFlagged"`
	PermittingCount        int                `json:"permittingCount"`
	PermittingCountFlagged int                `json:"permittingCountFlagged"`
	DelayedCount           int                `json:"delayedCount"`
	ScheduledCount         int                `json:"scheduledCount"`
	ScheduledCountFlagged  int                `json:"scheduledCountFlagged"`
	InProgressCount        int                `json:"inProgressCount"`
	InProgressCountFlagged int                `json:"inProgressCountFlagged"`
	InstalledCount         int                `json:"installedCount"`
	InstalledCountFlagged  int                `json:"installedCountFlagged"`
	InvoicedCount          int                `json:"invoicedCount"`
	InvoicedCountFlagged   int                `json:"invoicedCountFlagged"`
	Total                  int                `json:"total"`
	TotalFlagged           int                `json:"totalFlagged"`
}

type Job struct {
	ID                 string           `json:"id"` // invite ID
	OwnerFirstName     string           `json:"owner_first_name"`
	OwnerLastName      string           `json:"owner_last_name"`
	StreetNumber       string           `json:"street_number"`
	StreetName         string           `json:"street_name"`
	City               string           `json:"city"`
	State              string           `json:"state"`
	Region             string           `json:"region"`
	Zip                string           `json:"zip"`
	Latitude           float64          `json:"latitude"`
	Longitude          float64          `json:"longitude"`
	RepFirstName       string           `json:"rep_first_name"`
	RepLastName        string           `json:"rep_last_name"`
	RepEmail           string           `json:"rep_email"`
	RepMobile          string           `json:"rep_mobile"`
	CompanyName        string           `json:"company_name"`
	Status             enum.JobProgress `json:"status"`
	StatusAt           time.Time        `json:"status_at"`
	Notes              *string          `json:"notes"`
	Price              *float64         `json:"price"`
	ShingleColor       string           `json:"shingle_color"`
	PermitRequired     string           `json:"permit_required"`
	InspectionRequired string           `json:"inspection_required"`
	InstallDate        *time.Time       `json:"install_date"`
	InspectionDate     *time.Time       `json:"inspection_date"`
	CompletionDate     *time.Time       `json:"completion_date"`
}

type ContactUserInput struct {
	ID            *string                  `json:"id"`
	Type          enum.PartnerContact      `json:"type" validate:"required"`
	Role          *enum.PartnerContactRole `json:"role"`
	AccountStatus *enum.AccountStatus      `json:"accountStatus" `
	UserID        string                   `json:"userID"`
	FirstName     string                   `json:"firstName" validate:"required"`
	LastName      string                   `json:"lastName" validate:"required"`
	Phone         string                   `json:"phone" validate:"required"`
	Email         string                   `json:"email" validate:"required,email"`
	OtherEmail    string                   `json:"otherEmail" validate:"omitempty,email"`
	Title         *string                  `json:"title"`
	Description   *string                  `json:"description"`
}

type BasicDetail struct {
	ID              string `json:"id"`
	ExternalID      *string
	Type            enum.Partner        `json:"type"`
	Name            string              `json:"name" validate:"required"`
	Phone           *string             `json:"phone"`
	Address         *string             `json:"address"`
	Latitude        *float64            `json:"latitude"`
	Longitude       *float64            `json:"longitude"`
	Website         *string             `json:"website"`
	IsNationWide    bool                `json:"isNationWide"`
	CrewCount       int                 `json:"crewCount"`
	JobCapacity     int                 `json:"jobCapacity"`
	YearsInBusiness *int                `json:"yearsInBusiness"`
	Status          *enum.PartnerStatus `json:"status"`
}

type ServiceState struct {
	PartnerID      string `validate:"required"`
	Country        string `validate:"required"`
	State          string `validate:"required"`
	LicenseNo      *string
	LicenseExpDate *time.Time
	ProofDocID     *string
}

type ServiceCity struct {
	PartnerID  string `validate:"required"`
	PostalID   string `validate:"required"`
	Active     *bool
	LicenseNo  *string
	ProofDocID *string
}
