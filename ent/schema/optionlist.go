package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"roofix/pkg/enum"
)

// OptionList holds the schema definition for the OptionList entity.
type OptionList struct {
	ent.Schema
}

func (OptionList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndStamps{},
	}
}

// Annotations of the JobAssignmentHistory.
func (OptionList) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "option_list"},
	}
}

// Fields of the OptionList.
func (OptionList) Fields() []ent.Field {
	return []ent.Field{
		field.
			Enum("type").
			GoType(enum.OptionList("")).
			Annotations(entgql.Type("OptionListType")),
		field.
			String("name").
			MaxLen(60),
		field.
			String("display_name").
			MaxLen(60),
		field.
			Bool("active").
			Default(true),
		field.
			Int("order"),
	}
}

func (OptionList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("finance_options", Partner.Type).
			StorageKey(edge.Table("partner_finance_options"), edge.Columns("option_id", "partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Restrict}),
		edge.
			To("epc_options", Partner.Type).
			StorageKey(edge.Table("partner_epc_options"), edge.Columns("option_id", "partner_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Restrict}),

		edge.
			To("epc", Job.Type).
			StorageKey(edge.Column("epc_id")).
			Annotations(entsql.Annotation{OnDelete: entsql.Restrict}),
	}
}

func (Partner) OptionList() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("name"),
		index.Fields("active"),
		index.Fields("order"),
	}
}
