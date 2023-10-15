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

type JobProgressInput struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Order    int    `json:"order"`
}

// JobProgressHistory holds the schema definition for the JobProgressHistory entity.
type JobProgressHistory struct {
	ent.Schema
}

func (JobProgressHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Annotations of the JobProgressHistory.
func (JobProgressHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "job_progress_history"},
	}
}

// Fields of the JobProgressHistory.
func (JobProgressHistory) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("status").
			Immutable().
			GoType(enum.JobProgress("")).
			Annotations(entgql.Type("JobProgress")),
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

// Edges of the JobProgressHistory.
func (JobProgressHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("job", Job.Type).
			Ref("progress_history").
			Unique(),

		edge.
			From("creator", User.Type).
			Ref("job_progress_history").
			Unique(),

		edge.From("creator_api_user", ApiUser.Type).
			Ref("job_progress_history").
			Unique(),
	}
}
