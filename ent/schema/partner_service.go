package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PartnerService holds the schema definition for the PartnerServiceCity entity.
type PartnerService struct {
	ent.Schema
}

func (PartnerService) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the PartnerService.
func (PartnerService) Fields() []ent.Field {
	return []ent.Field{
		field.
			Uint8("service_id"),
		field.
			Bool("active"),
	}
}

// Edges of the PartnerService.
func (PartnerService) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("partner", Partner.Type).
			Ref("services").
			Unique().
			Required(),
	}
}

// Indexes of the PartnerService
func (PartnerService) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("service_id"),
		index.Fields("active"),
	}
}
