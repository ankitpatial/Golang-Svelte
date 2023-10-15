package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ChannelMessageRead holds the schema definition for the ChatMessage entity.
type ChannelMessageRead struct {
	ent.Schema
}

// Fields of the ChatMessage.
func (ChannelMessageRead) Fields() []ent.Field {
	return []ent.Field{
		fieldID,
		fieldCreated,
		field.
			Bool("read").
			Default(true),
	}
}

// Edges of the ChatMessage.
func (ChannelMessageRead) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("channel_message", ChannelMessage.Type).
			Ref("reads").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("channel_message_read").
			Unique().
			Required(),
	}
}

func (ChannelMessageRead) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at").Annotations(entsql.Desc()),
	}
}
