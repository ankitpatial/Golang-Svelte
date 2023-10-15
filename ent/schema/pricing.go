package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Pricing holds the schema definition for the Pricing entity.
type Pricing struct {
	ent.Schema
}

func (Pricing) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pricing"},
	}
}

func (Pricing) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the Pricing.
func (Pricing) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("postal_country").
			MaxLen(20),
		field.
			String("postal_code").
			MaxLen(20),
		field.
			Uint8("product_id"),
		field.
			String("description").
			Optional(),
		field.
			Float("price").
			SchemaType(map[string]string{dialect.MySQL: "decimal(8,2)"}),
		field.
			String("price_per").
			MaxLen(20),
	}
}

// Edges of the Pricing.
func (Pricing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("postal", PostalCode.Type).
			Ref("pricing").
			Unique().
			Required(),
	}
}

// Indexes of the Pricing.
func (Pricing) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("postal_country"),
		index.Fields("postal_code"),
	}
}
