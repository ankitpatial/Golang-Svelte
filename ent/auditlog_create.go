// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/apiuser"
	"roofix/ent/auditlog"
	"roofix/ent/user"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuditLogCreate is the builder for creating a AuditLog entity.
type AuditLogCreate struct {
	config
	mutation *AuditLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (alc *AuditLogCreate) SetCreatedAt(t time.Time) *AuditLogCreate {
	alc.mutation.SetCreatedAt(t)
	return alc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableCreatedAt(t *time.Time) *AuditLogCreate {
	if t != nil {
		alc.SetCreatedAt(*t)
	}
	return alc
}

// SetAction sets the "action" field.
func (alc *AuditLogCreate) SetAction(s string) *AuditLogCreate {
	alc.mutation.SetAction(s)
	return alc
}

// SetDescription sets the "description" field.
func (alc *AuditLogCreate) SetDescription(s string) *AuditLogCreate {
	alc.mutation.SetDescription(s)
	return alc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableDescription(s *string) *AuditLogCreate {
	if s != nil {
		alc.SetDescription(*s)
	}
	return alc
}

// SetIP sets the "ip" field.
func (alc *AuditLogCreate) SetIP(s string) *AuditLogCreate {
	alc.mutation.SetIP(s)
	return alc
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableIP(s *string) *AuditLogCreate {
	if s != nil {
		alc.SetIP(*s)
	}
	return alc
}

// SetID sets the "id" field.
func (alc *AuditLogCreate) SetID(s string) *AuditLogCreate {
	alc.mutation.SetID(s)
	return alc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (alc *AuditLogCreate) SetNillableID(s *string) *AuditLogCreate {
	if s != nil {
		alc.SetID(*s)
	}
	return alc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (alc *AuditLogCreate) SetUserID(id string) *AuditLogCreate {
	alc.mutation.SetUserID(id)
	return alc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (alc *AuditLogCreate) SetNillableUserID(id *string) *AuditLogCreate {
	if id != nil {
		alc = alc.SetUserID(*id)
	}
	return alc
}

// SetUser sets the "user" edge to the User entity.
func (alc *AuditLogCreate) SetUser(u *User) *AuditLogCreate {
	return alc.SetUserID(u.ID)
}

// SetAPIUserID sets the "api_user" edge to the ApiUser entity by ID.
func (alc *AuditLogCreate) SetAPIUserID(id string) *AuditLogCreate {
	alc.mutation.SetAPIUserID(id)
	return alc
}

// SetNillableAPIUserID sets the "api_user" edge to the ApiUser entity by ID if the given value is not nil.
func (alc *AuditLogCreate) SetNillableAPIUserID(id *string) *AuditLogCreate {
	if id != nil {
		alc = alc.SetAPIUserID(*id)
	}
	return alc
}

// SetAPIUser sets the "api_user" edge to the ApiUser entity.
func (alc *AuditLogCreate) SetAPIUser(a *ApiUser) *AuditLogCreate {
	return alc.SetAPIUserID(a.ID)
}

// Mutation returns the AuditLogMutation object of the builder.
func (alc *AuditLogCreate) Mutation() *AuditLogMutation {
	return alc.mutation
}

// Save creates the AuditLog in the database.
func (alc *AuditLogCreate) Save(ctx context.Context) (*AuditLog, error) {
	alc.defaults()
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AuditLogCreate) SaveX(ctx context.Context) *AuditLog {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AuditLogCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AuditLogCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alc *AuditLogCreate) defaults() {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		v := auditlog.DefaultCreatedAt()
		alc.mutation.SetCreatedAt(v)
	}
	if _, ok := alc.mutation.ID(); !ok {
		v := auditlog.DefaultID()
		alc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AuditLogCreate) check() error {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuditLog.created_at"`)}
	}
	if _, ok := alc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "AuditLog.action"`)}
	}
	if v, ok := alc.mutation.Action(); ok {
		if err := auditlog.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "AuditLog.action": %w`, err)}
		}
	}
	if v, ok := alc.mutation.Description(); ok {
		if err := auditlog.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "AuditLog.description": %w`, err)}
		}
	}
	if v, ok := alc.mutation.IP(); ok {
		if err := auditlog.IPValidator(v); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`ent: validator failed for field "AuditLog.ip": %w`, err)}
		}
	}
	if v, ok := alc.mutation.ID(); ok {
		if err := auditlog.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "AuditLog.id": %w`, err)}
		}
	}
	return nil
}

func (alc *AuditLogCreate) sqlSave(ctx context.Context) (*AuditLog, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AuditLog.ID type: %T", _spec.ID.Value)
		}
	}
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *AuditLogCreate) createSpec() (*AuditLog, *sqlgraph.CreateSpec) {
	var (
		_node = &AuditLog{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(auditlog.Table, sqlgraph.NewFieldSpec(auditlog.FieldID, field.TypeString))
	)
	_spec.OnConflict = alc.conflict
	if id, ok := alc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := alc.mutation.CreatedAt(); ok {
		_spec.SetField(auditlog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := alc.mutation.Action(); ok {
		_spec.SetField(auditlog.FieldAction, field.TypeString, value)
		_node.Action = value
	}
	if value, ok := alc.mutation.Description(); ok {
		_spec.SetField(auditlog.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := alc.mutation.IP(); ok {
		_spec.SetField(auditlog.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if nodes := alc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   auditlog.UserTable,
			Columns: []string{auditlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := alc.mutation.APIUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   auditlog.APIUserTable,
			Columns: []string{auditlog.APIUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apiuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.api_user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuditLog.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuditLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (alc *AuditLogCreate) OnConflict(opts ...sql.ConflictOption) *AuditLogUpsertOne {
	alc.conflict = opts
	return &AuditLogUpsertOne{
		create: alc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuditLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (alc *AuditLogCreate) OnConflictColumns(columns ...string) *AuditLogUpsertOne {
	alc.conflict = append(alc.conflict, sql.ConflictColumns(columns...))
	return &AuditLogUpsertOne{
		create: alc,
	}
}

type (
	// AuditLogUpsertOne is the builder for "upsert"-ing
	//  one AuditLog node.
	AuditLogUpsertOne struct {
		create *AuditLogCreate
	}

	// AuditLogUpsert is the "OnConflict" setter.
	AuditLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetAction sets the "action" field.
func (u *AuditLogUpsert) SetAction(v string) *AuditLogUpsert {
	u.Set(auditlog.FieldAction, v)
	return u
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *AuditLogUpsert) UpdateAction() *AuditLogUpsert {
	u.SetExcluded(auditlog.FieldAction)
	return u
}

// SetDescription sets the "description" field.
func (u *AuditLogUpsert) SetDescription(v string) *AuditLogUpsert {
	u.Set(auditlog.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AuditLogUpsert) UpdateDescription() *AuditLogUpsert {
	u.SetExcluded(auditlog.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *AuditLogUpsert) ClearDescription() *AuditLogUpsert {
	u.SetNull(auditlog.FieldDescription)
	return u
}

// SetIP sets the "ip" field.
func (u *AuditLogUpsert) SetIP(v string) *AuditLogUpsert {
	u.Set(auditlog.FieldIP, v)
	return u
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *AuditLogUpsert) UpdateIP() *AuditLogUpsert {
	u.SetExcluded(auditlog.FieldIP)
	return u
}

// ClearIP clears the value of the "ip" field.
func (u *AuditLogUpsert) ClearIP() *AuditLogUpsert {
	u.SetNull(auditlog.FieldIP)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AuditLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auditlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuditLogUpsertOne) UpdateNewValues() *AuditLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(auditlog.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(auditlog.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuditLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuditLogUpsertOne) Ignore() *AuditLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuditLogUpsertOne) DoNothing() *AuditLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuditLogCreate.OnConflict
// documentation for more info.
func (u *AuditLogUpsertOne) Update(set func(*AuditLogUpsert)) *AuditLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuditLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetAction sets the "action" field.
func (u *AuditLogUpsertOne) SetAction(v string) *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *AuditLogUpsertOne) UpdateAction() *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.UpdateAction()
	})
}

// SetDescription sets the "description" field.
func (u *AuditLogUpsertOne) SetDescription(v string) *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AuditLogUpsertOne) UpdateDescription() *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AuditLogUpsertOne) ClearDescription() *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.ClearDescription()
	})
}

// SetIP sets the "ip" field.
func (u *AuditLogUpsertOne) SetIP(v string) *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *AuditLogUpsertOne) UpdateIP() *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.UpdateIP()
	})
}

// ClearIP clears the value of the "ip" field.
func (u *AuditLogUpsertOne) ClearIP() *AuditLogUpsertOne {
	return u.Update(func(s *AuditLogUpsert) {
		s.ClearIP()
	})
}

// Exec executes the query.
func (u *AuditLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuditLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuditLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuditLogUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AuditLogUpsertOne.ID is not supported by MySQL driver. Use AuditLogUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuditLogUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuditLogCreateBulk is the builder for creating many AuditLog entities in bulk.
type AuditLogCreateBulk struct {
	config
	builders []*AuditLogCreate
	conflict []sql.ConflictOption
}

// Save creates the AuditLog entities in the database.
func (alcb *AuditLogCreateBulk) Save(ctx context.Context) ([]*AuditLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AuditLog, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuditLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = alcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AuditLogCreateBulk) SaveX(ctx context.Context) []*AuditLog {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AuditLogCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AuditLogCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuditLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuditLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (alcb *AuditLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuditLogUpsertBulk {
	alcb.conflict = opts
	return &AuditLogUpsertBulk{
		create: alcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuditLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (alcb *AuditLogCreateBulk) OnConflictColumns(columns ...string) *AuditLogUpsertBulk {
	alcb.conflict = append(alcb.conflict, sql.ConflictColumns(columns...))
	return &AuditLogUpsertBulk{
		create: alcb,
	}
}

// AuditLogUpsertBulk is the builder for "upsert"-ing
// a bulk of AuditLog nodes.
type AuditLogUpsertBulk struct {
	create *AuditLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AuditLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auditlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuditLogUpsertBulk) UpdateNewValues() *AuditLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(auditlog.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(auditlog.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuditLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuditLogUpsertBulk) Ignore() *AuditLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuditLogUpsertBulk) DoNothing() *AuditLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuditLogCreateBulk.OnConflict
// documentation for more info.
func (u *AuditLogUpsertBulk) Update(set func(*AuditLogUpsert)) *AuditLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuditLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetAction sets the "action" field.
func (u *AuditLogUpsertBulk) SetAction(v string) *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *AuditLogUpsertBulk) UpdateAction() *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.UpdateAction()
	})
}

// SetDescription sets the "description" field.
func (u *AuditLogUpsertBulk) SetDescription(v string) *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AuditLogUpsertBulk) UpdateDescription() *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AuditLogUpsertBulk) ClearDescription() *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.ClearDescription()
	})
}

// SetIP sets the "ip" field.
func (u *AuditLogUpsertBulk) SetIP(v string) *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *AuditLogUpsertBulk) UpdateIP() *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.UpdateIP()
	})
}

// ClearIP clears the value of the "ip" field.
func (u *AuditLogUpsertBulk) ClearIP() *AuditLogUpsertBulk {
	return u.Update(func(s *AuditLogUpsert) {
		s.ClearIP()
	})
}

// Exec executes the query.
func (u *AuditLogUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuditLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuditLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuditLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}