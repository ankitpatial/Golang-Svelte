package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"roofix/pkg/enum"
)

// InstallationJobProgress holds the schema definition for the InstallationJob entity.
type InstallationJobProgress struct {
	ent.Schema
}

func (InstallationJobProgress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

func (InstallationJobProgress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "installation_job_progress"},
	}
}

// Fields of the InstallationJob.
func (InstallationJobProgress) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("status").
			Immutable().
			GoType(enum.InstallationStatus("")).
			Annotations(entgql.Type("InstallationStatus")),
		field.
			Bool("complete").
			Default(false).
			Optional().
			Immutable(),

		field.
			String("note").
			Optional().
			Immutable(),
	}
}

// Edges of the InstallationJob.
func (InstallationJobProgress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("job", InstallationJob.Type).
			Ref("progress_history").
			Unique(),

		edge.
			From("creator", User.Type).
			Ref("installation_job_status_changer").
			Unique(),
	}
}
