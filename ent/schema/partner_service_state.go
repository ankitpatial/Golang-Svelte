package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PartnerServiceState holds the schema definition for the PartnerServiceCity entity.
type PartnerServiceState struct {
	ent.Schema
}

func (PartnerServiceState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

func (PartnerServiceState) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Fields of the PartnerServiceCity.
func (PartnerServiceState) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("country").
			MaxLen(20),
		field.
			String("state").
			MaxLen(50),
		field.
			String("license_no").
			MaxLen(50).
			Optional().
			Nillable(),
		field.
			Time("license_exp_date").
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
func (PartnerServiceState) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("partner", Partner.Type).
			Ref("service_states").
			Unique().
			Required(),
	}
}

func (PartnerServiceState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country"),
		index.Fields("state"),
	}
}
