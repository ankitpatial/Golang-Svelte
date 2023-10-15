// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"roofix/ent/channelmessage"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChannelMessageDelete is the builder for deleting a ChannelMessage entity.
type ChannelMessageDelete struct {
	config
	hooks    []Hook
	mutation *ChannelMessageMutation
}

// Where appends a list predicates to the ChannelMessageDelete builder.
func (cmd *ChannelMessageDelete) Where(ps ...predicate.ChannelMessage) *ChannelMessageDelete {
	cmd.mutation.Where(ps...)
	return cmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmd *ChannelMessageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cmd.sqlExec, cmd.mutation, cmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cmd *ChannelMessageDelete) ExecX(ctx context.Context) int {
	n, err := cmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmd *ChannelMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(channelmessage.Table, sqlgraph.NewFieldSpec(channelmessage.FieldID, field.TypeString))
	if ps := cmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cmd.mutation.done = true
	return affected, err
}

// ChannelMessageDeleteOne is the builder for deleting a single ChannelMessage entity.
type ChannelMessageDeleteOne struct {
	cmd *ChannelMessageDelete
}

// Where appends a list predicates to the ChannelMessageDelete builder.
func (cmdo *ChannelMessageDeleteOne) Where(ps ...predicate.ChannelMessage) *ChannelMessageDeleteOne {
	cmdo.cmd.mutation.Where(ps...)
	return cmdo
}

// Exec executes the deletion query.
func (cmdo *ChannelMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := cmdo.cmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{channelmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmdo *ChannelMessageDeleteOne) ExecX(ctx context.Context) {
	if err := cmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
