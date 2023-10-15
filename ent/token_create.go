// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/token"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tc *TokenCreate) SetCreatedAt(t time.Time) *TokenCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TokenCreate) SetNillableCreatedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetConfirmedAt sets the "confirmed_at" field.
func (tc *TokenCreate) SetConfirmedAt(t time.Time) *TokenCreate {
	tc.mutation.SetConfirmedAt(t)
	return tc
}

// SetNillableConfirmedAt sets the "confirmed_at" field if the given value is not nil.
func (tc *TokenCreate) SetNillableConfirmedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetConfirmedAt(*t)
	}
	return tc
}

// SetAction sets the "action" field.
func (tc *TokenCreate) SetAction(s string) *TokenCreate {
	tc.mutation.SetAction(s)
	return tc
}

// SetData sets the "data" field.
func (tc *TokenCreate) SetData(m map[string]interface{}) *TokenCreate {
	tc.mutation.SetData(m)
	return tc
}

// SetID sets the "id" field.
func (tc *TokenCreate) SetID(s string) *TokenCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TokenCreate) SetNillableID(s *string) *TokenCreate {
	if s != nil {
		tc.SetID(*s)
	}
	return tc
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := token.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := token.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Token.created_at"`)}
	}
	if _, ok := tc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "Token.action"`)}
	}
	if v, ok := tc.mutation.Action(); ok {
		if err := token.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Token.action": %w`, err)}
		}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := token.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Token.id": %w`, err)}
		}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Token.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(token.Table, sqlgraph.NewFieldSpec(token.FieldID, field.TypeString))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(token.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.ConfirmedAt(); ok {
		_spec.SetField(token.FieldConfirmedAt, field.TypeTime, value)
		_node.ConfirmedAt = &value
	}
	if value, ok := tc.mutation.Action(); ok {
		_spec.SetField(token.FieldAction, field.TypeString, value)
		_node.Action = value
	}
	if value, ok := tc.mutation.Data(); ok {
		_spec.SetField(token.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tc *TokenCreate) OnConflict(opts ...sql.ConflictOption) *TokenUpsertOne {
	tc.conflict = opts
	return &TokenUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TokenCreate) OnConflictColumns(columns ...string) *TokenUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertOne{
		create: tc,
	}
}

type (
	// TokenUpsertOne is the builder for "upsert"-ing
	//  one Token node.
	TokenUpsertOne struct {
		create *TokenCreate
	}

	// TokenUpsert is the "OnConflict" setter.
	TokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetConfirmedAt sets the "confirmed_at" field.
func (u *TokenUpsert) SetConfirmedAt(v time.Time) *TokenUpsert {
	u.Set(token.FieldConfirmedAt, v)
	return u
}

// UpdateConfirmedAt sets the "confirmed_at" field to the value that was provided on create.
func (u *TokenUpsert) UpdateConfirmedAt() *TokenUpsert {
	u.SetExcluded(token.FieldConfirmedAt)
	return u
}

// ClearConfirmedAt clears the value of the "confirmed_at" field.
func (u *TokenUpsert) ClearConfirmedAt() *TokenUpsert {
	u.SetNull(token.FieldConfirmedAt)
	return u
}

// SetAction sets the "action" field.
func (u *TokenUpsert) SetAction(v string) *TokenUpsert {
	u.Set(token.FieldAction, v)
	return u
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *TokenUpsert) UpdateAction() *TokenUpsert {
	u.SetExcluded(token.FieldAction)
	return u
}

// SetData sets the "data" field.
func (u *TokenUpsert) SetData(v map[string]interface{}) *TokenUpsert {
	u.Set(token.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *TokenUpsert) UpdateData() *TokenUpsert {
	u.SetExcluded(token.FieldData)
	return u
}

// ClearData clears the value of the "data" field.
func (u *TokenUpsert) ClearData() *TokenUpsert {
	u.SetNull(token.FieldData)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenUpsertOne) UpdateNewValues() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(token.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(token.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TokenUpsertOne) Ignore() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertOne) DoNothing() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreate.OnConflict
// documentation for more info.
func (u *TokenUpsertOne) Update(set func(*TokenUpsert)) *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetConfirmedAt sets the "confirmed_at" field.
func (u *TokenUpsertOne) SetConfirmedAt(v time.Time) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetConfirmedAt(v)
	})
}

// UpdateConfirmedAt sets the "confirmed_at" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateConfirmedAt() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateConfirmedAt()
	})
}

// ClearConfirmedAt clears the value of the "confirmed_at" field.
func (u *TokenUpsertOne) ClearConfirmedAt() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.ClearConfirmedAt()
	})
}

// SetAction sets the "action" field.
func (u *TokenUpsertOne) SetAction(v string) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateAction() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateAction()
	})
}

// SetData sets the "data" field.
func (u *TokenUpsertOne) SetData(v map[string]interface{}) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateData() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *TokenUpsertOne) ClearData() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.ClearData()
	})
}

// Exec executes the query.
func (u *TokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TokenUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TokenUpsertOne.ID is not supported by MySQL driver. Use TokenUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TokenUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	builders []*TokenCreate
	conflict []sql.ConflictOption
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *TokenUpsertBulk {
	tcb.conflict = opts
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflictColumns(columns ...string) *TokenUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// TokenUpsertBulk is the builder for "upsert"-ing
// a bulk of Token nodes.
type TokenUpsertBulk struct {
	create *TokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenUpsertBulk) UpdateNewValues() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(token.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(token.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TokenUpsertBulk) Ignore() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertBulk) DoNothing() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreateBulk.OnConflict
// documentation for more info.
func (u *TokenUpsertBulk) Update(set func(*TokenUpsert)) *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetConfirmedAt sets the "confirmed_at" field.
func (u *TokenUpsertBulk) SetConfirmedAt(v time.Time) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetConfirmedAt(v)
	})
}

// UpdateConfirmedAt sets the "confirmed_at" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateConfirmedAt() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateConfirmedAt()
	})
}

// ClearConfirmedAt clears the value of the "confirmed_at" field.
func (u *TokenUpsertBulk) ClearConfirmedAt() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.ClearConfirmedAt()
	})
}

// SetAction sets the "action" field.
func (u *TokenUpsertBulk) SetAction(v string) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateAction() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateAction()
	})
}

// SetData sets the "data" field.
func (u *TokenUpsertBulk) SetData(v map[string]interface{}) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateData() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *TokenUpsertBulk) ClearData() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.ClearData()
	})
}

// Exec executes the query.
func (u *TokenUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
