package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobActivity holds the schema definition for the EstimateUpdate entity.
// Response we will get from provider after submitting estimate request.
type JobActivity struct {
	ent.Schema
}

func (JobActivity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the EstimateUpdate.
func (JobActivity) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("description"),
		field.
			JSON("raw", map[string]interface{}{}).
			Optional().
			Comment("any related data"),
	}
}

// Edges of the EstimateUpdate.
func (JobActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("job", Job.Type).Ref("activities").Unique(),
		edge.From("creator", User.Type).Ref("job_activities").Unique(),
		edge.From("creator_api", ApiUser.Type).Ref("job_activities").Unique(),
	}
}
