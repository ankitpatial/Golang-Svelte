// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/contactus"
	"roofix/ent/partner"
	"roofix/ent/user"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContactUsCreate is the builder for creating a ContactUs entity.
type ContactUsCreate struct {
	config
	mutation *ContactUsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cuc *ContactUsCreate) SetCreatedAt(t time.Time) *ContactUsCreate {
	cuc.mutation.SetCreatedAt(t)
	return cuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuc *ContactUsCreate) SetNillableCreatedAt(t *time.Time) *ContactUsCreate {
	if t != nil {
		cuc.SetCreatedAt(*t)
	}
	return cuc
}

// SetUpdatedAt sets the "updated_at" field.
func (cuc *ContactUsCreate) SetUpdatedAt(t time.Time) *ContactUsCreate {
	cuc.mutation.SetUpdatedAt(t)
	return cuc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuc *ContactUsCreate) SetNillableUpdatedAt(t *time.Time) *ContactUsCreate {
	if t != nil {
		cuc.SetUpdatedAt(*t)
	}
	return cuc
}

// SetReason sets the "reason" field.
func (cuc *ContactUsCreate) SetReason(s string) *ContactUsCreate {
	cuc.mutation.SetReason(s)
	return cuc
}

// SetID sets the "id" field.
func (cuc *ContactUsCreate) SetID(s string) *ContactUsCreate {
	cuc.mutation.SetID(s)
	return cuc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cuc *ContactUsCreate) SetNillableID(s *string) *ContactUsCreate {
	if s != nil {
		cuc.SetID(*s)
	}
	return cuc
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (cuc *ContactUsCreate) SetPartnerID(id string) *ContactUsCreate {
	cuc.mutation.SetPartnerID(id)
	return cuc
}

// SetNillablePartnerID sets the "partner" edge to the Partner entity by ID if the given value is not nil.
func (cuc *ContactUsCreate) SetNillablePartnerID(id *string) *ContactUsCreate {
	if id != nil {
		cuc = cuc.SetPartnerID(*id)
	}
	return cuc
}

// SetPartner sets the "partner" edge to the Partner entity.
func (cuc *ContactUsCreate) SetPartner(p *Partner) *ContactUsCreate {
	return cuc.SetPartnerID(p.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (cuc *ContactUsCreate) SetCreatorID(id string) *ContactUsCreate {
	cuc.mutation.SetCreatorID(id)
	return cuc
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (cuc *ContactUsCreate) SetNillableCreatorID(id *string) *ContactUsCreate {
	if id != nil {
		cuc = cuc.SetCreatorID(*id)
	}
	return cuc
}

// SetCreator sets the "creator" edge to the User entity.
func (cuc *ContactUsCreate) SetCreator(u *User) *ContactUsCreate {
	return cuc.SetCreatorID(u.ID)
}

// Mutation returns the ContactUsMutation object of the builder.
func (cuc *ContactUsCreate) Mutation() *ContactUsMutation {
	return cuc.mutation
}

// Save creates the ContactUs in the database.
func (cuc *ContactUsCreate) Save(ctx context.Context) (*ContactUs, error) {
	cuc.defaults()
	return withHooks(ctx, cuc.sqlSave, cuc.mutation, cuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cuc *ContactUsCreate) SaveX(ctx context.Context) *ContactUs {
	v, err := cuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cuc *ContactUsCreate) Exec(ctx context.Context) error {
	_, err := cuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuc *ContactUsCreate) ExecX(ctx context.Context) {
	if err := cuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuc *ContactUsCreate) defaults() {
	if _, ok := cuc.mutation.CreatedAt(); !ok {
		v := contactus.DefaultCreatedAt()
		cuc.mutation.SetCreatedAt(v)
	}
	if _, ok := cuc.mutation.UpdatedAt(); !ok {
		v := contactus.DefaultUpdatedAt()
		cuc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cuc.mutation.ID(); !ok {
		v := contactus.DefaultID()
		cuc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuc *ContactUsCreate) check() error {
	if _, ok := cuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ContactUs.created_at"`)}
	}
	if _, ok := cuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ContactUs.updated_at"`)}
	}
	if _, ok := cuc.mutation.Reason(); !ok {
		return &ValidationError{Name: "reason", err: errors.New(`ent: missing required field "ContactUs.reason"`)}
	}
	if v, ok := cuc.mutation.Reason(); ok {
		if err := contactus.ReasonValidator(v); err != nil {
			return &ValidationError{Name: "reason", err: fmt.Errorf(`ent: validator failed for field "ContactUs.reason": %w`, err)}
		}
	}
	if v, ok := cuc.mutation.ID(); ok {
		if err := contactus.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "ContactUs.id": %w`, err)}
		}
	}
	return nil
}

func (cuc *ContactUsCreate) sqlSave(ctx context.Context) (*ContactUs, error) {
	if err := cuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ContactUs.ID type: %T", _spec.ID.Value)
		}
	}
	cuc.mutation.id = &_node.ID
	cuc.mutation.done = true
	return _node, nil
}

func (cuc *ContactUsCreate) createSpec() (*ContactUs, *sqlgraph.CreateSpec) {
	var (
		_node = &ContactUs{config: cuc.config}
		_spec = sqlgraph.NewCreateSpec(contactus.Table, sqlgraph.NewFieldSpec(contactus.FieldID, field.TypeString))
	)
	_spec.OnConflict = cuc.conflict
	if id, ok := cuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cuc.mutation.CreatedAt(); ok {
		_spec.SetField(contactus.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cuc.mutation.UpdatedAt(); ok {
		_spec.SetField(contactus.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cuc.mutation.Reason(); ok {
		_spec.SetField(contactus.FieldReason, field.TypeString, value)
		_node.Reason = value
	}
	if nodes := cuc.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contactus.PartnerTable,
			Columns: []string{contactus.PartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partner.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.partner_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cuc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contactus.CreatorTable,
			Columns: []string{contactus.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.creator_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ContactUs.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContactUsUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cuc *ContactUsCreate) OnConflict(opts ...sql.ConflictOption) *ContactUsUpsertOne {
	cuc.conflict = opts
	return &ContactUsUpsertOne{
		create: cuc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContactUs.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cuc *ContactUsCreate) OnConflictColumns(columns ...string) *ContactUsUpsertOne {
	cuc.conflict = append(cuc.conflict, sql.ConflictColumns(columns...))
	return &ContactUsUpsertOne{
		create: cuc,
	}
}

type (
	// ContactUsUpsertOne is the builder for "upsert"-ing
	//  one ContactUs node.
	ContactUsUpsertOne struct {
		create *ContactUsCreate
	}

	// ContactUsUpsert is the "OnConflict" setter.
	ContactUsUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ContactUsUpsert) SetUpdatedAt(v time.Time) *ContactUsUpsert {
	u.Set(contactus.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContactUsUpsert) UpdateUpdatedAt() *ContactUsUpsert {
	u.SetExcluded(contactus.FieldUpdatedAt)
	return u
}

// SetReason sets the "reason" field.
func (u *ContactUsUpsert) SetReason(v string) *ContactUsUpsert {
	u.Set(contactus.FieldReason, v)
	return u
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *ContactUsUpsert) UpdateReason() *ContactUsUpsert {
	u.SetExcluded(contactus.FieldReason)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ContactUs.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contactus.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContactUsUpsertOne) UpdateNewValues() *ContactUsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(contactus.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(contactus.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContactUs.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ContactUsUpsertOne) Ignore() *ContactUsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContactUsUpsertOne) DoNothing() *ContactUsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContactUsCreate.OnConflict
// documentation for more info.
func (u *ContactUsUpsertOne) Update(set func(*ContactUsUpsert)) *ContactUsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContactUsUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContactUsUpsertOne) SetUpdatedAt(v time.Time) *ContactUsUpsertOne {
	return u.Update(func(s *ContactUsUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContactUsUpsertOne) UpdateUpdatedAt() *ContactUsUpsertOne {
	return u.Update(func(s *ContactUsUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetReason sets the "reason" field.
func (u *ContactUsUpsertOne) SetReason(v string) *ContactUsUpsertOne {
	return u.Update(func(s *ContactUsUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *ContactUsUpsertOne) UpdateReason() *ContactUsUpsertOne {
	return u.Update(func(s *ContactUsUpsert) {
		s.UpdateReason()
	})
}

// Exec executes the query.
func (u *ContactUsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContactUsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContactUsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ContactUsUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ContactUsUpsertOne.ID is not supported by MySQL driver. Use ContactUsUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ContactUsUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ContactUsCreateBulk is the builder for creating many ContactUs entities in bulk.
type ContactUsCreateBulk struct {
	config
	builders []*ContactUsCreate
	conflict []sql.ConflictOption
}

// Save creates the ContactUs entities in the database.
func (cucb *ContactUsCreateBulk) Save(ctx context.Context) ([]*ContactUs, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cucb.builders))
	nodes := make([]*ContactUs, len(cucb.builders))
	mutators := make([]Mutator, len(cucb.builders))
	for i := range cucb.builders {
		func(i int, root context.Context) {
			builder := cucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContactUsMutation)
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
					_, err = mutators[i+1].Mutate(root, cucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cucb *ContactUsCreateBulk) SaveX(ctx context.Context) []*ContactUs {
	v, err := cucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cucb *ContactUsCreateBulk) Exec(ctx context.Context) error {
	_, err := cucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cucb *ContactUsCreateBulk) ExecX(ctx context.Context) {
	if err := cucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ContactUs.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContactUsUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cucb *ContactUsCreateBulk) OnConflict(opts ...sql.ConflictOption) *ContactUsUpsertBulk {
	cucb.conflict = opts
	return &ContactUsUpsertBulk{
		create: cucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContactUs.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cucb *ContactUsCreateBulk) OnConflictColumns(columns ...string) *ContactUsUpsertBulk {
	cucb.conflict = append(cucb.conflict, sql.ConflictColumns(columns...))
	return &ContactUsUpsertBulk{
		create: cucb,
	}
}

// ContactUsUpsertBulk is the builder for "upsert"-ing
// a bulk of ContactUs nodes.
type ContactUsUpsertBulk struct {
	create *ContactUsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ContactUs.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contactus.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContactUsUpsertBulk) UpdateNewValues() *ContactUsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(contactus.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(contactus.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContactUs.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ContactUsUpsertBulk) Ignore() *ContactUsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContactUsUpsertBulk) DoNothing() *ContactUsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContactUsCreateBulk.OnConflict
// documentation for more info.
func (u *ContactUsUpsertBulk) Update(set func(*ContactUsUpsert)) *ContactUsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContactUsUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContactUsUpsertBulk) SetUpdatedAt(v time.Time) *ContactUsUpsertBulk {
	return u.Update(func(s *ContactUsUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContactUsUpsertBulk) UpdateUpdatedAt() *ContactUsUpsertBulk {
	return u.Update(func(s *ContactUsUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetReason sets the "reason" field.
func (u *ContactUsUpsertBulk) SetReason(v string) *ContactUsUpsertBulk {
	return u.Update(func(s *ContactUsUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *ContactUsUpsertBulk) UpdateReason() *ContactUsUpsertBulk {
	return u.Update(func(s *ContactUsUpsert) {
		s.UpdateReason()
	})
}

// Exec executes the query.
func (u *ContactUsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ContactUsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContactUsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContactUsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
