package homeOwner

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"roofix/ent"
	"roofix/ent/homeowner"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/str"
)

type Input struct {
	FirstName        string  `json:"firstName" validate:"required"`
	LastName         string  `json:"lastName" validate:"required"`
	Email            *string `json:"email"`
	Phone            *string `json:"phone"`
	StreetNumber     string  `json:"streetNumber" validate:"required"`
	StreetName       string  `json:"streetName" validate:"required"`
	City             string  `json:"city" validate:"required"`
	State            string  `json:"state" validate:"required"`
	StateAbbr        *string `json:"stateAbbr" validate:"required"`
	Zip              string  `json:"zip" validate:"required"`
	FormattedAddress *string `json:"formattedAddress"`
	Latitude         float64 `json:"latitude" validate:"latitude,omitempty"`
	Longitude        float64 `json:"longitude" validate:"longitude,omitempty"`
}

func CreateTx(ctx context.Context, tx *ent.Tx, partnerID *string, inp *Input, homeOwnerHashCheck bool) (string, error) {
	if tx == nil {
		log.Warn("tx param is nil")
		return "", nil
	}

	if inp == nil {
		log.Warn("inp param is nil")
		return "", nil
	}

	inpHash := Hash(inp)
	// check duplicate
	if first, err := tx.HomeOwner.Query().Where(homeowner.Hash(inpHash)).Select(homeowner.FieldID).First(ctx); err != nil {
		if !ent.IsNotFound(err) {
			log.Error(err)
		}
	} else if first != nil {
		return first.ID, err
	}

	if v, err := CreateQry(tx.HomeOwner.Create(), inp, &inpHash, partnerID).Save(ctx); err != nil {
		return "", err
	} else {
		return v.ID, nil
	}
}

func CreateQry(qry *ent.HomeOwnerCreate, inp *Input, inpHash, partnerID *string) *ent.HomeOwnerCreate {
	qry.SetFirstName(inp.FirstName).
		SetLastName(inp.LastName).
		SetNillableEmail(inp.Email).
		SetNillablePhone(inp.Phone).
		SetStreetNumber(inp.StreetNumber).
		SetStreetName(inp.StreetName).
		SetCity(inp.City).
		SetState(inp.State).
		SetNillableStateAbbr(inp.StateAbbr).
		SetZip(inp.Zip).
		SetNillableFormattedAddress(inp.FormattedAddress).
		SetLatitude(inp.Latitude).
		SetLongitude(inp.Longitude).
		SetNillableHash(inpHash)

	if partnerID != nil {
		qry.SetPartnerID(*partnerID)
	}

	return qry
}

func Hash(inp *Input) string {
	return crypt.MD5Hash(fmt.Sprintf(
		"%s%s%s%s%s%s%s%s%s%f%f",
		inp.FirstName,
		inp.LastName,
		str.Val(inp.Email),
		str.Val(inp.Phone),
		inp.StreetNumber,
		inp.StreetName,
		inp.City,
		inp.State,
		inp.Zip,
		inp.Latitude,
		inp.Longitude,
	))
}
