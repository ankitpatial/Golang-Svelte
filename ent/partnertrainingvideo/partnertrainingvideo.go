// Code generated by ent, DO NOT EDIT.

package partnertrainingvideo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the partnertrainingvideo type in the database.
	Label = "partner_training_video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// EdgePartner holds the string denoting the partner edge name in mutations.
	EdgePartner = "partner"
	// Table holds the table name of the partnertrainingvideo in the database.
	Table = "partner_training_videos"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "partner_training_videos"
	// VideoInverseTable is the table name for the TrainingVideo entity.
	// It exists in this package in order to avoid circular dependency with the "trainingvideo" package.
	VideoInverseTable = "training_videos"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_id"
	// PartnerTable is the table that holds the partner relation/edge.
	PartnerTable = "partner_training_videos"
	// PartnerInverseTable is the table name for the Partner entity.
	// It exists in this package in order to avoid circular dependency with the "partner" package.
	PartnerInverseTable = "partners"
	// PartnerColumn is the table column denoting the partner relation/edge.
	PartnerColumn = "partner_id"
)

// Columns holds all SQL columns for partnertrainingvideo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEnabled,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "partner_training_videos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"partner_id",
	"video_id",
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
	// DefaultEnabled holds the default value on creation for the "enabled" field.
	DefaultEnabled bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the PartnerTrainingVideo queries.
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

// ByEnabled orders the results by the enabled field.
func ByEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnabled, opts...).ToFunc()
}

// ByVideoField orders the results by video field.
func ByVideoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoStep(), sql.OrderByField(field, opts...))
	}
}

// ByPartnerField orders the results by partner field.
func ByPartnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPartnerStep(), sql.OrderByField(field, opts...))
	}
}
func newVideoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
	)
}
func newPartnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PartnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PartnerTable, PartnerColumn),
	)
}
