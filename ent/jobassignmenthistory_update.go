// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/job"
	"roofix/ent/jobassignmenthistory"
	"roofix/ent/partner"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobAssignmentHistoryUpdate is the builder for updating JobAssignmentHistory entities.
type JobAssignmentHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *JobAssignmentHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the JobAssignmentHistoryUpdate builder.
func (jahu *JobAssignmentHistoryUpdate) Where(ps ...predicate.JobAssignmentHistory) *JobAssignmentHistoryUpdate {
	jahu.mutation.Where(ps...)
	return jahu
}

// SetNote sets the "Note" field.
func (jahu *JobAssignmentHistoryUpdate) SetNote(s string) *JobAssignmentHistoryUpdate {
	jahu.mutation.SetNote(s)
	return jahu
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jahu *JobAssignmentHistoryUpdate) SetNillableNote(s *string) *JobAssignmentHistoryUpdate {
	if s != nil {
		jahu.SetNote(*s)
	}
	return jahu
}

// ClearNote clears the value of the "Note" field.
func (jahu *JobAssignmentHistoryUpdate) ClearNote() *JobAssignmentHistoryUpdate {
	jahu.mutation.ClearNote()
	return jahu
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jahu *JobAssignmentHistoryUpdate) SetJobID(id string) *JobAssignmentHistoryUpdate {
	jahu.mutation.SetJobID(id)
	return jahu
}

// SetNillableJobID sets the "job" edge to the Job entity by ID if the given value is not nil.
func (jahu *JobAssignmentHistoryUpdate) SetNillableJobID(id *string) *JobAssignmentHistoryUpdate {
	if id != nil {
		jahu = jahu.SetJobID(*id)
	}
	return jahu
}

// SetJob sets the "job" edge to the Job entity.
func (jahu *JobAssignmentHistoryUpdate) SetJob(j *Job) *JobAssignmentHistoryUpdate {
	return jahu.SetJobID(j.ID)
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (jahu *JobAssignmentHistoryUpdate) SetPartnerID(id string) *JobAssignmentHistoryUpdate {
	jahu.mutation.SetPartnerID(id)
	return jahu
}

// SetNillablePartnerID sets the "partner" edge to the Partner entity by ID if the given value is not nil.
func (jahu *JobAssignmentHistoryUpdate) SetNillablePartnerID(id *string) *JobAssignmentHistoryUpdate {
	if id != nil {
		jahu = jahu.SetPartnerID(*id)
	}
	return jahu
}

// SetPartner sets the "partner" edge to the Partner entity.
func (jahu *JobAssignmentHistoryUpdate) SetPartner(p *Partner) *JobAssignmentHistoryUpdate {
	return jahu.SetPartnerID(p.ID)
}

// Mutation returns the JobAssignmentHistoryMutation object of the builder.
func (jahu *JobAssignmentHistoryUpdate) Mutation() *JobAssignmentHistoryMutation {
	return jahu.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jahu *JobAssignmentHistoryUpdate) ClearJob() *JobAssignmentHistoryUpdate {
	jahu.mutation.ClearJob()
	return jahu
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (jahu *JobAssignmentHistoryUpdate) ClearPartner() *JobAssignmentHistoryUpdate {
	jahu.mutation.ClearPartner()
	return jahu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jahu *JobAssignmentHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, jahu.sqlSave, jahu.mutation, jahu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jahu *JobAssignmentHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := jahu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jahu *JobAssignmentHistoryUpdate) Exec(ctx context.Context) error {
	_, err := jahu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jahu *JobAssignmentHistoryUpdate) ExecX(ctx context.Context) {
	if err := jahu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jahu *JobAssignmentHistoryUpdate) check() error {
	if v, ok := jahu.mutation.Note(); ok {
		if err := jobassignmenthistory.NoteValidator(v); err != nil {
			return &ValidationError{Name: "Note", err: fmt.Errorf(`ent: validator failed for field "JobAssignmentHistory.Note": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (jahu *JobAssignmentHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *JobAssignmentHistoryUpdate {
	jahu.modifiers = append(jahu.modifiers, modifiers...)
	return jahu
}

func (jahu *JobAssignmentHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := jahu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(jobassignmenthistory.Table, jobassignmenthistory.Columns, sqlgraph.NewFieldSpec(jobassignmenthistory.FieldID, field.TypeString))
	if ps := jahu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jahu.mutation.Note(); ok {
		_spec.SetField(jobassignmenthistory.FieldNote, field.TypeString, value)
	}
	if jahu.mutation.NoteCleared() {
		_spec.ClearField(jobassignmenthistory.FieldNote, field.TypeString)
	}
	if jahu.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.JobTable,
			Columns: []string{jobassignmenthistory.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jahu.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.JobTable,
			Columns: []string{jobassignmenthistory.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if jahu.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.PartnerTable,
			Columns: []string{jobassignmenthistory.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jahu.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.PartnerTable,
			Columns: []string{jobassignmenthistory.PartnerColumn},
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
	_spec.AddModifiers(jahu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, jahu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobassignmenthistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jahu.mutation.done = true
	return n, nil
}

// JobAssignmentHistoryUpdateOne is the builder for updating a single JobAssignmentHistory entity.
type JobAssignmentHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *JobAssignmentHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetNote sets the "Note" field.
func (jahuo *JobAssignmentHistoryUpdateOne) SetNote(s string) *JobAssignmentHistoryUpdateOne {
	jahuo.mutation.SetNote(s)
	return jahuo
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jahuo *JobAssignmentHistoryUpdateOne) SetNillableNote(s *string) *JobAssignmentHistoryUpdateOne {
	if s != nil {
		jahuo.SetNote(*s)
	}
	return jahuo
}

// ClearNote clears the value of the "Note" field.
func (jahuo *JobAssignmentHistoryUpdateOne) ClearNote() *JobAssignmentHistoryUpdateOne {
	jahuo.mutation.ClearNote()
	return jahuo
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jahuo *JobAssignmentHistoryUpdateOne) SetJobID(id string) *JobAssignmentHistoryUpdateOne {
	jahuo.mutation.SetJobID(id)
	return jahuo
}

// SetNillableJobID sets the "job" edge to the Job entity by ID if the given value is not nil.
func (jahuo *JobAssignmentHistoryUpdateOne) SetNillableJobID(id *string) *JobAssignmentHistoryUpdateOne {
	if id != nil {
		jahuo = jahuo.SetJobID(*id)
	}
	return jahuo
}

// SetJob sets the "job" edge to the Job entity.
func (jahuo *JobAssignmentHistoryUpdateOne) SetJob(j *Job) *JobAssignmentHistoryUpdateOne {
	return jahuo.SetJobID(j.ID)
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (jahuo *JobAssignmentHistoryUpdateOne) SetPartnerID(id string) *JobAssignmentHistoryUpdateOne {
	jahuo.mutation.SetPartnerID(id)
	return jahuo
}

// SetNillablePartnerID sets the "partner" edge to the Partner entity by ID if the given value is not nil.
func (jahuo *JobAssignmentHistoryUpdateOne) SetNillablePartnerID(id *string) *JobAssignmentHistoryUpdateOne {
	if id != nil {
		jahuo = jahuo.SetPartnerID(*id)
	}
	return jahuo
}

// SetPartner sets the "partner" edge to the Partner entity.
func (jahuo *JobAssignmentHistoryUpdateOne) SetPartner(p *Partner) *JobAssignmentHistoryUpdateOne {
	return jahuo.SetPartnerID(p.ID)
}

// Mutation returns the JobAssignmentHistoryMutation object of the builder.
func (jahuo *JobAssignmentHistoryUpdateOne) Mutation() *JobAssignmentHistoryMutation {
	return jahuo.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jahuo *JobAssignmentHistoryUpdateOne) ClearJob() *JobAssignmentHistoryUpdateOne {
	jahuo.mutation.ClearJob()
	return jahuo
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (jahuo *JobAssignmentHistoryUpdateOne) ClearPartner() *JobAssignmentHistoryUpdateOne {
	jahuo.mutation.ClearPartner()
	return jahuo
}

// Where appends a list predicates to the JobAssignmentHistoryUpdate builder.
func (jahuo *JobAssignmentHistoryUpdateOne) Where(ps ...predicate.JobAssignmentHistory) *JobAssignmentHistoryUpdateOne {
	jahuo.mutation.Where(ps...)
	return jahuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jahuo *JobAssignmentHistoryUpdateOne) Select(field string, fields ...string) *JobAssignmentHistoryUpdateOne {
	jahuo.fields = append([]string{field}, fields...)
	return jahuo
}

// Save executes the query and returns the updated JobAssignmentHistory entity.
func (jahuo *JobAssignmentHistoryUpdateOne) Save(ctx context.Context) (*JobAssignmentHistory, error) {
	return withHooks(ctx, jahuo.sqlSave, jahuo.mutation, jahuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jahuo *JobAssignmentHistoryUpdateOne) SaveX(ctx context.Context) *JobAssignmentHistory {
	node, err := jahuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jahuo *JobAssignmentHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := jahuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jahuo *JobAssignmentHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := jahuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jahuo *JobAssignmentHistoryUpdateOne) check() error {
	if v, ok := jahuo.mutation.Note(); ok {
		if err := jobassignmenthistory.NoteValidator(v); err != nil {
			return &ValidationError{Name: "Note", err: fmt.Errorf(`ent: validator failed for field "JobAssignmentHistory.Note": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (jahuo *JobAssignmentHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *JobAssignmentHistoryUpdateOne {
	jahuo.modifiers = append(jahuo.modifiers, modifiers...)
	return jahuo
}

func (jahuo *JobAssignmentHistoryUpdateOne) sqlSave(ctx context.Context) (_node *JobAssignmentHistory, err error) {
	if err := jahuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(jobassignmenthistory.Table, jobassignmenthistory.Columns, sqlgraph.NewFieldSpec(jobassignmenthistory.FieldID, field.TypeString))
	id, ok := jahuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobAssignmentHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jahuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobassignmenthistory.FieldID)
		for _, f := range fields {
			if !jobassignmenthistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobassignmenthistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jahuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jahuo.mutation.Note(); ok {
		_spec.SetField(jobassignmenthistory.FieldNote, field.TypeString, value)
	}
	if jahuo.mutation.NoteCleared() {
		_spec.ClearField(jobassignmenthistory.FieldNote, field.TypeString)
	}
	if jahuo.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.JobTable,
			Columns: []string{jobassignmenthistory.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jahuo.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.JobTable,
			Columns: []string{jobassignmenthistory.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if jahuo.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.PartnerTable,
			Columns: []string{jobassignmenthistory.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jahuo.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobassignmenthistory.PartnerTable,
			Columns: []string{jobassignmenthistory.PartnerColumn},
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
	_spec.AddModifiers(jahuo.modifiers...)
	_node = &JobAssignmentHistory{config: jahuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jahuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobassignmenthistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jahuo.mutation.done = true
	return _node, nil
}
