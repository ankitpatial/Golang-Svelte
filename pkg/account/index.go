package account

import (
	"fmt"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/model"
	"time"
)

const MaxWrongAttempts = 5
const AccLocDuration = time.Minute

var sessionUserField = []string{
	user.FieldEmail,
	user.FieldFirstName,
	user.FieldLastName,
	user.FieldRole,
	user.FieldStatus,
	user.FieldPicture,
	user.FieldPhone,
}

type User struct {
	ID             string             `json:"id"`
	SessionID      string             `json:"-"`
	Token          string             `json:"token"`
	Email          string             `json:"email"`
	FirstName      string             `json:"firstName"`
	LastName       string             `json:"lastName"`
	Role           enum.Role          `json:"role"`
	Status         enum.AccountStatus `json:"status"`
	Phone          string             `json:"phone"`
	Picture        *string            `json:"picture"`
	Partner        *UserPartnerInfo   `json:"partner"`
	IsAdmin        bool               `json:"isAdmin"`
	IsCompanyAdmin bool               `json:"isCompanyAdmin"`
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type PartnerUser struct {
	ID              string                  `json:"id"`
	FirstName       string                  `json:"firstName"`
	LastName        string                  `json:"lastName"`
	Email           string                  `json:"email"`
	Phone           string                  `json:"phone"`
	Picture         string                  `json:"picture"`
	PartnerName     string                  `json:"partnerName"`
	PartnerUserType enum.PartnerContact     `json:"partnerUserType"`
	PartnerUserRole enum.PartnerContactRole `json:"partnerUserRole"`
}

type UserPartnerInfo struct {
	ID                string                  `json:"id"`
	Type              enum.Partner            `json:"type"`
	PartnerName       string                  `json:"partnerName"`
	IsPrimary         bool                    `json:"isPrimary"`
	Status            enum.PartnerStatus      `json:"status"`
	ContactID         string                  `json:"-"`
	ContactType       enum.PartnerContact     `json:"contactType"`
	Role              enum.PartnerContactRole `json:"role"`
	MobileAppSettings model.MobileAppSettings `json:"mobileAppSettings,omitempty"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateUserInput struct {
	Email     string    `json:"email" validate:"required,email"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Phone     *string   `json:"phone"`
	Role      enum.Role `json:"role"`
	Note      *string   `json:"note"`
}

type EmailAddress struct {
	Name  string
	Email string
}

type SetPasswordInput struct {
	Token           string `json:"token" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

func (upi UserPartnerInfo) IsAdmin() bool {
	return upi.Role == enum.PartnerContactRoleAdmin
}

func (upi UserPartnerInfo) IsCoreUser() bool {
	switch upi.ContactType {
	case enum.PartnerContactPrimary,
		enum.PartnerContactOperations,
		enum.PartnerContactAccounting,
		enum.PartnerContactCustomerService,
		enum.PartnerContactInvoicing:
		return true
	default:
		return false
	}
}
