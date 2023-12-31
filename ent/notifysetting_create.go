// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/notifysetting"
	"roofix/ent/user"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotifySettingCreate is the builder for creating a NotifySetting entity.
type NotifySettingCreate struct {
	config
	mutation *NotifySettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (nsc *NotifySettingCreate) SetCreatedAt(t time.Time) *NotifySettingCreate {
	nsc.mutation.SetCreatedAt(t)
	return nsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nsc *NotifySettingCreate) SetNillableCreatedAt(t *time.Time) *NotifySettingCreate {
	if t != nil {
		nsc.SetCreatedAt(*t)
	}
	return nsc
}

// SetTopicID sets the "topic_id" field.
func (nsc *NotifySettingCreate) SetTopicID(s string) *NotifySettingCreate {
	nsc.mutation.SetTopicID(s)
	return nsc
}

// SetReceiveEmail sets the "receive_email" field.
func (nsc *NotifySettingCreate) SetReceiveEmail(b bool) *NotifySettingCreate {
	nsc.mutation.SetReceiveEmail(b)
	return nsc
}

// SetNillableReceiveEmail sets the "receive_email" field if the given value is not nil.
func (nsc *NotifySettingCreate) SetNillableReceiveEmail(b *bool) *NotifySettingCreate {
	if b != nil {
		nsc.SetReceiveEmail(*b)
	}
	return nsc
}

// SetReceiveSms sets the "receive_sms" field.
func (nsc *NotifySettingCreate) SetReceiveSms(b bool) *NotifySettingCreate {
	nsc.mutation.SetReceiveSms(b)
	return nsc
}

// SetNillableReceiveSms sets the "receive_sms" field if the given value is not nil.
func (nsc *NotifySettingCreate) SetNillableReceiveSms(b *bool) *NotifySettingCreate {
	if b != nil {
		nsc.SetReceiveSms(*b)
	}
	return nsc
}

// SetID sets the "id" field.
func (nsc *NotifySettingCreate) SetID(s string) *NotifySettingCreate {
	nsc.mutation.SetID(s)
	return nsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nsc *NotifySettingCreate) SetNillableID(s *string) *NotifySettingCreate {
	if s != nil {
		nsc.SetID(*s)
	}
	return nsc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (nsc *NotifySettingCreate) SetUserID(id string) *NotifySettingCreate {
	nsc.mutation.SetUserID(id)
	return nsc
}

// SetUser sets the "user" edge to the User entity.
func (nsc *NotifySettingCreate) SetUser(u *User) *NotifySettingCreate {
	return nsc.SetUserID(u.ID)
}

// Mutation returns the NotifySettingMutation object of the builder.
func (nsc *NotifySettingCreate) Mutation() *NotifySettingMutation {
	return nsc.mutation
}

// Save creates the NotifySetting in the database.
func (nsc *NotifySettingCreate) Save(ctx context.Context) (*NotifySetting, error) {
	nsc.defaults()
	return withHooks(ctx, nsc.sqlSave, nsc.mutation, nsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nsc *NotifySettingCreate) SaveX(ctx context.Context) *NotifySetting {
	v, err := nsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nsc *NotifySettingCreate) Exec(ctx context.Context) error {
	_, err := nsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsc *NotifySettingCreate) ExecX(ctx context.Context) {
	if err := nsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nsc *NotifySettingCreate) defaults() {
	if _, ok := nsc.mutation.CreatedAt(); !ok {
		v := notifysetting.DefaultCreatedAt()
		nsc.mutation.SetCreatedAt(v)
	}
	if _, ok := nsc.mutation.ReceiveEmail(); !ok {
		v := notifysetting.DefaultReceiveEmail
		nsc.mutation.SetReceiveEmail(v)
	}
	if _, ok := nsc.mutation.ReceiveSms(); !ok {
		v := notifysetting.DefaultReceiveSms
		nsc.mutation.SetReceiveSms(v)
	}
	if _, ok := nsc.mutation.ID(); !ok {
		v := notifysetting.DefaultID()
		nsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nsc *NotifySettingCreate) check() error {
	if _, ok := nsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotifySetting.created_at"`)}
	}
	if _, ok := nsc.mutation.TopicID(); !ok {
		return &ValidationError{Name: "topic_id", err: errors.New(`ent: missing required field "NotifySetting.topic_id"`)}
	}
	if v, ok := nsc.mutation.TopicID(); ok {
		if err := notifysetting.TopicIDValidator(v); err != nil {
			return &ValidationError{Name: "topic_id", err: fmt.Errorf(`ent: validator failed for field "NotifySetting.topic_id": %w`, err)}
		}
	}
	if v, ok := nsc.mutation.ID(); ok {
		if err := notifysetting.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "NotifySetting.id": %w`, err)}
		}
	}
	if _, ok := nsc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "NotifySetting.user"`)}
	}
	return nil
}

func (nsc *NotifySettingCreate) sqlSave(ctx context.Context) (*NotifySetting, error) {
	if err := nsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected NotifySetting.ID type: %T", _spec.ID.Value)
		}
	}
	nsc.mutation.id = &_node.ID
	nsc.mutation.done = true
	return _node, nil
}

func (nsc *NotifySettingCreate) createSpec() (*NotifySetting, *sqlgraph.CreateSpec) {
	var (
		_node = &NotifySetting{config: nsc.config}
		_spec = sqlgraph.NewCreateSpec(notifysetting.Table, sqlgraph.NewFieldSpec(notifysetting.FieldID, field.TypeString))
	)
	_spec.OnConflict = nsc.conflict
	if id, ok := nsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nsc.mutation.CreatedAt(); ok {
		_spec.SetField(notifysetting.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nsc.mutation.TopicID(); ok {
		_spec.SetField(notifysetting.FieldTopicID, field.TypeString, value)
		_node.TopicID = value
	}
	if value, ok := nsc.mutation.ReceiveEmail(); ok {
		_spec.SetField(notifysetting.FieldReceiveEmail, field.TypeBool, value)
		_node.ReceiveEmail = value
	}
	if value, ok := nsc.mutation.ReceiveSms(); ok {
		_spec.SetField(notifysetting.FieldReceiveSms, field.TypeBool, value)
		_node.ReceiveSms = value
	}
	if nodes := nsc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysetting.UserTable,
			Columns: []string{notifysetting.UserColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifySetting.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifySettingUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (nsc *NotifySettingCreate) OnConflict(opts ...sql.ConflictOption) *NotifySettingUpsertOne {
	nsc.conflict = opts
	return &NotifySettingUpsertOne{
		create: nsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifySetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nsc *NotifySettingCreate) OnConflictColumns(columns ...string) *NotifySettingUpsertOne {
	nsc.conflict = append(nsc.conflict, sql.ConflictColumns(columns...))
	return &NotifySettingUpsertOne{
		create: nsc,
	}
}

type (
	// NotifySettingUpsertOne is the builder for "upsert"-ing
	//  one NotifySetting node.
	NotifySettingUpsertOne struct {
		create *NotifySettingCreate
	}

	// NotifySettingUpsert is the "OnConflict" setter.
	NotifySettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetTopicID sets the "topic_id" field.
func (u *NotifySettingUpsert) SetTopicID(v string) *NotifySettingUpsert {
	u.Set(notifysetting.FieldTopicID, v)
	return u
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *NotifySettingUpsert) UpdateTopicID() *NotifySettingUpsert {
	u.SetExcluded(notifysetting.FieldTopicID)
	return u
}

// SetReceiveEmail sets the "receive_email" field.
func (u *NotifySettingUpsert) SetReceiveEmail(v bool) *NotifySettingUpsert {
	u.Set(notifysetting.FieldReceiveEmail, v)
	return u
}

// UpdateReceiveEmail sets the "receive_email" field to the value that was provided on create.
func (u *NotifySettingUpsert) UpdateReceiveEmail() *NotifySettingUpsert {
	u.SetExcluded(notifysetting.FieldReceiveEmail)
	return u
}

// ClearReceiveEmail clears the value of the "receive_email" field.
func (u *NotifySettingUpsert) ClearReceiveEmail() *NotifySettingUpsert {
	u.SetNull(notifysetting.FieldReceiveEmail)
	return u
}

// SetReceiveSms sets the "receive_sms" field.
func (u *NotifySettingUpsert) SetReceiveSms(v bool) *NotifySettingUpsert {
	u.Set(notifysetting.FieldReceiveSms, v)
	return u
}

// UpdateReceiveSms sets the "receive_sms" field to the value that was provided on create.
func (u *NotifySettingUpsert) UpdateReceiveSms() *NotifySettingUpsert {
	u.SetExcluded(notifysetting.FieldReceiveSms)
	return u
}

// ClearReceiveSms clears the value of the "receive_sms" field.
func (u *NotifySettingUpsert) ClearReceiveSms() *NotifySettingUpsert {
	u.SetNull(notifysetting.FieldReceiveSms)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.NotifySetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notifysetting.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotifySettingUpsertOne) UpdateNewValues() *NotifySettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(notifysetting.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(notifysetting.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifySetting.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NotifySettingUpsertOne) Ignore() *NotifySettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifySettingUpsertOne) DoNothing() *NotifySettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifySettingCreate.OnConflict
// documentation for more info.
func (u *NotifySettingUpsertOne) Update(set func(*NotifySettingUpsert)) *NotifySettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifySettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetTopicID sets the "topic_id" field.
func (u *NotifySettingUpsertOne) SetTopicID(v string) *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.SetTopicID(v)
	})
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *NotifySettingUpsertOne) UpdateTopicID() *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.UpdateTopicID()
	})
}

// SetReceiveEmail sets the "receive_email" field.
func (u *NotifySettingUpsertOne) SetReceiveEmail(v bool) *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.SetReceiveEmail(v)
	})
}

// UpdateReceiveEmail sets the "receive_email" field to the value that was provided on create.
func (u *NotifySettingUpsertOne) UpdateReceiveEmail() *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.UpdateReceiveEmail()
	})
}

// ClearReceiveEmail clears the value of the "receive_email" field.
func (u *NotifySettingUpsertOne) ClearReceiveEmail() *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.ClearReceiveEmail()
	})
}

// SetReceiveSms sets the "receive_sms" field.
func (u *NotifySettingUpsertOne) SetReceiveSms(v bool) *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.SetReceiveSms(v)
	})
}

// UpdateReceiveSms sets the "receive_sms" field to the value that was provided on create.
func (u *NotifySettingUpsertOne) UpdateReceiveSms() *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.UpdateReceiveSms()
	})
}

// ClearReceiveSms clears the value of the "receive_sms" field.
func (u *NotifySettingUpsertOne) ClearReceiveSms() *NotifySettingUpsertOne {
	return u.Update(func(s *NotifySettingUpsert) {
		s.ClearReceiveSms()
	})
}

// Exec executes the query.
func (u *NotifySettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifySettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifySettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotifySettingUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NotifySettingUpsertOne.ID is not supported by MySQL driver. Use NotifySettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotifySettingUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotifySettingCreateBulk is the builder for creating many NotifySetting entities in bulk.
type NotifySettingCreateBulk struct {
	config
	builders []*NotifySettingCreate
	conflict []sql.ConflictOption
}

// Save creates the NotifySetting entities in the database.
func (nscb *NotifySettingCreateBulk) Save(ctx context.Context) ([]*NotifySetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nscb.builders))
	nodes := make([]*NotifySetting, len(nscb.builders))
	mutators := make([]Mutator, len(nscb.builders))
	for i := range nscb.builders {
		func(i int, root context.Context) {
			builder := nscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifySettingMutation)
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
					_, err = mutators[i+1].Mutate(root, nscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nscb *NotifySettingCreateBulk) SaveX(ctx context.Context) []*NotifySetting {
	v, err := nscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nscb *NotifySettingCreateBulk) Exec(ctx context.Context) error {
	_, err := nscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nscb *NotifySettingCreateBulk) ExecX(ctx context.Context) {
	if err := nscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifySetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifySettingUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (nscb *NotifySettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotifySettingUpsertBulk {
	nscb.conflict = opts
	return &NotifySettingUpsertBulk{
		create: nscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifySetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nscb *NotifySettingCreateBulk) OnConflictColumns(columns ...string) *NotifySettingUpsertBulk {
	nscb.conflict = append(nscb.conflict, sql.ConflictColumns(columns...))
	return &NotifySettingUpsertBulk{
		create: nscb,
	}
}

// NotifySettingUpsertBulk is the builder for "upsert"-ing
// a bulk of NotifySetting nodes.
type NotifySettingUpsertBulk struct {
	create *NotifySettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NotifySetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notifysetting.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotifySettingUpsertBulk) UpdateNewValues() *NotifySettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(notifysetting.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(notifysetting.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifySetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NotifySettingUpsertBulk) Ignore() *NotifySettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifySettingUpsertBulk) DoNothing() *NotifySettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifySettingCreateBulk.OnConflict
// documentation for more info.
func (u *NotifySettingUpsertBulk) Update(set func(*NotifySettingUpsert)) *NotifySettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifySettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetTopicID sets the "topic_id" field.
func (u *NotifySettingUpsertBulk) SetTopicID(v string) *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.SetTopicID(v)
	})
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *NotifySettingUpsertBulk) UpdateTopicID() *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.UpdateTopicID()
	})
}

// SetReceiveEmail sets the "receive_email" field.
func (u *NotifySettingUpsertBulk) SetReceiveEmail(v bool) *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.SetReceiveEmail(v)
	})
}

// UpdateReceiveEmail sets the "receive_email" field to the value that was provided on create.
func (u *NotifySettingUpsertBulk) UpdateReceiveEmail() *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.UpdateReceiveEmail()
	})
}

// ClearReceiveEmail clears the value of the "receive_email" field.
func (u *NotifySettingUpsertBulk) ClearReceiveEmail() *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.ClearReceiveEmail()
	})
}

// SetReceiveSms sets the "receive_sms" field.
func (u *NotifySettingUpsertBulk) SetReceiveSms(v bool) *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.SetReceiveSms(v)
	})
}

// UpdateReceiveSms sets the "receive_sms" field to the value that was provided on create.
func (u *NotifySettingUpsertBulk) UpdateReceiveSms() *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.UpdateReceiveSms()
	})
}

// ClearReceiveSms clears the value of the "receive_sms" field.
func (u *NotifySettingUpsertBulk) ClearReceiveSms() *NotifySettingUpsertBulk {
	return u.Update(func(s *NotifySettingUpsert) {
		s.ClearReceiveSms()
	})
}

// Exec executes the query.
func (u *NotifySettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotifySettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifySettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifySettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
