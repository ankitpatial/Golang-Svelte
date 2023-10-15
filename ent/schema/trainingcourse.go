package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TrainingCourse holds the schema definition for the TrainingCourse entity.
type TrainingCourse struct {
	ent.Schema
}

func (TrainingCourse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the TrainingCourse.
func (TrainingCourse) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			MaxLen(100).
			Annotations(
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
				entgql.OrderField("NAME"),
			),
	}
}

// Edges of the TrainingCourse.
func (TrainingCourse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("training_videos", TrainingVideo.Type).
			StorageKey(edge.Column("course_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Restrict}),

		edge.
			From("creator", User.Type).
			Ref("created_training_courses").
			Unique(),
	}
}
