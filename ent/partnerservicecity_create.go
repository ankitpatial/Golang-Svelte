// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/partnerservicecity"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PartnerServiceCityCreate is the builder for creating a PartnerServiceCity entity.
type PartnerServiceCityCreate struct {
	config
	mutation *PartnerServiceCityMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pscc *PartnerServiceCityCreate) SetCreatedAt(t time.Time) *PartnerServiceCityCreate {
	pscc.mutation.SetCreatedAt(t)
	return pscc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pscc *PartnerServiceCityCreate) SetNillableCreatedAt(t *time.Time) *PartnerServiceCityCreate {
	if t != nil {
		pscc.SetCreatedAt(*t)
	}
	return pscc
}

// SetUpdatedAt sets the "updated_at" field.
func (pscc *PartnerServiceCityCreate) SetUpdatedAt(t time.Time) *PartnerServiceCityCreate {
	pscc.mutation.SetUpdatedAt(t)
	return pscc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pscc *PartnerServiceCityCreate) SetNillableUpdatedAt(t *time.Time) *PartnerServiceCityCreate {
	if t != nil {
		pscc.SetUpdatedAt(*t)
	}
	return pscc
}

// SetPostalID sets the "postal_id" field.
func (pscc *PartnerServiceCityCreate) SetPostalID(s string) *PartnerServiceCityCreate {
	pscc.mutation.SetPostalID(s)
	return pscc
}

// SetActive sets the "active" field.
func (pscc *PartnerServiceCityCreate) SetActive(b bool) *PartnerServiceCityCreate {
	pscc.mutation.SetActive(b)
	return pscc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (pscc *PartnerServiceCityCreate) SetNillableActive(b *bool) *PartnerServiceCityCreate {
	if b != nil {
		pscc.SetActive(*b)
	}
	return pscc
}

// SetName sets the "name" field.
func (pscc *PartnerServiceCityCreate) SetName(s string) *PartnerServiceCityCreate {
	pscc.mutation.SetName(s)
	return pscc
}

// SetNaicsCode sets the "naics_code" field.
func (pscc *PartnerServiceCityCreate) SetNaicsCode(u uint) *PartnerServiceCityCreate {
	pscc.mutation.SetNaicsCode(u)
	return pscc
}

// SetLicenseNo sets the "license_no" field.
func (pscc *PartnerServiceCityCreate) SetLicenseNo(s string) *PartnerServiceCityCreate {
	pscc.mutation.SetLicenseNo(s)
	return pscc
}

// SetNillableLicenseNo sets the "license_no" field if the given value is not nil.
func (pscc *PartnerServiceCityCreate) SetNillableLicenseNo(s *string) *PartnerServiceCityCreate {
	if s != nil {
		pscc.SetLicenseNo(*s)
	}
	return pscc
}

// SetProofDocID sets the "proof_doc_id" field.
func (pscc *PartnerServiceCityCreate) SetProofDocID(s string) *PartnerServiceCityCreate {
	pscc.mutation.SetProofDocID(s)
	return pscc
}

// SetNillableProofDocID sets the "proof_doc_id" field if the given value is not nil.
func (pscc *PartnerServiceCityCreate) SetNillableProofDocID(s *string) *PartnerServiceCityCreate {
	if s != nil {
		pscc.SetProofDocID(*s)
	}
	return pscc
}

// SetID sets the "id" field.
func (pscc *PartnerServiceCityCreate) SetID(s string) *PartnerServiceCityCreate {
	pscc.mutation.SetID(s)
	return pscc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pscc *PartnerServiceCityCreate) SetNillableID(s *string) *PartnerServiceCityCreate {
	if s != nil {
		pscc.SetID(*s)
	}
	return pscc
}

// SetPartnerID sets the "partner" edge to the Partner entity by ID.
func (pscc *PartnerServiceCityCreate) SetPartnerID(id string) *PartnerServiceCityCreate {
	pscc.mutation.SetPartnerID(id)
	return pscc
}

// SetPartner sets the "partner" edge to the Partner entity.
func (pscc *PartnerServiceCityCreate) SetPartner(p *Partner) *PartnerServiceCityCreate {
	return pscc.SetPartnerID(p.ID)
}

// Mutation returns the PartnerServiceCityMutation object of the builder.
func (pscc *PartnerServiceCityCreate) Mutation() *PartnerServiceCityMutation {
	return pscc.mutation
}

// Save creates the PartnerServiceCity in the database.
func (pscc *PartnerServiceCityCreate) Save(ctx context.Context) (*PartnerServiceCity, error) {
	pscc.defaults()
	return withHooks(ctx, pscc.sqlSave, pscc.mutation, pscc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pscc *PartnerServiceCityCreate) SaveX(ctx context.Context) *PartnerServiceCity {
	v, err := pscc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscc *PartnerServiceCityCreate) Exec(ctx context.Context) error {
	_, err := pscc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscc *PartnerServiceCityCreate) ExecX(ctx context.Context) {
	if err := pscc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pscc *PartnerServiceCityCreate) defaults() {
	if _, ok := pscc.mutation.CreatedAt(); !ok {
		v := partnerservicecity.DefaultCreatedAt()
		pscc.mutation.SetCreatedAt(v)
	}
	if _, ok := pscc.mutation.UpdatedAt(); !ok {
		v := partnerservicecity.DefaultUpdatedAt()
		pscc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pscc.mutation.Active(); !ok {
		v := partnerservicecity.DefaultActive
		pscc.mutation.SetActive(v)
	}
	if _, ok := pscc.mutation.ID(); !ok {
		v := partnerservicecity.DefaultID()
		pscc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pscc *PartnerServiceCityCreate) check() error {
	if _, ok := pscc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PartnerServiceCity.created_at"`)}
	}
	if _, ok := pscc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PartnerServiceCity.updated_at"`)}
	}
	if _, ok := pscc.mutation.PostalID(); !ok {
		return &ValidationError{Name: "postal_id", err: errors.New(`ent: missing required field "PartnerServiceCity.postal_id"`)}
	}
	if v, ok := pscc.mutation.PostalID(); ok {
		if err := partnerservicecity.PostalIDValidator(v); err != nil {
			return &ValidationError{Name: "postal_id", err: fmt.Errorf(`ent: validator failed for field "PartnerServiceCity.postal_id": %w`, err)}
		}
	}
	if _, ok := pscc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "PartnerServiceCity.active"`)}
	}
	if _, ok := pscc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "PartnerServiceCity.name"`)}
	}
	if v, ok := pscc.mutation.Name(); ok {
		if err := partnerservicecity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PartnerServiceCity.name": %w`, err)}
		}
	}
	if _, ok := pscc.mutation.NaicsCode(); !ok {
		return &ValidationError{Name: "naics_code", err: errors.New(`ent: missing required field "PartnerServiceCity.naics_code"`)}
	}
	if v, ok := pscc.mutation.LicenseNo(); ok {
		if err := partnerservicecity.LicenseNoValidator(v); err != nil {
			return &ValidationError{Name: "license_no", err: fmt.Errorf(`ent: validator failed for field "PartnerServiceCity.license_no": %w`, err)}
		}
	}
	if v, ok := pscc.mutation.ProofDocID(); ok {
		if err := partnerservicecity.ProofDocIDValidator(v); err != nil {
			return &ValidationError{Name: "proof_doc_id", err: fmt.Errorf(`ent: validator failed for field "PartnerServiceCity.proof_doc_id": %w`, err)}
		}
	}
	if v, ok := pscc.mutation.ID(); ok {
		if err := partnerservicecity.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "PartnerServiceCity.id": %w`, err)}
		}
	}
	if _, ok := pscc.mutation.PartnerID(); !ok {
		return &ValidationError{Name: "partner", err: errors.New(`ent: missing required edge "PartnerServiceCity.partner"`)}
	}
	return nil
}

func (pscc *PartnerServiceCityCreate) sqlSave(ctx context.Context) (*PartnerServiceCity, error) {
	if err := pscc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pscc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pscc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected PartnerServiceCity.ID type: %T", _spec.ID.Value)
		}
	}
	pscc.mutation.id = &_node.ID
	pscc.mutation.done = true
	return _node, nil
}

func (pscc *PartnerServiceCityCreate) createSpec() (*PartnerServiceCity, *sqlgraph.CreateSpec) {
	var (
		_node = &PartnerServiceCity{config: pscc.config}
		_spec = sqlgraph.NewCreateSpec(partnerservicecity.Table, sqlgraph.NewFieldSpec(partnerservicecity.FieldID, field.TypeString))
	)
	_spec.OnConflict = pscc.conflict
	if id, ok := pscc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pscc.mutation.CreatedAt(); ok {
		_spec.SetField(partnerservicecity.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pscc.mutation.UpdatedAt(); ok {
		_spec.SetField(partnerservicecity.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pscc.mutation.PostalID(); ok {
		_spec.SetField(partnerservicecity.FieldPostalID, field.TypeString, value)
		_node.PostalID = value
	}
	if value, ok := pscc.mutation.Active(); ok {
		_spec.SetField(partnerservicecity.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := pscc.mutation.Name(); ok {
		_spec.SetField(partnerservicecity.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pscc.mutation.NaicsCode(); ok {
		_spec.SetField(partnerservicecity.FieldNaicsCode, field.TypeUint, value)
		_node.NaicsCode = value
	}
	if value, ok := pscc.mutation.LicenseNo(); ok {
		_spec.SetField(partnerservicecity.FieldLicenseNo, field.TypeString, value)
		_node.LicenseNo = &value
	}
	if value, ok := pscc.mutation.ProofDocID(); ok {
		_spec.SetField(partnerservicecity.FieldProofDocID, field.TypeString, value)
		_node.ProofDocID = &value
	}
	if nodes := pscc.mutation.PartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partnerservicecity.PartnerTable,
			Columns: []string{partnerservicecity.PartnerColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PartnerServiceCity.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PartnerServiceCityUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pscc *PartnerServiceCityCreate) OnConflict(opts ...sql.ConflictOption) *PartnerServiceCityUpsertOne {
	pscc.conflict = opts
	return &PartnerServiceCityUpsertOne{
		create: pscc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PartnerServiceCity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pscc *PartnerServiceCityCreate) OnConflictColumns(columns ...string) *PartnerServiceCityUpsertOne {
	pscc.conflict = append(pscc.conflict, sql.ConflictColumns(columns...))
	return &PartnerServiceCityUpsertOne{
		create: pscc,
	}
}

type (
	// PartnerServiceCityUpsertOne is the builder for "upsert"-ing
	//  one PartnerServiceCity node.
	PartnerServiceCityUpsertOne struct {
		create *PartnerServiceCityCreate
	}

	// PartnerServiceCityUpsert is the "OnConflict" setter.
	PartnerServiceCityUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *PartnerServiceCityUpsert) SetUpdatedAt(v time.Time) *PartnerServiceCityUpsert {
	u.Set(partnerservicecity.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PartnerServiceCityUpsert) UpdateUpdatedAt() *PartnerServiceCityUpsert {
	u.SetExcluded(partnerservicecity.FieldUpdatedAt)
	return u
}

// SetPostalID sets the "postal_id" field.
func (u *PartnerServiceCityUpsert) SetPostalID(v string) *PartnerServiceCityUpsert {
	u.Set(partnerservicecity.FieldPostalID, v)
	return u
}

// UpdatePostalID sets the "postal_id" field to the value that was provided on create.
func (u *PartnerServiceCityUpsert) UpdatePostalID() *PartnerServiceCityUpsert {
	u.SetExcluded(partnerservicecity.FieldPostalID)
	return u
}

// SetActive sets the "active" field.
func (u *PartnerServiceCityUpsert) SetActive(v bool) *PartnerServiceCityUpsert {
	u.Set(partnerservicecity.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *PartnerServiceCityUpsert) UpdateActive() *PartnerServiceCityUpsert {
	u.SetExcluded(partnerservicecity.FieldActive)
	return u
}

// SetName sets the "name" field.
func (u *PartnerServiceCityUpsert) SetName(v string) *PartnerServiceCityUpsert {
	u.Set(partnerservicecity.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PartnerServiceCityUpsert) UpdateName() *PartnerServiceCityUpsert {
	u.SetExcluded(partnerservicecity.FieldName)
	return u
}

// SetNaicsCode sets the "naics_code" field.
func (u *PartnerServiceCityUpsert) SetNaicsCode(v uint) *PartnerServiceCityUpsert {
	u.Set(partnerservicecity.FieldNaicsCode, v)
	return u
}

// UpdateNaicsCode sets the "naics_code" field to the value that was provided on create.
func (u *PartnerServiceCityUpsert) UpdateNaicsCode() *PartnerServiceCityUpsert {
	u.SetExcluded(partnerservicecity.FieldNaicsCode)
	return u
}

// AddNaicsCode adds v to the "naics_code" field.
func (u *PartnerServiceCityUpsert) AddNaicsCode(v uint) *PartnerServiceCityUpsert {
	u.Add(partnerservicecity.FieldNaicsCode, v)
	return u
}

// SetLicenseNo sets the "license_no" field.
func (u *PartnerServiceCityUpsert) SetLicenseNo(v string) *PartnerServiceCityUpsert {
	u.Set(partnerservicecity.FieldLicenseNo, v)
	return u
}

// UpdateLicenseNo sets the "license_no" field to the value that was provided on create.
func (u *PartnerServiceCityUpsert) UpdateLicenseNo() *PartnerServiceCityUpsert {
	u.SetExcluded(partnerservicecity.FieldLicenseNo)
	return u
}

// ClearLicenseNo clears the value of the "license_no" field.
func (u *PartnerServiceCityUpsert) ClearLicenseNo() *PartnerServiceCityUpsert {
	u.SetNull(partnerservicecity.FieldLicenseNo)
	return u
}

// SetProofDocID sets the "proof_doc_id" field.
func (u *PartnerServiceCityUpsert) SetProofDocID(v string) *PartnerServiceCityUpsert {
	u.Set(partnerservicecity.FieldProofDocID, v)
	return u
}

// UpdateProofDocID sets the "proof_doc_id" field to the value that was provided on create.
func (u *PartnerServiceCityUpsert) UpdateProofDocID() *PartnerServiceCityUpsert {
	u.SetExcluded(partnerservicecity.FieldProofDocID)
	return u
}

// ClearProofDocID clears the value of the "proof_doc_id" field.
func (u *PartnerServiceCityUpsert) ClearProofDocID() *PartnerServiceCityUpsert {
	u.SetNull(partnerservicecity.FieldProofDocID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PartnerServiceCity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(partnerservicecity.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PartnerServiceCityUpsertOne) UpdateNewValues() *PartnerServiceCityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(partnerservicecity.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(partnerservicecity.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PartnerServiceCity.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PartnerServiceCityUpsertOne) Ignore() *PartnerServiceCityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PartnerServiceCityUpsertOne) DoNothing() *PartnerServiceCityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PartnerServiceCityCreate.OnConflict
// documentation for more info.
func (u *PartnerServiceCityUpsertOne) Update(set func(*PartnerServiceCityUpsert)) *PartnerServiceCityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PartnerServiceCityUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PartnerServiceCityUpsertOne) SetUpdatedAt(v time.Time) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertOne) UpdateUpdatedAt() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetPostalID sets the "postal_id" field.
func (u *PartnerServiceCityUpsertOne) SetPostalID(v string) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetPostalID(v)
	})
}

// UpdatePostalID sets the "postal_id" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertOne) UpdatePostalID() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdatePostalID()
	})
}

// SetActive sets the "active" field.
func (u *PartnerServiceCityUpsertOne) SetActive(v bool) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertOne) UpdateActive() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateActive()
	})
}

// SetName sets the "name" field.
func (u *PartnerServiceCityUpsertOne) SetName(v string) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertOne) UpdateName() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateName()
	})
}

// SetNaicsCode sets the "naics_code" field.
func (u *PartnerServiceCityUpsertOne) SetNaicsCode(v uint) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetNaicsCode(v)
	})
}

// AddNaicsCode adds v to the "naics_code" field.
func (u *PartnerServiceCityUpsertOne) AddNaicsCode(v uint) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.AddNaicsCode(v)
	})
}

// UpdateNaicsCode sets the "naics_code" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertOne) UpdateNaicsCode() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateNaicsCode()
	})
}

// SetLicenseNo sets the "license_no" field.
func (u *PartnerServiceCityUpsertOne) SetLicenseNo(v string) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetLicenseNo(v)
	})
}

// UpdateLicenseNo sets the "license_no" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertOne) UpdateLicenseNo() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateLicenseNo()
	})
}

// ClearLicenseNo clears the value of the "license_no" field.
func (u *PartnerServiceCityUpsertOne) ClearLicenseNo() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.ClearLicenseNo()
	})
}

// SetProofDocID sets the "proof_doc_id" field.
func (u *PartnerServiceCityUpsertOne) SetProofDocID(v string) *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetProofDocID(v)
	})
}

// UpdateProofDocID sets the "proof_doc_id" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertOne) UpdateProofDocID() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateProofDocID()
	})
}

// ClearProofDocID clears the value of the "proof_doc_id" field.
func (u *PartnerServiceCityUpsertOne) ClearProofDocID() *PartnerServiceCityUpsertOne {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.ClearProofDocID()
	})
}

// Exec executes the query.
func (u *PartnerServiceCityUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PartnerServiceCityCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PartnerServiceCityUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PartnerServiceCityUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PartnerServiceCityUpsertOne.ID is not supported by MySQL driver. Use PartnerServiceCityUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PartnerServiceCityUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PartnerServiceCityCreateBulk is the builder for creating many PartnerServiceCity entities in bulk.
type PartnerServiceCityCreateBulk struct {
	config
	builders []*PartnerServiceCityCreate
	conflict []sql.ConflictOption
}

// Save creates the PartnerServiceCity entities in the database.
func (psccb *PartnerServiceCityCreateBulk) Save(ctx context.Context) ([]*PartnerServiceCity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(psccb.builders))
	nodes := make([]*PartnerServiceCity, len(psccb.builders))
	mutators := make([]Mutator, len(psccb.builders))
	for i := range psccb.builders {
		func(i int, root context.Context) {
			builder := psccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PartnerServiceCityMutation)
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
					_, err = mutators[i+1].Mutate(root, psccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = psccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, psccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, psccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (psccb *PartnerServiceCityCreateBulk) SaveX(ctx context.Context) []*PartnerServiceCity {
	v, err := psccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psccb *PartnerServiceCityCreateBulk) Exec(ctx context.Context) error {
	_, err := psccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psccb *PartnerServiceCityCreateBulk) ExecX(ctx context.Context) {
	if err := psccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PartnerServiceCity.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PartnerServiceCityUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (psccb *PartnerServiceCityCreateBulk) OnConflict(opts ...sql.ConflictOption) *PartnerServiceCityUpsertBulk {
	psccb.conflict = opts
	return &PartnerServiceCityUpsertBulk{
		create: psccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PartnerServiceCity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (psccb *PartnerServiceCityCreateBulk) OnConflictColumns(columns ...string) *PartnerServiceCityUpsertBulk {
	psccb.conflict = append(psccb.conflict, sql.ConflictColumns(columns...))
	return &PartnerServiceCityUpsertBulk{
		create: psccb,
	}
}

// PartnerServiceCityUpsertBulk is the builder for "upsert"-ing
// a bulk of PartnerServiceCity nodes.
type PartnerServiceCityUpsertBulk struct {
	create *PartnerServiceCityCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PartnerServiceCity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(partnerservicecity.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PartnerServiceCityUpsertBulk) UpdateNewValues() *PartnerServiceCityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(partnerservicecity.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(partnerservicecity.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PartnerServiceCity.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PartnerServiceCityUpsertBulk) Ignore() *PartnerServiceCityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PartnerServiceCityUpsertBulk) DoNothing() *PartnerServiceCityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PartnerServiceCityCreateBulk.OnConflict
// documentation for more info.
func (u *PartnerServiceCityUpsertBulk) Update(set func(*PartnerServiceCityUpsert)) *PartnerServiceCityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PartnerServiceCityUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PartnerServiceCityUpsertBulk) SetUpdatedAt(v time.Time) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertBulk) UpdateUpdatedAt() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetPostalID sets the "postal_id" field.
func (u *PartnerServiceCityUpsertBulk) SetPostalID(v string) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetPostalID(v)
	})
}

// UpdatePostalID sets the "postal_id" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertBulk) UpdatePostalID() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdatePostalID()
	})
}

// SetActive sets the "active" field.
func (u *PartnerServiceCityUpsertBulk) SetActive(v bool) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertBulk) UpdateActive() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateActive()
	})
}

// SetName sets the "name" field.
func (u *PartnerServiceCityUpsertBulk) SetName(v string) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertBulk) UpdateName() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateName()
	})
}

// SetNaicsCode sets the "naics_code" field.
func (u *PartnerServiceCityUpsertBulk) SetNaicsCode(v uint) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetNaicsCode(v)
	})
}

// AddNaicsCode adds v to the "naics_code" field.
func (u *PartnerServiceCityUpsertBulk) AddNaicsCode(v uint) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.AddNaicsCode(v)
	})
}

// UpdateNaicsCode sets the "naics_code" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertBulk) UpdateNaicsCode() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateNaicsCode()
	})
}

// SetLicenseNo sets the "license_no" field.
func (u *PartnerServiceCityUpsertBulk) SetLicenseNo(v string) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetLicenseNo(v)
	})
}

// UpdateLicenseNo sets the "license_no" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertBulk) UpdateLicenseNo() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateLicenseNo()
	})
}

// ClearLicenseNo clears the value of the "license_no" field.
func (u *PartnerServiceCityUpsertBulk) ClearLicenseNo() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.ClearLicenseNo()
	})
}

// SetProofDocID sets the "proof_doc_id" field.
func (u *PartnerServiceCityUpsertBulk) SetProofDocID(v string) *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.SetProofDocID(v)
	})
}

// UpdateProofDocID sets the "proof_doc_id" field to the value that was provided on create.
func (u *PartnerServiceCityUpsertBulk) UpdateProofDocID() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.UpdateProofDocID()
	})
}

// ClearProofDocID clears the value of the "proof_doc_id" field.
func (u *PartnerServiceCityUpsertBulk) ClearProofDocID() *PartnerServiceCityUpsertBulk {
	return u.Update(func(s *PartnerServiceCityUpsert) {
		s.ClearProofDocID()
	})
}

// Exec executes the query.
func (u *PartnerServiceCityUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PartnerServiceCityCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PartnerServiceCityCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PartnerServiceCityUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}