// Code generated by ent, DO NOT EDIT.

package channelmessage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the channelmessage type in the database.
	Label = "channel_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldFromName holds the string denoting the from_name field in the database.
	FieldFromName = "from_name"
	// FieldToName holds the string denoting the to_name field in the database.
	FieldToName = "to_name"
	// FieldPrivate holds the string denoting the private field in the database.
	FieldPrivate = "private"
	// FieldUnread holds the string denoting the unread field in the database.
	FieldUnread = "unread"
	// EdgeReads holds the string denoting the reads edge name in mutations.
	EdgeReads = "reads"
	// EdgeChannel holds the string denoting the channel edge name in mutations.
	EdgeChannel = "channel"
	// EdgeFrom holds the string denoting the from edge name in mutations.
	EdgeFrom = "from"
	// EdgeFromAPIUser holds the string denoting the from_api_user edge name in mutations.
	EdgeFromAPIUser = "from_api_user"
	// EdgeTo holds the string denoting the to edge name in mutations.
	EdgeTo = "to"
	// Table holds the table name of the channelmessage in the database.
	Table = "channel_messages"
	// ReadsTable is the table that holds the reads relation/edge.
	ReadsTable = "channel_message_reads"
	// ReadsInverseTable is the table name for the ChannelMessageRead entity.
	// It exists in this package in order to avoid circular dependency with the "channelmessageread" package.
	ReadsInverseTable = "channel_message_reads"
	// ReadsColumn is the table column denoting the reads relation/edge.
	ReadsColumn = "channel_message_id"
	// ChannelTable is the table that holds the channel relation/edge.
	ChannelTable = "channel_messages"
	// ChannelInverseTable is the table name for the Channel entity.
	// It exists in this package in order to avoid circular dependency with the "channel" package.
	ChannelInverseTable = "channels"
	// ChannelColumn is the table column denoting the channel relation/edge.
	ChannelColumn = "channel_id"
	// FromTable is the table that holds the from relation/edge.
	FromTable = "channel_messages"
	// FromInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	FromInverseTable = "users"
	// FromColumn is the table column denoting the from relation/edge.
	FromColumn = "from_user_id"
	// FromAPIUserTable is the table that holds the from_api_user relation/edge.
	FromAPIUserTable = "channel_messages"
	// FromAPIUserInverseTable is the table name for the ApiUser entity.
	// It exists in this package in order to avoid circular dependency with the "apiuser" package.
	FromAPIUserInverseTable = "api_users"
	// FromAPIUserColumn is the table column denoting the from_api_user relation/edge.
	FromAPIUserColumn = "from_api_user_id"
	// ToTable is the table that holds the to relation/edge.
	ToTable = "channel_messages"
	// ToInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ToInverseTable = "users"
	// ToColumn is the table column denoting the to relation/edge.
	ToColumn = "to_user_id"
)

// Columns holds all SQL columns for channelmessage fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldMessage,
	FieldFromName,
	FieldToName,
	FieldPrivate,
	FieldUnread,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "channel_messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"from_api_user_id",
	"channel_id",
	"from_user_id",
	"to_user_id",
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// FromNameValidator is a validator for the "from_name" field. It is called by the builders before save.
	FromNameValidator func(string) error
	// ToNameValidator is a validator for the "to_name" field. It is called by the builders before save.
	ToNameValidator func(string) error
	// DefaultPrivate holds the default value on creation for the "private" field.
	DefaultPrivate bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the ChannelMessage queries.
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

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByFromName orders the results by the from_name field.
func ByFromName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromName, opts...).ToFunc()
}

// ByToName orders the results by the to_name field.
func ByToName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToName, opts...).ToFunc()
}

// ByPrivate orders the results by the private field.
func ByPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivate, opts...).ToFunc()
}

// ByUnread orders the results by the unread field.
func ByUnread(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnread, opts...).ToFunc()
}

// ByReadsCount orders the results by reads count.
func ByReadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReadsStep(), opts...)
	}
}

// ByReads orders the results by reads terms.
func ByReads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByChannelField orders the results by channel field.
func ByChannelField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChannelStep(), sql.OrderByField(field, opts...))
	}
}

// ByFromField orders the results by from field.
func ByFromField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromStep(), sql.OrderByField(field, opts...))
	}
}

// ByFromAPIUserField orders the results by from_api_user field.
func ByFromAPIUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromAPIUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByToField orders the results by to field.
func ByToField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToStep(), sql.OrderByField(field, opts...))
	}
}
func newReadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReadsTable, ReadsColumn),
	)
}
func newChannelStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChannelInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
	)
}
func newFromStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FromTable, FromColumn),
	)
}
func newFromAPIUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromAPIUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FromAPIUserTable, FromAPIUserColumn),
	)
}
func newToStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ToTable, ToColumn),
	)
}
