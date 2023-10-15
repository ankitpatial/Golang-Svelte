/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  29/04/22, 1:05 PM
 */

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"roofix/pkg/util/uid"
	"time"
)

var fieldID = field.String("id").
	MaxLen(36).
	NotEmpty().
	Immutable().
	DefaultFunc(uid.ULID).
	Unique()

var fieldCreated = field.Time("created_at").
	Default(time.Now).
	Immutable().
	StructTag(`json:"createdAt"`).
	Annotations(
		entgql.OrderField("CREATED"),
	)

var fieldUpdated = field.Time("updated_at").
	Default(time.Now).
	UpdateDefault(time.Now).
	StructTag(`json:"updatedAt"`).
	Annotations(
		entgql.OrderField("UPDATED"),
	)

var fieldDeleted = field.Time("deleted_at").
	Optional().
	Nillable().
	StructTag(`json:"deletedAt"`).
	Annotations(
		entgql.OrderField("DELETED"),
	)

type MixinIdAndAllStamps struct {
	mixin.Schema
}

func (MixinIdAndAllStamps) Fields() []ent.Field {
	return []ent.Field{
		fieldID,
		fieldCreated,
		fieldDeleted,
		fieldUpdated,
	}
}

func (MixinIdAndAllStamps) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at").Annotations(entsql.Desc()),
		index.Fields("updated_at").Annotations(entsql.Desc()),
		index.Fields("deleted_at").Annotations(entsql.Desc()),
	}
}

// == ID, CreatedAt & UpdatedAt

type MixinIdAndStamps struct {
	mixin.Schema
}

func (MixinIdAndStamps) Fields() []ent.Field {
	return []ent.Field{
		fieldID,
		fieldCreated,
		fieldUpdated,
	}
}

func (MixinIdAndStamps) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at").Annotations(entsql.Desc()),
		index.Fields("updated_at").Annotations(entsql.Desc()),
	}
}

// == ID & CreatedAt

type MixinIdAndCreateAt struct {
	mixin.Schema
}

func (MixinIdAndCreateAt) Fields() []ent.Field {
	return []ent.Field{
		fieldID,
		fieldCreated,
	}
}

func (MixinIdAndCreateAt) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at").Annotations(entsql.Desc()),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*MixinIdAndAllStamps)(nil)
var _ ent.Mixin = (*MixinIdAndStamps)(nil)
