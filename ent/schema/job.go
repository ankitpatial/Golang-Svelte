/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package schema

import (
	"roofix/pkg/enum"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

func (Job) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationUpdate()),
	}
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("progress").
			GoType(enum.JobProgress("")).
			Optional().
			Nillable().
			Annotations(entgql.Type("JobProgress")),
		field.
			Time("progress_at").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("PROGRESS_AT")),
		field.
			Time("progress_flag_at").
			Optional().
			Nillable(),
		field.
			Uint8("region_id").
			Default(0).Optional(),
		field.
			String("company_ref_id").
			MaxLen(36).
			Optional().
			Nillable(),
		field.
			String("company_name").
			MaxLen(50).
			Optional().
			Annotations(
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			).Optional(),
		field.
			Float("price").
			Default(0).
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.
			Float("work_order_price").
			Default(0).
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.
			Float("contract_price").
			Optional().
			Default(0).
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.
			Float("change_order_price").
			Optional().
			Default(0).
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.
			Text("note").
			Optional(),
		field.
			String("shingle_color").
			Optional().
			Nillable(),
		field.
			Bool("permit_required").
			Optional().
			Nillable(),
		field.
			Bool("inspection_required").
			Optional().
			Nillable().
			Comment("final inspection required flag"),
		field.
			Time("inspection_date").
			Optional().
			Nillable().
			Comment("final inspection date"),
		field.
			Time("progress_inspection_date").
			Optional().
			Nillable(),
		field.
			Time("install_date").
			Optional().
			Nillable(),
		field.
			Time("completion_date").
			Optional().
			Nillable(),
		field.
			Time("material_delivery_date").
			Optional().
			Nillable(),
		field.
			Bool("agree").
			Optional().
			Nillable().
			Comment("solar approved job to production and agrees to the terms and conditions"),
		field.
			Time("agree_at").
			Optional().
			Nillable(),
		field.
			String("po_number").
			Optional(),
		field.
			Time("roofing_partner_assigned_at").
			Optional().
			Nillable(),
		field.
			Time("roofing_partner_flag_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payments", Payment.Type).StorageKey(edge.Column("job_id")),
		edge.
			To("activities", JobActivity.Type).
			StorageKey(edge.Column("job_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("assignment_history", JobAssignmentHistory.Type).
			StorageKey(edge.Column("job_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("progress_history", JobProgressHistory.Type).
			StorageKey(edge.Column("job_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("notes", JobNote.Type).
			StorageKey(edge.Column("job_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("doc_urls", JobDocURL.Type).
			StorageKey(edge.Column("job_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("estimate", Estimate.Type).Ref("job_info").Unique(),
		edge.From("creator_api", ApiUser.Type).Ref("created_jobs").Unique(),
		edge.From("creator", User.Type).Ref("created_jobs").Unique(),
		edge.From("home_owner", HomeOwner.Type).Ref("jobs").Unique(),
		edge.From("sales_rep", User.Type).Ref("sales").Unique(),
		edge.From("requester", Partner.Type).Ref("job_requests").Unique(),
		edge.
			From("roofing_partner", Partner.Type).
			Ref("roofing_jobs").
			Unique().
			Comment("a.k.a contractor"),
		edge.
			From("integration_partner", Partner.Type).
			Ref("integration_jobs").
			Unique().
			Comment("job contractor"),
		edge.
			From("epc_partner", Partner.Type).
			Ref("epc_jobs").
			Unique().
			Comment("job contractor"),
		edge.
			From("epc", OptionList.Type).
			Ref("epc").
			Unique(),
		edge.From("estimate_pdf", Document.Type).Ref("job_estimate_pdf").Unique(),
	}
}

func (receiver Job) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("progress"),
		index.Fields("progress_at"),
		index.Fields("progress_flag_at"),
		index.Fields("region_id"),
	}
}
