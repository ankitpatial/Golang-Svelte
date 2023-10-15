package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// UserSession holds the schema definition for the UserSession entity.
type UserSession struct {
	ent.Schema
}

func (UserSession) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the UserSession.
func (UserSession) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("expires_at").
			Default(time.Now),
		field.
			String("ip").
			MaxLen(50),
	}
}

// Edges of the UserSession.
func (UserSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("sockets", UserSessionSocket.Type).
			StorageKey(edge.Column("sessions_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			From("user", User.Type).
			Ref("sessions").
			Unique().
			Required(),
		edge.
			From("partner", Partner.Type).
			Ref("sessions").
			Unique(),
		edge.
			From("partner_contact", PartnerContact.Type).
			Ref("sessions").
			Unique(),
	}
}

func (UserSession) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("expires_at"),
	}
}
