package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotifySetting holds the schema definition for the NotifySetting entity.
type NotifySetting struct {
	ent.Schema
}

func (NotifySetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the NotifySetting.
func (NotifySetting) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("topic_id").
			MaxLen(36),
		field.
			Bool("receive_email").
			Optional().
			Default(false),
		field.
			Bool("receive_sms").
			Optional().
			Default(false),
	}
}

// Edges of the NotifySetting.
func (NotifySetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", User.Type).
			Ref("notify").
			Unique().
			Required(),
	}
}

func (NotifySetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic_id"),
	}
}
