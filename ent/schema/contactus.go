package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ContactUs holds the schema definition for the ContactUs entity.
type ContactUs struct {
	ent.Schema
}

func (ContactUs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the ContactUs.
func (ContactUs) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("reason").
			MaxLen(1000),
	}
}

// Edges of the ContactUs.
func (ContactUs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("partner", Partner.Type).
			Ref("contact_us_requests").
			Unique(),
		edge.
			From("creator", User.Type).
			Ref("contact_us_requests").
			Unique(),
	}
}
