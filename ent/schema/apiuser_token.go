package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ApiUserToken holds the schema definition for the ApiUserToken entity.
type ApiUserToken struct {
	ent.Schema
}

func (ApiUserToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinIdAndCreateAt{},
	}
}

// Fields of the ApiUserToken.
func (ApiUserToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token_type").MaxLen(10),
		field.String("token_id").Unique().MaxLen(500),
		field.String("refresh_token_id").MaxLen(500),
	}
}

func (ApiUserToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token_type"),
		index.Fields("token_id"),
	}
}

// Edges of the ApiUserToken.
func (ApiUserToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("api_user", ApiUser.Type).Ref("tokens").Unique(),
	}
}
