// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/partnerservice"
	"roofix/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PartnerServiceUpdate is the builder for updating PartnerService entities.
type PartnerServiceUpdate struct {
	config
	hooks     []Hook
	mutation  *PartnerServiceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PartnerServiceUpdate builder.
func (psu *PartnerServiceUpdate) Where(ps ...predicate.PartnerService) *PartnerServiceUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetUpdatedAt sets the "updated_at" field.
func (psu *PartnerServiceUpdate) SetUpdatedAt(t time.Time) *PartnerServiceUpdate {
	psu.mutation.SetUpdatedAt(t)
	return psu
}

// SetServiceID sets the "service_id" field.
func (psu *PartnerServiceUpdate) SetServiceID(u uint8) *PartnerServiceUpdate {
	psu.mutation.ResetServiceID()
	psu.mutation.SetServiceID(u)
	return psu
}

// AddServiceID adds u to the "service_id" field.
func (psu *PartnerServiceUpdate) AddServiceID(u int8) *PartnerServiceUpdate {
	psu.mutation.AddServiceID(u)
	return psu
}

// SetActive sets the "active" field.
func (psu *PartnerServiceUpdate) SetActive(b bool) *PartnerServiceUpdate {
	psu.mutation.SetActive(b)
	return psu
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (psu *PartnerServiceUpdate) SetPartnerID(id string) *PartnerServiceUpdate {
	psu.mutation.SetPartnerID(id)
	return psu
}

// SetPartner sets the "partner" edge to the Partner entity.
func (psu *PartnerServiceUpdate) SetPartner(p *Partner) *PartnerServiceUpdate {
	return psu.SetPartnerID(p.ID)
}

// Mutation returns the PartnerServiceMutation object of the builder.
func (psu *PartnerServiceUpdate) Mutation() *PartnerServiceMutation {
	return psu.mutation
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (psu *PartnerServiceUpdate) ClearPartner() *PartnerServiceUpdate {
	psu.mutation.ClearPartner()
	return psu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PartnerServiceUpdate) Save(ctx context.Context) (int, error) {
	psu.defaults()
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PartnerServiceUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PartnerServiceUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PartnerServiceUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psu *PartnerServiceUpdate) defaults() {
	if _, ok := psu.mutation.UpdatedAt(); !ok {
		v := partnerservice.UpdateDefaultUpdatedAt()
		psu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psu *PartnerServiceUpdate) check() error {
	if _, ok := psu.mutation.PartnerID(); psu.mutation.PartnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PartnerService.partner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (psu *PartnerServiceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PartnerServiceUpdate {
	psu.modifiers = append(psu.modifiers, modifiers...)
	return psu
}

func (psu *PartnerServiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := psu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(partnerservice.Table, partnerservice.Columns, sqlgraph.NewFieldSpec(partnerservice.FieldID, field.TypeString))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.UpdatedAt(); ok {
		_spec.SetField(partnerservice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := psu.mutation.ServiceID(); ok {
		_spec.SetField(partnerservice.FieldServiceID, field.TypeUint8, value)
	}
	if value, ok := psu.mutation.AddedServiceID(); ok {
		_spec.AddField(partnerservice.FieldServiceID, field.TypeUint8, value)
	}
	if value, ok := psu.mutation.Active(); ok {
		_spec.SetField(partnerservice.FieldActive, field.TypeBool, value)
	}
	if psu.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partnerservice.PartnerTable,
			Columns: []string{partnerservice.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partnerservice.PartnerTable,
			Columns: []string{partnerservice.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(psu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partnerservice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PartnerServiceUpdateOne is the builder for updating a single PartnerService entity.
type PartnerServiceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PartnerServiceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (psuo *PartnerServiceUpdateOne) SetUpdatedAt(t time.Time) *PartnerServiceUpdateOne {
	psuo.mutation.SetUpdatedAt(t)
	return psuo
}

// SetServiceID sets the "service_id" field.
func (psuo *PartnerServiceUpdateOne) SetServiceID(u uint8) *PartnerServiceUpdateOne {
	psuo.mutation.ResetServiceID()
	psuo.mutation.SetServiceID(u)
	return psuo
}

// AddServiceID adds u to the "service_id" field.
func (psuo *PartnerServiceUpdateOne) AddServiceID(u int8) *PartnerServiceUpdateOne {
	psuo.mutation.AddServiceID(u)
	return psuo
}

// SetActive sets the "active" field.
func (psuo *PartnerServiceUpdateOne) SetActive(b bool) *PartnerServiceUpdateOne {
	psuo.mutation.SetActive(b)
	return psuo
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (psuo *PartnerServiceUpdateOne) SetPartnerID(id string) *PartnerServiceUpdateOne {
	psuo.mutation.SetPartnerID(id)
	return psuo
}

// SetPartner sets the "partner" edge to the Partner entity.
func (psuo *PartnerServiceUpdateOne) SetPartner(p *Partner) *PartnerServiceUpdateOne {
	return psuo.SetPartnerID(p.ID)
}

// Mutation returns the PartnerServiceMutation object of the builder.
func (psuo *PartnerServiceUpdateOne) Mutation() *PartnerServiceMutation {
	return psuo.mutation
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (psuo *PartnerServiceUpdateOne) ClearPartner() *PartnerServiceUpdateOne {
	psuo.mutation.ClearPartner()
	return psuo
}

// Where appends a list predicates to the PartnerServiceUpdate builder.
func (psuo *PartnerServiceUpdateOne) Where(ps ...predicate.PartnerService) *PartnerServiceUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PartnerServiceUpdateOne) Select(field string, fields ...string) *PartnerServiceUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PartnerService entity.
func (psuo *PartnerServiceUpdateOne) Save(ctx context.Context) (*PartnerService, error) {
	psuo.defaults()
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PartnerServiceUpdateOne) SaveX(ctx context.Context) *PartnerService {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PartnerServiceUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PartnerServiceUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psuo *PartnerServiceUpdateOne) defaults() {
	if _, ok := psuo.mutation.UpdatedAt(); !ok {
		v := partnerservice.UpdateDefaultUpdatedAt()
		psuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psuo *PartnerServiceUpdateOne) check() error {
	if _, ok := psuo.mutation.PartnerID(); psuo.mutation.PartnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PartnerService.partner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (psuo *PartnerServiceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PartnerServiceUpdateOne {
	psuo.modifiers = append(psuo.modifiers, modifiers...)
	return psuo
}

func (psuo *PartnerServiceUpdateOne) sqlSave(ctx context.Context) (_node *PartnerService, err error) {
	if err := psuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(partnerservice.Table, partnerservice.Columns, sqlgraph.NewFieldSpec(partnerservice.FieldID, field.TypeString))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PartnerService.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partnerservice.FieldID)
		for _, f := range fields {
			if !partnerservice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != partnerservice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.UpdatedAt(); ok {
		_spec.SetField(partnerservice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := psuo.mutation.ServiceID(); ok {
		_spec.SetField(partnerservice.FieldServiceID, field.TypeUint8, value)
	}
	if value, ok := psuo.mutation.AddedServiceID(); ok {
		_spec.AddField(partnerservice.FieldServiceID, field.TypeUint8, value)
	}
	if value, ok := psuo.mutation.Active(); ok {
		_spec.SetField(partnerservice.FieldActive, field.TypeBool, value)
	}
	if psuo.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partnerservice.PartnerTable,
			Columns: []string{partnerservice.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partnerservice.PartnerTable,
			Columns: []string{partnerservice.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(psuo.modifiers...)
	_node = &PartnerService{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partnerservice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}
