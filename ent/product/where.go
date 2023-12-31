// Code generated by ent, DO NOT EDIT.

package product

import (
	"roofix/ent/predicate"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// SpecialNote applies equality check predicate on the "special_note" field. It's identical to SpecialNoteEQ.
func SpecialNote(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSpecialNote, v))
}

// UnitPrice applies equality check predicate on the "unit_price" field. It's identical to UnitPriceEQ.
func UnitPrice(v float64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitPrice, v))
}

// UnitsInStock applies equality check predicate on the "units_in_stock" field. It's identical to UnitsInStockEQ.
func UnitsInStock(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitsInStock, v))
}

// UnitsOnOrder applies equality check predicate on the "units_on_order" field. It's identical to UnitsOnOrderEQ.
func UnitsOnOrder(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitsOnOrder, v))
}

// Discontinued applies equality check predicate on the "discontinued" field. It's identical to DiscontinuedEQ.
func Discontinued(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDiscontinued, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enum.Product) predicate.Product {
	vc := v
	return predicate.Product(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enum.Product) predicate.Product {
	vc := v
	return predicate.Product(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enum.Product) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enum.Product) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(sql.FieldNotIn(FieldType, v...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldDescription, v))
}

// SpecialNoteEQ applies the EQ predicate on the "special_note" field.
func SpecialNoteEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSpecialNote, v))
}

// SpecialNoteNEQ applies the NEQ predicate on the "special_note" field.
func SpecialNoteNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldSpecialNote, v))
}

// SpecialNoteIn applies the In predicate on the "special_note" field.
func SpecialNoteIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldSpecialNote, vs...))
}

// SpecialNoteNotIn applies the NotIn predicate on the "special_note" field.
func SpecialNoteNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldSpecialNote, vs...))
}

// SpecialNoteGT applies the GT predicate on the "special_note" field.
func SpecialNoteGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldSpecialNote, v))
}

// SpecialNoteGTE applies the GTE predicate on the "special_note" field.
func SpecialNoteGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldSpecialNote, v))
}

// SpecialNoteLT applies the LT predicate on the "special_note" field.
func SpecialNoteLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldSpecialNote, v))
}

// SpecialNoteLTE applies the LTE predicate on the "special_note" field.
func SpecialNoteLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldSpecialNote, v))
}

// SpecialNoteContains applies the Contains predicate on the "special_note" field.
func SpecialNoteContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldSpecialNote, v))
}

// SpecialNoteHasPrefix applies the HasPrefix predicate on the "special_note" field.
func SpecialNoteHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldSpecialNote, v))
}

// SpecialNoteHasSuffix applies the HasSuffix predicate on the "special_note" field.
func SpecialNoteHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldSpecialNote, v))
}

// SpecialNoteIsNil applies the IsNil predicate on the "special_note" field.
func SpecialNoteIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldSpecialNote))
}

// SpecialNoteNotNil applies the NotNil predicate on the "special_note" field.
func SpecialNoteNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldSpecialNote))
}

// SpecialNoteEqualFold applies the EqualFold predicate on the "special_note" field.
func SpecialNoteEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldSpecialNote, v))
}

// SpecialNoteContainsFold applies the ContainsFold predicate on the "special_note" field.
func SpecialNoteContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldSpecialNote, v))
}

// UnitPriceEQ applies the EQ predicate on the "unit_price" field.
func UnitPriceEQ(v float64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitPrice, v))
}

// UnitPriceNEQ applies the NEQ predicate on the "unit_price" field.
func UnitPriceNEQ(v float64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUnitPrice, v))
}

// UnitPriceIn applies the In predicate on the "unit_price" field.
func UnitPriceIn(vs ...float64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUnitPrice, vs...))
}

// UnitPriceNotIn applies the NotIn predicate on the "unit_price" field.
func UnitPriceNotIn(vs ...float64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUnitPrice, vs...))
}

// UnitPriceGT applies the GT predicate on the "unit_price" field.
func UnitPriceGT(v float64) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUnitPrice, v))
}

// UnitPriceGTE applies the GTE predicate on the "unit_price" field.
func UnitPriceGTE(v float64) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUnitPrice, v))
}

// UnitPriceLT applies the LT predicate on the "unit_price" field.
func UnitPriceLT(v float64) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUnitPrice, v))
}

// UnitPriceLTE applies the LTE predicate on the "unit_price" field.
func UnitPriceLTE(v float64) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUnitPrice, v))
}

// UnitsInStockEQ applies the EQ predicate on the "units_in_stock" field.
func UnitsInStockEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitsInStock, v))
}

// UnitsInStockNEQ applies the NEQ predicate on the "units_in_stock" field.
func UnitsInStockNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUnitsInStock, v))
}

// UnitsInStockIn applies the In predicate on the "units_in_stock" field.
func UnitsInStockIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUnitsInStock, vs...))
}

// UnitsInStockNotIn applies the NotIn predicate on the "units_in_stock" field.
func UnitsInStockNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUnitsInStock, vs...))
}

// UnitsInStockGT applies the GT predicate on the "units_in_stock" field.
func UnitsInStockGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUnitsInStock, v))
}

// UnitsInStockGTE applies the GTE predicate on the "units_in_stock" field.
func UnitsInStockGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUnitsInStock, v))
}

// UnitsInStockLT applies the LT predicate on the "units_in_stock" field.
func UnitsInStockLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUnitsInStock, v))
}

// UnitsInStockLTE applies the LTE predicate on the "units_in_stock" field.
func UnitsInStockLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUnitsInStock, v))
}

// UnitsInStockIsNil applies the IsNil predicate on the "units_in_stock" field.
func UnitsInStockIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldUnitsInStock))
}

// UnitsInStockNotNil applies the NotNil predicate on the "units_in_stock" field.
func UnitsInStockNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldUnitsInStock))
}

// UnitsOnOrderEQ applies the EQ predicate on the "units_on_order" field.
func UnitsOnOrderEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUnitsOnOrder, v))
}

// UnitsOnOrderNEQ applies the NEQ predicate on the "units_on_order" field.
func UnitsOnOrderNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUnitsOnOrder, v))
}

// UnitsOnOrderIn applies the In predicate on the "units_on_order" field.
func UnitsOnOrderIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUnitsOnOrder, vs...))
}

// UnitsOnOrderNotIn applies the NotIn predicate on the "units_on_order" field.
func UnitsOnOrderNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUnitsOnOrder, vs...))
}

// UnitsOnOrderGT applies the GT predicate on the "units_on_order" field.
func UnitsOnOrderGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUnitsOnOrder, v))
}

// UnitsOnOrderGTE applies the GTE predicate on the "units_on_order" field.
func UnitsOnOrderGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUnitsOnOrder, v))
}

// UnitsOnOrderLT applies the LT predicate on the "units_on_order" field.
func UnitsOnOrderLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUnitsOnOrder, v))
}

// UnitsOnOrderLTE applies the LTE predicate on the "units_on_order" field.
func UnitsOnOrderLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUnitsOnOrder, v))
}

// UnitsOnOrderIsNil applies the IsNil predicate on the "units_on_order" field.
func UnitsOnOrderIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldUnitsOnOrder))
}

// UnitsOnOrderNotNil applies the NotNil predicate on the "units_on_order" field.
func UnitsOnOrderNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldUnitsOnOrder))
}

// DiscontinuedEQ applies the EQ predicate on the "discontinued" field.
func DiscontinuedEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDiscontinued, v))
}

// DiscontinuedNEQ applies the NEQ predicate on the "discontinued" field.
func DiscontinuedNEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldDiscontinued, v))
}

// DiscontinuedIsNil applies the IsNil predicate on the "discontinued" field.
func DiscontinuedIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldDiscontinued))
}

// DiscontinuedNotNil applies the NotNil predicate on the "discontinued" field.
func DiscontinuedNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldDiscontinued))
}

// HasPackage applies the HasEdge predicate on the "package" edge.
func HasPackage() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PackageTable, PackagePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPackageWith applies the HasEdge predicate on the "package" edge with a given conditions (other predicates).
func HasPackageWith(preds ...predicate.ProductPackage) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newPackageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImage applies the HasEdge predicate on the "image" edge.
func HasImage() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ImageTable, ImageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImageWith applies the HasEdge predicate on the "image" edge with a given conditions (other predicates).
func HasImageWith(preds ...predicate.Document) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newImageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
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
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
