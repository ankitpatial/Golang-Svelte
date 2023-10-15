package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/pkg/enum"
)

// Channel holds the schema definition for the ChatChannel entity.
type Channel struct {
	ent.Schema
}

func (Channel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the ChatChannel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("name").
			GoType(enum.Channel("")).
			Immutable().
			Annotations(entgql.Type("Channel")),
		field.
			Enum("topic").
			GoType(enum.Topic("")).
			Annotations(entgql.Type("Topic")),
		field.
			String("ref_id").
			MaxLen(36).
			Immutable().
			Optional(),
	}
}

// Edges of the ChatChannel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("subscriptions", ChannelSub.Type).
			StorageKey(edge.Column("channel_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("messages", ChannelMessage.Type).
			StorageKey(edge.Column("channel_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
func (c Channel) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("topic"),
		index.Fields("ref_id"),
	}
}
