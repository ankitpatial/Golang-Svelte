// Code generated by ent, DO NOT EDIT.

package partnerservicecity

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the partnerservicecity type in the database.
	Label = "partner_service_city"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPostalID holds the string denoting the postal_id field in the database.
	FieldPostalID = "postal_id"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNaicsCode holds the string denoting the naics_code field in the database.
	FieldNaicsCode = "naics_code"
	// FieldLicenseNo holds the string denoting the license_no field in the database.
	FieldLicenseNo = "license_no"
	// FieldProofDocID holds the string denoting the proof_doc_id field in the database.
	FieldProofDocID = "proof_doc_id"
	// EdgePartner holds the string denoting the partner edge name in mutations.
	EdgePartner = "partner"
	// Table holds the table name of the partnerservicecity in the database.
	Table = "partner_service_cities"
	// PartnerTable is the table that holds the partner relation/edge.
	PartnerTable = "partner_service_cities"
	// PartnerInverseTable is the table name for the Partner entity.
	// It exists in this package in order to avoid circular dependency with the "partner" package.
	PartnerInverseTable = "partners"
	// PartnerColumn is the table column denoting the partner relation/edge.
	PartnerColumn = "partner_id"
)

// Columns holds all SQL columns for partnerservicecity fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPostalID,
	FieldActive,
	FieldName,
	FieldNaicsCode,
	FieldLicenseNo,
	FieldProofDocID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "partner_service_cities"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"partner_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// PostalIDValidator is a validator for the "postal_id" field. It is called by the builders before save.
	PostalIDValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LicenseNoValidator is a validator for the "license_no" field. It is called by the builders before save.
	LicenseNoValidator func(string) error
	// ProofDocIDValidator is a validator for the "proof_doc_id" field. It is called by the builders before save.
	ProofDocIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the PartnerServiceCity queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPostalID orders the results by the postal_id field.
func ByPostalID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalID, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByNaicsCode orders the results by the naics_code field.
func ByNaicsCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNaicsCode, opts...).ToFunc()
}

// ByLicenseNo orders the results by the license_no field.
func ByLicenseNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseNo, opts...).ToFunc()
}

// ByProofDocID orders the results by the proof_doc_id field.
func ByProofDocID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProofDocID, opts...).ToFunc()
}

// ByPartnerField orders the results by partner field.
func ByPartnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPartnerStep(), sql.OrderByField(field, opts...))
	}
}
func newPartnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PartnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PartnerTable, PartnerColumn),
	)
}
