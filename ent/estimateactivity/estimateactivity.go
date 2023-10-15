// Code generated by ent, DO NOT EDIT.

package estimateactivity

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the estimateactivity type in the database.
	Label = "estimate_activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// EdgeEstimate holds the string denoting the estimate edge name in mutations.
	EdgeEstimate = "estimate"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeCreatorAPI holds the string denoting the creator_api edge name in mutations.
	EdgeCreatorAPI = "creator_api"
	// Table holds the table name of the estimateactivity in the database.
	Table = "estimate_activities"
	// EstimateTable is the table that holds the estimate relation/edge.
	EstimateTable = "estimate_activities"
	// EstimateInverseTable is the table name for the Estimate entity.
	// It exists in this package in order to avoid circular dependency with the "estimate" package.
	EstimateInverseTable = "estimates"
	// EstimateColumn is the table column denoting the estimate relation/edge.
	EstimateColumn = "estimate_id"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "estimate_activities"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "creator_id"
	// CreatorAPITable is the table that holds the creator_api relation/edge.
	CreatorAPITable = "estimate_activities"
	// CreatorAPIInverseTable is the table name for the ApiUser entity.
	// It exists in this package in order to avoid circular dependency with the "apiuser" package.
	CreatorAPIInverseTable = "api_users"
	// CreatorAPIColumn is the table column denoting the creator_api relation/edge.
	CreatorAPIColumn = "api_user_id"
)

// Columns holds all SQL columns for estimateactivity fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldDescription,
	FieldRaw,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "estimate_activities"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"api_user_id",
	"estimate_id",
	"creator_id",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the EstimateActivity queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByEstimateField orders the results by estimate field.
func ByEstimateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEstimateStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreatorField orders the results by creator field.
func ByCreatorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreatorAPIField orders the results by creator_api field.
func ByCreatorAPIField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorAPIStep(), sql.OrderByField(field, opts...))
	}
}
func newEstimateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EstimateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EstimateTable, EstimateColumn),
	)
}
func newCreatorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
	)
}
func newCreatorAPIStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorAPIInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatorAPITable, CreatorAPIColumn),
	)
}
