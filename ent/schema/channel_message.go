package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChannelMessage holds the schema definition for the ChatMessage entity.
type ChannelMessage struct {
	ent.Schema
}

func (ChannelMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the ChatMessage.
func (ChannelMessage) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title").
			MaxLen(50).
			Optional(),
		field.
			String("message").
			MaxLen(500),
		field.
			String("from_name").
			MaxLen(100),
		field.
			String("to_name").
			MaxLen(100).
			Optional(),
		field.
			Bool("private").
			Default(true),
		field.
			Bool("unread").
			Optional().
			Comment("not to be used to save data, will be used in query in ORM"),
	}
}

// Edges of the ChatMessage.
func (ChannelMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("reads", ChannelMessageRead.Type).
			StorageKey(edge.Column("channel_message_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("channel", Channel.Type).
			Ref("messages").
			Unique().
			Required(),
		edge.From("from", User.Type).
			Ref("sent_messages").
			Unique(),
		edge.From("from_api_user", ApiUser.Type).
			Ref("notifications").
			Unique(),
		edge.From("to", User.Type).
			Ref("received_messages").
			Unique(),
	}
}
