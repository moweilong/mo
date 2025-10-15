package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*CreatorId)(nil)

type CreatorId struct {
	mixin.Schema
}

func (CreatorId) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("creator_id").
			Comment("创建者用户ID").
			Immutable().
			Optional().
			Nillable(),
	}
}
