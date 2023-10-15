package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PartnerServiceCity holds the schema definition for the PartnerServiceCity entity.
type PartnerServiceCity struct {
	ent.Schema
}

func (PartnerServiceCity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

func (PartnerServiceCity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Fields of the PartnerServiceCity.
func (PartnerServiceCity) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("postal_id").
			MaxLen(36),
		field.
			Bool("active").
			Default(false),
		field.
			String("name").
			MaxLen(50),
		field.
			Uint("naics_code"),
		field.
			String("license_no").
			MaxLen(20).
			Optional().
			Nillable(),
		field.
			String("proof_doc_id").
			MaxLen(36).
			Optional().
			Nillable(),
	}
}

// Edges of the PartnerServiceCity.
func (PartnerServiceCity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("partner", Partner.Type).
			Ref("service_cities").
			Unique().
			Required(),
	}
}

func (PartnerServiceCity) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("postal_id"),
	}
}
