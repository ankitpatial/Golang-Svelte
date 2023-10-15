// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/document"
	"roofix/ent/partnertrainingvideo"
	"roofix/ent/trainingcourse"
	"roofix/ent/trainingvideo"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrainingVideoCreate is the builder for creating a TrainingVideo entity.
type TrainingVideoCreate struct {
	config
	mutation *TrainingVideoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tvc *TrainingVideoCreate) SetCreatedAt(t time.Time) *TrainingVideoCreate {
	tvc.mutation.SetCreatedAt(t)
	return tvc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tvc *TrainingVideoCreate) SetNillableCreatedAt(t *time.Time) *TrainingVideoCreate {
	if t != nil {
		tvc.SetCreatedAt(*t)
	}
	return tvc
}

// SetUpdatedAt sets the "updated_at" field.
func (tvc *TrainingVideoCreate) SetUpdatedAt(t time.Time) *TrainingVideoCreate {
	tvc.mutation.SetUpdatedAt(t)
	return tvc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tvc *TrainingVideoCreate) SetNillableUpdatedAt(t *time.Time) *TrainingVideoCreate {
	if t != nil {
		tvc.SetUpdatedAt(*t)
	}
	return tvc
}

// SetKind sets the "kind" field.
func (tvc *TrainingVideoCreate) SetKind(et enum.TrainingType) *TrainingVideoCreate {
	tvc.mutation.SetKind(et)
	return tvc
}

// SetTitle sets the "title" field.
func (tvc *TrainingVideoCreate) SetTitle(s string) *TrainingVideoCreate {
	tvc.mutation.SetTitle(s)
	return tvc
}

// SetDescription sets the "description" field.
func (tvc *TrainingVideoCreate) SetDescription(s string) *TrainingVideoCreate {
	tvc.mutation.SetDescription(s)
	return tvc
}

// SetID sets the "id" field.
func (tvc *TrainingVideoCreate) SetID(s string) *TrainingVideoCreate {
	tvc.mutation.SetID(s)
	return tvc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tvc *TrainingVideoCreate) SetNillableID(s *string) *TrainingVideoCreate {
	if s != nil {
		tvc.SetID(*s)
	}
	return tvc
}

// AddTrainingVideoIDs adds the "training_videos" edge to the PartnerTrainingVideo entity by IDs.
func (tvc *TrainingVideoCreate) AddTrainingVideoIDs(ids ...string) *TrainingVideoCreate {
	tvc.mutation.AddTrainingVideoIDs(ids...)
	return tvc
}

// AddTrainingVideos adds the "training_videos" edges to the PartnerTrainingVideo entity.
func (tvc *TrainingVideoCreate) AddTrainingVideos(p ...*PartnerTrainingVideo) *TrainingVideoCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tvc.AddTrainingVideoIDs(ids...)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (tvc *TrainingVideoCreate) SetCreatorID(id string) *TrainingVideoCreate {
	tvc.mutation.SetCreatorID(id)
	return tvc
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (tvc *TrainingVideoCreate) SetNillableCreatorID(id *string) *TrainingVideoCreate {
	if id != nil {
		tvc = tvc.SetCreatorID(*id)
	}
	return tvc
}

// SetCreator sets the "creator" edge to the User entity.
func (tvc *TrainingVideoCreate) SetCreator(u *User) *TrainingVideoCreate {
	return tvc.SetCreatorID(u.ID)
}

// SetCourseID sets the "course" edge to the TrainingCourse entity by ID.
func (tvc *TrainingVideoCreate) SetCourseID(id string) *TrainingVideoCreate {
	tvc.mutation.SetCourseID(id)
	return tvc
}

// SetNillableCourseID sets the "course" edge to the TrainingCourse entity by ID if the given value is not nil.
func (tvc *TrainingVideoCreate) SetNillableCourseID(id *string) *TrainingVideoCreate {
	if id != nil {
		tvc = tvc.SetCourseID(*id)
	}
	return tvc
}

// SetCourse sets the "course" edge to the TrainingCourse entity.
func (tvc *TrainingVideoCreate) SetCourse(t *TrainingCourse) *TrainingVideoCreate {
	return tvc.SetCourseID(t.ID)
}

// SetPosterID sets the "poster" edge to the Document entity by ID.
func (tvc *TrainingVideoCreate) SetPosterID(id string) *TrainingVideoCreate {
	tvc.mutation.SetPosterID(id)
	return tvc
}

// SetNillablePosterID sets the "poster" edge to the Document entity by ID if the given value is not nil.
func (tvc *TrainingVideoCreate) SetNillablePosterID(id *string) *TrainingVideoCreate {
	if id != nil {
		tvc = tvc.SetPosterID(*id)
	}
	return tvc
}

// SetPoster sets the "poster" edge to the Document entity.
func (tvc *TrainingVideoCreate) SetPoster(d *Document) *TrainingVideoCreate {
	return tvc.SetPosterID(d.ID)
}

// SetVideoID sets the "video" edge to the Document entity by ID.
func (tvc *TrainingVideoCreate) SetVideoID(id string) *TrainingVideoCreate {
	tvc.mutation.SetVideoID(id)
	return tvc
}

// SetNillableVideoID sets the "video" edge to the Document entity by ID if the given value is not nil.
func (tvc *TrainingVideoCreate) SetNillableVideoID(id *string) *TrainingVideoCreate {
	if id != nil {
		tvc = tvc.SetVideoID(*id)
	}
	return tvc
}

// SetVideo sets the "video" edge to the Document entity.
func (tvc *TrainingVideoCreate) SetVideo(d *Document) *TrainingVideoCreate {
	return tvc.SetVideoID(d.ID)
}

// Mutation returns the TrainingVideoMutation object of the builder.
func (tvc *TrainingVideoCreate) Mutation() *TrainingVideoMutation {
	return tvc.mutation
}

// Save creates the TrainingVideo in the database.
func (tvc *TrainingVideoCreate) Save(ctx context.Context) (*TrainingVideo, error) {
	tvc.defaults()
	return withHooks(ctx, tvc.sqlSave, tvc.mutation, tvc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tvc *TrainingVideoCreate) SaveX(ctx context.Context) *TrainingVideo {
	v, err := tvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tvc *TrainingVideoCreate) Exec(ctx context.Context) error {
	_, err := tvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvc *TrainingVideoCreate) ExecX(ctx context.Context) {
	if err := tvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tvc *TrainingVideoCreate) defaults() {
	if _, ok := tvc.mutation.CreatedAt(); !ok {
		v := trainingvideo.DefaultCreatedAt()
		tvc.mutation.SetCreatedAt(v)
	}
	if _, ok := tvc.mutation.UpdatedAt(); !ok {
		v := trainingvideo.DefaultUpdatedAt()
		tvc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tvc.mutation.ID(); !ok {
		v := trainingvideo.DefaultID()
		tvc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tvc *TrainingVideoCreate) check() error {
	if _, ok := tvc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TrainingVideo.created_at"`)}
	}
	if _, ok := tvc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TrainingVideo.updated_at"`)}
	}
	if _, ok := tvc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "TrainingVideo.kind"`)}
	}
	if v, ok := tvc.mutation.Kind(); ok {
		if err := trainingvideo.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.kind": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "TrainingVideo.title"`)}
	}
	if v, ok := tvc.mutation.Title(); ok {
		if err := trainingvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.title": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "TrainingVideo.description"`)}
	}
	if v, ok := tvc.mutation.Description(); ok {
		if err := trainingvideo.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.description": %w`, err)}
		}
	}
	if v, ok := tvc.mutation.ID(); ok {
		if err := trainingvideo.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.id": %w`, err)}
		}
	}
	return nil
}

func (tvc *TrainingVideoCreate) sqlSave(ctx context.Context) (*TrainingVideo, error) {
	if err := tvc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TrainingVideo.ID type: %T", _spec.ID.Value)
		}
	}
	tvc.mutation.id = &_node.ID
	tvc.mutation.done = true
	return _node, nil
}

func (tvc *TrainingVideoCreate) createSpec() (*TrainingVideo, *sqlgraph.CreateSpec) {
	var (
		_node = &TrainingVideo{config: tvc.config}
		_spec = sqlgraph.NewCreateSpec(trainingvideo.Table, sqlgraph.NewFieldSpec(trainingvideo.FieldID, field.TypeString))
	)
	_spec.OnConflict = tvc.conflict
	if id, ok := tvc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tvc.mutation.CreatedAt(); ok {
		_spec.SetField(trainingvideo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tvc.mutation.UpdatedAt(); ok {
		_spec.SetField(trainingvideo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tvc.mutation.Kind(); ok {
		_spec.SetField(trainingvideo.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := tvc.mutation.Title(); ok {
		_spec.SetField(trainingvideo.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tvc.mutation.Description(); ok {
		_spec.SetField(trainingvideo.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := tvc.mutation.TrainingVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   trainingvideo.TrainingVideosTable,
			Columns: []string{trainingvideo.TrainingVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(partnertrainingvideo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tvc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   trainingvideo.CreatorTable,
			Columns: []string{trainingvideo.CreatorColumn},
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
	if nodes := tvc.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   trainingvideo.CourseTable,
			Columns: []string{trainingvideo.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trainingcourse.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.course_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tvc.mutation.PosterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   trainingvideo.PosterTable,
			Columns: []string{trainingvideo.PosterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.poster_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tvc.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   trainingvideo.VideoTable,
			Columns: []string{trainingvideo.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.video_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TrainingVideo.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TrainingVideoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tvc *TrainingVideoCreate) OnConflict(opts ...sql.ConflictOption) *TrainingVideoUpsertOne {
	tvc.conflict = opts
	return &TrainingVideoUpsertOne{
		create: tvc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TrainingVideo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tvc *TrainingVideoCreate) OnConflictColumns(columns ...string) *TrainingVideoUpsertOne {
	tvc.conflict = append(tvc.conflict, sql.ConflictColumns(columns...))
	return &TrainingVideoUpsertOne{
		create: tvc,
	}
}

type (
	// TrainingVideoUpsertOne is the builder for "upsert"-ing
	//  one TrainingVideo node.
	TrainingVideoUpsertOne struct {
		create *TrainingVideoCreate
	}

	// TrainingVideoUpsert is the "OnConflict" setter.
	TrainingVideoUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *TrainingVideoUpsert) SetUpdatedAt(v time.Time) *TrainingVideoUpsert {
	u.Set(trainingvideo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TrainingVideoUpsert) UpdateUpdatedAt() *TrainingVideoUpsert {
	u.SetExcluded(trainingvideo.FieldUpdatedAt)
	return u
}

// SetKind sets the "kind" field.
func (u *TrainingVideoUpsert) SetKind(v enum.TrainingType) *TrainingVideoUpsert {
	u.Set(trainingvideo.FieldKind, v)
	return u
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *TrainingVideoUpsert) UpdateKind() *TrainingVideoUpsert {
	u.SetExcluded(trainingvideo.FieldKind)
	return u
}

// SetTitle sets the "title" field.
func (u *TrainingVideoUpsert) SetTitle(v string) *TrainingVideoUpsert {
	u.Set(trainingvideo.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TrainingVideoUpsert) UpdateTitle() *TrainingVideoUpsert {
	u.SetExcluded(trainingvideo.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *TrainingVideoUpsert) SetDescription(v string) *TrainingVideoUpsert {
	u.Set(trainingvideo.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TrainingVideoUpsert) UpdateDescription() *TrainingVideoUpsert {
	u.SetExcluded(trainingvideo.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TrainingVideo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(trainingvideo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TrainingVideoUpsertOne) UpdateNewValues() *TrainingVideoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(trainingvideo.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(trainingvideo.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TrainingVideo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TrainingVideoUpsertOne) Ignore() *TrainingVideoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TrainingVideoUpsertOne) DoNothing() *TrainingVideoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TrainingVideoCreate.OnConflict
// documentation for more info.
func (u *TrainingVideoUpsertOne) Update(set func(*TrainingVideoUpsert)) *TrainingVideoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TrainingVideoUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TrainingVideoUpsertOne) SetUpdatedAt(v time.Time) *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TrainingVideoUpsertOne) UpdateUpdatedAt() *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetKind sets the "kind" field.
func (u *TrainingVideoUpsertOne) SetKind(v enum.TrainingType) *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *TrainingVideoUpsertOne) UpdateKind() *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateKind()
	})
}

// SetTitle sets the "title" field.
func (u *TrainingVideoUpsertOne) SetTitle(v string) *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TrainingVideoUpsertOne) UpdateTitle() *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *TrainingVideoUpsertOne) SetDescription(v string) *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TrainingVideoUpsertOne) UpdateDescription() *TrainingVideoUpsertOne {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *TrainingVideoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TrainingVideoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TrainingVideoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TrainingVideoUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TrainingVideoUpsertOne.ID is not supported by MySQL driver. Use TrainingVideoUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TrainingVideoUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TrainingVideoCreateBulk is the builder for creating many TrainingVideo entities in bulk.
type TrainingVideoCreateBulk struct {
	config
	builders []*TrainingVideoCreate
	conflict []sql.ConflictOption
}

// Save creates the TrainingVideo entities in the database.
func (tvcb *TrainingVideoCreateBulk) Save(ctx context.Context) ([]*TrainingVideo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tvcb.builders))
	nodes := make([]*TrainingVideo, len(tvcb.builders))
	mutators := make([]Mutator, len(tvcb.builders))
	for i := range tvcb.builders {
		func(i int, root context.Context) {
			builder := tvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TrainingVideoMutation)
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
					_, err = mutators[i+1].Mutate(root, tvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tvcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tvcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tvcb *TrainingVideoCreateBulk) SaveX(ctx context.Context) []*TrainingVideo {
	v, err := tvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tvcb *TrainingVideoCreateBulk) Exec(ctx context.Context) error {
	_, err := tvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvcb *TrainingVideoCreateBulk) ExecX(ctx context.Context) {
	if err := tvcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TrainingVideo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TrainingVideoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tvcb *TrainingVideoCreateBulk) OnConflict(opts ...sql.ConflictOption) *TrainingVideoUpsertBulk {
	tvcb.conflict = opts
	return &TrainingVideoUpsertBulk{
		create: tvcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TrainingVideo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tvcb *TrainingVideoCreateBulk) OnConflictColumns(columns ...string) *TrainingVideoUpsertBulk {
	tvcb.conflict = append(tvcb.conflict, sql.ConflictColumns(columns...))
	return &TrainingVideoUpsertBulk{
		create: tvcb,
	}
}

// TrainingVideoUpsertBulk is the builder for "upsert"-ing
// a bulk of TrainingVideo nodes.
type TrainingVideoUpsertBulk struct {
	create *TrainingVideoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TrainingVideo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(trainingvideo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TrainingVideoUpsertBulk) UpdateNewValues() *TrainingVideoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(trainingvideo.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(trainingvideo.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TrainingVideo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TrainingVideoUpsertBulk) Ignore() *TrainingVideoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TrainingVideoUpsertBulk) DoNothing() *TrainingVideoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TrainingVideoCreateBulk.OnConflict
// documentation for more info.
func (u *TrainingVideoUpsertBulk) Update(set func(*TrainingVideoUpsert)) *TrainingVideoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TrainingVideoUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TrainingVideoUpsertBulk) SetUpdatedAt(v time.Time) *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TrainingVideoUpsertBulk) UpdateUpdatedAt() *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetKind sets the "kind" field.
func (u *TrainingVideoUpsertBulk) SetKind(v enum.TrainingType) *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *TrainingVideoUpsertBulk) UpdateKind() *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateKind()
	})
}

// SetTitle sets the "title" field.
func (u *TrainingVideoUpsertBulk) SetTitle(v string) *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TrainingVideoUpsertBulk) UpdateTitle() *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *TrainingVideoUpsertBulk) SetDescription(v string) *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TrainingVideoUpsertBulk) UpdateDescription() *TrainingVideoUpsertBulk {
	return u.Update(func(s *TrainingVideoUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *TrainingVideoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TrainingVideoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TrainingVideoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TrainingVideoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}