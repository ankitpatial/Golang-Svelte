// Code generated by ent, DO NOT EDIT.

package jobprogresshistory

import (
	"roofix/ent/predicate"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// Complete applies equality check predicate on the "complete" field. It's identical to CompleteEQ.
func Complete(v bool) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldComplete, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enum.JobProgress) predicate.JobProgressHistory {
	vc := v
	return predicate.JobProgressHistory(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enum.JobProgress) predicate.JobProgressHistory {
	vc := v
	return predicate.JobProgressHistory(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enum.JobProgress) predicate.JobProgressHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobProgressHistory(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enum.JobProgress) predicate.JobProgressHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobProgressHistory(sql.FieldNotIn(FieldStatus, v...))
}

// CompleteEQ applies the EQ predicate on the "complete" field.
func CompleteEQ(v bool) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldComplete, v))
}

// CompleteNEQ applies the NEQ predicate on the "complete" field.
func CompleteNEQ(v bool) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNEQ(FieldComplete, v))
}

// CompleteIsNil applies the IsNil predicate on the "complete" field.
func CompleteIsNil() predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldIsNull(FieldComplete))
}

// CompleteNotNil applies the NotNil predicate on the "complete" field.
func CompleteNotNil() predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNotNull(FieldComplete))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(sql.FieldContainsFold(FieldNote, v))
}

// HasJob applies the HasEdge predicate on the "job" edge.
func HasJob() predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobWith applies the HasEdge predicate on the "job" edge with a given conditions (other predicates).
func HasJobWith(preds ...predicate.Job) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		step := newJobStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatorAPIUser applies the HasEdge predicate on the "creator_api_user" edge.
func HasCreatorAPIUser() predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorAPIUserTable, CreatorAPIUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorAPIUserWith applies the HasEdge predicate on the "creator_api_user" edge with a given conditions (other predicates).
func HasCreatorAPIUserWith(preds ...predicate.ApiUser) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		step := newCreatorAPIUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobProgressHistory) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobProgressHistory) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
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
func Not(p predicate.JobProgressHistory) predicate.JobProgressHistory {
	return predicate.JobProgressHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}
