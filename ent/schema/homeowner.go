package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// HomeOwner holds the schema definition for the HomeOwner entity.
type HomeOwner struct {
	ent.Schema
}

func (HomeOwner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Fields of the HomeOwner.
func (HomeOwner) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").MaxLen(50).Annotations(entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"}),
		field.String("last_name").MaxLen(50).Annotations(entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"}),
		field.String("email").Optional().Annotations(entsql.Annotation{Collation: "utf8mb4_0900_ai_ci"}),
		field.String("phone").Optional().MaxLen(20),
		field.String("street_number"),
		field.String("street_name"),
		field.String("city").MaxLen(50),
		field.String("state").MaxLen(50),
		field.String("state_abbr").Optional().MaxLen(10),
		field.String("zip").MaxLen(20),
		field.String("formatted_address").Optional(),
		field.
			Float("latitude").
			Optional().
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,8)"}),
		field.
			Float("longitude").
			Optional().
			SchemaType(map[string]string{dialect.MySQL: "decimal(11,8)"}),
		field.String("hash").MaxLen(36).Optional(),
	}
}

func (e HomeOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("estimates", Estimate.Type).StorageKey(edge.Column("home_owner_id")),
		edge.To("jobs", Job.Type).StorageKey(edge.Column("home_owner_id")),
		edge.From("partner", Partner.Type).Ref("estimate_home_owners").Unique(),
	}
}

func (HomeOwner) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
		index.Fields("first_name", "last_name"),
		index.Fields("phone"),
		index.Fields("formatted_address"),
		index.Fields("hash"),
	}
}
