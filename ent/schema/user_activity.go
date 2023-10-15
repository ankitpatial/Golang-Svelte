package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserActivity struct {
	ent.Schema
}

func (UserActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the EstimateUpdate.
func (UserActivity) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("description"),
		field.
			JSON("raw", map[string]interface{}{}).
			Optional().
			Comment("any related data"),
	}
}

// Edges of the EstimateUpdate.
func (UserActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("activities").Unique(),
		edge.From("creator", User.Type).Ref("user_activities").Unique(),
		edge.From("creator_api", ApiUser.Type).Ref("user_activities").Unique(),
	}
}
