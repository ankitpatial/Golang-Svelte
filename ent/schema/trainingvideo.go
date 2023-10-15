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

// TrainingVideo holds the schema definition for the TrainingVideo entity.
type TrainingVideo struct {
	ent.Schema
}

func (TrainingVideo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the TrainingVideo.
func (TrainingVideo) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("kind").
			GoType(enum.TrainingType("")).
			Annotations(entgql.Type("TrainingType")),
		field.
			String("title").
			MaxLen(100).
			Annotations(
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
				entgql.OrderField("TITLE"),
			),
		field.
			String("description").
			MaxLen(500),
	}
}

// Edges of the TrainingVideo.
func (TrainingVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("training_videos", PartnerTrainingVideo.Type).
			StorageKey(edge.Column("video_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.
			From("creator", User.Type).
			Ref("created_training_videos").
			Unique(),

		edge.
			From("course", TrainingCourse.Type).
			Ref("training_videos").
			Unique(),

		edge.
			From("poster", Document.Type).
			Ref("training_video_poster").
			Unique(),

		edge.
			From("video", Document.Type).
			Ref("training_video").
			Unique(),
	}
}

// Indexes of the Survey.
func (TrainingVideo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("kind"),
		index.Fields("title"),
		index.Fields("description"),
	}
}
