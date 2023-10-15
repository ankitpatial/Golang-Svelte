// Code generated by ent, DO NOT EDIT.

package jobnote

import (
	"roofix/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.JobNote {
	return predicate.JobNote(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.JobNote {
	return predicate.JobNote(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.JobNote {
	return predicate.JobNote(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldUpdatedAt, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.JobNote {
	return predicate.JobNote(sql.FieldLTE(FieldUpdatedAt, v))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.JobNote {
	return predicate.JobNote(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.JobNote {
	return predicate.JobNote(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldHasSuffix(FieldNote, v))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.JobNote {
	return predicate.JobNote(sql.FieldContainsFold(FieldNote, v))
}

// HasJob applies the HasEdge predicate on the "job" edge.
func HasJob() predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobWith applies the HasEdge predicate on the "job" edge with a given conditions (other predicates).
func HasJobWith(preds ...predicate.Job) predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		step := newJobStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPartner applies the HasEdge predicate on the "partner" edge.
func HasPartner() predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PartnerTable, PartnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPartnerWith applies the HasEdge predicate on the "partner" edge with a given conditions (other predicates).
func HasPartnerWith(preds ...predicate.Partner) predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		step := newPartnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobNote) predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobNote) predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
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
func Not(p predicate.JobNote) predicate.JobNote {
	return predicate.JobNote(func(s *sql.Selector) {
		p(s.Not())
	})
}
