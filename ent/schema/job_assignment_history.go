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

// JobAssignmentHistory history
type JobAssignmentHistory struct {
	ent.Schema
}

func (JobAssignmentHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Annotations of the JobAssignmentHistory.
func (JobAssignmentHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "job_assignment_history"},
	}
}

// Fields of the JobAssignmentHistory.
func (JobAssignmentHistory) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("status").
			Immutable().
			GoType(enum.JobAssignmentStatus("")).
			Default(enum.JobAssignmentStatusAssigned.String()).
			Annotations(entgql.Type("JobAssignmentStatus")),
		field.
			String("Note").
			MaxLen(1000).
			Optional(),
	}
}

// Edges of the JobAssignmentHistory.
func (JobAssignmentHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("job", Job.Type).
			Ref("assignment_history").
			Unique(),
		edge.
			From("partner", Partner.Type).
			Ref("job_assignment_history").
			Unique(),
	}
}
