// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"roofix/ent/estimate"
	"roofix/ent/homeowner"
	"roofix/ent/job"
	"roofix/ent/partner"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HomeOwnerQuery is the builder for querying HomeOwner entities.
type HomeOwnerQuery struct {
	config
	ctx                *QueryContext
	order              []homeowner.OrderOption
	inters             []Interceptor
	predicates         []predicate.HomeOwner
	withEstimates      *EstimateQuery
	withJobs           *JobQuery
	withPartner        *PartnerQuery
	withFKs            bool
	loadTotal          []func(context.Context, []*HomeOwner) error
	modifiers          []func(*sql.Selector)
	withNamedEstimates map[string]*EstimateQuery
	withNamedJobs      map[string]*JobQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HomeOwnerQuery builder.
func (hoq *HomeOwnerQuery) Where(ps ...predicate.HomeOwner) *HomeOwnerQuery {
	hoq.predicates = append(hoq.predicates, ps...)
	return hoq
}

// Limit the number of records to be returned by this query.
func (hoq *HomeOwnerQuery) Limit(limit int) *HomeOwnerQuery {
	hoq.ctx.Limit = &limit
	return hoq
}

// Offset to start from.
func (hoq *HomeOwnerQuery) Offset(offset int) *HomeOwnerQuery {
	hoq.ctx.Offset = &offset
	return hoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hoq *HomeOwnerQuery) Unique(unique bool) *HomeOwnerQuery {
	hoq.ctx.Unique = &unique
	return hoq
}

// Order specifies how the records should be ordered.
func (hoq *HomeOwnerQuery) Order(o ...homeowner.OrderOption) *HomeOwnerQuery {
	hoq.order = append(hoq.order, o...)
	return hoq
}

// QueryEstimates chains the current query on the "estimates" edge.
func (hoq *HomeOwnerQuery) QueryEstimates() *EstimateQuery {
	query := (&EstimateClient{config: hoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(homeowner.Table, homeowner.FieldID, selector),
			sqlgraph.To(estimate.Table, estimate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, homeowner.EstimatesTable, homeowner.EstimatesColumn),
		)
		fromU = sqlgraph.SetNeighbors(hoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryJobs chains the current query on the "jobs" edge.
func (hoq *HomeOwnerQuery) QueryJobs() *JobQuery {
	query := (&JobClient{config: hoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(homeowner.Table, homeowner.FieldID, selector),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, homeowner.JobsTable, homeowner.JobsColumn),
		)
		fromU = sqlgraph.SetNeighbors(hoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPartner chains the current query on the "partner" edge.
func (hoq *HomeOwnerQuery) QueryPartner() *PartnerQuery {
	query := (&PartnerClient{config: hoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(homeowner.Table, homeowner.FieldID, selector),
			sqlgraph.To(partner.Table, partner.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, homeowner.PartnerTable, homeowner.PartnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(hoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HomeOwner entity from the query.
// Returns a *NotFoundError when no HomeOwner was found.
func (hoq *HomeOwnerQuery) First(ctx context.Context) (*HomeOwner, error) {
	nodes, err := hoq.Limit(1).All(setContextOp(ctx, hoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{homeowner.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hoq *HomeOwnerQuery) FirstX(ctx context.Context) *HomeOwner {
	node, err := hoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HomeOwner ID from the query.
// Returns a *NotFoundError when no HomeOwner ID was found.
func (hoq *HomeOwnerQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = hoq.Limit(1).IDs(setContextOp(ctx, hoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{homeowner.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hoq *HomeOwnerQuery) FirstIDX(ctx context.Context) string {
	id, err := hoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HomeOwner entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HomeOwner entity is found.
// Returns a *NotFoundError when no HomeOwner entities are found.
func (hoq *HomeOwnerQuery) Only(ctx context.Context) (*HomeOwner, error) {
	nodes, err := hoq.Limit(2).All(setContextOp(ctx, hoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{homeowner.Label}
	default:
		return nil, &NotSingularError{homeowner.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hoq *HomeOwnerQuery) OnlyX(ctx context.Context) *HomeOwner {
	node, err := hoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HomeOwner ID in the query.
// Returns a *NotSingularError when more than one HomeOwner ID is found.
// Returns a *NotFoundError when no entities are found.
func (hoq *HomeOwnerQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = hoq.Limit(2).IDs(setContextOp(ctx, hoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{homeowner.Label}
	default:
		err = &NotSingularError{homeowner.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hoq *HomeOwnerQuery) OnlyIDX(ctx context.Context) string {
	id, err := hoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HomeOwners.
func (hoq *HomeOwnerQuery) All(ctx context.Context) ([]*HomeOwner, error) {
	ctx = setContextOp(ctx, hoq.ctx, "All")
	if err := hoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HomeOwner, *HomeOwnerQuery]()
	return withInterceptors[[]*HomeOwner](ctx, hoq, qr, hoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hoq *HomeOwnerQuery) AllX(ctx context.Context) []*HomeOwner {
	nodes, err := hoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HomeOwner IDs.
func (hoq *HomeOwnerQuery) IDs(ctx context.Context) (ids []string, err error) {
	if hoq.ctx.Unique == nil && hoq.path != nil {
		hoq.Unique(true)
	}
	ctx = setContextOp(ctx, hoq.ctx, "IDs")
	if err = hoq.Select(homeowner.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hoq *HomeOwnerQuery) IDsX(ctx context.Context) []string {
	ids, err := hoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hoq *HomeOwnerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hoq.ctx, "Count")
	if err := hoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hoq, querierCount[*HomeOwnerQuery](), hoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hoq *HomeOwnerQuery) CountX(ctx context.Context) int {
	count, err := hoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hoq *HomeOwnerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hoq.ctx, "Exist")
	switch _, err := hoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hoq *HomeOwnerQuery) ExistX(ctx context.Context) bool {
	exist, err := hoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HomeOwnerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hoq *HomeOwnerQuery) Clone() *HomeOwnerQuery {
	if hoq == nil {
		return nil
	}
	return &HomeOwnerQuery{
		config:        hoq.config,
		ctx:           hoq.ctx.Clone(),
		order:         append([]homeowner.OrderOption{}, hoq.order...),
		inters:        append([]Interceptor{}, hoq.inters...),
		predicates:    append([]predicate.HomeOwner{}, hoq.predicates...),
		withEstimates: hoq.withEstimates.Clone(),
		withJobs:      hoq.withJobs.Clone(),
		withPartner:   hoq.withPartner.Clone(),
		// clone intermediate query.
		sql:  hoq.sql.Clone(),
		path: hoq.path,
	}
}

// WithEstimates tells the query-builder to eager-load the nodes that are connected to
// the "estimates" edge. The optional arguments are used to configure the query builder of the edge.
func (hoq *HomeOwnerQuery) WithEstimates(opts ...func(*EstimateQuery)) *HomeOwnerQuery {
	query := (&EstimateClient{config: hoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hoq.withEstimates = query
	return hoq
}

// WithJobs tells the query-builder to eager-load the nodes that are connected to
// the "jobs" edge. The optional arguments are used to configure the query builder of the edge.
func (hoq *HomeOwnerQuery) WithJobs(opts ...func(*JobQuery)) *HomeOwnerQuery {
	query := (&JobClient{config: hoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hoq.withJobs = query
	return hoq
}

// WithPartner tells the query-builder to eager-load the nodes that are connected to
// the "partner" edge. The optional arguments are used to configure the query builder of the edge.
func (hoq *HomeOwnerQuery) WithPartner(opts ...func(*PartnerQuery)) *HomeOwnerQuery {
	query := (&PartnerClient{config: hoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hoq.withPartner = query
	return hoq
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
//	client.HomeOwner.Query().
//		GroupBy(homeowner.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hoq *HomeOwnerQuery) GroupBy(field string, fields ...string) *HomeOwnerGroupBy {
	hoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HomeOwnerGroupBy{build: hoq}
	grbuild.flds = &hoq.ctx.Fields
	grbuild.label = homeowner.Label
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
//	client.HomeOwner.Query().
//		Select(homeowner.FieldCreatedAt).
//		Scan(ctx, &v)
func (hoq *HomeOwnerQuery) Select(fields ...string) *HomeOwnerSelect {
	hoq.ctx.Fields = append(hoq.ctx.Fields, fields...)
	sbuild := &HomeOwnerSelect{HomeOwnerQuery: hoq}
	sbuild.label = homeowner.Label
	sbuild.flds, sbuild.scan = &hoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HomeOwnerSelect configured with the given aggregations.
func (hoq *HomeOwnerQuery) Aggregate(fns ...AggregateFunc) *HomeOwnerSelect {
	return hoq.Select().Aggregate(fns...)
}

func (hoq *HomeOwnerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hoq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hoq); err != nil {
				return err
			}
		}
	}
	for _, f := range hoq.ctx.Fields {
		if !homeowner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hoq.path != nil {
		prev, err := hoq.path(ctx)
		if err != nil {
			return err
		}
		hoq.sql = prev
	}
	return nil
}

func (hoq *HomeOwnerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HomeOwner, error) {
	var (
		nodes       = []*HomeOwner{}
		withFKs     = hoq.withFKs
		_spec       = hoq.querySpec()
		loadedTypes = [3]bool{
			hoq.withEstimates != nil,
			hoq.withJobs != nil,
			hoq.withPartner != nil,
		}
	)
	if hoq.withPartner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, homeowner.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HomeOwner).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HomeOwner{config: hoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(hoq.modifiers) > 0 {
		_spec.Modifiers = hoq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hoq.withEstimates; query != nil {
		if err := hoq.loadEstimates(ctx, query, nodes,
			func(n *HomeOwner) { n.Edges.Estimates = []*Estimate{} },
			func(n *HomeOwner, e *Estimate) { n.Edges.Estimates = append(n.Edges.Estimates, e) }); err != nil {
			return nil, err
		}
	}
	if query := hoq.withJobs; query != nil {
		if err := hoq.loadJobs(ctx, query, nodes,
			func(n *HomeOwner) { n.Edges.Jobs = []*Job{} },
			func(n *HomeOwner, e *Job) { n.Edges.Jobs = append(n.Edges.Jobs, e) }); err != nil {
			return nil, err
		}
	}
	if query := hoq.withPartner; query != nil {
		if err := hoq.loadPartner(ctx, query, nodes, nil,
			func(n *HomeOwner, e *Partner) { n.Edges.Partner = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range hoq.withNamedEstimates {
		if err := hoq.loadEstimates(ctx, query, nodes,
			func(n *HomeOwner) { n.appendNamedEstimates(name) },
			func(n *HomeOwner, e *Estimate) { n.appendNamedEstimates(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hoq.withNamedJobs {
		if err := hoq.loadJobs(ctx, query, nodes,
			func(n *HomeOwner) { n.appendNamedJobs(name) },
			func(n *HomeOwner, e *Job) { n.appendNamedJobs(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range hoq.loadTotal {
		if err := hoq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hoq *HomeOwnerQuery) loadEstimates(ctx context.Context, query *EstimateQuery, nodes []*HomeOwner, init func(*HomeOwner), assign func(*HomeOwner, *Estimate)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*HomeOwner)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Estimate(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(homeowner.EstimatesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.home_owner_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "home_owner_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "home_owner_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (hoq *HomeOwnerQuery) loadJobs(ctx context.Context, query *JobQuery, nodes []*HomeOwner, init func(*HomeOwner), assign func(*HomeOwner, *Job)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*HomeOwner)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Job(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(homeowner.JobsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.home_owner_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "home_owner_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "home_owner_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (hoq *HomeOwnerQuery) loadPartner(ctx context.Context, query *PartnerQuery, nodes []*HomeOwner, init func(*HomeOwner), assign func(*HomeOwner, *Partner)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*HomeOwner)
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

func (hoq *HomeOwnerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hoq.querySpec()
	if len(hoq.modifiers) > 0 {
		_spec.Modifiers = hoq.modifiers
	}
	_spec.Node.Columns = hoq.ctx.Fields
	if len(hoq.ctx.Fields) > 0 {
		_spec.Unique = hoq.ctx.Unique != nil && *hoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hoq.driver, _spec)
}

func (hoq *HomeOwnerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(homeowner.Table, homeowner.Columns, sqlgraph.NewFieldSpec(homeowner.FieldID, field.TypeString))
	_spec.From = hoq.sql
	if unique := hoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hoq.path != nil {
		_spec.Unique = true
	}
	if fields := hoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, homeowner.FieldID)
		for i := range fields {
			if fields[i] != homeowner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hoq *HomeOwnerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hoq.driver.Dialect())
	t1 := builder.Table(homeowner.Table)
	columns := hoq.ctx.Fields
	if len(columns) == 0 {
		columns = homeowner.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hoq.sql != nil {
		selector = hoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hoq.ctx.Unique != nil && *hoq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range hoq.modifiers {
		m(selector)
	}
	for _, p := range hoq.predicates {
		p(selector)
	}
	for _, p := range hoq.order {
		p(selector)
	}
	if offset := hoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (hoq *HomeOwnerQuery) Modify(modifiers ...func(s *sql.Selector)) *HomeOwnerSelect {
	hoq.modifiers = append(hoq.modifiers, modifiers...)
	return hoq.Select()
}

// WithNamedEstimates tells the query-builder to eager-load the nodes that are connected to the "estimates"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hoq *HomeOwnerQuery) WithNamedEstimates(name string, opts ...func(*EstimateQuery)) *HomeOwnerQuery {
	query := (&EstimateClient{config: hoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hoq.withNamedEstimates == nil {
		hoq.withNamedEstimates = make(map[string]*EstimateQuery)
	}
	hoq.withNamedEstimates[name] = query
	return hoq
}

// WithNamedJobs tells the query-builder to eager-load the nodes that are connected to the "jobs"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hoq *HomeOwnerQuery) WithNamedJobs(name string, opts ...func(*JobQuery)) *HomeOwnerQuery {
	query := (&JobClient{config: hoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hoq.withNamedJobs == nil {
		hoq.withNamedJobs = make(map[string]*JobQuery)
	}
	hoq.withNamedJobs[name] = query
	return hoq
}

// HomeOwnerGroupBy is the group-by builder for HomeOwner entities.
type HomeOwnerGroupBy struct {
	selector
	build *HomeOwnerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hogb *HomeOwnerGroupBy) Aggregate(fns ...AggregateFunc) *HomeOwnerGroupBy {
	hogb.fns = append(hogb.fns, fns...)
	return hogb
}

// Scan applies the selector query and scans the result into the given value.
func (hogb *HomeOwnerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hogb.build.ctx, "GroupBy")
	if err := hogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HomeOwnerQuery, *HomeOwnerGroupBy](ctx, hogb.build, hogb, hogb.build.inters, v)
}

func (hogb *HomeOwnerGroupBy) sqlScan(ctx context.Context, root *HomeOwnerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hogb.fns))
	for _, fn := range hogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hogb.flds)+len(hogb.fns))
		for _, f := range *hogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HomeOwnerSelect is the builder for selecting fields of HomeOwner entities.
type HomeOwnerSelect struct {
	*HomeOwnerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hos *HomeOwnerSelect) Aggregate(fns ...AggregateFunc) *HomeOwnerSelect {
	hos.fns = append(hos.fns, fns...)
	return hos
}

// Scan applies the selector query and scans the result into the given value.
func (hos *HomeOwnerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hos.ctx, "Select")
	if err := hos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HomeOwnerQuery, *HomeOwnerSelect](ctx, hos.HomeOwnerQuery, hos, hos.inters, v)
}

func (hos *HomeOwnerSelect) sqlScan(ctx context.Context, root *HomeOwnerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hos.fns))
	for _, fn := range hos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (hos *HomeOwnerSelect) Modify(modifiers ...func(s *sql.Selector)) *HomeOwnerSelect {
	hos.modifiers = append(hos.modifiers, modifiers...)
	return hos
}