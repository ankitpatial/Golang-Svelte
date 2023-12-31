// Code generated by ent, DO NOT EDIT.

package useractivity

import (
	"roofix/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEQ(FieldCreatedAt, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEQ(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldLTE(FieldCreatedAt, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.UserActivity {
	return predicate.UserActivity(sql.FieldContainsFold(FieldDescription, v))
}

// RawIsNil applies the IsNil predicate on the "raw" field.
func RawIsNil() predicate.UserActivity {
	return predicate.UserActivity(sql.FieldIsNull(FieldRaw))
}

// RawNotNil applies the NotNil predicate on the "raw" field.
func RawNotNil() predicate.UserActivity {
	return predicate.UserActivity(sql.FieldNotNull(FieldRaw))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatorAPI applies the HasEdge predicate on the "creator_api" edge.
func HasCreatorAPI() predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorAPITable, CreatorAPIColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorAPIWith applies the HasEdge predicate on the "creator_api" edge with a given conditions (other predicates).
func HasCreatorAPIWith(preds ...predicate.ApiUser) predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		step := newCreatorAPIStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserActivity) predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserActivity) predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
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
func Not(p predicate.UserActivity) predicate.UserActivity {
	return predicate.UserActivity(func(s *sql.Selector) {
		p(s.Not())
	})
}
