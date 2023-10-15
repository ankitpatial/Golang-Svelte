package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/pkg/enum"
)

// ChannelSub holds the schema definition for the ChannelSub entity.
type ChannelSub struct {
	ent.Schema
}

func (ChannelSub) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the ChatChannelUser.
func (ChannelSub) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("role").
			GoType(enum.Role("")).
			Optional().
			Nillable().
			Annotations(entgql.Type("Role")),
	}
}

// Edges of the ChatChannelUser.
func (ChannelSub) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("channel", Channel.Type).
			Ref("subscriptions").
			Unique().
			Required(),
		edge.
			From("user", User.Type).
			Ref("chat_channels").
			Unique(),
		edge.
			From("partner", Partner.Type).
			Ref("channels").
			Unique(),
	}
}

func (ChannelSub) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role"),
	}
}
