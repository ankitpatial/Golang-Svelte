package rfx

import (
	"context"
	"roofix/pkg/util/validate"
)

type CreateUserInput struct {
	ID                      string
	Email                   string
	FirstName               string
	LastName                string
	Location                string
	Phone                   string
	Status                  uint8
	Role                    uint8
	InternalRole            *uint8
	AcceptedGeneralTerms    *bool
	AcceptedTermsAndPrivacy *bool
}

func CreateUser(ctx context.Context, apiUID string, inp *CreateUserInput) error {
	if err := validate.Struct(inp); err != nil {
		return err
	}

	return nil
}
