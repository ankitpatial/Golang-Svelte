package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// UserSessionSocket holds the schema definition for the UserSession entity.
type UserSessionSocket struct {
	ent.Schema
}

func (UserSessionSocket) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Edges of the UserSession.
func (UserSessionSocket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("session", UserSession.Type).
			Ref("sockets").
			Unique().
			Required(),
	}
}
