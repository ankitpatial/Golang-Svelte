// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/apiuser"
	"roofix/ent/estimate"
	"roofix/ent/estimateactivity"
	"roofix/ent/user"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EstimateActivityCreate is the builder for creating a EstimateActivity entity.
type EstimateActivityCreate struct {
	config
	mutation *EstimateActivityMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (eac *EstimateActivityCreate) SetCreatedAt(t time.Time) *EstimateActivityCreate {
	eac.mutation.SetCreatedAt(t)
	return eac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eac *EstimateActivityCreate) SetNillableCreatedAt(t *time.Time) *EstimateActivityCreate {
	if t != nil {
		eac.SetCreatedAt(*t)
	}
	return eac
}

// SetDescription sets the "description" field.
func (eac *EstimateActivityCreate) SetDescription(s string) *EstimateActivityCreate {
	eac.mutation.SetDescription(s)
	return eac
}

// SetRaw sets the "raw" field.
func (eac *EstimateActivityCreate) SetRaw(m map[string]interface{}) *EstimateActivityCreate {
	eac.mutation.SetRaw(m)
	return eac
}

// SetID sets the "id" field.
func (eac *EstimateActivityCreate) SetID(s string) *EstimateActivityCreate {
	eac.mutation.SetID(s)
	return eac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (eac *EstimateActivityCreate) SetNillableID(s *string) *EstimateActivityCreate {
	if s != nil {
		eac.SetID(*s)
	}
	return eac
}

// SetEstimateID sets the "estimate" edge to the Estimate entity by ID.
func (eac *EstimateActivityCreate) SetEstimateID(id string) *EstimateActivityCreate {
	eac.mutation.SetEstimateID(id)
	return eac
}

// SetNillableEstimateID sets the "estimate" edge to the Estimate entity by ID if the given value is not nil.
func (eac *EstimateActivityCreate) SetNillableEstimateID(id *string) *EstimateActivityCreate {
	if id != nil {
		eac = eac.SetEstimateID(*id)
	}
	return eac
}

// SetEstimate sets the "estimate" edge to the Estimate entity.
func (eac *EstimateActivityCreate) SetEstimate(e *Estimate) *EstimateActivityCreate {
	return eac.SetEstimateID(e.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (eac *EstimateActivityCreate) SetCreatorID(id string) *EstimateActivityCreate {
	eac.mutation.SetCreatorID(id)
	return eac
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (eac *EstimateActivityCreate) SetNillableCreatorID(id *string) *EstimateActivityCreate {
	if id != nil {
		eac = eac.SetCreatorID(*id)
	}
	return eac
}

// SetCreator sets the "creator" edge to the User entity.
func (eac *EstimateActivityCreate) SetCreator(u *User) *EstimateActivityCreate {
	return eac.SetCreatorID(u.ID)
}

// SetCreatorAPIID sets the "creator_api" edge to the ApiUser entity by ID.
func (eac *EstimateActivityCreate) SetCreatorAPIID(id string) *EstimateActivityCreate {
	eac.mutation.SetCreatorAPIID(id)
	return eac
}

// SetNillableCreatorAPIID sets the "creator_api" edge to the ApiUser entity by ID if the given value is not nil.
func (eac *EstimateActivityCreate) SetNillableCreatorAPIID(id *string) *EstimateActivityCreate {
	if id != nil {
		eac = eac.SetCreatorAPIID(*id)
	}
	return eac
}

// SetCreatorAPI sets the "creator_api" edge to the ApiUser entity.
func (eac *EstimateActivityCreate) SetCreatorAPI(a *ApiUser) *EstimateActivityCreate {
	return eac.SetCreatorAPIID(a.ID)
}

// Mutation returns the EstimateActivityMutation object of the builder.
func (eac *EstimateActivityCreate) Mutation() *EstimateActivityMutation {
	return eac.mutation
}

// Save creates the EstimateActivity in the database.
func (eac *EstimateActivityCreate) Save(ctx context.Context) (*EstimateActivity, error) {
	eac.defaults()
	return withHooks(ctx, eac.sqlSave, eac.mutation, eac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eac *EstimateActivityCreate) SaveX(ctx context.Context) *EstimateActivity {
	v, err := eac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eac *EstimateActivityCreate) Exec(ctx context.Context) error {
	_, err := eac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eac *EstimateActivityCreate) ExecX(ctx context.Context) {
	if err := eac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eac *EstimateActivityCreate) defaults() {
	if _, ok := eac.mutation.CreatedAt(); !ok {
		v := estimateactivity.DefaultCreatedAt()
		eac.mutation.SetCreatedAt(v)
	}
	if _, ok := eac.mutation.ID(); !ok {
		v := estimateactivity.DefaultID()
		eac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eac *EstimateActivityCreate) check() error {
	if _, ok := eac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EstimateActivity.created_at"`)}
	}
	if _, ok := eac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "EstimateActivity.description"`)}
	}
	if v, ok := eac.mutation.ID(); ok {
		if err := estimateactivity.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "EstimateActivity.id": %w`, err)}
		}
	}
	return nil
}

func (eac *EstimateActivityCreate) sqlSave(ctx context.Context) (*EstimateActivity, error) {
	if err := eac.check(); err != nil {
		return nil, err
	}
	_node, _spec := eac.createSpec()
	if err := sqlgraph.CreateNode(ctx, eac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EstimateActivity.ID type: %T", _spec.ID.Value)
		}
	}
	eac.mutation.id = &_node.ID
	eac.mutation.done = true
	return _node, nil
}

func (eac *EstimateActivityCreate) createSpec() (*EstimateActivity, *sqlgraph.CreateSpec) {
	var (
		_node = &EstimateActivity{config: eac.config}
		_spec = sqlgraph.NewCreateSpec(estimateactivity.Table, sqlgraph.NewFieldSpec(estimateactivity.FieldID, field.TypeString))
	)
	_spec.OnConflict = eac.conflict
	if id, ok := eac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := eac.mutation.CreatedAt(); ok {
		_spec.SetField(estimateactivity.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := eac.mutation.Description(); ok {
		_spec.SetField(estimateactivity.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := eac.mutation.Raw(); ok {
		_spec.SetField(estimateactivity.FieldRaw, field.TypeJSON, value)
		_node.Raw = value
	}
	if nodes := eac.mutation.EstimateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   estimateactivity.EstimateTable,
			Columns: []string{estimateactivity.EstimateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(estimate.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.estimate_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := eac.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   estimateactivity.CreatorTable,
			Columns: []string{estimateactivity.CreatorColumn},
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
	if nodes := eac.mutation.CreatorAPIIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   estimateactivity.CreatorAPITable,
			Columns: []string{estimateactivity.CreatorAPIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apiuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.api_user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EstimateActivity.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EstimateActivityUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (eac *EstimateActivityCreate) OnConflict(opts ...sql.ConflictOption) *EstimateActivityUpsertOne {
	eac.conflict = opts
	return &EstimateActivityUpsertOne{
		create: eac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EstimateActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eac *EstimateActivityCreate) OnConflictColumns(columns ...string) *EstimateActivityUpsertOne {
	eac.conflict = append(eac.conflict, sql.ConflictColumns(columns...))
	return &EstimateActivityUpsertOne{
		create: eac,
	}
}

type (
	// EstimateActivityUpsertOne is the builder for "upsert"-ing
	//  one EstimateActivity node.
	EstimateActivityUpsertOne struct {
		create *EstimateActivityCreate
	}

	// EstimateActivityUpsert is the "OnConflict" setter.
	EstimateActivityUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *EstimateActivityUpsert) SetDescription(v string) *EstimateActivityUpsert {
	u.Set(estimateactivity.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *EstimateActivityUpsert) UpdateDescription() *EstimateActivityUpsert {
	u.SetExcluded(estimateactivity.FieldDescription)
	return u
}

// SetRaw sets the "raw" field.
func (u *EstimateActivityUpsert) SetRaw(v map[string]interface{}) *EstimateActivityUpsert {
	u.Set(estimateactivity.FieldRaw, v)
	return u
}

// UpdateRaw sets the "raw" field to the value that was provided on create.
func (u *EstimateActivityUpsert) UpdateRaw() *EstimateActivityUpsert {
	u.SetExcluded(estimateactivity.FieldRaw)
	return u
}

// ClearRaw clears the value of the "raw" field.
func (u *EstimateActivityUpsert) ClearRaw() *EstimateActivityUpsert {
	u.SetNull(estimateactivity.FieldRaw)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.EstimateActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(estimateactivity.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EstimateActivityUpsertOne) UpdateNewValues() *EstimateActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(estimateactivity.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(estimateactivity.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EstimateActivity.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EstimateActivityUpsertOne) Ignore() *EstimateActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EstimateActivityUpsertOne) DoNothing() *EstimateActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EstimateActivityCreate.OnConflict
// documentation for more info.
func (u *EstimateActivityUpsertOne) Update(set func(*EstimateActivityUpsert)) *EstimateActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EstimateActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *EstimateActivityUpsertOne) SetDescription(v string) *EstimateActivityUpsertOne {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *EstimateActivityUpsertOne) UpdateDescription() *EstimateActivityUpsertOne {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.UpdateDescription()
	})
}

// SetRaw sets the "raw" field.
func (u *EstimateActivityUpsertOne) SetRaw(v map[string]interface{}) *EstimateActivityUpsertOne {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.SetRaw(v)
	})
}

// UpdateRaw sets the "raw" field to the value that was provided on create.
func (u *EstimateActivityUpsertOne) UpdateRaw() *EstimateActivityUpsertOne {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.UpdateRaw()
	})
}

// ClearRaw clears the value of the "raw" field.
func (u *EstimateActivityUpsertOne) ClearRaw() *EstimateActivityUpsertOne {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.ClearRaw()
	})
}

// Exec executes the query.
func (u *EstimateActivityUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EstimateActivityCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EstimateActivityUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EstimateActivityUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EstimateActivityUpsertOne.ID is not supported by MySQL driver. Use EstimateActivityUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EstimateActivityUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EstimateActivityCreateBulk is the builder for creating many EstimateActivity entities in bulk.
type EstimateActivityCreateBulk struct {
	config
	builders []*EstimateActivityCreate
	conflict []sql.ConflictOption
}

// Save creates the EstimateActivity entities in the database.
func (eacb *EstimateActivityCreateBulk) Save(ctx context.Context) ([]*EstimateActivity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eacb.builders))
	nodes := make([]*EstimateActivity, len(eacb.builders))
	mutators := make([]Mutator, len(eacb.builders))
	for i := range eacb.builders {
		func(i int, root context.Context) {
			builder := eacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EstimateActivityMutation)
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
					_, err = mutators[i+1].Mutate(root, eacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = eacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eacb *EstimateActivityCreateBulk) SaveX(ctx context.Context) []*EstimateActivity {
	v, err := eacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eacb *EstimateActivityCreateBulk) Exec(ctx context.Context) error {
	_, err := eacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eacb *EstimateActivityCreateBulk) ExecX(ctx context.Context) {
	if err := eacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EstimateActivity.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EstimateActivityUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (eacb *EstimateActivityCreateBulk) OnConflict(opts ...sql.ConflictOption) *EstimateActivityUpsertBulk {
	eacb.conflict = opts
	return &EstimateActivityUpsertBulk{
		create: eacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EstimateActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eacb *EstimateActivityCreateBulk) OnConflictColumns(columns ...string) *EstimateActivityUpsertBulk {
	eacb.conflict = append(eacb.conflict, sql.ConflictColumns(columns...))
	return &EstimateActivityUpsertBulk{
		create: eacb,
	}
}

// EstimateActivityUpsertBulk is the builder for "upsert"-ing
// a bulk of EstimateActivity nodes.
type EstimateActivityUpsertBulk struct {
	create *EstimateActivityCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.EstimateActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(estimateactivity.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EstimateActivityUpsertBulk) UpdateNewValues() *EstimateActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(estimateactivity.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(estimateactivity.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EstimateActivity.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EstimateActivityUpsertBulk) Ignore() *EstimateActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EstimateActivityUpsertBulk) DoNothing() *EstimateActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EstimateActivityCreateBulk.OnConflict
// documentation for more info.
func (u *EstimateActivityUpsertBulk) Update(set func(*EstimateActivityUpsert)) *EstimateActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EstimateActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *EstimateActivityUpsertBulk) SetDescription(v string) *EstimateActivityUpsertBulk {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *EstimateActivityUpsertBulk) UpdateDescription() *EstimateActivityUpsertBulk {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.UpdateDescription()
	})
}

// SetRaw sets the "raw" field.
func (u *EstimateActivityUpsertBulk) SetRaw(v map[string]interface{}) *EstimateActivityUpsertBulk {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.SetRaw(v)
	})
}

// UpdateRaw sets the "raw" field to the value that was provided on create.
func (u *EstimateActivityUpsertBulk) UpdateRaw() *EstimateActivityUpsertBulk {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.UpdateRaw()
	})
}

// ClearRaw clears the value of the "raw" field.
func (u *EstimateActivityUpsertBulk) ClearRaw() *EstimateActivityUpsertBulk {
	return u.Update(func(s *EstimateActivityUpsert) {
		s.ClearRaw()
	})
}

// Exec executes the query.
func (u *EstimateActivityUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EstimateActivityCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EstimateActivityCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EstimateActivityUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}