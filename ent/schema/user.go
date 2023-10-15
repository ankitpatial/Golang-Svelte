package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/ent/partnercontact"
	"roofix/pkg/enum"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndAllStamps{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").MaxLen(36).Optional(),
		field.
			String("email").
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.String("phone").Optional().MaxLen(20),
		field.String("pwd").Sensitive().NotEmpty().MaxLen(150),
		field.
			String("first_name").
			MaxLen(50).
			Optional().
			Annotations(
				entgql.OrderField("FIRST_NAME"),
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.
			String("last_name").
			MaxLen(50).
			Optional().
			Annotations(
				entgql.OrderField("LAST_NAME"),
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.Bool("email_verified").Default(false),
		field.Bool("phone_verified").Default(false),
		field.String("picture").MaxLen(250).Optional().Nillable(),
		field.
			Enum("status").
			GoType(enum.AccountStatus("")).
			Default(enum.AccountStatusPending.String()).
			Annotations(entgql.Type("AccountStatus")),
		field.
			Enum("role").
			GoType(enum.Role("")).
			Default(enum.RoleSiteUser.String()).
			Annotations(entgql.Type("Role")),
		field.String("note").MaxLen(500).Optional(),
		field.Uint8("wrong_attempts").Optional().Default(0),
		field.Time("wrong_attempt_at").Optional().Nillable(),
		field.Time("locked_until").Optional().Nillable(),
		field.String("location").Optional(),
		field.Bool("accepted_general_terms").Optional().Default(false),
		field.Bool("accepted_terms_n_privacy").Optional().Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("auths", UserAuth.Type).
			Unique().
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("sessions", UserSession.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("audit_logs", AuditLog.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("notify", NotifySetting.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("created_jobs", Job.Type).StorageKey(edge.Column("creator_id")),
		edge.To("created_estimates", Estimate.Type).StorageKey(edge.Column("creator_id")),
		edge.To("sales_rep_estimates", Estimate.Type).StorageKey(edge.Column("sales_rep_id")),
		edge.To("sales", Job.Type).StorageKey(edge.Column("sales_rep_id")),
		edge.To("created_installation_jobs", InstallationJob.Type).StorageKey(edge.Column("creator_id")),
		edge.To("installation_leads", InstallationJob.Type).StorageKey(edge.Column("sales_rep_id")),

		edge.To("estimate_activities", EstimateActivity.Type).StorageKey(edge.Column("creator_id")),
		edge.To("job_activities", JobActivity.Type).StorageKey(edge.Column("creator_id")),
		edge.To("partner_activities", PartnerActivity.Type).StorageKey(edge.Column("creator_id")),
		edge.To("user_activities", UserActivity.Type).StorageKey(edge.Column("creator_id")),
		edge.
			To("activities", UserActivity.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("job_progress_history", JobProgressHistory.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("installation_job_status_changer", InstallationJobProgress.Type).
			StorageKey(edge.Column("creator_id")),
		edge.
			To("contact_us_requests", ContactUs.Type).
			StorageKey(edge.Column("creator_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			From("partner", Partner.Type).
			Ref("contacts").
			Through(partnercontact.Table, PartnerContact.Type),
		edge.To("surveys", Survey.Type).StorageKey(edge.Column("user_id")),
		edge.To("survey_progress", SurveyProgress.Type).StorageKey(edge.Column("creator")),
		edge.To("created_training_videos", TrainingVideo.Type).StorageKey(edge.Column("creator_id")),
		edge.To("created_training_courses", TrainingCourse.Type).StorageKey(edge.Column("creator_id")),
		edge.
			To("job_notes", JobNote.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("chat_channels", ChannelSub.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("sent_messages", ChannelMessage.Type).
			StorageKey(edge.Column("from_user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("received_messages", ChannelMessage.Type).
			StorageKey(edge.Column("to_user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("channel_message_read", ChannelMessageRead.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("products_created", Product.Type).StorageKey(edge.Column("creator_id")),
		edge.To("product_pkg_created", ProductPackage.Type).StorageKey(edge.Column("creator_id")),
		edge.To("job_doc_urls", JobDocURL.Type).StorageKey(edge.Column("creator_id")),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("external_id"),
		index.Fields("first_name"),
		index.Fields("last_name"),
		index.Fields("phone"),
		index.Fields("role"),
		index.Fields("status"),
	}
}
