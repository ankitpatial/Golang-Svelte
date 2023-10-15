package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InstallationJobItem holds the schema definition for the InstallationJobItem entity.
type InstallationJobItem struct {
	ent.Schema
}

// Fields of the InstallationJobItem.
func (InstallationJobItem) Fields() []ent.Field {
	return []ent.Field{
		fieldID,
		field.
			String("name").
			MaxLen(100),
		field.
			String("description").
			MaxLen(500),
		field.
			JSON("features", []string{}).
			Default([]string{}),
		field.
			Float("price").
			SchemaType(map[string]string{dialect.MySQL: "decimal(10,2)"}),
		field.
			String("img_key"),
		field.
			String("special_note").
			Optional().
			MaxLen(100),
	}
}

// Edges of the InstallationJobItem.
func (InstallationJobItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("job", InstallationJob.Type).
			Ref("items").
			Unique().
			Required(),
	}
}
