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

// PartnerContact holds the schema definition for the PartnerContact entity.
type PartnerContact struct {
	ent.Schema
}

func (PartnerContact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the PartnerContact.
func (PartnerContact) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("partner_id"),
		field.
			String("user_id"),
		field.
			Enum("role").
			GoType(enum.PartnerContactRole("")).
			Annotations(entgql.Type("PartnerContactRole")),
		field.
			Enum("type").
			GoType(enum.PartnerContact("")).
			Annotations(entgql.Type("PartnerContactType")),
		field.
			String("title").
			MaxLen(50).
			Optional().
			Nillable(),
		field.
			String("description").
			MaxLen(250).
			Optional().
			Nillable(),
		field.
			String("invoicing_email").
			Optional().
			Nillable(),
	}
}

// Edges of the PartnerContact.
func (PartnerContact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("user", User.Type).
			Field("user_id").
			Required().
			Unique(),
		edge.
			To("partner", Partner.Type).
			Field("partner_id").
			Required().
			Unique(),
		edge.
			To("sessions", UserSession.Type).
			StorageKey(edge.Column("partner_contact_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.SetNull}),
	}
}

func (PartnerContact) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
	}
}
