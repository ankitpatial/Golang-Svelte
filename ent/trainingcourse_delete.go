// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"roofix/ent/predicate"
	"roofix/ent/trainingcourse"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrainingCourseDelete is the builder for deleting a TrainingCourse entity.
type TrainingCourseDelete struct {
	config
	hooks    []Hook
	mutation *TrainingCourseMutation
}

// Where appends a list predicates to the TrainingCourseDelete builder.
func (tcd *TrainingCourseDelete) Where(ps ...predicate.TrainingCourse) *TrainingCourseDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TrainingCourseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tcd.sqlExec, tcd.mutation, tcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TrainingCourseDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TrainingCourseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(trainingcourse.Table, sqlgraph.NewFieldSpec(trainingcourse.FieldID, field.TypeString))
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tcd.mutation.done = true
	return affected, err
}

// TrainingCourseDeleteOne is the builder for deleting a single TrainingCourse entity.
type TrainingCourseDeleteOne struct {
	tcd *TrainingCourseDelete
}

// Where appends a list predicates to the TrainingCourseDelete builder.
func (tcdo *TrainingCourseDeleteOne) Where(ps ...predicate.TrainingCourse) *TrainingCourseDeleteOne {
	tcdo.tcd.mutation.Where(ps...)
	return tcdo
}

// Exec executes the deletion query.
func (tcdo *TrainingCourseDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{trainingcourse.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TrainingCourseDeleteOne) ExecX(ctx context.Context) {
	if err := tcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
