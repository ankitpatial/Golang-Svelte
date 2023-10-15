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
	"roofix/pkg/model"
)

// Estimate holds the schema definition for the Estimate entity.
type Estimate struct {
	ent.Schema
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (Estimate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the Estimate.
func (Estimate) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("region_id").Default(0),
		field.Enum("status").GoType(enum.EstimateStatus("")).Annotations(entgql.Type("EstimateStatus")),
		field.String("current_material").Optional().MaxLen(50),
		field.String("new_roofing_material").Optional().MaxLen(50),
		field.Bool("lowSlope").Default(false),
		field.String("current_material_low_slope").Optional().MaxLen(50),
		field.String("new_roofing_material_low_slope").Optional().MaxLen(50),
		field.Bool("redeck").Default(false),
		field.Uint8("layers"),
		field.String("layer2_material").Optional().MaxLen(50),
		field.String("layer3_material").Optional().MaxLen(50),
		field.Float("partial_percentage").Optional().Default(0),
		field.String("material_mapping_note").MaxLen(500).Optional(),
		field.Enum("measure_type").GoType(enum.Measure("")).Annotations(entgql.Type("Measure")),
		field.
			Enum("extra_charge_type").
			GoType(enum.ExtraCharge("")).
			Optional().
			Default(enum.ExtraChargeNone.String()).
			Annotations(entgql.Type("ExtraCharge")),
		field.
			Float("extra_charges").
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}).
			Optional().
			Default(0),
		field.JSON("extra_charge_cond", []*model.ExtraChargeCondition{}).Optional(),
		field.String("extra_charge_note").Optional().Nillable(),
		field.String("estimator").Optional().Comment("estimate provider name"),
		field.Uint("estimator_order_id").Optional().Default(0),
		field.Uint("estimator_report_id").Optional().Default(0),
		field.
			Float("total_squares").
			Default(0).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(8,2)",
			}),
		field.
			Float("primary_pitch").
			Default(0).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(8,2)",
			}),
		field.Float("price").Default(0).SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.Text("price_summary").Optional(),
		field.JSON("bounds", []Point{}).Optional().Comment("geo boundary of property"),
		field.
			JSON("estimator_raw_response", map[string]interface{}{}).
			Optional().
			Comment("data received from estimate provider"),

		field.Bool("override").Default(false).Optional(),
		field.
			Float("override_total_squares").
			Default(0).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(8,2)",
			}),
		field.
			Float("override_primary_pitch").
			Default(0).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(8,2)",
			}),
		field.Float("override_price").Default(0).SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.Text("override_price_summary").Optional(),
		field.String("company_ref_id").MaxLen(36).Optional(),
		field.String("company_ref_name").MaxLen(100).Optional().Default(enum.MasterCompanyName),
		field.String("failure_reason").MaxLen(500).Optional().Comment("reason for estimate Failed status"),
	}
}

// Edges of the Estimate.
func (Estimate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("activities", EstimateActivity.Type).
			StorageKey(edge.Column("estimate_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("job_info", Job.Type).StorageKey(edge.Column("estimate_id")).Unique(),
		edge.From("partner", Partner.Type).Ref("requested_estimates").Unique(),
		edge.From("home_owner", HomeOwner.Type).Ref("estimates").Unique(),
		edge.From("sales_rep", User.Type).Ref("sales_rep_estimates").Unique(),
		edge.From("creator", User.Type).Ref("created_estimates").Unique(),
		edge.From("creator_api", ApiUser.Type).Ref("created_estimates").Unique(),
		edge.From("pdf", Document.Type).Ref("estimate_pdf").Unique(),
	}
}

func (e Estimate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
		index.Fields("region_id"),
		index.Fields("estimator_order_id"),
	}
}
