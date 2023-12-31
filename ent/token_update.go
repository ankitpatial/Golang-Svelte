// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/predicate"
	"roofix/ent/token"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks     []Hook
	mutation  *TokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetConfirmedAt sets the "confirmed_at" field.
func (tu *TokenUpdate) SetConfirmedAt(t time.Time) *TokenUpdate {
	tu.mutation.SetConfirmedAt(t)
	return tu
}

// SetNillableConfirmedAt sets the "confirmed_at" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableConfirmedAt(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetConfirmedAt(*t)
	}
	return tu
}

// ClearConfirmedAt clears the value of the "confirmed_at" field.
func (tu *TokenUpdate) ClearConfirmedAt() *TokenUpdate {
	tu.mutation.ClearConfirmedAt()
	return tu
}

// SetAction sets the "action" field.
func (tu *TokenUpdate) SetAction(s string) *TokenUpdate {
	tu.mutation.SetAction(s)
	return tu
}

// SetData sets the "data" field.
func (tu *TokenUpdate) SetData(m map[string]interface{}) *TokenUpdate {
	tu.mutation.SetData(m)
	return tu
}

// ClearData clears the value of the "data" field.
func (tu *TokenUpdate) ClearData() *TokenUpdate {
	tu.mutation.ClearData()
	return tu
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TokenUpdate) check() error {
	if v, ok := tu.mutation.Action(); ok {
		if err := token.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Token.action": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TokenUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.ConfirmedAt(); ok {
		_spec.SetField(token.FieldConfirmedAt, field.TypeTime, value)
	}
	if tu.mutation.ConfirmedAtCleared() {
		_spec.ClearField(token.FieldConfirmedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Action(); ok {
		_spec.SetField(token.FieldAction, field.TypeString, value)
	}
	if value, ok := tu.mutation.Data(); ok {
		_spec.SetField(token.FieldData, field.TypeJSON, value)
	}
	if tu.mutation.DataCleared() {
		_spec.ClearField(token.FieldData, field.TypeJSON)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetConfirmedAt sets the "confirmed_at" field.
func (tuo *TokenUpdateOne) SetConfirmedAt(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetConfirmedAt(t)
	return tuo
}

// SetNillableConfirmedAt sets the "confirmed_at" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableConfirmedAt(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetConfirmedAt(*t)
	}
	return tuo
}

// ClearConfirmedAt clears the value of the "confirmed_at" field.
func (tuo *TokenUpdateOne) ClearConfirmedAt() *TokenUpdateOne {
	tuo.mutation.ClearConfirmedAt()
	return tuo
}

// SetAction sets the "action" field.
func (tuo *TokenUpdateOne) SetAction(s string) *TokenUpdateOne {
	tuo.mutation.SetAction(s)
	return tuo
}

// SetData sets the "data" field.
func (tuo *TokenUpdateOne) SetData(m map[string]interface{}) *TokenUpdateOne {
	tuo.mutation.SetData(m)
	return tuo
}

// ClearData clears the value of the "data" field.
func (tuo *TokenUpdateOne) ClearData() *TokenUpdateOne {
	tuo.mutation.ClearData()
	return tuo
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tuo *TokenUpdateOne) Where(ps ...predicate.Token) *TokenUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TokenUpdateOne) check() error {
	if v, ok := tuo.mutation.Action(); ok {
		if err := token.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Token.action": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TokenUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.ConfirmedAt(); ok {
		_spec.SetField(token.FieldConfirmedAt, field.TypeTime, value)
	}
	if tuo.mutation.ConfirmedAtCleared() {
		_spec.ClearField(token.FieldConfirmedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Action(); ok {
		_spec.SetField(token.FieldAction, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Data(); ok {
		_spec.SetField(token.FieldData, field.TypeJSON, value)
	}
	if tuo.mutation.DataCleared() {
		_spec.ClearField(token.FieldData, field.TypeJSON)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
