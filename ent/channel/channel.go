// Code generated by ent, DO NOT EDIT.

package channel

import (
	"fmt"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
)

const (
	// Label holds the string label denoting the channel type in the database.
	Label = "channel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTopic holds the string denoting the topic field in the database.
	FieldTopic = "topic"
	// FieldRefID holds the string denoting the ref_id field in the database.
	FieldRefID = "ref_id"
	// EdgeSubscriptions holds the string denoting the subscriptions edge name in mutations.
	EdgeSubscriptions = "subscriptions"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the channel in the database.
	Table = "channels"
	// SubscriptionsTable is the table that holds the subscriptions relation/edge.
	SubscriptionsTable = "channel_subs"
	// SubscriptionsInverseTable is the table name for the ChannelSub entity.
	// It exists in this package in order to avoid circular dependency with the "channelsub" package.
	SubscriptionsInverseTable = "channel_subs"
	// SubscriptionsColumn is the table column denoting the subscriptions relation/edge.
	SubscriptionsColumn = "channel_id"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "channel_messages"
	// MessagesInverseTable is the table name for the ChannelMessage entity.
	// It exists in this package in order to avoid circular dependency with the "channelmessage" package.
	MessagesInverseTable = "channel_messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "channel_id"
)

// Columns holds all SQL columns for channel fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldTopic,
	FieldRefID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// RefIDValidator is a validator for the "ref_id" field. It is called by the builders before save.
	RefIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// NameValidator is a validator for the "name" field enum values. It is called by the builders before save.
func NameValidator(n enum.Channel) error {
	switch n.String() {
	case "PING", "ESTIMATE", "JOB", "JOB_CHAT", "TASK", "TRAINING_VIDEO", "PARTNER", "PARTNER_USER":
		return nil
	default:
		return fmt.Errorf("channel: invalid enum value for name field: %q", n)
	}
}

// TopicValidator is a validator for the "topic" field enum values. It is called by the builders before save.
func TopicValidator(t enum.Topic) error {
	switch t.String() {
	case "BROADCAST", "CHANGE", "STATUS_CHANGE", "PROGRESS", "FILE_UPLOAD", "ASSIGNED", "CREATED", "UPDATED", "NEW_MESSAGE":
		return nil
	default:
		return fmt.Errorf("channel: invalid enum value for topic field: %q", t)
	}
}

// OrderOption defines the ordering options for the Channel queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTopic orders the results by the topic field.
func ByTopic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTopic, opts...).ToFunc()
}

// ByRefID orders the results by the ref_id field.
func ByRefID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefID, opts...).ToFunc()
}

// BySubscriptionsCount orders the results by subscriptions count.
func BySubscriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscriptionsStep(), opts...)
	}
}

// BySubscriptions orders the results by subscriptions terms.
func BySubscriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMessagesCount orders the results by messages count.
func ByMessagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMessagesStep(), opts...)
	}
}

// ByMessages orders the results by messages terms.
func ByMessages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMessagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubscriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubscriptionsTable, SubscriptionsColumn),
	)
}
func newMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
	)
}

var (
	// enum.Channel must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enum.Channel)(nil)
	// enum.Channel must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enum.Channel)(nil)
)

var (
	// enum.Topic must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enum.Topic)(nil)
	// enum.Topic must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enum.Topic)(nil)
)
