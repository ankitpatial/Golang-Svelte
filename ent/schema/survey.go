package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/pkg/enum"
)

// Survey holds the schema definition for the Survey entity.
type Survey struct {
	ent.Schema
}

// Mixin of the Survey.
func (Survey) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the Survey.
func (Survey) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("date").
			MaxLen(10).
			Immutable(),
		field.
			String("slot").
			MaxLen(10).
			Immutable(),
		field.
			String("slot_id").
			MaxLen(36).
			Immutable(),
		field.
			Time("from").
			Immutable(),
		field.
			Time("to").
			Immutable(),
		field.
			Time("until").
			Optional().
			Nillable().
			Comment("will be used with status REQUESTING and will denote expiry of request"),
		field.
			String("name").
			MaxLen(100).
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.
			String("address").
			MaxLen(200).
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("ADDRESS"),
			),
		field.
			String("phone").
			MaxLen(20).
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("PHONE"),
			),
		field.
			String("notes").
			MaxLen(500).
			Optional().
			Nillable(),

		field.
			Enum("status").
			GoType(enum.SurveyStatus("")).
			Annotations(entgql.Type("SurveyStatus")),
		field.
			Enum("progress").
			GoType(enum.SurveyProgress("")).
			Optional().
			Nillable().
			Annotations(entgql.Type("SurveyProgress")),
		field.
			Time("progress_at").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("PROGRESS_AT")),
		field.
			Time("progress_flag_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Survey.
func (Survey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("progress_history", SurveyProgress.Type).
			StorageKey(edge.Column("survey_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			From("created_by", User.Type).
			Ref("surveys").
			Unique(),
		edge.
			From("partner", Partner.Type).
			Ref("surveys").
			Unique(),
	}
}

// Indexes of the Survey.
func (Survey) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("date"),
		index.Fields("slot_id"),
		index.Fields("from"),
		index.Fields("to"),
		index.Fields("name"),
		index.Fields("address"),
		index.Fields("phone"),
	}
}
