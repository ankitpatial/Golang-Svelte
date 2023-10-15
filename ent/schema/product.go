package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/pkg/enum"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("type").
			GoType(enum.Product("")).
			Annotations(entgql.Type("ProductType")),
		field.
			String("name").
			MaxLen(100).
			Annotations(entgql.OrderField("NAME")),
		field.
			String("description").
			MaxLen(500),
		field.
			JSON("features", []string{}).
			Default([]string{}),
		field.
			String("special_note").
			Optional().
			MaxLen(100),
		field.
			Float("unit_price").
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.
			Int("units_in_stock").
			Optional().
			Default(0),
		field.
			Int("units_on_order").
			Optional().
			Default(0),
		field.
			Bool("discontinued").
			Optional().
			Default(false),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("package", ProductPackage.Type).
			Ref("items"),
		edge.
			From("creator", User.Type).
			Ref("products_created").
			Unique().
			Required(),
		edge.
			From("image", Document.Type).
			Ref("products_image").
			Unique(),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("name"),
		index.Fields("discontinued"),
	}
}
