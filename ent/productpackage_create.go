// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/product"
	"roofix/ent/productpackage"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductPackageCreate is the builder for creating a ProductPackage entity.
type ProductPackageCreate struct {
	config
	mutation *ProductPackageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ppc *ProductPackageCreate) SetCreatedAt(t time.Time) *ProductPackageCreate {
	ppc.mutation.SetCreatedAt(t)
	return ppc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ppc *ProductPackageCreate) SetNillableCreatedAt(t *time.Time) *ProductPackageCreate {
	if t != nil {
		ppc.SetCreatedAt(*t)
	}
	return ppc
}

// SetUpdatedAt sets the "updated_at" field.
func (ppc *ProductPackageCreate) SetUpdatedAt(t time.Time) *ProductPackageCreate {
	ppc.mutation.SetUpdatedAt(t)
	return ppc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ppc *ProductPackageCreate) SetNillableUpdatedAt(t *time.Time) *ProductPackageCreate {
	if t != nil {
		ppc.SetUpdatedAt(*t)
	}
	return ppc
}

// SetType sets the "type" field.
func (ppc *ProductPackageCreate) SetType(e enum.Product) *ProductPackageCreate {
	ppc.mutation.SetType(e)
	return ppc
}

// SetSoldAs sets the "sold_as" field.
func (ppc *ProductPackageCreate) SetSoldAs(ea enum.SoldAs) *ProductPackageCreate {
	ppc.mutation.SetSoldAs(ea)
	return ppc
}

// SetName sets the "name" field.
func (ppc *ProductPackageCreate) SetName(s string) *ProductPackageCreate {
	ppc.mutation.SetName(s)
	return ppc
}

// SetDescription sets the "description" field.
func (ppc *ProductPackageCreate) SetDescription(s string) *ProductPackageCreate {
	ppc.mutation.SetDescription(s)
	return ppc
}

// SetFeatures sets the "features" field.
func (ppc *ProductPackageCreate) SetFeatures(s []string) *ProductPackageCreate {
	ppc.mutation.SetFeatures(s)
	return ppc
}

// SetPrice sets the "price" field.
func (ppc *ProductPackageCreate) SetPrice(f float64) *ProductPackageCreate {
	ppc.mutation.SetPrice(f)
	return ppc
}

// SetDiscontinued sets the "discontinued" field.
func (ppc *ProductPackageCreate) SetDiscontinued(b bool) *ProductPackageCreate {
	ppc.mutation.SetDiscontinued(b)
	return ppc
}

// SetNillableDiscontinued sets the "discontinued" field if the given value is not nil.
func (ppc *ProductPackageCreate) SetNillableDiscontinued(b *bool) *ProductPackageCreate {
	if b != nil {
		ppc.SetDiscontinued(*b)
	}
	return ppc
}

// SetID sets the "id" field.
func (ppc *ProductPackageCreate) SetID(s string) *ProductPackageCreate {
	ppc.mutation.SetID(s)
	return ppc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ppc *ProductPackageCreate) SetNillableID(s *string) *ProductPackageCreate {
	if s != nil {
		ppc.SetID(*s)
	}
	return ppc
}

// AddItemIDs adds the "items" edge to the Product entity by IDs.
func (ppc *ProductPackageCreate) AddItemIDs(ids ...string) *ProductPackageCreate {
	ppc.mutation.AddItemIDs(ids...)
	return ppc
}

// AddItems adds the "items" edges to the Product entity.
func (ppc *ProductPackageCreate) AddItems(p ...*Product) *ProductPackageCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppc.AddItemIDs(ids...)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (ppc *ProductPackageCreate) SetCreatorID(id string) *ProductPackageCreate {
	ppc.mutation.SetCreatorID(id)
	return ppc
}

// SetCreator sets the "creator" edge to the User entity.
func (ppc *ProductPackageCreate) SetCreator(u *User) *ProductPackageCreate {
	return ppc.SetCreatorID(u.ID)
}

// Mutation returns the ProductPackageMutation object of the builder.
func (ppc *ProductPackageCreate) Mutation() *ProductPackageMutation {
	return ppc.mutation
}

// Save creates the ProductPackage in the database.
func (ppc *ProductPackageCreate) Save(ctx context.Context) (*ProductPackage, error) {
	ppc.defaults()
	return withHooks(ctx, ppc.sqlSave, ppc.mutation, ppc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ppc *ProductPackageCreate) SaveX(ctx context.Context) *ProductPackage {
	v, err := ppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ppc *ProductPackageCreate) Exec(ctx context.Context) error {
	_, err := ppc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppc *ProductPackageCreate) ExecX(ctx context.Context) {
	if err := ppc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ppc *ProductPackageCreate) defaults() {
	if _, ok := ppc.mutation.CreatedAt(); !ok {
		v := productpackage.DefaultCreatedAt()
		ppc.mutation.SetCreatedAt(v)
	}
	if _, ok := ppc.mutation.UpdatedAt(); !ok {
		v := productpackage.DefaultUpdatedAt()
		ppc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ppc.mutation.Features(); !ok {
		v := productpackage.DefaultFeatures
		ppc.mutation.SetFeatures(v)
	}
	if _, ok := ppc.mutation.Discontinued(); !ok {
		v := productpackage.DefaultDiscontinued
		ppc.mutation.SetDiscontinued(v)
	}
	if _, ok := ppc.mutation.ID(); !ok {
		v := productpackage.DefaultID()
		ppc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppc *ProductPackageCreate) check() error {
	if _, ok := ppc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProductPackage.created_at"`)}
	}
	if _, ok := ppc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProductPackage.updated_at"`)}
	}
	if _, ok := ppc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ProductPackage.type"`)}
	}
	if v, ok := ppc.mutation.GetType(); ok {
		if err := productpackage.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ProductPackage.type": %w`, err)}
		}
	}
	if _, ok := ppc.mutation.SoldAs(); !ok {
		return &ValidationError{Name: "sold_as", err: errors.New(`ent: missing required field "ProductPackage.sold_as"`)}
	}
	if v, ok := ppc.mutation.SoldAs(); ok {
		if err := productpackage.SoldAsValidator(v); err != nil {
			return &ValidationError{Name: "sold_as", err: fmt.Errorf(`ent: validator failed for field "ProductPackage.sold_as": %w`, err)}
		}
	}
	if _, ok := ppc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProductPackage.name"`)}
	}
	if v, ok := ppc.mutation.Name(); ok {
		if err := productpackage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ProductPackage.name": %w`, err)}
		}
	}
	if _, ok := ppc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ProductPackage.description"`)}
	}
	if v, ok := ppc.mutation.Description(); ok {
		if err := productpackage.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "ProductPackage.description": %w`, err)}
		}
	}
	if _, ok := ppc.mutation.Features(); !ok {
		return &ValidationError{Name: "features", err: errors.New(`ent: missing required field "ProductPackage.features"`)}
	}
	if _, ok := ppc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "ProductPackage.price"`)}
	}
	if v, ok := ppc.mutation.ID(); ok {
		if err := productpackage.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "ProductPackage.id": %w`, err)}
		}
	}
	if _, ok := ppc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required edge "ProductPackage.creator"`)}
	}
	return nil
}

func (ppc *ProductPackageCreate) sqlSave(ctx context.Context) (*ProductPackage, error) {
	if err := ppc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ppc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ProductPackage.ID type: %T", _spec.ID.Value)
		}
	}
	ppc.mutation.id = &_node.ID
	ppc.mutation.done = true
	return _node, nil
}

func (ppc *ProductPackageCreate) createSpec() (*ProductPackage, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductPackage{config: ppc.config}
		_spec = sqlgraph.NewCreateSpec(productpackage.Table, sqlgraph.NewFieldSpec(productpackage.FieldID, field.TypeString))
	)
	_spec.OnConflict = ppc.conflict
	if id, ok := ppc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ppc.mutation.CreatedAt(); ok {
		_spec.SetField(productpackage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ppc.mutation.UpdatedAt(); ok {
		_spec.SetField(productpackage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ppc.mutation.GetType(); ok {
		_spec.SetField(productpackage.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ppc.mutation.SoldAs(); ok {
		_spec.SetField(productpackage.FieldSoldAs, field.TypeEnum, value)
		_node.SoldAs = value
	}
	if value, ok := ppc.mutation.Name(); ok {
		_spec.SetField(productpackage.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ppc.mutation.Description(); ok {
		_spec.SetField(productpackage.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ppc.mutation.Features(); ok {
		_spec.SetField(productpackage.FieldFeatures, field.TypeJSON, value)
		_node.Features = value
	}
	if value, ok := ppc.mutation.Price(); ok {
		_spec.SetField(productpackage.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := ppc.mutation.Discontinued(); ok {
		_spec.SetField(productpackage.FieldDiscontinued, field.TypeBool, value)
		_node.Discontinued = value
	}
	if nodes := ppc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   productpackage.ItemsTable,
			Columns: productpackage.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productpackage.CreatorTable,
			Columns: []string{productpackage.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.creator_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProductPackage.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProductPackageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ppc *ProductPackageCreate) OnConflict(opts ...sql.ConflictOption) *ProductPackageUpsertOne {
	ppc.conflict = opts
	return &ProductPackageUpsertOne{
		create: ppc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProductPackage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ppc *ProductPackageCreate) OnConflictColumns(columns ...string) *ProductPackageUpsertOne {
	ppc.conflict = append(ppc.conflict, sql.ConflictColumns(columns...))
	return &ProductPackageUpsertOne{
		create: ppc,
	}
}

type (
	// ProductPackageUpsertOne is the builder for "upsert"-ing
	//  one ProductPackage node.
	ProductPackageUpsertOne struct {
		create *ProductPackageCreate
	}

	// ProductPackageUpsert is the "OnConflict" setter.
	ProductPackageUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ProductPackageUpsert) SetUpdatedAt(v time.Time) *ProductPackageUpsert {
	u.Set(productpackage.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdateUpdatedAt() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldUpdatedAt)
	return u
}

// SetType sets the "type" field.
func (u *ProductPackageUpsert) SetType(v enum.Product) *ProductPackageUpsert {
	u.Set(productpackage.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdateType() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldType)
	return u
}

// SetSoldAs sets the "sold_as" field.
func (u *ProductPackageUpsert) SetSoldAs(v enum.SoldAs) *ProductPackageUpsert {
	u.Set(productpackage.FieldSoldAs, v)
	return u
}

// UpdateSoldAs sets the "sold_as" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdateSoldAs() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldSoldAs)
	return u
}

// SetName sets the "name" field.
func (u *ProductPackageUpsert) SetName(v string) *ProductPackageUpsert {
	u.Set(productpackage.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdateName() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *ProductPackageUpsert) SetDescription(v string) *ProductPackageUpsert {
	u.Set(productpackage.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdateDescription() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldDescription)
	return u
}

// SetFeatures sets the "features" field.
func (u *ProductPackageUpsert) SetFeatures(v []string) *ProductPackageUpsert {
	u.Set(productpackage.FieldFeatures, v)
	return u
}

// UpdateFeatures sets the "features" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdateFeatures() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldFeatures)
	return u
}

// SetPrice sets the "price" field.
func (u *ProductPackageUpsert) SetPrice(v float64) *ProductPackageUpsert {
	u.Set(productpackage.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdatePrice() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldPrice)
	return u
}

// AddPrice adds v to the "price" field.
func (u *ProductPackageUpsert) AddPrice(v float64) *ProductPackageUpsert {
	u.Add(productpackage.FieldPrice, v)
	return u
}

// SetDiscontinued sets the "discontinued" field.
func (u *ProductPackageUpsert) SetDiscontinued(v bool) *ProductPackageUpsert {
	u.Set(productpackage.FieldDiscontinued, v)
	return u
}

// UpdateDiscontinued sets the "discontinued" field to the value that was provided on create.
func (u *ProductPackageUpsert) UpdateDiscontinued() *ProductPackageUpsert {
	u.SetExcluded(productpackage.FieldDiscontinued)
	return u
}

// ClearDiscontinued clears the value of the "discontinued" field.
func (u *ProductPackageUpsert) ClearDiscontinued() *ProductPackageUpsert {
	u.SetNull(productpackage.FieldDiscontinued)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ProductPackage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(productpackage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProductPackageUpsertOne) UpdateNewValues() *ProductPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(productpackage.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(productpackage.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProductPackage.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProductPackageUpsertOne) Ignore() *ProductPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProductPackageUpsertOne) DoNothing() *ProductPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProductPackageCreate.OnConflict
// documentation for more info.
func (u *ProductPackageUpsertOne) Update(set func(*ProductPackageUpsert)) *ProductPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProductPackageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProductPackageUpsertOne) SetUpdatedAt(v time.Time) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdateUpdatedAt() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetType sets the "type" field.
func (u *ProductPackageUpsertOne) SetType(v enum.Product) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdateType() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateType()
	})
}

// SetSoldAs sets the "sold_as" field.
func (u *ProductPackageUpsertOne) SetSoldAs(v enum.SoldAs) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetSoldAs(v)
	})
}

// UpdateSoldAs sets the "sold_as" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdateSoldAs() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateSoldAs()
	})
}

// SetName sets the "name" field.
func (u *ProductPackageUpsertOne) SetName(v string) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdateName() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ProductPackageUpsertOne) SetDescription(v string) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdateDescription() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateDescription()
	})
}

// SetFeatures sets the "features" field.
func (u *ProductPackageUpsertOne) SetFeatures(v []string) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetFeatures(v)
	})
}

// UpdateFeatures sets the "features" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdateFeatures() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateFeatures()
	})
}

// SetPrice sets the "price" field.
func (u *ProductPackageUpsertOne) SetPrice(v float64) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *ProductPackageUpsertOne) AddPrice(v float64) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdatePrice() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdatePrice()
	})
}

// SetDiscontinued sets the "discontinued" field.
func (u *ProductPackageUpsertOne) SetDiscontinued(v bool) *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetDiscontinued(v)
	})
}

// UpdateDiscontinued sets the "discontinued" field to the value that was provided on create.
func (u *ProductPackageUpsertOne) UpdateDiscontinued() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateDiscontinued()
	})
}

// ClearDiscontinued clears the value of the "discontinued" field.
func (u *ProductPackageUpsertOne) ClearDiscontinued() *ProductPackageUpsertOne {
	return u.Update(func(s *ProductPackageUpsert) {
		s.ClearDiscontinued()
	})
}

// Exec executes the query.
func (u *ProductPackageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProductPackageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProductPackageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProductPackageUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ProductPackageUpsertOne.ID is not supported by MySQL driver. Use ProductPackageUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProductPackageUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProductPackageCreateBulk is the builder for creating many ProductPackage entities in bulk.
type ProductPackageCreateBulk struct {
	config
	builders []*ProductPackageCreate
	conflict []sql.ConflictOption
}

// Save creates the ProductPackage entities in the database.
func (ppcb *ProductPackageCreateBulk) Save(ctx context.Context) ([]*ProductPackage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ppcb.builders))
	nodes := make([]*ProductPackage, len(ppcb.builders))
	mutators := make([]Mutator, len(ppcb.builders))
	for i := range ppcb.builders {
		func(i int, root context.Context) {
			builder := ppcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductPackageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ppcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ppcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ppcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ppcb *ProductPackageCreateBulk) SaveX(ctx context.Context) []*ProductPackage {
	v, err := ppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ppcb *ProductPackageCreateBulk) Exec(ctx context.Context) error {
	_, err := ppcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppcb *ProductPackageCreateBulk) ExecX(ctx context.Context) {
	if err := ppcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProductPackage.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProductPackageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ppcb *ProductPackageCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProductPackageUpsertBulk {
	ppcb.conflict = opts
	return &ProductPackageUpsertBulk{
		create: ppcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProductPackage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ppcb *ProductPackageCreateBulk) OnConflictColumns(columns ...string) *ProductPackageUpsertBulk {
	ppcb.conflict = append(ppcb.conflict, sql.ConflictColumns(columns...))
	return &ProductPackageUpsertBulk{
		create: ppcb,
	}
}

// ProductPackageUpsertBulk is the builder for "upsert"-ing
// a bulk of ProductPackage nodes.
type ProductPackageUpsertBulk struct {
	create *ProductPackageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ProductPackage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(productpackage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProductPackageUpsertBulk) UpdateNewValues() *ProductPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(productpackage.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(productpackage.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProductPackage.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProductPackageUpsertBulk) Ignore() *ProductPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProductPackageUpsertBulk) DoNothing() *ProductPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProductPackageCreateBulk.OnConflict
// documentation for more info.
func (u *ProductPackageUpsertBulk) Update(set func(*ProductPackageUpsert)) *ProductPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProductPackageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProductPackageUpsertBulk) SetUpdatedAt(v time.Time) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdateUpdatedAt() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetType sets the "type" field.
func (u *ProductPackageUpsertBulk) SetType(v enum.Product) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdateType() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateType()
	})
}

// SetSoldAs sets the "sold_as" field.
func (u *ProductPackageUpsertBulk) SetSoldAs(v enum.SoldAs) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetSoldAs(v)
	})
}

// UpdateSoldAs sets the "sold_as" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdateSoldAs() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateSoldAs()
	})
}

// SetName sets the "name" field.
func (u *ProductPackageUpsertBulk) SetName(v string) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdateName() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ProductPackageUpsertBulk) SetDescription(v string) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdateDescription() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateDescription()
	})
}

// SetFeatures sets the "features" field.
func (u *ProductPackageUpsertBulk) SetFeatures(v []string) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetFeatures(v)
	})
}

// UpdateFeatures sets the "features" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdateFeatures() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateFeatures()
	})
}

// SetPrice sets the "price" field.
func (u *ProductPackageUpsertBulk) SetPrice(v float64) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *ProductPackageUpsertBulk) AddPrice(v float64) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdatePrice() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdatePrice()
	})
}

// SetDiscontinued sets the "discontinued" field.
func (u *ProductPackageUpsertBulk) SetDiscontinued(v bool) *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.SetDiscontinued(v)
	})
}

// UpdateDiscontinued sets the "discontinued" field to the value that was provided on create.
func (u *ProductPackageUpsertBulk) UpdateDiscontinued() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.UpdateDiscontinued()
	})
}

// ClearDiscontinued clears the value of the "discontinued" field.
func (u *ProductPackageUpsertBulk) ClearDiscontinued() *ProductPackageUpsertBulk {
	return u.Update(func(s *ProductPackageUpsert) {
		s.ClearDiscontinued()
	})
}

// Exec executes the query.
func (u *ProductPackageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProductPackageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProductPackageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProductPackageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
