// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/partner"
	"roofix/ent/partnertrainingvideo"
	"roofix/ent/predicate"
	"roofix/ent/trainingvideo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PartnerTrainingVideoQuery is the builder for querying PartnerTrainingVideo entities.
type PartnerTrainingVideoQuery struct {
	config
	ctx         *QueryContext
	order       []partnertrainingvideo.OrderOption
	inters      []Interceptor
	predicates  []predicate.PartnerTrainingVideo
	withVideo   *TrainingVideoQuery
	withPartner *PartnerQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*PartnerTrainingVideo) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PartnerTrainingVideoQuery builder.
func (ptvq *PartnerTrainingVideoQuery) Where(ps ...predicate.PartnerTrainingVideo) *PartnerTrainingVideoQuery {
	ptvq.predicates = append(ptvq.predicates, ps...)
	return ptvq
}

// Limit the number of records to be returned by this query.
func (ptvq *PartnerTrainingVideoQuery) Limit(limit int) *PartnerTrainingVideoQuery {
	ptvq.ctx.Limit = &limit
	return ptvq
}

// Offset to start from.
func (ptvq *PartnerTrainingVideoQuery) Offset(offset int) *PartnerTrainingVideoQuery {
	ptvq.ctx.Offset = &offset
	return ptvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptvq *PartnerTrainingVideoQuery) Unique(unique bool) *PartnerTrainingVideoQuery {
	ptvq.ctx.Unique = &unique
	return ptvq
}

// Order specifies how the records should be ordered.
func (ptvq *PartnerTrainingVideoQuery) Order(o ...partnertrainingvideo.OrderOption) *PartnerTrainingVideoQuery {
	ptvq.order = append(ptvq.order, o...)
	return ptvq
}

// QueryVideo chains the current query on the "video" edge.
func (ptvq *PartnerTrainingVideoQuery) QueryVideo() *TrainingVideoQuery {
	query := (&TrainingVideoClient{config: ptvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partnertrainingvideo.Table, partnertrainingvideo.FieldID, selector),
			sqlgraph.To(trainingvideo.Table, trainingvideo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, partnertrainingvideo.VideoTable, partnertrainingvideo.VideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPartner chains the current query on the "partner" edge.
func (ptvq *PartnerTrainingVideoQuery) QueryPartner() *PartnerQuery {
	query := (&PartnerClient{config: ptvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partnertrainingvideo.Table, partnertrainingvideo.FieldID, selector),
			sqlgraph.To(partner.Table, partner.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, partnertrainingvideo.PartnerTable, partnertrainingvideo.PartnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PartnerTrainingVideo entity from the query.
// Returns a *NotFoundError when no PartnerTrainingVideo was found.
func (ptvq *PartnerTrainingVideoQuery) First(ctx context.Context) (*PartnerTrainingVideo, error) {
	nodes, err := ptvq.Limit(1).All(setContextOp(ctx, ptvq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{partnertrainingvideo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) FirstX(ctx context.Context) *PartnerTrainingVideo {
	node, err := ptvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PartnerTrainingVideo ID from the query.
// Returns a *NotFoundError when no PartnerTrainingVideo ID was found.
func (ptvq *PartnerTrainingVideoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ptvq.Limit(1).IDs(setContextOp(ctx, ptvq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{partnertrainingvideo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) FirstIDX(ctx context.Context) string {
	id, err := ptvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PartnerTrainingVideo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PartnerTrainingVideo entity is found.
// Returns a *NotFoundError when no PartnerTrainingVideo entities are found.
func (ptvq *PartnerTrainingVideoQuery) Only(ctx context.Context) (*PartnerTrainingVideo, error) {
	nodes, err := ptvq.Limit(2).All(setContextOp(ctx, ptvq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{partnertrainingvideo.Label}
	default:
		return nil, &NotSingularError{partnertrainingvideo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) OnlyX(ctx context.Context) *PartnerTrainingVideo {
	node, err := ptvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PartnerTrainingVideo ID in the query.
// Returns a *NotSingularError when more than one PartnerTrainingVideo ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptvq *PartnerTrainingVideoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ptvq.Limit(2).IDs(setContextOp(ctx, ptvq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{partnertrainingvideo.Label}
	default:
		err = &NotSingularError{partnertrainingvideo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) OnlyIDX(ctx context.Context) string {
	id, err := ptvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PartnerTrainingVideos.
func (ptvq *PartnerTrainingVideoQuery) All(ctx context.Context) ([]*PartnerTrainingVideo, error) {
	ctx = setContextOp(ctx, ptvq.ctx, "All")
	if err := ptvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PartnerTrainingVideo, *PartnerTrainingVideoQuery]()
	return withInterceptors[[]*PartnerTrainingVideo](ctx, ptvq, qr, ptvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) AllX(ctx context.Context) []*PartnerTrainingVideo {
	nodes, err := ptvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PartnerTrainingVideo IDs.
func (ptvq *PartnerTrainingVideoQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ptvq.ctx.Unique == nil && ptvq.path != nil {
		ptvq.Unique(true)
	}
	ctx = setContextOp(ctx, ptvq.ctx, "IDs")
	if err = ptvq.Select(partnertrainingvideo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) IDsX(ctx context.Context) []string {
	ids, err := ptvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptvq *PartnerTrainingVideoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptvq.ctx, "Count")
	if err := ptvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptvq, querierCount[*PartnerTrainingVideoQuery](), ptvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) CountX(ctx context.Context) int {
	count, err := ptvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptvq *PartnerTrainingVideoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptvq.ctx, "Exist")
	switch _, err := ptvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptvq *PartnerTrainingVideoQuery) ExistX(ctx context.Context) bool {
	exist, err := ptvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PartnerTrainingVideoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptvq *PartnerTrainingVideoQuery) Clone() *PartnerTrainingVideoQuery {
	if ptvq == nil {
		return nil
	}
	return &PartnerTrainingVideoQuery{
		config:      ptvq.config,
		ctx:         ptvq.ctx.Clone(),
		order:       append([]partnertrainingvideo.OrderOption{}, ptvq.order...),
		inters:      append([]Interceptor{}, ptvq.inters...),
		predicates:  append([]predicate.PartnerTrainingVideo{}, ptvq.predicates...),
		withVideo:   ptvq.withVideo.Clone(),
		withPartner: ptvq.withPartner.Clone(),
		// clone intermediate query.
		sql:  ptvq.sql.Clone(),
		path: ptvq.path,
	}
}

// WithVideo tells the query-builder to eager-load the nodes that are connected to
// the "video" edge. The optional arguments are used to configure the query builder of the edge.
func (ptvq *PartnerTrainingVideoQuery) WithVideo(opts ...func(*TrainingVideoQuery)) *PartnerTrainingVideoQuery {
	query := (&TrainingVideoClient{config: ptvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptvq.withVideo = query
	return ptvq
}

// WithPartner tells the query-builder to eager-load the nodes that are connected to
// the "partner" edge. The optional arguments are used to configure the query builder of the edge.
func (ptvq *PartnerTrainingVideoQuery) WithPartner(opts ...func(*PartnerQuery)) *PartnerTrainingVideoQuery {
	query := (&PartnerClient{config: ptvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptvq.withPartner = query
	return ptvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PartnerTrainingVideo.Query().
//		GroupBy(partnertrainingvideo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptvq *PartnerTrainingVideoQuery) GroupBy(field string, fields ...string) *PartnerTrainingVideoGroupBy {
	ptvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PartnerTrainingVideoGroupBy{build: ptvq}
	grbuild.flds = &ptvq.ctx.Fields
	grbuild.label = partnertrainingvideo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//	}
//
//	client.PartnerTrainingVideo.Query().
//		Select(partnertrainingvideo.FieldCreatedAt).
//		Scan(ctx, &v)
func (ptvq *PartnerTrainingVideoQuery) Select(fields ...string) *PartnerTrainingVideoSelect {
	ptvq.ctx.Fields = append(ptvq.ctx.Fields, fields...)
	sbuild := &PartnerTrainingVideoSelect{PartnerTrainingVideoQuery: ptvq}
	sbuild.label = partnertrainingvideo.Label
	sbuild.flds, sbuild.scan = &ptvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PartnerTrainingVideoSelect configured with the given aggregations.
func (ptvq *PartnerTrainingVideoQuery) Aggregate(fns ...AggregateFunc) *PartnerTrainingVideoSelect {
	return ptvq.Select().Aggregate(fns...)
}

func (ptvq *PartnerTrainingVideoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptvq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptvq.ctx.Fields {
		if !partnertrainingvideo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptvq.path != nil {
		prev, err := ptvq.path(ctx)
		if err != nil {
			return err
		}
		ptvq.sql = prev
	}
	return nil
}

func (ptvq *PartnerTrainingVideoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PartnerTrainingVideo, error) {
	var (
		nodes       = []*PartnerTrainingVideo{}
		withFKs     = ptvq.withFKs
		_spec       = ptvq.querySpec()
		loadedTypes = [2]bool{
			ptvq.withVideo != nil,
			ptvq.withPartner != nil,
		}
	)
	if ptvq.withVideo != nil || ptvq.withPartner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, partnertrainingvideo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PartnerTrainingVideo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PartnerTrainingVideo{config: ptvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ptvq.modifiers) > 0 {
		_spec.Modifiers = ptvq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptvq.withVideo; query != nil {
		if err := ptvq.loadVideo(ctx, query, nodes, nil,
			func(n *PartnerTrainingVideo, e *TrainingVideo) { n.Edges.Video = e }); err != nil {
			return nil, err
		}
	}
	if query := ptvq.withPartner; query != nil {
		if err := ptvq.loadPartner(ctx, query, nodes, nil,
			func(n *PartnerTrainingVideo, e *Partner) { n.Edges.Partner = e }); err != nil {
			return nil, err
		}
	}
	for i := range ptvq.loadTotal {
		if err := ptvq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptvq *PartnerTrainingVideoQuery) loadVideo(ctx context.Context, query *TrainingVideoQuery, nodes []*PartnerTrainingVideo, init func(*PartnerTrainingVideo), assign func(*PartnerTrainingVideo, *TrainingVideo)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PartnerTrainingVideo)
	for i := range nodes {
		if nodes[i].video_id == nil {
			continue
		}
		fk := *nodes[i].video_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(trainingvideo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "video_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ptvq *PartnerTrainingVideoQuery) loadPartner(ctx context.Context, query *PartnerQuery, nodes []*PartnerTrainingVideo, init func(*PartnerTrainingVideo), assign func(*PartnerTrainingVideo, *Partner)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PartnerTrainingVideo)
	for i := range nodes {
		if nodes[i].partner_id == nil {
			continue
		}
		fk := *nodes[i].partner_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(partner.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "partner_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ptvq *PartnerTrainingVideoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptvq.querySpec()
	if len(ptvq.modifiers) > 0 {
		_spec.Modifiers = ptvq.modifiers
	}
	_spec.Node.Columns = ptvq.ctx.Fields
	if len(ptvq.ctx.Fields) > 0 {
		_spec.Unique = ptvq.ctx.Unique != nil && *ptvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptvq.driver, _spec)
}

func (ptvq *PartnerTrainingVideoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(partnertrainingvideo.Table, partnertrainingvideo.Columns, sqlgraph.NewFieldSpec(partnertrainingvideo.FieldID, field.TypeString))
	_spec.From = ptvq.sql
	if unique := ptvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptvq.path != nil {
		_spec.Unique = true
	}
	if fields := ptvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partnertrainingvideo.FieldID)
		for i := range fields {
			if fields[i] != partnertrainingvideo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptvq *PartnerTrainingVideoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptvq.driver.Dialect())
	t1 := builder.Table(partnertrainingvideo.Table)
	columns := ptvq.ctx.Fields
	if len(columns) == 0 {
		columns = partnertrainingvideo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptvq.sql != nil {
		selector = ptvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptvq.ctx.Unique != nil && *ptvq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ptvq.modifiers {
		m(selector)
	}
	for _, p := range ptvq.predicates {
		p(selector)
	}
	for _, p := range ptvq.order {
		p(selector)
	}
	if offset := ptvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptvq *PartnerTrainingVideoQuery) Modify(modifiers ...func(s *sql.Selector)) *PartnerTrainingVideoSelect {
	ptvq.modifiers = append(ptvq.modifiers, modifiers...)
	return ptvq.Select()
}

// PartnerTrainingVideoGroupBy is the group-by builder for PartnerTrainingVideo entities.
type PartnerTrainingVideoGroupBy struct {
	selector
	build *PartnerTrainingVideoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptvgb *PartnerTrainingVideoGroupBy) Aggregate(fns ...AggregateFunc) *PartnerTrainingVideoGroupBy {
	ptvgb.fns = append(ptvgb.fns, fns...)
	return ptvgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptvgb *PartnerTrainingVideoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptvgb.build.ctx, "GroupBy")
	if err := ptvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerTrainingVideoQuery, *PartnerTrainingVideoGroupBy](ctx, ptvgb.build, ptvgb, ptvgb.build.inters, v)
}

func (ptvgb *PartnerTrainingVideoGroupBy) sqlScan(ctx context.Context, root *PartnerTrainingVideoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptvgb.fns))
	for _, fn := range ptvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptvgb.flds)+len(ptvgb.fns))
		for _, f := range *ptvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PartnerTrainingVideoSelect is the builder for selecting fields of PartnerTrainingVideo entities.
type PartnerTrainingVideoSelect struct {
	*PartnerTrainingVideoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ptvs *PartnerTrainingVideoSelect) Aggregate(fns ...AggregateFunc) *PartnerTrainingVideoSelect {
	ptvs.fns = append(ptvs.fns, fns...)
	return ptvs
}

// Scan applies the selector query and scans the result into the given value.
func (ptvs *PartnerTrainingVideoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptvs.ctx, "Select")
	if err := ptvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerTrainingVideoQuery, *PartnerTrainingVideoSelect](ctx, ptvs.PartnerTrainingVideoQuery, ptvs, ptvs.inters, v)
}

func (ptvs *PartnerTrainingVideoSelect) sqlScan(ctx context.Context, root *PartnerTrainingVideoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ptvs.fns))
	for _, fn := range ptvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ptvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptvs *PartnerTrainingVideoSelect) Modify(modifiers ...func(s *sql.Selector)) *PartnerTrainingVideoSelect {
	ptvs.modifiers = append(ptvs.modifiers, modifiers...)
	return ptvs
}
