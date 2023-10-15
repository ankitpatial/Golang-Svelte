package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/pkg/enum"
	"time"
)

// InstallationJob holds the schema definition for the InstallationJob entity.
type InstallationJob struct {
	ent.Schema
}

func (InstallationJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the InstallationJob.
func (InstallationJob) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("type").
			GoType(enum.InstallationType("")).
			Annotations(entgql.Type("InstallationType")),
		field.
			String("owner_name").
			MaxLen(60),
		field.
			String("owner_phone").
			MaxLen(15),
		field.
			String("owner_email").
			Optional(),
		field.
			String("owner_address").
			MaxLen(250),
		field.
			Float("owner_address_lat").
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,8)"}).
			Optional().
			Nillable(),
		field.
			Float("owner_address_lng").
			SchemaType(map[string]string{dialect.MySQL: "decimal(11,8)"}).
			Optional().
			Nillable(),
		field.
			String("special_note").
			MaxLen(500).
			Optional().
			Nillable(),
		field.
			String("pkg").
			MaxLen(100),
		field.
			String("pkg_description").
			Optional().
			MaxLen(500),
		field.
			JSON("pkg_features", []string{}).
			Optional().
			Default([]string{}),
		field.
			Float("price").
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.
			Enum("approval").
			GoType(enum.Approval("")).
			Default(enum.ApprovalPending.String()).
			Annotations(entgql.Type("Approval")),
		field.
			Time("approval_at").
			Default(time.Now().UTC()).
			Annotations(entgql.OrderField("APPROVAL_AT")),
		field.
			String("deny_reason").
			Optional(),
		field.
			Enum("status").
			GoType(enum.InstallationStatus("")).
			Default(enum.InstallationStatusPending.String()).
			Annotations(entgql.Type("InstallationStatus")),
		field.
			Time("status_at").
			Default(time.Now().UTC()),
	}
}

// Edges of the InstallationJob.
func (InstallationJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("items", InstallationJobItem.Type).
			StorageKey(edge.Column("job_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("progress_history", InstallationJobProgress.Type).
			StorageKey(edge.Column("job_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			From("requesting_partner", Partner.Type).
			Ref("requested_installation_jobs").
			Unique(),
		edge.
			From("assigned_partner", Partner.Type).
			Ref("assigned_installation_jobs").
			Unique(),
		edge.
			From("creator", User.Type).
			Ref("created_installation_jobs").
			Unique(),
		edge.
			From("sales_rep", User.Type).
			Ref("installation_leads").
			Unique(),
	}
}

func (InstallationJob) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_name"),
		index.Fields("owner_email"),
		index.Fields("owner_phone"),
		index.Fields("owner_address"),
		index.Fields("approval", "approval_at"),
		index.Fields("status", "status_at"),
	}
}
