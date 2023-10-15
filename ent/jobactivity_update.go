// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/apiuser"
	"roofix/ent/job"
	"roofix/ent/jobactivity"
	"roofix/ent/predicate"
	"roofix/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobActivityUpdate is the builder for updating JobActivity entities.
type JobActivityUpdate struct {
	config
	hooks     []Hook
	mutation  *JobActivityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the JobActivityUpdate builder.
func (jau *JobActivityUpdate) Where(ps ...predicate.JobActivity) *JobActivityUpdate {
	jau.mutation.Where(ps...)
	return jau
}

// SetDescription sets the "description" field.
func (jau *JobActivityUpdate) SetDescription(s string) *JobActivityUpdate {
	jau.mutation.SetDescription(s)
	return jau
}

// SetRaw sets the "raw" field.
func (jau *JobActivityUpdate) SetRaw(m map[string]interface{}) *JobActivityUpdate {
	jau.mutation.SetRaw(m)
	return jau
}

// ClearRaw clears the value of the "raw" field.
func (jau *JobActivityUpdate) ClearRaw() *JobActivityUpdate {
	jau.mutation.ClearRaw()
	return jau
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jau *JobActivityUpdate) SetJobID(id string) *JobActivityUpdate {
	jau.mutation.SetJobID(id)
	return jau
}

// SetNillableJobID sets the "job" edge to the Job entity by ID if the given value is not nil.
func (jau *JobActivityUpdate) SetNillableJobID(id *string) *JobActivityUpdate {
	if id != nil {
		jau = jau.SetJobID(*id)
	}
	return jau
}

// SetJob sets the "job" edge to the Job entity.
func (jau *JobActivityUpdate) SetJob(j *Job) *JobActivityUpdate {
	return jau.SetJobID(j.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (jau *JobActivityUpdate) SetCreatorID(id string) *JobActivityUpdate {
	jau.mutation.SetCreatorID(id)
	return jau
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (jau *JobActivityUpdate) SetNillableCreatorID(id *string) *JobActivityUpdate {
	if id != nil {
		jau = jau.SetCreatorID(*id)
	}
	return jau
}

// SetCreator sets the "creator" edge to the User entity.
func (jau *JobActivityUpdate) SetCreator(u *User) *JobActivityUpdate {
	return jau.SetCreatorID(u.ID)
}

// SetCreatorAPIID sets the "creator_api" edge to the ApiUser entity by ID.
func (jau *JobActivityUpdate) SetCreatorAPIID(id string) *JobActivityUpdate {
	jau.mutation.SetCreatorAPIID(id)
	return jau
}

// SetNillableCreatorAPIID sets the "creator_api" edge to the ApiUser entity by ID if the given value is not nil.
func (jau *JobActivityUpdate) SetNillableCreatorAPIID(id *string) *JobActivityUpdate {
	if id != nil {
		jau = jau.SetCreatorAPIID(*id)
	}
	return jau
}

// SetCreatorAPI sets the "creator_api" edge to the ApiUser entity.
func (jau *JobActivityUpdate) SetCreatorAPI(a *ApiUser) *JobActivityUpdate {
	return jau.SetCreatorAPIID(a.ID)
}

// Mutation returns the JobActivityMutation object of the builder.
func (jau *JobActivityUpdate) Mutation() *JobActivityMutation {
	return jau.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jau *JobActivityUpdate) ClearJob() *JobActivityUpdate {
	jau.mutation.ClearJob()
	return jau
}

// ClearCreator clears the "creator" edge to the User entity.
func (jau *JobActivityUpdate) ClearCreator() *JobActivityUpdate {
	jau.mutation.ClearCreator()
	return jau
}

// ClearCreatorAPI clears the "creator_api" edge to the ApiUser entity.
func (jau *JobActivityUpdate) ClearCreatorAPI() *JobActivityUpdate {
	jau.mutation.ClearCreatorAPI()
	return jau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jau *JobActivityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, jau.sqlSave, jau.mutation, jau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jau *JobActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := jau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jau *JobActivityUpdate) Exec(ctx context.Context) error {
	_, err := jau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jau *JobActivityUpdate) ExecX(ctx context.Context) {
	if err := jau.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (jau *JobActivityUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *JobActivityUpdate {
	jau.modifiers = append(jau.modifiers, modifiers...)
	return jau
}

func (jau *JobActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobactivity.Table, jobactivity.Columns, sqlgraph.NewFieldSpec(jobactivity.FieldID, field.TypeString))
	if ps := jau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jau.mutation.Description(); ok {
		_spec.SetField(jobactivity.FieldDescription, field.TypeString, value)
	}
	if value, ok := jau.mutation.Raw(); ok {
		_spec.SetField(jobactivity.FieldRaw, field.TypeJSON, value)
	}
	if jau.mutation.RawCleared() {
		_spec.ClearField(jobactivity.FieldRaw, field.TypeJSON)
	}
	if jau.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.JobTable,
			Columns: []string{jobactivity.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jau.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.JobTable,
			Columns: []string{jobactivity.JobColumn},
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
	if jau.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorTable,
			Columns: []string{jobactivity.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jau.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorTable,
			Columns: []string{jobactivity.CreatorColumn},
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
	if jau.mutation.CreatorAPICleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorAPITable,
			Columns: []string{jobactivity.CreatorAPIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apiuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jau.mutation.CreatorAPIIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorAPITable,
			Columns: []string{jobactivity.CreatorAPIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apiuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(jau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, jau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobactivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jau.mutation.done = true
	return n, nil
}

// JobActivityUpdateOne is the builder for updating a single JobActivity entity.
type JobActivityUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *JobActivityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetDescription sets the "description" field.
func (jauo *JobActivityUpdateOne) SetDescription(s string) *JobActivityUpdateOne {
	jauo.mutation.SetDescription(s)
	return jauo
}

// SetRaw sets the "raw" field.
func (jauo *JobActivityUpdateOne) SetRaw(m map[string]interface{}) *JobActivityUpdateOne {
	jauo.mutation.SetRaw(m)
	return jauo
}

// ClearRaw clears the value of the "raw" field.
func (jauo *JobActivityUpdateOne) ClearRaw() *JobActivityUpdateOne {
	jauo.mutation.ClearRaw()
	return jauo
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jauo *JobActivityUpdateOne) SetJobID(id string) *JobActivityUpdateOne {
	jauo.mutation.SetJobID(id)
	return jauo
}

// SetNillableJobID sets the "job" edge to the Job entity by ID if the given value is not nil.
func (jauo *JobActivityUpdateOne) SetNillableJobID(id *string) *JobActivityUpdateOne {
	if id != nil {
		jauo = jauo.SetJobID(*id)
	}
	return jauo
}

// SetJob sets the "job" edge to the Job entity.
func (jauo *JobActivityUpdateOne) SetJob(j *Job) *JobActivityUpdateOne {
	return jauo.SetJobID(j.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (jauo *JobActivityUpdateOne) SetCreatorID(id string) *JobActivityUpdateOne {
	jauo.mutation.SetCreatorID(id)
	return jauo
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (jauo *JobActivityUpdateOne) SetNillableCreatorID(id *string) *JobActivityUpdateOne {
	if id != nil {
		jauo = jauo.SetCreatorID(*id)
	}
	return jauo
}

// SetCreator sets the "creator" edge to the User entity.
func (jauo *JobActivityUpdateOne) SetCreator(u *User) *JobActivityUpdateOne {
	return jauo.SetCreatorID(u.ID)
}

// SetCreatorAPIID sets the "creator_api" edge to the ApiUser entity by ID.
func (jauo *JobActivityUpdateOne) SetCreatorAPIID(id string) *JobActivityUpdateOne {
	jauo.mutation.SetCreatorAPIID(id)
	return jauo
}

// SetNillableCreatorAPIID sets the "creator_api" edge to the ApiUser entity by ID if the given value is not nil.
func (jauo *JobActivityUpdateOne) SetNillableCreatorAPIID(id *string) *JobActivityUpdateOne {
	if id != nil {
		jauo = jauo.SetCreatorAPIID(*id)
	}
	return jauo
}

// SetCreatorAPI sets the "creator_api" edge to the ApiUser entity.
func (jauo *JobActivityUpdateOne) SetCreatorAPI(a *ApiUser) *JobActivityUpdateOne {
	return jauo.SetCreatorAPIID(a.ID)
}

// Mutation returns the JobActivityMutation object of the builder.
func (jauo *JobActivityUpdateOne) Mutation() *JobActivityMutation {
	return jauo.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jauo *JobActivityUpdateOne) ClearJob() *JobActivityUpdateOne {
	jauo.mutation.ClearJob()
	return jauo
}

// ClearCreator clears the "creator" edge to the User entity.
func (jauo *JobActivityUpdateOne) ClearCreator() *JobActivityUpdateOne {
	jauo.mutation.ClearCreator()
	return jauo
}

// ClearCreatorAPI clears the "creator_api" edge to the ApiUser entity.
func (jauo *JobActivityUpdateOne) ClearCreatorAPI() *JobActivityUpdateOne {
	jauo.mutation.ClearCreatorAPI()
	return jauo
}

// Where appends a list predicates to the JobActivityUpdate builder.
func (jauo *JobActivityUpdateOne) Where(ps ...predicate.JobActivity) *JobActivityUpdateOne {
	jauo.mutation.Where(ps...)
	return jauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jauo *JobActivityUpdateOne) Select(field string, fields ...string) *JobActivityUpdateOne {
	jauo.fields = append([]string{field}, fields...)
	return jauo
}

// Save executes the query and returns the updated JobActivity entity.
func (jauo *JobActivityUpdateOne) Save(ctx context.Context) (*JobActivity, error) {
	return withHooks(ctx, jauo.sqlSave, jauo.mutation, jauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jauo *JobActivityUpdateOne) SaveX(ctx context.Context) *JobActivity {
	node, err := jauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jauo *JobActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := jauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jauo *JobActivityUpdateOne) ExecX(ctx context.Context) {
	if err := jauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (jauo *JobActivityUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *JobActivityUpdateOne {
	jauo.modifiers = append(jauo.modifiers, modifiers...)
	return jauo
}

func (jauo *JobActivityUpdateOne) sqlSave(ctx context.Context) (_node *JobActivity, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobactivity.Table, jobactivity.Columns, sqlgraph.NewFieldSpec(jobactivity.FieldID, field.TypeString))
	id, ok := jauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobActivity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobactivity.FieldID)
		for _, f := range fields {
			if !jobactivity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobactivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jauo.mutation.Description(); ok {
		_spec.SetField(jobactivity.FieldDescription, field.TypeString, value)
	}
	if value, ok := jauo.mutation.Raw(); ok {
		_spec.SetField(jobactivity.FieldRaw, field.TypeJSON, value)
	}
	if jauo.mutation.RawCleared() {
		_spec.ClearField(jobactivity.FieldRaw, field.TypeJSON)
	}
	if jauo.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.JobTable,
			Columns: []string{jobactivity.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jauo.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.JobTable,
			Columns: []string{jobactivity.JobColumn},
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
	if jauo.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorTable,
			Columns: []string{jobactivity.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jauo.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorTable,
			Columns: []string{jobactivity.CreatorColumn},
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
	if jauo.mutation.CreatorAPICleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorAPITable,
			Columns: []string{jobactivity.CreatorAPIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apiuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jauo.mutation.CreatorAPIIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobactivity.CreatorAPITable,
			Columns: []string{jobactivity.CreatorAPIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apiuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(jauo.modifiers...)
	_node = &JobActivity{config: jauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobactivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jauo.mutation.done = true
	return _node, nil
}