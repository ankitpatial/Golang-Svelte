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

// SurveyProgress holds the schema definition for the JobProgressHistory entity.
type SurveyProgress struct {
	ent.Schema
}

func (SurveyProgress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Annotations of the JobProgressHistory.
func (SurveyProgress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey_progress"},
	}
}

// Fields of the JobProgressHistory.
func (SurveyProgress) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("status").
			GoType(enum.SurveyProgress("")).
			Immutable().
			Annotations(entgql.Type("SurveyProgress")),
		field.
			Bool("complete").
			Default(false).
			Immutable().
			Optional(),

		field.
			String("note").
			Optional().
			Immutable(),
	}
}

// Edges of the JobProgressHistory.
func (SurveyProgress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("survey", Survey.Type).
			Ref("progress_history").
			Unique(),
		edge.
			From("creator", User.Type).
			Ref("survey_progress").
			Unique(),
		edge.
			From("creator_api", ApiUser.Type).
			Ref("survey_progress").
			Unique(),
	}
}
