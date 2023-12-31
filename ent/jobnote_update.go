// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/job"
	"roofix/ent/jobnote"
	"roofix/ent/partner"
	"roofix/ent/predicate"
	"roofix/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobNoteUpdate is the builder for updating JobNote entities.
type JobNoteUpdate struct {
	config
	hooks     []Hook
	mutation  *JobNoteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the JobNoteUpdate builder.
func (jnu *JobNoteUpdate) Where(ps ...predicate.JobNote) *JobNoteUpdate {
	jnu.mutation.Where(ps...)
	return jnu
}

// SetUpdatedAt sets the "updated_at" field.
func (jnu *JobNoteUpdate) SetUpdatedAt(t time.Time) *JobNoteUpdate {
	jnu.mutation.SetUpdatedAt(t)
	return jnu
}

// SetNote sets the "note" field.
func (jnu *JobNoteUpdate) SetNote(s string) *JobNoteUpdate {
	jnu.mutation.SetNote(s)
	return jnu
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jnu *JobNoteUpdate) SetJobID(id string) *JobNoteUpdate {
	jnu.mutation.SetJobID(id)
	return jnu
}

// SetJob sets the "job" edge to the Job entity.
func (jnu *JobNoteUpdate) SetJob(j *Job) *JobNoteUpdate {
	return jnu.SetJobID(j.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (jnu *JobNoteUpdate) SetUserID(id string) *JobNoteUpdate {
	jnu.mutation.SetUserID(id)
	return jnu
}

// SetUser sets the "user" edge to the User entity.
func (jnu *JobNoteUpdate) SetUser(u *User) *JobNoteUpdate {
	return jnu.SetUserID(u.ID)
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (jnu *JobNoteUpdate) SetPartnerID(id string) *JobNoteUpdate {
	jnu.mutation.SetPartnerID(id)
	return jnu
}

// SetNillablePartnerID sets the "partner" edge to the Partner entity by ID if the given value is not nil.
func (jnu *JobNoteUpdate) SetNillablePartnerID(id *string) *JobNoteUpdate {
	if id != nil {
		jnu = jnu.SetPartnerID(*id)
	}
	return jnu
}

// SetPartner sets the "partner" edge to the Partner entity.
func (jnu *JobNoteUpdate) SetPartner(p *Partner) *JobNoteUpdate {
	return jnu.SetPartnerID(p.ID)
}

// Mutation returns the JobNoteMutation object of the builder.
func (jnu *JobNoteUpdate) Mutation() *JobNoteMutation {
	return jnu.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jnu *JobNoteUpdate) ClearJob() *JobNoteUpdate {
	jnu.mutation.ClearJob()
	return jnu
}

// ClearUser clears the "user" edge to the User entity.
func (jnu *JobNoteUpdate) ClearUser() *JobNoteUpdate {
	jnu.mutation.ClearUser()
	return jnu
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (jnu *JobNoteUpdate) ClearPartner() *JobNoteUpdate {
	jnu.mutation.ClearPartner()
	return jnu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jnu *JobNoteUpdate) Save(ctx context.Context) (int, error) {
	jnu.defaults()
	return withHooks(ctx, jnu.sqlSave, jnu.mutation, jnu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jnu *JobNoteUpdate) SaveX(ctx context.Context) int {
	affected, err := jnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jnu *JobNoteUpdate) Exec(ctx context.Context) error {
	_, err := jnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jnu *JobNoteUpdate) ExecX(ctx context.Context) {
	if err := jnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jnu *JobNoteUpdate) defaults() {
	if _, ok := jnu.mutation.UpdatedAt(); !ok {
		v := jobnote.UpdateDefaultUpdatedAt()
		jnu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jnu *JobNoteUpdate) check() error {
	if v, ok := jnu.mutation.Note(); ok {
		if err := jobnote.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "JobNote.note": %w`, err)}
		}
	}
	if _, ok := jnu.mutation.JobID(); jnu.mutation.JobCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "JobNote.job"`)
	}
	if _, ok := jnu.mutation.UserID(); jnu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "JobNote.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (jnu *JobNoteUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *JobNoteUpdate {
	jnu.modifiers = append(jnu.modifiers, modifiers...)
	return jnu
}

func (jnu *JobNoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := jnu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(jobnote.Table, jobnote.Columns, sqlgraph.NewFieldSpec(jobnote.FieldID, field.TypeString))
	if ps := jnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jnu.mutation.UpdatedAt(); ok {
		_spec.SetField(jobnote.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := jnu.mutation.Note(); ok {
		_spec.SetField(jobnote.FieldNote, field.TypeString, value)
	}
	if jnu.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.JobTable,
			Columns: []string{jobnote.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jnu.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.JobTable,
			Columns: []string{jobnote.JobColumn},
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
	if jnu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.UserTable,
			Columns: []string{jobnote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jnu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.UserTable,
			Columns: []string{jobnote.UserColumn},
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
	if jnu.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.PartnerTable,
			Columns: []string{jobnote.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jnu.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.PartnerTable,
			Columns: []string{jobnote.PartnerColumn},
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
	_spec.AddModifiers(jnu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, jnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobnote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jnu.mutation.done = true
	return n, nil
}

// JobNoteUpdateOne is the builder for updating a single JobNote entity.
type JobNoteUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *JobNoteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (jnuo *JobNoteUpdateOne) SetUpdatedAt(t time.Time) *JobNoteUpdateOne {
	jnuo.mutation.SetUpdatedAt(t)
	return jnuo
}

// SetNote sets the "note" field.
func (jnuo *JobNoteUpdateOne) SetNote(s string) *JobNoteUpdateOne {
	jnuo.mutation.SetNote(s)
	return jnuo
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jnuo *JobNoteUpdateOne) SetJobID(id string) *JobNoteUpdateOne {
	jnuo.mutation.SetJobID(id)
	return jnuo
}

// SetJob sets the "job" edge to the Job entity.
func (jnuo *JobNoteUpdateOne) SetJob(j *Job) *JobNoteUpdateOne {
	return jnuo.SetJobID(j.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (jnuo *JobNoteUpdateOne) SetUserID(id string) *JobNoteUpdateOne {
	jnuo.mutation.SetUserID(id)
	return jnuo
}

// SetUser sets the "user" edge to the User entity.
func (jnuo *JobNoteUpdateOne) SetUser(u *User) *JobNoteUpdateOne {
	return jnuo.SetUserID(u.ID)
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (jnuo *JobNoteUpdateOne) SetPartnerID(id string) *JobNoteUpdateOne {
	jnuo.mutation.SetPartnerID(id)
	return jnuo
}

// SetNillablePartnerID sets the "partner" edge to the Partner entity by ID if the given value is not nil.
func (jnuo *JobNoteUpdateOne) SetNillablePartnerID(id *string) *JobNoteUpdateOne {
	if id != nil {
		jnuo = jnuo.SetPartnerID(*id)
	}
	return jnuo
}

// SetPartner sets the "partner" edge to the Partner entity.
func (jnuo *JobNoteUpdateOne) SetPartner(p *Partner) *JobNoteUpdateOne {
	return jnuo.SetPartnerID(p.ID)
}

// Mutation returns the JobNoteMutation object of the builder.
func (jnuo *JobNoteUpdateOne) Mutation() *JobNoteMutation {
	return jnuo.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jnuo *JobNoteUpdateOne) ClearJob() *JobNoteUpdateOne {
	jnuo.mutation.ClearJob()
	return jnuo
}

// ClearUser clears the "user" edge to the User entity.
func (jnuo *JobNoteUpdateOne) ClearUser() *JobNoteUpdateOne {
	jnuo.mutation.ClearUser()
	return jnuo
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (jnuo *JobNoteUpdateOne) ClearPartner() *JobNoteUpdateOne {
	jnuo.mutation.ClearPartner()
	return jnuo
}

// Where appends a list predicates to the JobNoteUpdate builder.
func (jnuo *JobNoteUpdateOne) Where(ps ...predicate.JobNote) *JobNoteUpdateOne {
	jnuo.mutation.Where(ps...)
	return jnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jnuo *JobNoteUpdateOne) Select(field string, fields ...string) *JobNoteUpdateOne {
	jnuo.fields = append([]string{field}, fields...)
	return jnuo
}

// Save executes the query and returns the updated JobNote entity.
func (jnuo *JobNoteUpdateOne) Save(ctx context.Context) (*JobNote, error) {
	jnuo.defaults()
	return withHooks(ctx, jnuo.sqlSave, jnuo.mutation, jnuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jnuo *JobNoteUpdateOne) SaveX(ctx context.Context) *JobNote {
	node, err := jnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jnuo *JobNoteUpdateOne) Exec(ctx context.Context) error {
	_, err := jnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jnuo *JobNoteUpdateOne) ExecX(ctx context.Context) {
	if err := jnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jnuo *JobNoteUpdateOne) defaults() {
	if _, ok := jnuo.mutation.UpdatedAt(); !ok {
		v := jobnote.UpdateDefaultUpdatedAt()
		jnuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jnuo *JobNoteUpdateOne) check() error {
	if v, ok := jnuo.mutation.Note(); ok {
		if err := jobnote.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "JobNote.note": %w`, err)}
		}
	}
	if _, ok := jnuo.mutation.JobID(); jnuo.mutation.JobCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "JobNote.job"`)
	}
	if _, ok := jnuo.mutation.UserID(); jnuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "JobNote.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (jnuo *JobNoteUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *JobNoteUpdateOne {
	jnuo.modifiers = append(jnuo.modifiers, modifiers...)
	return jnuo
}

func (jnuo *JobNoteUpdateOne) sqlSave(ctx context.Context) (_node *JobNote, err error) {
	if err := jnuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(jobnote.Table, jobnote.Columns, sqlgraph.NewFieldSpec(jobnote.FieldID, field.TypeString))
	id, ok := jnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobNote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobnote.FieldID)
		for _, f := range fields {
			if !jobnote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobnote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jnuo.mutation.UpdatedAt(); ok {
		_spec.SetField(jobnote.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := jnuo.mutation.Note(); ok {
		_spec.SetField(jobnote.FieldNote, field.TypeString, value)
	}
	if jnuo.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.JobTable,
			Columns: []string{jobnote.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jnuo.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.JobTable,
			Columns: []string{jobnote.JobColumn},
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
	if jnuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.UserTable,
			Columns: []string{jobnote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jnuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.UserTable,
			Columns: []string{jobnote.UserColumn},
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
	if jnuo.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.PartnerTable,
			Columns: []string{jobnote.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jnuo.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobnote.PartnerTable,
			Columns: []string{jobnote.PartnerColumn},
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
	_spec.AddModifiers(jnuo.modifiers...)
	_node = &JobNote{config: jnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobnote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jnuo.mutation.done = true
	return _node, nil
}
