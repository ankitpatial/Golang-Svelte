package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EstimateActivity holds the schema definition for the EstimateUpdate entity.
// Response we will get from provider after submitting estimate request.
type EstimateActivity struct {
	ent.Schema
}

func (EstimateActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the EstimateUpdate.
func (EstimateActivity) Fields() []ent.Field {
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
func (EstimateActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("estimate", Estimate.Type).Ref("activities").Unique(),
		edge.From("creator", User.Type).Ref("estimate_activities").Unique(),
		edge.From("creator_api", ApiUser.Type).Ref("estimate_activities").Unique(),
	}
}
