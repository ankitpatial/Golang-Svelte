// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"roofix/ent/jobassignmenthistory"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobAssignmentHistoryDelete is the builder for deleting a JobAssignmentHistory entity.
type JobAssignmentHistoryDelete struct {
	config
	hooks    []Hook
	mutation *JobAssignmentHistoryMutation
}

// Where appends a list predicates to the JobAssignmentHistoryDelete builder.
func (jahd *JobAssignmentHistoryDelete) Where(ps ...predicate.JobAssignmentHistory) *JobAssignmentHistoryDelete {
	jahd.mutation.Where(ps...)
	return jahd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jahd *JobAssignmentHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, jahd.sqlExec, jahd.mutation, jahd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jahd *JobAssignmentHistoryDelete) ExecX(ctx context.Context) int {
	n, err := jahd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jahd *JobAssignmentHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(jobassignmenthistory.Table, sqlgraph.NewFieldSpec(jobassignmenthistory.FieldID, field.TypeString))
	if ps := jahd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jahd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jahd.mutation.done = true
	return affected, err
}

// JobAssignmentHistoryDeleteOne is the builder for deleting a single JobAssignmentHistory entity.
type JobAssignmentHistoryDeleteOne struct {
	jahd *JobAssignmentHistoryDelete
}

// Where appends a list predicates to the JobAssignmentHistoryDelete builder.
func (jahdo *JobAssignmentHistoryDeleteOne) Where(ps ...predicate.JobAssignmentHistory) *JobAssignmentHistoryDeleteOne {
	jahdo.jahd.mutation.Where(ps...)
	return jahdo
}

// Exec executes the deletion query.
func (jahdo *JobAssignmentHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := jahdo.jahd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{jobassignmenthistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jahdo *JobAssignmentHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := jahdo.Exec(ctx); err != nil {
		panic(err)
	}
}
