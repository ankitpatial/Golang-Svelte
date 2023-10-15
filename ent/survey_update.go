// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/predicate"
	"roofix/ent/survey"
	"roofix/ent/surveyprogress"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SurveyUpdate is the builder for updating Survey entities.
type SurveyUpdate struct {
	config
	hooks     []Hook
	mutation  *SurveyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SurveyUpdate builder.
func (su *SurveyUpdate) Where(ps ...predicate.Survey) *SurveyUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SurveyUpdate) SetUpdatedAt(t time.Time) *SurveyUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetUntil sets the "until" field.
func (su *SurveyUpdate) SetUntil(t time.Time) *SurveyUpdate {
	su.mutation.SetUntil(t)
	return su
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableUntil(t *time.Time) *SurveyUpdate {
	if t != nil {
		su.SetUntil(*t)
	}
	return su
}

// ClearUntil clears the value of the "until" field.
func (su *SurveyUpdate) ClearUntil() *SurveyUpdate {
	su.mutation.ClearUntil()
	return su
}

// SetName sets the "name" field.
func (su *SurveyUpdate) SetName(s string) *SurveyUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableName(s *string) *SurveyUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// ClearName clears the value of the "name" field.
func (su *SurveyUpdate) ClearName() *SurveyUpdate {
	su.mutation.ClearName()
	return su
}

// SetAddress sets the "address" field.
func (su *SurveyUpdate) SetAddress(s string) *SurveyUpdate {
	su.mutation.SetAddress(s)
	return su
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableAddress(s *string) *SurveyUpdate {
	if s != nil {
		su.SetAddress(*s)
	}
	return su
}

// ClearAddress clears the value of the "address" field.
func (su *SurveyUpdate) ClearAddress() *SurveyUpdate {
	su.mutation.ClearAddress()
	return su
}

// SetPhone sets the "phone" field.
func (su *SurveyUpdate) SetPhone(s string) *SurveyUpdate {
	su.mutation.SetPhone(s)
	return su
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (su *SurveyUpdate) SetNillablePhone(s *string) *SurveyUpdate {
	if s != nil {
		su.SetPhone(*s)
	}
	return su
}

// ClearPhone clears the value of the "phone" field.
func (su *SurveyUpdate) ClearPhone() *SurveyUpdate {
	su.mutation.ClearPhone()
	return su
}

// SetNotes sets the "notes" field.
func (su *SurveyUpdate) SetNotes(s string) *SurveyUpdate {
	su.mutation.SetNotes(s)
	return su
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableNotes(s *string) *SurveyUpdate {
	if s != nil {
		su.SetNotes(*s)
	}
	return su
}

// ClearNotes clears the value of the "notes" field.
func (su *SurveyUpdate) ClearNotes() *SurveyUpdate {
	su.mutation.ClearNotes()
	return su
}

// SetStatus sets the "status" field.
func (su *SurveyUpdate) SetStatus(es enum.SurveyStatus) *SurveyUpdate {
	su.mutation.SetStatus(es)
	return su
}

// SetProgress sets the "progress" field.
func (su *SurveyUpdate) SetProgress(ep enum.SurveyProgress) *SurveyUpdate {
	su.mutation.SetProgress(ep)
	return su
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableProgress(ep *enum.SurveyProgress) *SurveyUpdate {
	if ep != nil {
		su.SetProgress(*ep)
	}
	return su
}

// ClearProgress clears the value of the "progress" field.
func (su *SurveyUpdate) ClearProgress() *SurveyUpdate {
	su.mutation.ClearProgress()
	return su
}

// SetProgressAt sets the "progress_at" field.
func (su *SurveyUpdate) SetProgressAt(t time.Time) *SurveyUpdate {
	su.mutation.SetProgressAt(t)
	return su
}

// SetNillableProgressAt sets the "progress_at" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableProgressAt(t *time.Time) *SurveyUpdate {
	if t != nil {
		su.SetProgressAt(*t)
	}
	return su
}

// ClearProgressAt clears the value of the "progress_at" field.
func (su *SurveyUpdate) ClearProgressAt() *SurveyUpdate {
	su.mutation.ClearProgressAt()
	return su
}

// SetProgressFlagAt sets the "progress_flag_at" field.
func (su *SurveyUpdate) SetProgressFlagAt(t time.Time) *SurveyUpdate {
	su.mutation.SetProgressFlagAt(t)
	return su
}

// SetNillableProgressFlagAt sets the "progress_flag_at" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableProgressFlagAt(t *time.Time) *SurveyUpdate {
	if t != nil {
		su.SetProgressFlagAt(*t)
	}
	return su
}

// ClearProgressFlagAt clears the value of the "progress_flag_at" field.
func (su *SurveyUpdate) ClearProgressFlagAt() *SurveyUpdate {
	su.mutation.ClearProgressFlagAt()
	return su
}

// AddProgressHistoryIDs adds the "progress_history" edge to the SurveyProgress entity by IDs.
func (su *SurveyUpdate) AddProgressHistoryIDs(ids ...string) *SurveyUpdate {
	su.mutation.AddProgressHistoryIDs(ids...)
	return su
}

// AddProgressHistory adds the "progress_history" edges to the SurveyProgress entity.
func (su *SurveyUpdate) AddProgressHistory(s ...*SurveyProgress) *SurveyUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddProgressHistoryIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (su *SurveyUpdate) SetCreatedByID(id string) *SurveyUpdate {
	su.mutation.SetCreatedByID(id)
	return su
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (su *SurveyUpdate) SetNillableCreatedByID(id *string) *SurveyUpdate {
	if id != nil {
		su = su.SetCreatedByID(*id)
	}
	return su
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (su *SurveyUpdate) SetCreatedBy(u *User) *SurveyUpdate {
	return su.SetCreatedByID(u.ID)
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (su *SurveyUpdate) SetPartnerID(id string) *SurveyUpdate {
	su.mutation.SetPartnerID(id)
	return su
}

// SetNillablePartnerID sets the "partner" edge to the Partner entity by ID if the given value is not nil.
func (su *SurveyUpdate) SetNillablePartnerID(id *string) *SurveyUpdate {
	if id != nil {
		su = su.SetPartnerID(*id)
	}
	return su
}

// SetPartner sets the "partner" edge to the Partner entity.
func (su *SurveyUpdate) SetPartner(p *Partner) *SurveyUpdate {
	return su.SetPartnerID(p.ID)
}

// Mutation returns the SurveyMutation object of the builder.
func (su *SurveyUpdate) Mutation() *SurveyMutation {
	return su.mutation
}

// ClearProgressHistory clears all "progress_history" edges to the SurveyProgress entity.
func (su *SurveyUpdate) ClearProgressHistory() *SurveyUpdate {
	su.mutation.ClearProgressHistory()
	return su
}

// RemoveProgressHistoryIDs removes the "progress_history" edge to SurveyProgress entities by IDs.
func (su *SurveyUpdate) RemoveProgressHistoryIDs(ids ...string) *SurveyUpdate {
	su.mutation.RemoveProgressHistoryIDs(ids...)
	return su
}

// RemoveProgressHistory removes "progress_history" edges to SurveyProgress entities.
func (su *SurveyUpdate) RemoveProgressHistory(s ...*SurveyProgress) *SurveyUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveProgressHistoryIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (su *SurveyUpdate) ClearCreatedBy() *SurveyUpdate {
	su.mutation.ClearCreatedBy()
	return su
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (su *SurveyUpdate) ClearPartner() *SurveyUpdate {
	su.mutation.ClearPartner()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SurveyUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SurveyUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SurveyUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SurveyUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SurveyUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := survey.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SurveyUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := survey.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Survey.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.Address(); ok {
		if err := survey.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Survey.address": %w`, err)}
		}
	}
	if v, ok := su.mutation.Phone(); ok {
		if err := survey.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Survey.phone": %w`, err)}
		}
	}
	if v, ok := su.mutation.Notes(); ok {
		if err := survey.NotesValidator(v); err != nil {
			return &ValidationError{Name: "notes", err: fmt.Errorf(`ent: validator failed for field "Survey.notes": %w`, err)}
		}
	}
	if v, ok := su.mutation.Status(); ok {
		if err := survey.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Survey.status": %w`, err)}
		}
	}
	if v, ok := su.mutation.Progress(); ok {
		if err := survey.ProgressValidator(v); err != nil {
			return &ValidationError{Name: "progress", err: fmt.Errorf(`ent: validator failed for field "Survey.progress": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SurveyUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SurveyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(survey.Table, survey.Columns, sqlgraph.NewFieldSpec(survey.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(survey.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Until(); ok {
		_spec.SetField(survey.FieldUntil, field.TypeTime, value)
	}
	if su.mutation.UntilCleared() {
		_spec.ClearField(survey.FieldUntil, field.TypeTime)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(survey.FieldName, field.TypeString, value)
	}
	if su.mutation.NameCleared() {
		_spec.ClearField(survey.FieldName, field.TypeString)
	}
	if value, ok := su.mutation.Address(); ok {
		_spec.SetField(survey.FieldAddress, field.TypeString, value)
	}
	if su.mutation.AddressCleared() {
		_spec.ClearField(survey.FieldAddress, field.TypeString)
	}
	if value, ok := su.mutation.Phone(); ok {
		_spec.SetField(survey.FieldPhone, field.TypeString, value)
	}
	if su.mutation.PhoneCleared() {
		_spec.ClearField(survey.FieldPhone, field.TypeString)
	}
	if value, ok := su.mutation.Notes(); ok {
		_spec.SetField(survey.FieldNotes, field.TypeString, value)
	}
	if su.mutation.NotesCleared() {
		_spec.ClearField(survey.FieldNotes, field.TypeString)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(survey.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := su.mutation.Progress(); ok {
		_spec.SetField(survey.FieldProgress, field.TypeEnum, value)
	}
	if su.mutation.ProgressCleared() {
		_spec.ClearField(survey.FieldProgress, field.TypeEnum)
	}
	if value, ok := su.mutation.ProgressAt(); ok {
		_spec.SetField(survey.FieldProgressAt, field.TypeTime, value)
	}
	if su.mutation.ProgressAtCleared() {
		_spec.ClearField(survey.FieldProgressAt, field.TypeTime)
	}
	if value, ok := su.mutation.ProgressFlagAt(); ok {
		_spec.SetField(survey.FieldProgressFlagAt, field.TypeTime, value)
	}
	if su.mutation.ProgressFlagAtCleared() {
		_spec.ClearField(survey.FieldProgressFlagAt, field.TypeTime)
	}
	if su.mutation.ProgressHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.ProgressHistoryTable,
			Columns: []string{survey.ProgressHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyprogress.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedProgressHistoryIDs(); len(nodes) > 0 && !su.mutation.ProgressHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.ProgressHistoryTable,
			Columns: []string{survey.ProgressHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyprogress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProgressHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.ProgressHistoryTable,
			Columns: []string{survey.ProgressHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyprogress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.CreatedByTable,
			Columns: []string{survey.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.CreatedByTable,
			Columns: []string{survey.CreatedByColumn},
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
	if su.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.PartnerTable,
			Columns: []string{survey.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.PartnerTable,
			Columns: []string{survey.PartnerColumn},
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
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{survey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SurveyUpdateOne is the builder for updating a single Survey entity.
type SurveyUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SurveyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SurveyUpdateOne) SetUpdatedAt(t time.Time) *SurveyUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetUntil sets the "until" field.
func (suo *SurveyUpdateOne) SetUntil(t time.Time) *SurveyUpdateOne {
	suo.mutation.SetUntil(t)
	return suo
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableUntil(t *time.Time) *SurveyUpdateOne {
	if t != nil {
		suo.SetUntil(*t)
	}
	return suo
}

// ClearUntil clears the value of the "until" field.
func (suo *SurveyUpdateOne) ClearUntil() *SurveyUpdateOne {
	suo.mutation.ClearUntil()
	return suo
}

// SetName sets the "name" field.
func (suo *SurveyUpdateOne) SetName(s string) *SurveyUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableName(s *string) *SurveyUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// ClearName clears the value of the "name" field.
func (suo *SurveyUpdateOne) ClearName() *SurveyUpdateOne {
	suo.mutation.ClearName()
	return suo
}

// SetAddress sets the "address" field.
func (suo *SurveyUpdateOne) SetAddress(s string) *SurveyUpdateOne {
	suo.mutation.SetAddress(s)
	return suo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableAddress(s *string) *SurveyUpdateOne {
	if s != nil {
		suo.SetAddress(*s)
	}
	return suo
}

// ClearAddress clears the value of the "address" field.
func (suo *SurveyUpdateOne) ClearAddress() *SurveyUpdateOne {
	suo.mutation.ClearAddress()
	return suo
}

// SetPhone sets the "phone" field.
func (suo *SurveyUpdateOne) SetPhone(s string) *SurveyUpdateOne {
	suo.mutation.SetPhone(s)
	return suo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillablePhone(s *string) *SurveyUpdateOne {
	if s != nil {
		suo.SetPhone(*s)
	}
	return suo
}

// ClearPhone clears the value of the "phone" field.
func (suo *SurveyUpdateOne) ClearPhone() *SurveyUpdateOne {
	suo.mutation.ClearPhone()
	return suo
}

// SetNotes sets the "notes" field.
func (suo *SurveyUpdateOne) SetNotes(s string) *SurveyUpdateOne {
	suo.mutation.SetNotes(s)
	return suo
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableNotes(s *string) *SurveyUpdateOne {
	if s != nil {
		suo.SetNotes(*s)
	}
	return suo
}

// ClearNotes clears the value of the "notes" field.
func (suo *SurveyUpdateOne) ClearNotes() *SurveyUpdateOne {
	suo.mutation.ClearNotes()
	return suo
}

// SetStatus sets the "status" field.
func (suo *SurveyUpdateOne) SetStatus(es enum.SurveyStatus) *SurveyUpdateOne {
	suo.mutation.SetStatus(es)
	return suo
}

// SetProgress sets the "progress" field.
func (suo *SurveyUpdateOne) SetProgress(ep enum.SurveyProgress) *SurveyUpdateOne {
	suo.mutation.SetProgress(ep)
	return suo
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableProgress(ep *enum.SurveyProgress) *SurveyUpdateOne {
	if ep != nil {
		suo.SetProgress(*ep)
	}
	return suo
}

// ClearProgress clears the value of the "progress" field.
func (suo *SurveyUpdateOne) ClearProgress() *SurveyUpdateOne {
	suo.mutation.ClearProgress()
	return suo
}

// SetProgressAt sets the "progress_at" field.
func (suo *SurveyUpdateOne) SetProgressAt(t time.Time) *SurveyUpdateOne {
	suo.mutation.SetProgressAt(t)
	return suo
}

// SetNillableProgressAt sets the "progress_at" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableProgressAt(t *time.Time) *SurveyUpdateOne {
	if t != nil {
		suo.SetProgressAt(*t)
	}
	return suo
}

// ClearProgressAt clears the value of the "progress_at" field.
func (suo *SurveyUpdateOne) ClearProgressAt() *SurveyUpdateOne {
	suo.mutation.ClearProgressAt()
	return suo
}

// SetProgressFlagAt sets the "progress_flag_at" field.
func (suo *SurveyUpdateOne) SetProgressFlagAt(t time.Time) *SurveyUpdateOne {
	suo.mutation.SetProgressFlagAt(t)
	return suo
}

// SetNillableProgressFlagAt sets the "progress_flag_at" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableProgressFlagAt(t *time.Time) *SurveyUpdateOne {
	if t != nil {
		suo.SetProgressFlagAt(*t)
	}
	return suo
}

// ClearProgressFlagAt clears the value of the "progress_flag_at" field.
func (suo *SurveyUpdateOne) ClearProgressFlagAt() *SurveyUpdateOne {
	suo.mutation.ClearProgressFlagAt()
	return suo
}

// AddProgressHistoryIDs adds the "progress_history" edge to the SurveyProgress entity by IDs.
func (suo *SurveyUpdateOne) AddProgressHistoryIDs(ids ...string) *SurveyUpdateOne {
	suo.mutation.AddProgressHistoryIDs(ids...)
	return suo
}

// AddProgressHistory adds the "progress_history" edges to the SurveyProgress entity.
func (suo *SurveyUpdateOne) AddProgressHistory(s ...*SurveyProgress) *SurveyUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddProgressHistoryIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (suo *SurveyUpdateOne) SetCreatedByID(id string) *SurveyUpdateOne {
	suo.mutation.SetCreatedByID(id)
	return suo
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableCreatedByID(id *string) *SurveyUpdateOne {
	if id != nil {
		suo = suo.SetCreatedByID(*id)
	}
	return suo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (suo *SurveyUpdateOne) SetCreatedBy(u *User) *SurveyUpdateOne {
	return suo.SetCreatedByID(u.ID)
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (suo *SurveyUpdateOne) SetPartnerID(id string) *SurveyUpdateOne {
	suo.mutation.SetPartnerID(id)
	return suo
}

// SetNillablePartnerID sets the "partner" edge to the Partner entity by ID if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillablePartnerID(id *string) *SurveyUpdateOne {
	if id != nil {
		suo = suo.SetPartnerID(*id)
	}
	return suo
}

// SetPartner sets the "partner" edge to the Partner entity.
func (suo *SurveyUpdateOne) SetPartner(p *Partner) *SurveyUpdateOne {
	return suo.SetPartnerID(p.ID)
}

// Mutation returns the SurveyMutation object of the builder.
func (suo *SurveyUpdateOne) Mutation() *SurveyMutation {
	return suo.mutation
}

// ClearProgressHistory clears all "progress_history" edges to the SurveyProgress entity.
func (suo *SurveyUpdateOne) ClearProgressHistory() *SurveyUpdateOne {
	suo.mutation.ClearProgressHistory()
	return suo
}

// RemoveProgressHistoryIDs removes the "progress_history" edge to SurveyProgress entities by IDs.
func (suo *SurveyUpdateOne) RemoveProgressHistoryIDs(ids ...string) *SurveyUpdateOne {
	suo.mutation.RemoveProgressHistoryIDs(ids...)
	return suo
}

// RemoveProgressHistory removes "progress_history" edges to SurveyProgress entities.
func (suo *SurveyUpdateOne) RemoveProgressHistory(s ...*SurveyProgress) *SurveyUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveProgressHistoryIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (suo *SurveyUpdateOne) ClearCreatedBy() *SurveyUpdateOne {
	suo.mutation.ClearCreatedBy()
	return suo
}

// ClearPartner clears the "partner" edge to the Partner entity.
func (suo *SurveyUpdateOne) ClearPartner() *SurveyUpdateOne {
	suo.mutation.ClearPartner()
	return suo
}

// Where appends a list predicates to the SurveyUpdate builder.
func (suo *SurveyUpdateOne) Where(ps ...predicate.Survey) *SurveyUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SurveyUpdateOne) Select(field string, fields ...string) *SurveyUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Survey entity.
func (suo *SurveyUpdateOne) Save(ctx context.Context) (*Survey, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SurveyUpdateOne) SaveX(ctx context.Context) *Survey {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SurveyUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SurveyUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SurveyUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := survey.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SurveyUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := survey.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Survey.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Address(); ok {
		if err := survey.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Survey.address": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Phone(); ok {
		if err := survey.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Survey.phone": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Notes(); ok {
		if err := survey.NotesValidator(v); err != nil {
			return &ValidationError{Name: "notes", err: fmt.Errorf(`ent: validator failed for field "Survey.notes": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Status(); ok {
		if err := survey.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Survey.status": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Progress(); ok {
		if err := survey.ProgressValidator(v); err != nil {
			return &ValidationError{Name: "progress", err: fmt.Errorf(`ent: validator failed for field "Survey.progress": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SurveyUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SurveyUpdateOne) sqlSave(ctx context.Context) (_node *Survey, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(survey.Table, survey.Columns, sqlgraph.NewFieldSpec(survey.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Survey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, survey.FieldID)
		for _, f := range fields {
			if !survey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != survey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(survey.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Until(); ok {
		_spec.SetField(survey.FieldUntil, field.TypeTime, value)
	}
	if suo.mutation.UntilCleared() {
		_spec.ClearField(survey.FieldUntil, field.TypeTime)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(survey.FieldName, field.TypeString, value)
	}
	if suo.mutation.NameCleared() {
		_spec.ClearField(survey.FieldName, field.TypeString)
	}
	if value, ok := suo.mutation.Address(); ok {
		_spec.SetField(survey.FieldAddress, field.TypeString, value)
	}
	if suo.mutation.AddressCleared() {
		_spec.ClearField(survey.FieldAddress, field.TypeString)
	}
	if value, ok := suo.mutation.Phone(); ok {
		_spec.SetField(survey.FieldPhone, field.TypeString, value)
	}
	if suo.mutation.PhoneCleared() {
		_spec.ClearField(survey.FieldPhone, field.TypeString)
	}
	if value, ok := suo.mutation.Notes(); ok {
		_spec.SetField(survey.FieldNotes, field.TypeString, value)
	}
	if suo.mutation.NotesCleared() {
		_spec.ClearField(survey.FieldNotes, field.TypeString)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(survey.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := suo.mutation.Progress(); ok {
		_spec.SetField(survey.FieldProgress, field.TypeEnum, value)
	}
	if suo.mutation.ProgressCleared() {
		_spec.ClearField(survey.FieldProgress, field.TypeEnum)
	}
	if value, ok := suo.mutation.ProgressAt(); ok {
		_spec.SetField(survey.FieldProgressAt, field.TypeTime, value)
	}
	if suo.mutation.ProgressAtCleared() {
		_spec.ClearField(survey.FieldProgressAt, field.TypeTime)
	}
	if value, ok := suo.mutation.ProgressFlagAt(); ok {
		_spec.SetField(survey.FieldProgressFlagAt, field.TypeTime, value)
	}
	if suo.mutation.ProgressFlagAtCleared() {
		_spec.ClearField(survey.FieldProgressFlagAt, field.TypeTime)
	}
	if suo.mutation.ProgressHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.ProgressHistoryTable,
			Columns: []string{survey.ProgressHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyprogress.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedProgressHistoryIDs(); len(nodes) > 0 && !suo.mutation.ProgressHistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.ProgressHistoryTable,
			Columns: []string{survey.ProgressHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyprogress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProgressHistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.ProgressHistoryTable,
			Columns: []string{survey.ProgressHistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyprogress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.CreatedByTable,
			Columns: []string{survey.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.CreatedByTable,
			Columns: []string{survey.CreatedByColumn},
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
	if suo.mutation.PartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.PartnerTable,
			Columns: []string{survey.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.PartnerTable,
			Columns: []string{survey.PartnerColumn},
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
	_spec.AddModifiers(suo.modifiers...)
	_node = &Survey{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{survey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
