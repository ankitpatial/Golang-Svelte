package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PostalCode holds the schema definition for the PostalCode entity.
type PostalCode struct {
	ent.Schema
}

func (PostalCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the PostalCode.
func (PostalCode) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("country").
			MaxLen(20).
			Annotations(
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.
			String("code").
			MaxLen(12),
		field.
			String("city").
			MaxLen(50),
		field.
			String("state").
			MaxLen(50).
			Annotations(
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.
			String("state_abr").
			MaxLen(5).
			Annotations(
				entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"},
			),
		field.
			Uint8("region_id"),
		field.
			Float("latitude").
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,8)"}),
		field.
			Float("longitude").
			SchemaType(map[string]string{dialect.MySQL: "decimal(11,8)"}),
		field.
			Uint8("accuracy").
			Comment("accuracy of lat/lng, 1=estimated, 4=geonameid, 6=centroid"),
		field.
			Bool("service_area").
			Default(false).
			Comment("roofix service area"),
	}
}

// Edges of the PostalCode.
func (PostalCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("pricing", Pricing.Type).
			StorageKey(edge.Column("postal_id")),
	}
}

// Indexes of the PostalCode.
func (PostalCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country"),
		index.Fields("code"),
		index.Fields("state"),
		index.Fields("state_abr"),
		index.Fields("region_id"),
	}
}
