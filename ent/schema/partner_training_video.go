package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PartnerTrainingVideo holds the schema definition for the PartnerTrainingVideo entity.
type PartnerTrainingVideo struct {
	ent.Schema
}

func (PartnerTrainingVideo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the PartnerTrainingVideo.
func (PartnerTrainingVideo) Fields() []ent.Field {
	return []ent.Field{
		field.
			Bool("enabled").
			Default(false),
	}
}

// Edges of the PartnerTrainingVideo.
func (PartnerTrainingVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("video", TrainingVideo.Type).
			Ref("training_videos").
			Unique().
			Required(),
		edge.
			From("partner", Partner.Type).
			Ref("training_videos").
			Unique().
			Required(),
	}
}
