package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserAuth holds the schema definition for the UserAuth entity.
type UserAuth struct {
	ent.Schema
}

// Fields of the UserAuth.
func (UserAuth) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(150).
			NotEmpty().
			Unique().
			Comment("user id from provider side"),

		field.Uint8("provider_id").Default(0),
		fieldCreated,
	}
}

// Edges of the UserAuth.
func (UserAuth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auths").
			Unique().
			Required(),
	}
}
