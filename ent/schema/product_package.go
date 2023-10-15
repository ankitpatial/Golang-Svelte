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
)

// ProductPackage holds the schema definition for the ProductPackage entity.
type ProductPackage struct {
	ent.Schema
}

func (ProductPackage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the ProductPackage.
func (ProductPackage) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("type").
			GoType(enum.Product("")).
			Annotations(entgql.Type("ProductType")),
		field.
			Enum("sold_as").
			GoType(enum.SoldAs("")).
			Annotations(entgql.Type("SoldAs")),
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
			Float("price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}),
		field.
			Bool("discontinued").
			Optional().
			Default(false),
	}
}

// Edges of the ProductPackage.
func (ProductPackage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("items", Product.Type).
			StorageKey(edge.Table("product_package_items"), edge.Columns("package_id", "product_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.
			From("creator", User.Type).
			Ref("product_pkg_created").
			Unique().
			Required(),
	}
}

func (ProductPackage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("name"),
		index.Fields("discontinued"),
	}
}
