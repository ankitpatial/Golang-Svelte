// Code generated by ent, DO NOT EDIT.

package partnercontact

import (
	"roofix/ent/predicate"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldUpdatedAt, v))
}

// PartnerID applies equality check predicate on the "partner_id" field. It's identical to PartnerIDEQ.
func PartnerID(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldPartnerID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldUserID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldDescription, v))
}

// InvoicingEmail applies equality check predicate on the "invoicing_email" field. It's identical to InvoicingEmailEQ.
func InvoicingEmail(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldInvoicingEmail, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldUpdatedAt, v))
}

// PartnerIDEQ applies the EQ predicate on the "partner_id" field.
func PartnerIDEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldPartnerID, v))
}

// PartnerIDNEQ applies the NEQ predicate on the "partner_id" field.
func PartnerIDNEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldPartnerID, v))
}

// PartnerIDIn applies the In predicate on the "partner_id" field.
func PartnerIDIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldPartnerID, vs...))
}

// PartnerIDNotIn applies the NotIn predicate on the "partner_id" field.
func PartnerIDNotIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldPartnerID, vs...))
}

// PartnerIDGT applies the GT predicate on the "partner_id" field.
func PartnerIDGT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldPartnerID, v))
}

// PartnerIDGTE applies the GTE predicate on the "partner_id" field.
func PartnerIDGTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldPartnerID, v))
}

// PartnerIDLT applies the LT predicate on the "partner_id" field.
func PartnerIDLT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldPartnerID, v))
}

// PartnerIDLTE applies the LTE predicate on the "partner_id" field.
func PartnerIDLTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldPartnerID, v))
}

// PartnerIDContains applies the Contains predicate on the "partner_id" field.
func PartnerIDContains(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContains(FieldPartnerID, v))
}

// PartnerIDHasPrefix applies the HasPrefix predicate on the "partner_id" field.
func PartnerIDHasPrefix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasPrefix(FieldPartnerID, v))
}

// PartnerIDHasSuffix applies the HasSuffix predicate on the "partner_id" field.
func PartnerIDHasSuffix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasSuffix(FieldPartnerID, v))
}

// PartnerIDEqualFold applies the EqualFold predicate on the "partner_id" field.
func PartnerIDEqualFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEqualFold(FieldPartnerID, v))
}

// PartnerIDContainsFold applies the ContainsFold predicate on the "partner_id" field.
func PartnerIDContainsFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContainsFold(FieldPartnerID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContainsFold(FieldUserID, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v enum.PartnerContactRole) predicate.PartnerContact {
	vc := v
	return predicate.PartnerContact(sql.FieldEQ(FieldRole, vc))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v enum.PartnerContactRole) predicate.PartnerContact {
	vc := v
	return predicate.PartnerContact(sql.FieldNEQ(FieldRole, vc))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...enum.PartnerContactRole) predicate.PartnerContact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PartnerContact(sql.FieldIn(FieldRole, v...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...enum.PartnerContactRole) predicate.PartnerContact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PartnerContact(sql.FieldNotIn(FieldRole, v...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enum.PartnerContact) predicate.PartnerContact {
	vc := v
	return predicate.PartnerContact(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enum.PartnerContact) predicate.PartnerContact {
	vc := v
	return predicate.PartnerContact(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enum.PartnerContact) predicate.PartnerContact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PartnerContact(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enum.PartnerContact) predicate.PartnerContact {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PartnerContact(sql.FieldNotIn(FieldType, v...))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContainsFold(FieldDescription, v))
}

// InvoicingEmailEQ applies the EQ predicate on the "invoicing_email" field.
func InvoicingEmailEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEQ(FieldInvoicingEmail, v))
}

// InvoicingEmailNEQ applies the NEQ predicate on the "invoicing_email" field.
func InvoicingEmailNEQ(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNEQ(FieldInvoicingEmail, v))
}

// InvoicingEmailIn applies the In predicate on the "invoicing_email" field.
func InvoicingEmailIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIn(FieldInvoicingEmail, vs...))
}

// InvoicingEmailNotIn applies the NotIn predicate on the "invoicing_email" field.
func InvoicingEmailNotIn(vs ...string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotIn(FieldInvoicingEmail, vs...))
}

// InvoicingEmailGT applies the GT predicate on the "invoicing_email" field.
func InvoicingEmailGT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGT(FieldInvoicingEmail, v))
}

// InvoicingEmailGTE applies the GTE predicate on the "invoicing_email" field.
func InvoicingEmailGTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldGTE(FieldInvoicingEmail, v))
}

// InvoicingEmailLT applies the LT predicate on the "invoicing_email" field.
func InvoicingEmailLT(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLT(FieldInvoicingEmail, v))
}

// InvoicingEmailLTE applies the LTE predicate on the "invoicing_email" field.
func InvoicingEmailLTE(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldLTE(FieldInvoicingEmail, v))
}

// InvoicingEmailContains applies the Contains predicate on the "invoicing_email" field.
func InvoicingEmailContains(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContains(FieldInvoicingEmail, v))
}

// InvoicingEmailHasPrefix applies the HasPrefix predicate on the "invoicing_email" field.
func InvoicingEmailHasPrefix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasPrefix(FieldInvoicingEmail, v))
}

// InvoicingEmailHasSuffix applies the HasSuffix predicate on the "invoicing_email" field.
func InvoicingEmailHasSuffix(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldHasSuffix(FieldInvoicingEmail, v))
}

// InvoicingEmailIsNil applies the IsNil predicate on the "invoicing_email" field.
func InvoicingEmailIsNil() predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldIsNull(FieldInvoicingEmail))
}

// InvoicingEmailNotNil applies the NotNil predicate on the "invoicing_email" field.
func InvoicingEmailNotNil() predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldNotNull(FieldInvoicingEmail))
}

// InvoicingEmailEqualFold applies the EqualFold predicate on the "invoicing_email" field.
func InvoicingEmailEqualFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldEqualFold(FieldInvoicingEmail, v))
}

// InvoicingEmailContainsFold applies the ContainsFold predicate on the "invoicing_email" field.
func InvoicingEmailContainsFold(v string) predicate.PartnerContact {
	return predicate.PartnerContact(sql.FieldContainsFold(FieldInvoicingEmail, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPartner applies the HasEdge predicate on the "partner" edge.
func HasPartner() predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PartnerTable, PartnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPartnerWith applies the HasEdge predicate on the "partner" edge with a given conditions (other predicates).
func HasPartnerWith(preds ...predicate.Partner) predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		step := newPartnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSessions applies the HasEdge predicate on the "sessions" edge.
func HasSessions() predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSessionsWith applies the HasEdge predicate on the "sessions" edge with a given conditions (other predicates).
func HasSessionsWith(preds ...predicate.UserSession) predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		step := newSessionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PartnerContact) predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PartnerContact) predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
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
func Not(p predicate.PartnerContact) predicate.PartnerContact {
	return predicate.PartnerContact(func(s *sql.Selector) {
		p(s.Not())
	})
}
