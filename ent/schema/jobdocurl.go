package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"roofix/pkg/enum"
)

// JobDocURL holds the schema definition for the JobDocURL entity.
type JobDocURL struct {
	ent.Schema
}

func (JobDocURL) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the JobDocURL.
func (JobDocURL) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enum.JobDocUrlType("")).Annotations(entgql.Type("JobDocUrlType")),
		field.String("url"),
	}
}

// Edges of the JobDocURL.
func (JobDocURL) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("job", Job.Type).Ref("doc_urls").Unique(),
		edge.From("creator", User.Type).Ref("job_doc_urls").Unique(),
	}
}
