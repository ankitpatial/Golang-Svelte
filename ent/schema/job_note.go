package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobNote holds the schema definition for the JobNote entity.
type JobNote struct {
	ent.Schema
}

func (JobNote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the JobNote.
func (JobNote) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("note").
			MaxLen(500),
	}
}

// Edges of the JobNote.
func (JobNote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("job", Job.Type).
			Ref("notes").
			Unique().
			Required(),
		edge.
			From("user", User.Type).
			Ref("job_notes").
			Unique().
			Required(),
		edge.
			From("partner", Partner.Type).
			Ref("job_notes").
			Unique(),
	}
}
