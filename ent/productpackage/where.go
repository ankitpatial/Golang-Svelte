// Code generated by ent, DO NOT EDIT.

package productpackage

import (
	"roofix/ent/predicate"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldDescription, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldPrice, v))
}

// Discontinued applies equality check predicate on the "discontinued" field. It's identical to DiscontinuedEQ.
func Discontinued(v bool) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldDiscontinued, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enum.Product) predicate.ProductPackage {
	vc := v
	return predicate.ProductPackage(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enum.Product) predicate.ProductPackage {
	vc := v
	return predicate.ProductPackage(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enum.Product) predicate.ProductPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductPackage(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enum.Product) predicate.ProductPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductPackage(sql.FieldNotIn(FieldType, v...))
}

// SoldAsEQ applies the EQ predicate on the "sold_as" field.
func SoldAsEQ(v enum.SoldAs) predicate.ProductPackage {
	vc := v
	return predicate.ProductPackage(sql.FieldEQ(FieldSoldAs, vc))
}

// SoldAsNEQ applies the NEQ predicate on the "sold_as" field.
func SoldAsNEQ(v enum.SoldAs) predicate.ProductPackage {
	vc := v
	return predicate.ProductPackage(sql.FieldNEQ(FieldSoldAs, vc))
}

// SoldAsIn applies the In predicate on the "sold_as" field.
func SoldAsIn(vs ...enum.SoldAs) predicate.ProductPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductPackage(sql.FieldIn(FieldSoldAs, v...))
}

// SoldAsNotIn applies the NotIn predicate on the "sold_as" field.
func SoldAsNotIn(vs ...enum.SoldAs) predicate.ProductPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductPackage(sql.FieldNotIn(FieldSoldAs, v...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldContainsFold(FieldDescription, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldLTE(FieldPrice, v))
}

// DiscontinuedEQ applies the EQ predicate on the "discontinued" field.
func DiscontinuedEQ(v bool) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldEQ(FieldDiscontinued, v))
}

// DiscontinuedNEQ applies the NEQ predicate on the "discontinued" field.
func DiscontinuedNEQ(v bool) predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNEQ(FieldDiscontinued, v))
}

// DiscontinuedIsNil applies the IsNil predicate on the "discontinued" field.
func DiscontinuedIsNil() predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldIsNull(FieldDiscontinued))
}

// DiscontinuedNotNil applies the NotNil predicate on the "discontinued" field.
func DiscontinuedNotNil() predicate.ProductPackage {
	return predicate.ProductPackage(sql.FieldNotNull(FieldDiscontinued))
}

// HasItems applies the HasEdge predicate on the "items" edge.
func HasItems() predicate.ProductPackage {
	return predicate.ProductPackage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ItemsTable, ItemsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemsWith applies the HasEdge predicate on the "items" edge with a given conditions (other predicates).
func HasItemsWith(preds ...predicate.Product) predicate.ProductPackage {
	return predicate.ProductPackage(func(s *sql.Selector) {
		step := newItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.ProductPackage {
	return predicate.ProductPackage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.ProductPackage {
	return predicate.ProductPackage(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductPackage) predicate.ProductPackage {
	return predicate.ProductPackage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductPackage) predicate.ProductPackage {
	return predicate.ProductPackage(func(s *sql.Selector) {
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
func Not(p predicate.ProductPackage) predicate.ProductPackage {
	return predicate.ProductPackage(func(s *sql.Selector) {
		p(s.Not())
	})
}
