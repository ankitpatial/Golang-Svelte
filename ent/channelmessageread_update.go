// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/channelmessage"
	"roofix/ent/channelmessageread"
	"roofix/ent/predicate"
	"roofix/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChannelMessageReadUpdate is the builder for updating ChannelMessageRead entities.
type ChannelMessageReadUpdate struct {
	config
	hooks     []Hook
	mutation  *ChannelMessageReadMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ChannelMessageReadUpdate builder.
func (cmru *ChannelMessageReadUpdate) Where(ps ...predicate.ChannelMessageRead) *ChannelMessageReadUpdate {
	cmru.mutation.Where(ps...)
	return cmru
}

// SetRead sets the "read" field.
func (cmru *ChannelMessageReadUpdate) SetRead(b bool) *ChannelMessageReadUpdate {
	cmru.mutation.SetRead(b)
	return cmru
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (cmru *ChannelMessageReadUpdate) SetNillableRead(b *bool) *ChannelMessageReadUpdate {
	if b != nil {
		cmru.SetRead(*b)
	}
	return cmru
}

// SetChannelMessageID sets the "channel_message" edge to the ChannelMessage entity by ID.
func (cmru *ChannelMessageReadUpdate) SetChannelMessageID(id string) *ChannelMessageReadUpdate {
	cmru.mutation.SetChannelMessageID(id)
	return cmru
}

// SetChannelMessage sets the "channel_message" edge to the ChannelMessage entity.
func (cmru *ChannelMessageReadUpdate) SetChannelMessage(c *ChannelMessage) *ChannelMessageReadUpdate {
	return cmru.SetChannelMessageID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cmru *ChannelMessageReadUpdate) SetUserID(id string) *ChannelMessageReadUpdate {
	cmru.mutation.SetUserID(id)
	return cmru
}

// SetUser sets the "user" edge to the User entity.
func (cmru *ChannelMessageReadUpdate) SetUser(u *User) *ChannelMessageReadUpdate {
	return cmru.SetUserID(u.ID)
}

// Mutation returns the ChannelMessageReadMutation object of the builder.
func (cmru *ChannelMessageReadUpdate) Mutation() *ChannelMessageReadMutation {
	return cmru.mutation
}

// ClearChannelMessage clears the "channel_message" edge to the ChannelMessage entity.
func (cmru *ChannelMessageReadUpdate) ClearChannelMessage() *ChannelMessageReadUpdate {
	cmru.mutation.ClearChannelMessage()
	return cmru
}

// ClearUser clears the "user" edge to the User entity.
func (cmru *ChannelMessageReadUpdate) ClearUser() *ChannelMessageReadUpdate {
	cmru.mutation.ClearUser()
	return cmru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmru *ChannelMessageReadUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cmru.sqlSave, cmru.mutation, cmru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmru *ChannelMessageReadUpdate) SaveX(ctx context.Context) int {
	affected, err := cmru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmru *ChannelMessageReadUpdate) Exec(ctx context.Context) error {
	_, err := cmru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmru *ChannelMessageReadUpdate) ExecX(ctx context.Context) {
	if err := cmru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmru *ChannelMessageReadUpdate) check() error {
	if _, ok := cmru.mutation.ChannelMessageID(); cmru.mutation.ChannelMessageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChannelMessageRead.channel_message"`)
	}
	if _, ok := cmru.mutation.UserID(); cmru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChannelMessageRead.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cmru *ChannelMessageReadUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChannelMessageReadUpdate {
	cmru.modifiers = append(cmru.modifiers, modifiers...)
	return cmru
}

func (cmru *ChannelMessageReadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cmru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(channelmessageread.Table, channelmessageread.Columns, sqlgraph.NewFieldSpec(channelmessageread.FieldID, field.TypeString))
	if ps := cmru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmru.mutation.Read(); ok {
		_spec.SetField(channelmessageread.FieldRead, field.TypeBool, value)
	}
	if cmru.mutation.ChannelMessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.ChannelMessageTable,
			Columns: []string{channelmessageread.ChannelMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(channelmessage.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmru.mutation.ChannelMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.ChannelMessageTable,
			Columns: []string{channelmessageread.ChannelMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(channelmessage.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.UserTable,
			Columns: []string{channelmessageread.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.UserTable,
			Columns: []string{channelmessageread.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cmru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cmru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channelmessageread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cmru.mutation.done = true
	return n, nil
}

// ChannelMessageReadUpdateOne is the builder for updating a single ChannelMessageRead entity.
type ChannelMessageReadUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ChannelMessageReadMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetRead sets the "read" field.
func (cmruo *ChannelMessageReadUpdateOne) SetRead(b bool) *ChannelMessageReadUpdateOne {
	cmruo.mutation.SetRead(b)
	return cmruo
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (cmruo *ChannelMessageReadUpdateOne) SetNillableRead(b *bool) *ChannelMessageReadUpdateOne {
	if b != nil {
		cmruo.SetRead(*b)
	}
	return cmruo
}

// SetChannelMessageID sets the "channel_message" edge to the ChannelMessage entity by ID.
func (cmruo *ChannelMessageReadUpdateOne) SetChannelMessageID(id string) *ChannelMessageReadUpdateOne {
	cmruo.mutation.SetChannelMessageID(id)
	return cmruo
}

// SetChannelMessage sets the "channel_message" edge to the ChannelMessage entity.
func (cmruo *ChannelMessageReadUpdateOne) SetChannelMessage(c *ChannelMessage) *ChannelMessageReadUpdateOne {
	return cmruo.SetChannelMessageID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cmruo *ChannelMessageReadUpdateOne) SetUserID(id string) *ChannelMessageReadUpdateOne {
	cmruo.mutation.SetUserID(id)
	return cmruo
}

// SetUser sets the "user" edge to the User entity.
func (cmruo *ChannelMessageReadUpdateOne) SetUser(u *User) *ChannelMessageReadUpdateOne {
	return cmruo.SetUserID(u.ID)
}

// Mutation returns the ChannelMessageReadMutation object of the builder.
func (cmruo *ChannelMessageReadUpdateOne) Mutation() *ChannelMessageReadMutation {
	return cmruo.mutation
}

// ClearChannelMessage clears the "channel_message" edge to the ChannelMessage entity.
func (cmruo *ChannelMessageReadUpdateOne) ClearChannelMessage() *ChannelMessageReadUpdateOne {
	cmruo.mutation.ClearChannelMessage()
	return cmruo
}

// ClearUser clears the "user" edge to the User entity.
func (cmruo *ChannelMessageReadUpdateOne) ClearUser() *ChannelMessageReadUpdateOne {
	cmruo.mutation.ClearUser()
	return cmruo
}

// Where appends a list predicates to the ChannelMessageReadUpdate builder.
func (cmruo *ChannelMessageReadUpdateOne) Where(ps ...predicate.ChannelMessageRead) *ChannelMessageReadUpdateOne {
	cmruo.mutation.Where(ps...)
	return cmruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmruo *ChannelMessageReadUpdateOne) Select(field string, fields ...string) *ChannelMessageReadUpdateOne {
	cmruo.fields = append([]string{field}, fields...)
	return cmruo
}

// Save executes the query and returns the updated ChannelMessageRead entity.
func (cmruo *ChannelMessageReadUpdateOne) Save(ctx context.Context) (*ChannelMessageRead, error) {
	return withHooks(ctx, cmruo.sqlSave, cmruo.mutation, cmruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmruo *ChannelMessageReadUpdateOne) SaveX(ctx context.Context) *ChannelMessageRead {
	node, err := cmruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmruo *ChannelMessageReadUpdateOne) Exec(ctx context.Context) error {
	_, err := cmruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmruo *ChannelMessageReadUpdateOne) ExecX(ctx context.Context) {
	if err := cmruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmruo *ChannelMessageReadUpdateOne) check() error {
	if _, ok := cmruo.mutation.ChannelMessageID(); cmruo.mutation.ChannelMessageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChannelMessageRead.channel_message"`)
	}
	if _, ok := cmruo.mutation.UserID(); cmruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChannelMessageRead.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cmruo *ChannelMessageReadUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChannelMessageReadUpdateOne {
	cmruo.modifiers = append(cmruo.modifiers, modifiers...)
	return cmruo
}

func (cmruo *ChannelMessageReadUpdateOne) sqlSave(ctx context.Context) (_node *ChannelMessageRead, err error) {
	if err := cmruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(channelmessageread.Table, channelmessageread.Columns, sqlgraph.NewFieldSpec(channelmessageread.FieldID, field.TypeString))
	id, ok := cmruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChannelMessageRead.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cmruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channelmessageread.FieldID)
		for _, f := range fields {
			if !channelmessageread.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != channelmessageread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmruo.mutation.Read(); ok {
		_spec.SetField(channelmessageread.FieldRead, field.TypeBool, value)
	}
	if cmruo.mutation.ChannelMessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.ChannelMessageTable,
			Columns: []string{channelmessageread.ChannelMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(channelmessage.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmruo.mutation.ChannelMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.ChannelMessageTable,
			Columns: []string{channelmessageread.ChannelMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(channelmessage.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.UserTable,
			Columns: []string{channelmessageread.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channelmessageread.UserTable,
			Columns: []string{channelmessageread.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cmruo.modifiers...)
	_node = &ChannelMessageRead{config: cmruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channelmessageread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cmruo.mutation.done = true
	return _node, nil
}
