// Code generated by ent, DO NOT EDIT.

package channelsub

import (
	"roofix/ent/predicate"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldLTE(FieldUpdatedAt, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v enum.Role) predicate.ChannelSub {
	vc := v
	return predicate.ChannelSub(sql.FieldEQ(FieldRole, vc))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v enum.Role) predicate.ChannelSub {
	vc := v
	return predicate.ChannelSub(sql.FieldNEQ(FieldRole, vc))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...enum.Role) predicate.ChannelSub {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelSub(sql.FieldIn(FieldRole, v...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...enum.Role) predicate.ChannelSub {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelSub(sql.FieldNotIn(FieldRole, v...))
}

// RoleIsNil applies the IsNil predicate on the "role" field.
func RoleIsNil() predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldIsNull(FieldRole))
}

// RoleNotNil applies the NotNil predicate on the "role" field.
func RoleNotNil() predicate.ChannelSub {
	return predicate.ChannelSub(sql.FieldNotNull(FieldRole))
}

// HasChannel applies the HasEdge predicate on the "channel" edge.
func HasChannel() predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChannelWith applies the HasEdge predicate on the "channel" edge with a given conditions (other predicates).
func HasChannelWith(preds ...predicate.Channel) predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		step := newChannelStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPartner applies the HasEdge predicate on the "partner" edge.
func HasPartner() predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PartnerTable, PartnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPartnerWith applies the HasEdge predicate on the "partner" edge with a given conditions (other predicates).
func HasPartnerWith(preds ...predicate.Partner) predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		step := newPartnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChannelSub) predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChannelSub) predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChannelSub) predicate.ChannelSub {
	return predicate.ChannelSub(func(s *sql.Selector) {
		p(s.Not())
	})
}
