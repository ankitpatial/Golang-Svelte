package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ApiUser holds the schema definition for the ApiUser entity.
type ApiUser struct {
	ent.Schema
}

func (ApiUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the ApiUser.
func (ApiUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("username").
			Unique().
			MaxLen(50).
			Annotations(
				entgql.OrderField("USER_NAME"),
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.
			String("pwd_hash").
			MaxLen(150),
		field.
			Bool("active").
			Default(true),
		field.
			String("cb_api_url").
			Optional().
			Nillable(),
		field.
			Enum("cb_api_auth").
			Values("NONE", "BASIC", "TOKEN", "OAUTH").
			Optional().
			Default("NONE"),
		field.
			String("cb_api_user").
			Optional(),
		field.
			String("cb_api_pwd").
			Optional(),
		field.
			String("cb_api_token").
			Optional(),
		field.
			JSON("cb_api_endpoints", map[string]string{}).
			Optional(),
	}
}

// Edges of the ApiUser.
func (ApiUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", ApiUserToken.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("audit_logs", AuditLog.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("created_estimates", Estimate.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("created_jobs", Job.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("created_partners", Partner.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("survey_progress", SurveyProgress.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("estimate_activities", EstimateActivity.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("user_activities", UserActivity.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("partner_activities", PartnerActivity.Type).StorageKey(edge.Column("api_user_id")),
		edge.To("job_activities", JobActivity.Type).StorageKey(edge.Column("api_user_id")),
		edge.
			To("notifications", ChannelMessage.Type).
			StorageKey(edge.Column("from_api_user_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("job_progress_history", JobProgressHistory.Type).StorageKey(edge.Column("api_user_id")),
	}
}

func (ApiUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("active"),
	}
}
