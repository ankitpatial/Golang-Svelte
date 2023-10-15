package model

import "roofix/pkg/enum"

type ExtraChargeCondition struct {
	Condition    enum.Condition
	ConditionVal any
	Val          float64
}
