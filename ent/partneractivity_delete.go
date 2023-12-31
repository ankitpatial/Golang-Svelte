// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"roofix/ent/partneractivity"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PartnerActivityDelete is the builder for deleting a PartnerActivity entity.
type PartnerActivityDelete struct {
	config
	hooks    []Hook
	mutation *PartnerActivityMutation
}

// Where appends a list predicates to the PartnerActivityDelete builder.
func (pad *PartnerActivityDelete) Where(ps ...predicate.PartnerActivity) *PartnerActivityDelete {
	pad.mutation.Where(ps...)
	return pad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pad *PartnerActivityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pad.sqlExec, pad.mutation, pad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pad *PartnerActivityDelete) ExecX(ctx context.Context) int {
	n, err := pad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pad *PartnerActivityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(partneractivity.Table, sqlgraph.NewFieldSpec(partneractivity.FieldID, field.TypeString))
	if ps := pad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pad.mutation.done = true
	return affected, err
}

// PartnerActivityDeleteOne is the builder for deleting a single PartnerActivity entity.
type PartnerActivityDeleteOne struct {
	pad *PartnerActivityDelete
}

// Where appends a list predicates to the PartnerActivityDelete builder.
func (pado *PartnerActivityDeleteOne) Where(ps ...predicate.PartnerActivity) *PartnerActivityDeleteOne {
	pado.pad.mutation.Where(ps...)
	return pado
}

// Exec executes the deletion query.
func (pado *PartnerActivityDeleteOne) Exec(ctx context.Context) error {
	n, err := pado.pad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{partneractivity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pado *PartnerActivityDeleteOne) ExecX(ctx context.Context) {
	if err := pado.Exec(ctx); err != nil {
		panic(err)
	}
}
