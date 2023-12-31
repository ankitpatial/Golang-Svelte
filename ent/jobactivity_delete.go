// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"roofix/ent/jobactivity"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobActivityDelete is the builder for deleting a JobActivity entity.
type JobActivityDelete struct {
	config
	hooks    []Hook
	mutation *JobActivityMutation
}

// Where appends a list predicates to the JobActivityDelete builder.
func (jad *JobActivityDelete) Where(ps ...predicate.JobActivity) *JobActivityDelete {
	jad.mutation.Where(ps...)
	return jad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jad *JobActivityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, jad.sqlExec, jad.mutation, jad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jad *JobActivityDelete) ExecX(ctx context.Context) int {
	n, err := jad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jad *JobActivityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(jobactivity.Table, sqlgraph.NewFieldSpec(jobactivity.FieldID, field.TypeString))
	if ps := jad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jad.mutation.done = true
	return affected, err
}

// JobActivityDeleteOne is the builder for deleting a single JobActivity entity.
type JobActivityDeleteOne struct {
	jad *JobActivityDelete
}

// Where appends a list predicates to the JobActivityDelete builder.
func (jado *JobActivityDeleteOne) Where(ps ...predicate.JobActivity) *JobActivityDeleteOne {
	jado.jad.mutation.Where(ps...)
	return jado
}

// Exec executes the deletion query.
func (jado *JobActivityDeleteOne) Exec(ctx context.Context) error {
	n, err := jado.jad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{jobactivity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jado *JobActivityDeleteOne) ExecX(ctx context.Context) {
	if err := jado.Exec(ctx); err != nil {
		panic(err)
	}
}
