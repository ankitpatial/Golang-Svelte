package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/ent/partnercontact"
	"roofix/pkg/enum"
	"roofix/pkg/model"
)

// Partner holds the schema definition for the Partner entity.
type Partner struct {
	ent.Schema
}

func (Partner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the Partner.
func (Partner) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").MaxLen(36).Optional(),
		field.
			String("creator_id").
			Optional().
			MaxLen(36),
		field.
			Enum("type").
			GoType(enum.Partner("")).
			Annotations(entgql.Type("PartnerType")),
		field.
			String("name").
			MaxLen(100).
			Annotations(
				entgql.OrderField("NAME"),
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.
			String("address").
			Optional(),
		field.
			String("website").
			Optional(),
		field.
			String("phone").
			MaxLen(20).
			Optional(),
		field.
			Float("latitude").
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,8)"}).
			Optional(),
		field.
			Float("longitude").
			SchemaType(map[string]string{dialect.MySQL: "decimal(11,8)"}).
			Optional(),
		field.
			Bool("is_nation_wide").
			Default(false),
		field.
			Uint16("crew_count").
			Default(0),
		field.
			Int("years_in_business").
			Optional().
			Nillable(),
		field.
			Uint16("job_capacity").
			Comment("weekly job capacity").
			Default(0),
		field.
			String("asphalt_lead_t").
			Optional().
			Comment("times to completely complete a Asphalt project"),
		field.
			String("metal_lead_t").
			Optional().
			Comment("times to completely complete a Metal project"),
		field.
			String("tile_lead_t").
			Optional().
			Comment("times to completely complete a Tile project"),
		field.
			Uint8("setup_steps_completed").
			Default(1),
		field.
			Int("sales_volume").
			Optional().
			Nillable().
			Comment("solar partner's monthly sales volume"),
		field.
			Int("down_payment").
			Optional().
			Nillable().
			Comment("solar partner's down payment %"),
		field.
			Int("pif").
			Optional().
			Nillable().
			Comment("solar partner's PIF date in days"),
		field.
			Bool("install_in_house").
			Optional().
			Nillable().
			Comment("solar partner's: Do you install your solar projects in house?"),
		field.
			Enum("status").
			GoType(enum.PartnerStatus("")).
			Default(enum.PartnerStatusInActive.String()).
			Annotations(entgql.Type("PartnerStatus")),
		field.
			Enum("epc_status").
			GoType(enum.EPCStatus("")).
			Annotations(entgql.Type("EPCStatus")).
			Optional(),
		field.
			JSON("mobile_app_settings", model.MobileAppSettings{}).
			Optional(),
	}
}

// Edges of the Partner.
func (Partner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("requested_estimates", Estimate.Type).
			StorageKey(edge.Column("partner_id")).
			Comment("company that requested estimate"),
		edge.
			To("estimate_home_owners", HomeOwner.Type).
			StorageKey(edge.Column("partner_id")).
			Comment("requested estimate home owner"),
		edge.To("roofing_jobs", Job.Type).StorageKey(edge.Column("roofing_partner_id")),
		edge.To("integration_jobs", Job.Type).StorageKey(edge.Column("integration_partner_id")),
		edge.To("epc_jobs", Job.Type).StorageKey(edge.Column("epc_partner_id")),
		edge.To("job_requests", Job.Type).StorageKey(edge.Column("requester_id")),
		edge.
			To("activities", PartnerActivity.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("job_assignment_history", JobAssignmentHistory.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("requested_installation_jobs", InstallationJob.Type).
			StorageKey(edge.Column("requesting_partner_id")),
		edge.
			To("assigned_installation_jobs", InstallationJob.Type).
			StorageKey(edge.Column("assigned_partner_id")),
		edge.
			To("contacts", User.Type).
			Through(partnercontact.Table, PartnerContact.Type),
		edge.
			To("creator", User.Type).
			Field("creator_id").
			Unique(),
		edge.
			From("creator_api", ApiUser.Type).
			Ref("created_partners").
			Unique(),
		edge.
			To("services", PartnerService.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("service_states", PartnerServiceState.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("service_cities", PartnerServiceCity.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("training_videos", PartnerTrainingVideo.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("contact_us_requests", ContactUs.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("surveys", Survey.Type).
			StorageKey(edge.Column("partner_id")),
		edge.
			To("sessions", UserSession.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.SetNull}),
		edge.
			To("job_notes", JobNote.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			To("channels", ChannelSub.Type).
			StorageKey(edge.Column("partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			From("finance_options", OptionList.Type).
			Ref("finance_options"),
		edge.
			From("epc_options", OptionList.Type).
			Ref("epc_options"),
	}
}

func (Partner) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("external_id"),
		index.Fields("name"),
		index.Fields("type"),
		index.Fields("status"),
	}
}
