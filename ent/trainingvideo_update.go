// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent/document"
	"roofix/ent/partnertrainingvideo"
	"roofix/ent/predicate"
	"roofix/ent/trainingcourse"
	"roofix/ent/trainingvideo"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrainingVideoUpdate is the builder for updating TrainingVideo entities.
type TrainingVideoUpdate struct {
	config
	hooks     []Hook
	mutation  *TrainingVideoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TrainingVideoUpdate builder.
func (tvu *TrainingVideoUpdate) Where(ps ...predicate.TrainingVideo) *TrainingVideoUpdate {
	tvu.mutation.Where(ps...)
	return tvu
}

// SetUpdatedAt sets the "updated_at" field.
func (tvu *TrainingVideoUpdate) SetUpdatedAt(t time.Time) *TrainingVideoUpdate {
	tvu.mutation.SetUpdatedAt(t)
	return tvu
}

// SetKind sets the "kind" field.
func (tvu *TrainingVideoUpdate) SetKind(et enum.TrainingType) *TrainingVideoUpdate {
	tvu.mutation.SetKind(et)
	return tvu
}

// SetTitle sets the "title" field.
func (tvu *TrainingVideoUpdate) SetTitle(s string) *TrainingVideoUpdate {
	tvu.mutation.SetTitle(s)
	return tvu
}

// SetDescription sets the "description" field.
func (tvu *TrainingVideoUpdate) SetDescription(s string) *TrainingVideoUpdate {
	tvu.mutation.SetDescription(s)
	return tvu
}

// AddTrainingVideoIDs adds the "training_videos" edge to the PartnerTrainingVideo entity by IDs.
func (tvu *TrainingVideoUpdate) AddTrainingVideoIDs(ids ...string) *TrainingVideoUpdate {
	tvu.mutation.AddTrainingVideoIDs(ids...)
	return tvu
}

// AddTrainingVideos adds the "training_videos" edges to the PartnerTrainingVideo entity.
func (tvu *TrainingVideoUpdate) AddTrainingVideos(p ...*PartnerTrainingVideo) *TrainingVideoUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tvu.AddTrainingVideoIDs(ids...)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (tvu *TrainingVideoUpdate) SetCreatorID(id string) *TrainingVideoUpdate {
	tvu.mutation.SetCreatorID(id)
	return tvu
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (tvu *TrainingVideoUpdate) SetNillableCreatorID(id *string) *TrainingVideoUpdate {
	if id != nil {
		tvu = tvu.SetCreatorID(*id)
	}
	return tvu
}

// SetCreator sets the "creator" edge to the User entity.
func (tvu *TrainingVideoUpdate) SetCreator(u *User) *TrainingVideoUpdate {
	return tvu.SetCreatorID(u.ID)
}

// SetCourseID sets the "course" edge to the TrainingCourse entity by ID.
func (tvu *TrainingVideoUpdate) SetCourseID(id string) *TrainingVideoUpdate {
	tvu.mutation.SetCourseID(id)
	return tvu
}

// SetNillableCourseID sets the "course" edge to the TrainingCourse entity by ID if the given value is not nil.
func (tvu *TrainingVideoUpdate) SetNillableCourseID(id *string) *TrainingVideoUpdate {
	if id != nil {
		tvu = tvu.SetCourseID(*id)
	}
	return tvu
}

// SetCourse sets the "course" edge to the TrainingCourse entity.
func (tvu *TrainingVideoUpdate) SetCourse(t *TrainingCourse) *TrainingVideoUpdate {
	return tvu.SetCourseID(t.ID)
}

// SetPosterID sets the "poster" edge to the Document entity by ID.
func (tvu *TrainingVideoUpdate) SetPosterID(id string) *TrainingVideoUpdate {
	tvu.mutation.SetPosterID(id)
	return tvu
}

// SetNillablePosterID sets the "poster" edge to the Document entity by ID if the given value is not nil.
func (tvu *TrainingVideoUpdate) SetNillablePosterID(id *string) *TrainingVideoUpdate {
	if id != nil {
		tvu = tvu.SetPosterID(*id)
	}
	return tvu
}

// SetPoster sets the "poster" edge to the Document entity.
func (tvu *TrainingVideoUpdate) SetPoster(d *Document) *TrainingVideoUpdate {
	return tvu.SetPosterID(d.ID)
}

// SetVideoID sets the "video" edge to the Document entity by ID.
func (tvu *TrainingVideoUpdate) SetVideoID(id string) *TrainingVideoUpdate {
	tvu.mutation.SetVideoID(id)
	return tvu
}

// SetNillableVideoID sets the "video" edge to the Document entity by ID if the given value is not nil.
func (tvu *TrainingVideoUpdate) SetNillableVideoID(id *string) *TrainingVideoUpdate {
	if id != nil {
		tvu = tvu.SetVideoID(*id)
	}
	return tvu
}

// SetVideo sets the "video" edge to the Document entity.
func (tvu *TrainingVideoUpdate) SetVideo(d *Document) *TrainingVideoUpdate {
	return tvu.SetVideoID(d.ID)
}

// Mutation returns the TrainingVideoMutation object of the builder.
func (tvu *TrainingVideoUpdate) Mutation() *TrainingVideoMutation {
	return tvu.mutation
}

// ClearTrainingVideos clears all "training_videos" edges to the PartnerTrainingVideo entity.
func (tvu *TrainingVideoUpdate) ClearTrainingVideos() *TrainingVideoUpdate {
	tvu.mutation.ClearTrainingVideos()
	return tvu
}

// RemoveTrainingVideoIDs removes the "training_videos" edge to PartnerTrainingVideo entities by IDs.
func (tvu *TrainingVideoUpdate) RemoveTrainingVideoIDs(ids ...string) *TrainingVideoUpdate {
	tvu.mutation.RemoveTrainingVideoIDs(ids...)
	return tvu
}

// RemoveTrainingVideos removes "training_videos" edges to PartnerTrainingVideo entities.
func (tvu *TrainingVideoUpdate) RemoveTrainingVideos(p ...*PartnerTrainingVideo) *TrainingVideoUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tvu.RemoveTrainingVideoIDs(ids...)
}

// ClearCreator clears the "creator" edge to the User entity.
func (tvu *TrainingVideoUpdate) ClearCreator() *TrainingVideoUpdate {
	tvu.mutation.ClearCreator()
	return tvu
}

// ClearCourse clears the "course" edge to the TrainingCourse entity.
func (tvu *TrainingVideoUpdate) ClearCourse() *TrainingVideoUpdate {
	tvu.mutation.ClearCourse()
	return tvu
}

// ClearPoster clears the "poster" edge to the Document entity.
func (tvu *TrainingVideoUpdate) ClearPoster() *TrainingVideoUpdate {
	tvu.mutation.ClearPoster()
	return tvu
}

// ClearVideo clears the "video" edge to the Document entity.
func (tvu *TrainingVideoUpdate) ClearVideo() *TrainingVideoUpdate {
	tvu.mutation.ClearVideo()
	return tvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tvu *TrainingVideoUpdate) Save(ctx context.Context) (int, error) {
	tvu.defaults()
	return withHooks(ctx, tvu.sqlSave, tvu.mutation, tvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tvu *TrainingVideoUpdate) SaveX(ctx context.Context) int {
	affected, err := tvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tvu *TrainingVideoUpdate) Exec(ctx context.Context) error {
	_, err := tvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvu *TrainingVideoUpdate) ExecX(ctx context.Context) {
	if err := tvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tvu *TrainingVideoUpdate) defaults() {
	if _, ok := tvu.mutation.UpdatedAt(); !ok {
		v := trainingvideo.UpdateDefaultUpdatedAt()
		tvu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tvu *TrainingVideoUpdate) check() error {
	if v, ok := tvu.mutation.Kind(); ok {
		if err := trainingvideo.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.kind": %w`, err)}
		}
	}
	if v, ok := tvu.mutation.Title(); ok {
		if err := trainingvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.title": %w`, err)}
		}
	}
	if v, ok := tvu.mutation.Description(); ok {
		if err := trainingvideo.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.description": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tvu *TrainingVideoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TrainingVideoUpdate {
	tvu.modifiers = append(tvu.modifiers, modifiers...)
	return tvu
}

func (tvu *TrainingVideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(trainingvideo.Table, trainingvideo.Columns, sqlgraph.NewFieldSpec(trainingvideo.FieldID, field.TypeString))
	if ps := tvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tvu.mutation.UpdatedAt(); ok {
		_spec.SetField(trainingvideo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tvu.mutation.Kind(); ok {
		_spec.SetField(trainingvideo.FieldKind, field.TypeEnum, value)
	}
	if value, ok := tvu.mutation.Title(); ok {
		_spec.SetField(trainingvideo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tvu.mutation.Description(); ok {
		_spec.SetField(trainingvideo.FieldDescription, field.TypeString, value)
	}
	if tvu.mutation.TrainingVideosCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.RemovedTrainingVideosIDs(); len(nodes) > 0 && !tvu.mutation.TrainingVideosCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.TrainingVideosIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvu.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvu.mutation.CourseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.CourseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvu.mutation.PosterCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.PosterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvu.mutation.VideoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvu.mutation.VideoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tvu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trainingvideo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tvu.mutation.done = true
	return n, nil
}

// TrainingVideoUpdateOne is the builder for updating a single TrainingVideo entity.
type TrainingVideoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TrainingVideoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tvuo *TrainingVideoUpdateOne) SetUpdatedAt(t time.Time) *TrainingVideoUpdateOne {
	tvuo.mutation.SetUpdatedAt(t)
	return tvuo
}

// SetKind sets the "kind" field.
func (tvuo *TrainingVideoUpdateOne) SetKind(et enum.TrainingType) *TrainingVideoUpdateOne {
	tvuo.mutation.SetKind(et)
	return tvuo
}

// SetTitle sets the "title" field.
func (tvuo *TrainingVideoUpdateOne) SetTitle(s string) *TrainingVideoUpdateOne {
	tvuo.mutation.SetTitle(s)
	return tvuo
}

// SetDescription sets the "description" field.
func (tvuo *TrainingVideoUpdateOne) SetDescription(s string) *TrainingVideoUpdateOne {
	tvuo.mutation.SetDescription(s)
	return tvuo
}

// AddTrainingVideoIDs adds the "training_videos" edge to the PartnerTrainingVideo entity by IDs.
func (tvuo *TrainingVideoUpdateOne) AddTrainingVideoIDs(ids ...string) *TrainingVideoUpdateOne {
	tvuo.mutation.AddTrainingVideoIDs(ids...)
	return tvuo
}

// AddTrainingVideos adds the "training_videos" edges to the PartnerTrainingVideo entity.
func (tvuo *TrainingVideoUpdateOne) AddTrainingVideos(p ...*PartnerTrainingVideo) *TrainingVideoUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tvuo.AddTrainingVideoIDs(ids...)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (tvuo *TrainingVideoUpdateOne) SetCreatorID(id string) *TrainingVideoUpdateOne {
	tvuo.mutation.SetCreatorID(id)
	return tvuo
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (tvuo *TrainingVideoUpdateOne) SetNillableCreatorID(id *string) *TrainingVideoUpdateOne {
	if id != nil {
		tvuo = tvuo.SetCreatorID(*id)
	}
	return tvuo
}

// SetCreator sets the "creator" edge to the User entity.
func (tvuo *TrainingVideoUpdateOne) SetCreator(u *User) *TrainingVideoUpdateOne {
	return tvuo.SetCreatorID(u.ID)
}

// SetCourseID sets the "course" edge to the TrainingCourse entity by ID.
func (tvuo *TrainingVideoUpdateOne) SetCourseID(id string) *TrainingVideoUpdateOne {
	tvuo.mutation.SetCourseID(id)
	return tvuo
}

// SetNillableCourseID sets the "course" edge to the TrainingCourse entity by ID if the given value is not nil.
func (tvuo *TrainingVideoUpdateOne) SetNillableCourseID(id *string) *TrainingVideoUpdateOne {
	if id != nil {
		tvuo = tvuo.SetCourseID(*id)
	}
	return tvuo
}

// SetCourse sets the "course" edge to the TrainingCourse entity.
func (tvuo *TrainingVideoUpdateOne) SetCourse(t *TrainingCourse) *TrainingVideoUpdateOne {
	return tvuo.SetCourseID(t.ID)
}

// SetPosterID sets the "poster" edge to the Document entity by ID.
func (tvuo *TrainingVideoUpdateOne) SetPosterID(id string) *TrainingVideoUpdateOne {
	tvuo.mutation.SetPosterID(id)
	return tvuo
}

// SetNillablePosterID sets the "poster" edge to the Document entity by ID if the given value is not nil.
func (tvuo *TrainingVideoUpdateOne) SetNillablePosterID(id *string) *TrainingVideoUpdateOne {
	if id != nil {
		tvuo = tvuo.SetPosterID(*id)
	}
	return tvuo
}

// SetPoster sets the "poster" edge to the Document entity.
func (tvuo *TrainingVideoUpdateOne) SetPoster(d *Document) *TrainingVideoUpdateOne {
	return tvuo.SetPosterID(d.ID)
}

// SetVideoID sets the "video" edge to the Document entity by ID.
func (tvuo *TrainingVideoUpdateOne) SetVideoID(id string) *TrainingVideoUpdateOne {
	tvuo.mutation.SetVideoID(id)
	return tvuo
}

// SetNillableVideoID sets the "video" edge to the Document entity by ID if the given value is not nil.
func (tvuo *TrainingVideoUpdateOne) SetNillableVideoID(id *string) *TrainingVideoUpdateOne {
	if id != nil {
		tvuo = tvuo.SetVideoID(*id)
	}
	return tvuo
}

// SetVideo sets the "video" edge to the Document entity.
func (tvuo *TrainingVideoUpdateOne) SetVideo(d *Document) *TrainingVideoUpdateOne {
	return tvuo.SetVideoID(d.ID)
}

// Mutation returns the TrainingVideoMutation object of the builder.
func (tvuo *TrainingVideoUpdateOne) Mutation() *TrainingVideoMutation {
	return tvuo.mutation
}

// ClearTrainingVideos clears all "training_videos" edges to the PartnerTrainingVideo entity.
func (tvuo *TrainingVideoUpdateOne) ClearTrainingVideos() *TrainingVideoUpdateOne {
	tvuo.mutation.ClearTrainingVideos()
	return tvuo
}

// RemoveTrainingVideoIDs removes the "training_videos" edge to PartnerTrainingVideo entities by IDs.
func (tvuo *TrainingVideoUpdateOne) RemoveTrainingVideoIDs(ids ...string) *TrainingVideoUpdateOne {
	tvuo.mutation.RemoveTrainingVideoIDs(ids...)
	return tvuo
}

// RemoveTrainingVideos removes "training_videos" edges to PartnerTrainingVideo entities.
func (tvuo *TrainingVideoUpdateOne) RemoveTrainingVideos(p ...*PartnerTrainingVideo) *TrainingVideoUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tvuo.RemoveTrainingVideoIDs(ids...)
}

// ClearCreator clears the "creator" edge to the User entity.
func (tvuo *TrainingVideoUpdateOne) ClearCreator() *TrainingVideoUpdateOne {
	tvuo.mutation.ClearCreator()
	return tvuo
}

// ClearCourse clears the "course" edge to the TrainingCourse entity.
func (tvuo *TrainingVideoUpdateOne) ClearCourse() *TrainingVideoUpdateOne {
	tvuo.mutation.ClearCourse()
	return tvuo
}

// ClearPoster clears the "poster" edge to the Document entity.
func (tvuo *TrainingVideoUpdateOne) ClearPoster() *TrainingVideoUpdateOne {
	tvuo.mutation.ClearPoster()
	return tvuo
}

// ClearVideo clears the "video" edge to the Document entity.
func (tvuo *TrainingVideoUpdateOne) ClearVideo() *TrainingVideoUpdateOne {
	tvuo.mutation.ClearVideo()
	return tvuo
}

// Where appends a list predicates to the TrainingVideoUpdate builder.
func (tvuo *TrainingVideoUpdateOne) Where(ps ...predicate.TrainingVideo) *TrainingVideoUpdateOne {
	tvuo.mutation.Where(ps...)
	return tvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tvuo *TrainingVideoUpdateOne) Select(field string, fields ...string) *TrainingVideoUpdateOne {
	tvuo.fields = append([]string{field}, fields...)
	return tvuo
}

// Save executes the query and returns the updated TrainingVideo entity.
func (tvuo *TrainingVideoUpdateOne) Save(ctx context.Context) (*TrainingVideo, error) {
	tvuo.defaults()
	return withHooks(ctx, tvuo.sqlSave, tvuo.mutation, tvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tvuo *TrainingVideoUpdateOne) SaveX(ctx context.Context) *TrainingVideo {
	node, err := tvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tvuo *TrainingVideoUpdateOne) Exec(ctx context.Context) error {
	_, err := tvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvuo *TrainingVideoUpdateOne) ExecX(ctx context.Context) {
	if err := tvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tvuo *TrainingVideoUpdateOne) defaults() {
	if _, ok := tvuo.mutation.UpdatedAt(); !ok {
		v := trainingvideo.UpdateDefaultUpdatedAt()
		tvuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tvuo *TrainingVideoUpdateOne) check() error {
	if v, ok := tvuo.mutation.Kind(); ok {
		if err := trainingvideo.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.kind": %w`, err)}
		}
	}
	if v, ok := tvuo.mutation.Title(); ok {
		if err := trainingvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.title": %w`, err)}
		}
	}
	if v, ok := tvuo.mutation.Description(); ok {
		if err := trainingvideo.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "TrainingVideo.description": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tvuo *TrainingVideoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TrainingVideoUpdateOne {
	tvuo.modifiers = append(tvuo.modifiers, modifiers...)
	return tvuo
}

func (tvuo *TrainingVideoUpdateOne) sqlSave(ctx context.Context) (_node *TrainingVideo, err error) {
	if err := tvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(trainingvideo.Table, trainingvideo.Columns, sqlgraph.NewFieldSpec(trainingvideo.FieldID, field.TypeString))
	id, ok := tvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TrainingVideo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trainingvideo.FieldID)
		for _, f := range fields {
			if !trainingvideo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != trainingvideo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tvuo.mutation.UpdatedAt(); ok {
		_spec.SetField(trainingvideo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tvuo.mutation.Kind(); ok {
		_spec.SetField(trainingvideo.FieldKind, field.TypeEnum, value)
	}
	if value, ok := tvuo.mutation.Title(); ok {
		_spec.SetField(trainingvideo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tvuo.mutation.Description(); ok {
		_spec.SetField(trainingvideo.FieldDescription, field.TypeString, value)
	}
	if tvuo.mutation.TrainingVideosCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.RemovedTrainingVideosIDs(); len(nodes) > 0 && !tvuo.mutation.TrainingVideosCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.TrainingVideosIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvuo.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvuo.mutation.CourseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.CourseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvuo.mutation.PosterCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.PosterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tvuo.mutation.VideoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tvuo.mutation.VideoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tvuo.modifiers...)
	_node = &TrainingVideo{config: tvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trainingvideo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tvuo.mutation.done = true
	return _node, nil
}
