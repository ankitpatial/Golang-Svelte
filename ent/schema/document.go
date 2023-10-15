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

// Document holds the schema definition for the Document entity.
type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the Document.
func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("bucket").
			MaxLen(50),
		field.
			String("key"),
		field.
			Enum("folder").
			GoType(enum.DocFolder("")).
			Annotations(entgql.Type("DocumentFolder")),
		field.
			String("dir").
			Optional().
			Nillable().
			MaxLen(36),
		field.
			Enum("section").
			GoType(enum.DocSection("")).
			Annotations(entgql.Type("DocumentSection")),
		field.
			String("name").
			MaxLen(150).
			Comment("name on storage"),
		field.
			String("filename").
			MaxLen(150).
			Comment("actual filename"),
		field.
			String("content_type").
			MaxLen(50).
			Optional().
			Nillable(),
		field.
			Int64("content_size"),
		field.
			Bool("ready").
			Default(false).
			Comment("ready to access. True means uploaded and saved ok to storage"),
		field.
			String("creator_id").
			MaxLen(36),
		field.
			String("updater_id").
			MaxLen(36).
			Optional().
			Nillable(),
	}
}

// Edges of the Partner.
func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("training_video", TrainingVideo.Type).
			StorageKey(edge.Column("video_id")).
			Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.SetNull}),
		edge.
			To("training_video_poster", TrainingVideo.Type).
			StorageKey(edge.Column("poster_id")).
			Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.SetNull}),
		edge.
			To("products_image", Product.Type).
			StorageKey(edge.Column("image_id")).
			Unique(),
		edge.
			To("installation_job_item_image", InstallationJobItem.Type).
			StorageKey(edge.Column("image_id")).
			Unique(),
		edge.To("estimate_pdf", Estimate.Type).StorageKey(edge.Column("pdf_doc_id")).Unique(),
		edge.To("job_estimate_pdf", Job.Type).StorageKey(edge.Column("estimate_pdf_id")).Unique(),
	}
}

// Indexes of the Document.
func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("bucket", "key"),
		index.Fields("folder", "dir"),
		index.Fields("section"),
		index.Fields("name"),
		index.Fields("ready"),
	}
}
