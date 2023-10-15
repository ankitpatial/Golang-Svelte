package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PartnerActivity struct {
	ent.Schema
}

func (PartnerActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the EstimateUpdate.
func (PartnerActivity) Fields() []ent.Field {
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
func (PartnerActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("partner", Partner.Type).Ref("activities").Unique(),
		edge.From("creator", User.Type).Ref("partner_activities").Unique(),
		edge.From("creator_api", ApiUser.Type).Ref("partner_activities").Unique(),
	}
}
