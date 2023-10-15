package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ApiAccess holds the schema definition for the ApiAccess entity.
type ApiAccess struct {
	ent.Schema
}

func (ApiAccess) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the ApiAccess.
func (ApiAccess) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("username").
			Default(""),
		field.String("password").
			Default(""),
		field.String("key").
			Default(""),
		field.String("secret").
			Default(""),
		field.String("access_token").
			MaxLen(800).
			Optional(),
		field.String("refresh_token").
			MaxLen(800).
			Optional(),
		field.Time("expires_at").
			Optional(),
	}
}

// Edges of the ApiAccess.
func (ApiAccess) Edges() []ent.Edge {
	return nil
}

func (ApiAccess) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("url"),
		index.Fields("username"),
	}
}
