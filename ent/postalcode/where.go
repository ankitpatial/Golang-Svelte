// Code generated by ent, DO NOT EDIT.

package postalcode

import (
	"roofix/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldUpdatedAt, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCountry, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCode, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCity, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldState, v))
}

// StateAbr applies equality check predicate on the "state_abr" field. It's identical to StateAbrEQ.
func StateAbr(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldStateAbr, v))
}

// RegionID applies equality check predicate on the "region_id" field. It's identical to RegionIDEQ.
func RegionID(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldRegionID, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldLongitude, v))
}

// Accuracy applies equality check predicate on the "accuracy" field. It's identical to AccuracyEQ.
func Accuracy(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldAccuracy, v))
}

// ServiceArea applies equality check predicate on the "service_area" field. It's identical to ServiceAreaEQ.
func ServiceArea(v bool) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldServiceArea, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldUpdatedAt, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContainsFold(FieldCountry, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContainsFold(FieldCode, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContainsFold(FieldCity, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasSuffix(FieldState, v))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContainsFold(FieldState, v))
}

// StateAbrEQ applies the EQ predicate on the "state_abr" field.
func StateAbrEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldStateAbr, v))
}

// StateAbrNEQ applies the NEQ predicate on the "state_abr" field.
func StateAbrNEQ(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldStateAbr, v))
}

// StateAbrIn applies the In predicate on the "state_abr" field.
func StateAbrIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldStateAbr, vs...))
}

// StateAbrNotIn applies the NotIn predicate on the "state_abr" field.
func StateAbrNotIn(vs ...string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldStateAbr, vs...))
}

// StateAbrGT applies the GT predicate on the "state_abr" field.
func StateAbrGT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldStateAbr, v))
}

// StateAbrGTE applies the GTE predicate on the "state_abr" field.
func StateAbrGTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldStateAbr, v))
}

// StateAbrLT applies the LT predicate on the "state_abr" field.
func StateAbrLT(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldStateAbr, v))
}

// StateAbrLTE applies the LTE predicate on the "state_abr" field.
func StateAbrLTE(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldStateAbr, v))
}

// StateAbrContains applies the Contains predicate on the "state_abr" field.
func StateAbrContains(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContains(FieldStateAbr, v))
}

// StateAbrHasPrefix applies the HasPrefix predicate on the "state_abr" field.
func StateAbrHasPrefix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasPrefix(FieldStateAbr, v))
}

// StateAbrHasSuffix applies the HasSuffix predicate on the "state_abr" field.
func StateAbrHasSuffix(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldHasSuffix(FieldStateAbr, v))
}

// StateAbrEqualFold applies the EqualFold predicate on the "state_abr" field.
func StateAbrEqualFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEqualFold(FieldStateAbr, v))
}

// StateAbrContainsFold applies the ContainsFold predicate on the "state_abr" field.
func StateAbrContainsFold(v string) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldContainsFold(FieldStateAbr, v))
}

// RegionIDEQ applies the EQ predicate on the "region_id" field.
func RegionIDEQ(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldRegionID, v))
}

// RegionIDNEQ applies the NEQ predicate on the "region_id" field.
func RegionIDNEQ(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldRegionID, v))
}

// RegionIDIn applies the In predicate on the "region_id" field.
func RegionIDIn(vs ...uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldRegionID, vs...))
}

// RegionIDNotIn applies the NotIn predicate on the "region_id" field.
func RegionIDNotIn(vs ...uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldRegionID, vs...))
}

// RegionIDGT applies the GT predicate on the "region_id" field.
func RegionIDGT(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldRegionID, v))
}

// RegionIDGTE applies the GTE predicate on the "region_id" field.
func RegionIDGTE(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldRegionID, v))
}

// RegionIDLT applies the LT predicate on the "region_id" field.
func RegionIDLT(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldRegionID, v))
}

// RegionIDLTE applies the LTE predicate on the "region_id" field.
func RegionIDLTE(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldRegionID, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldLatitude, v))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldLongitude, v))
}

// AccuracyEQ applies the EQ predicate on the "accuracy" field.
func AccuracyEQ(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldAccuracy, v))
}

// AccuracyNEQ applies the NEQ predicate on the "accuracy" field.
func AccuracyNEQ(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldAccuracy, v))
}

// AccuracyIn applies the In predicate on the "accuracy" field.
func AccuracyIn(vs ...uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldIn(FieldAccuracy, vs...))
}

// AccuracyNotIn applies the NotIn predicate on the "accuracy" field.
func AccuracyNotIn(vs ...uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNotIn(FieldAccuracy, vs...))
}

// AccuracyGT applies the GT predicate on the "accuracy" field.
func AccuracyGT(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGT(FieldAccuracy, v))
}

// AccuracyGTE applies the GTE predicate on the "accuracy" field.
func AccuracyGTE(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldGTE(FieldAccuracy, v))
}

// AccuracyLT applies the LT predicate on the "accuracy" field.
func AccuracyLT(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLT(FieldAccuracy, v))
}

// AccuracyLTE applies the LTE predicate on the "accuracy" field.
func AccuracyLTE(v uint8) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldLTE(FieldAccuracy, v))
}

// ServiceAreaEQ applies the EQ predicate on the "service_area" field.
func ServiceAreaEQ(v bool) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldEQ(FieldServiceArea, v))
}

// ServiceAreaNEQ applies the NEQ predicate on the "service_area" field.
func ServiceAreaNEQ(v bool) predicate.PostalCode {
	return predicate.PostalCode(sql.FieldNEQ(FieldServiceArea, v))
}

// HasPricing applies the HasEdge predicate on the "pricing" edge.
func HasPricing() predicate.PostalCode {
	return predicate.PostalCode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PricingTable, PricingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPricingWith applies the HasEdge predicate on the "pricing" edge with a given conditions (other predicates).
func HasPricingWith(preds ...predicate.Pricing) predicate.PostalCode {
	return predicate.PostalCode(func(s *sql.Selector) {
		step := newPricingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PostalCode) predicate.PostalCode {
	return predicate.PostalCode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PostalCode) predicate.PostalCode {
	return predicate.PostalCode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PostalCode) predicate.PostalCode {
	return predicate.PostalCode(func(s *sql.Selector) {
		p(s.Not())
	})
}
