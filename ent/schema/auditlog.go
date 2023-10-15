package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AuditLog holds the schema definition for the AuditLog entity.
type AuditLog struct {
	ent.Schema
}

func (AuditLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the AuditLog.
func (AuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("action").
			MaxLen(50),
		field.
			String("description").
			Optional().
			MaxLen(500).
			Annotations(
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.
			String("ip").
			Optional().
			MaxLen(50),
	}
}

// Edges of the AuditLog.
func (AuditLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Unique().Ref("audit_logs"),
		edge.From("api_user", ApiUser.Type).Unique().Ref("audit_logs"),
	}
}

func (AuditLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ip"),
		index.Fields("action"),
		index.Fields("description"),
	}
}
