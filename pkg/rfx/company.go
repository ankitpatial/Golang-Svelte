package rfx

import (
	"roofix/pkg/enum"
	"roofix/pkg/model"
	"strings"
)

type Company struct {
	ID               string
	Name             string
	ChargeType       enum.ExtraCharge
	ChargeVal        float64
	ChargeNote       string
	ChargeConditions []*model.ExtraChargeCondition
}

func CompanyExtraCharges(id string) *Company {
	if strings.TrimSpace(id) == "" {
		return nil
	}

	for _, c := range chargeThem() {
		if c.ID == id {
			return c
		}
	}
	return nil
}

func chargeThem() []*Company {
	return []*Company{
		{
			ID:         "1682443216511x829605057140883500",
			Name:       "WW Solar",
			ChargeType: enum.ExtraChargePercent,
			ChargeVal:  15,
			ChargeNote: "WW Solar, Additional Charges 15%",
		},
		{
			ID:         "1665613316086x177779357767696400",
			Name:       "Blue Raven",
			ChargeType: enum.ExtraChargePercent,
			ChargeVal:  4,
			ChargeNote: "Blue Raven, Additional Charges 4%",
		},
		{
			ID:         "1666565417200x503394046838833150",
			Name:       "Lumio Central",
			ChargeType: enum.ExtraChargeAmount,
			ChargeVal:  1000,
			ChargeNote: "Lumio, Additional Charges $1,000",
		},
		{
			ID:         "1665081039337x127793381322784770",
			Name:       "Lumio West",
			ChargeType: enum.ExtraChargeAmount,
			ChargeVal:  1000,
			ChargeNote: "Lumio, Additional Charges $1,000"},
		{
			ID:         "1666375028799x332596612424269800",
			Name:       "Lumio East",
			ChargeType: enum.ExtraChargeAmount,
			ChargeVal:  1000,
			ChargeNote: "Lumio, Additional Charges $1,000",
		},
		{
			ID:         "1666039317748x589698696859877400",
			Name:       "LGCY",
			ChargeType: enum.ExtraChargeAmount,
			ChargeVal:  1000,
			ChargeNote: "Lumio, Additional Charges $1,000",
		},
		{
			ID:         "1684881454514x649937849821954000",
			Name:       "EnergyPal",
			ChargeType: enum.ExtraChargePerSQ,
			ChargeVal:  40,
		},
		//{
		//	ID:         "abc",
		//	Name:       "ABC Company",
		//	ChargeType: enum.ExtraChargeConditional,
		//	ChargeConditions: []*model.ExtraChargeCondition{
		//		{
		//			Condition:    enum.ConditionSubtractIfStateIn,
		//			ConditionVal: []string{"Texas", "TX"},
		//			Val:          250,
		//		},
		//		{
		//			Condition:    enum.ConditionAddIfStateIn,
		//			ConditionVal: []string{"Oklahoma", "OK", "Arkansas", "AR"},
		//			Val:          1000,
		//		},
		//	},
		//},
	}
}
