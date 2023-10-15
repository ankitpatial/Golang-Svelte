// Code generated by ent, DO NOT EDIT.

package pricing

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the pricing type in the database.
	Label = "pricing"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPostalCountry holds the string denoting the postal_country field in the database.
	FieldPostalCountry = "postal_country"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldPricePer holds the string denoting the price_per field in the database.
	FieldPricePer = "price_per"
	// EdgePostal holds the string denoting the postal edge name in mutations.
	EdgePostal = "postal"
	// Table holds the table name of the pricing in the database.
	Table = "pricing"
	// PostalTable is the table that holds the postal relation/edge.
	PostalTable = "pricing"
	// PostalInverseTable is the table name for the PostalCode entity.
	// It exists in this package in order to avoid circular dependency with the "postalcode" package.
	PostalInverseTable = "postal_codes"
	// PostalColumn is the table column denoting the postal relation/edge.
	PostalColumn = "postal_id"
)

// Columns holds all SQL columns for pricing fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPostalCountry,
	FieldPostalCode,
	FieldProductID,
	FieldDescription,
	FieldPrice,
	FieldPricePer,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pricing"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"postal_id",
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
	// PostalCountryValidator is a validator for the "postal_country" field. It is called by the builders before save.
	PostalCountryValidator func(string) error
	// PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
	// PricePerValidator is a validator for the "price_per" field. It is called by the builders before save.
	PricePerValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Pricing queries.
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

// ByPostalCountry orders the results by the postal_country field.
func ByPostalCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCountry, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByPricePer orders the results by the price_per field.
func ByPricePer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPricePer, opts...).ToFunc()
}

// ByPostalField orders the results by postal field.
func ByPostalField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostalStep(), sql.OrderByField(field, opts...))
	}
}
func newPostalStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostalInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostalTable, PostalColumn),
	)
}