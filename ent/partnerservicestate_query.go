// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/partner"
	"roofix/ent/partnerservicestate"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PartnerServiceStateQuery is the builder for querying PartnerServiceState entities.
type PartnerServiceStateQuery struct {
	config
	ctx         *QueryContext
	order       []partnerservicestate.OrderOption
	inters      []Interceptor
	predicates  []predicate.PartnerServiceState
	withPartner *PartnerQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*PartnerServiceState) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PartnerServiceStateQuery builder.
func (pssq *PartnerServiceStateQuery) Where(ps ...predicate.PartnerServiceState) *PartnerServiceStateQuery {
	pssq.predicates = append(pssq.predicates, ps...)
	return pssq
}

// Limit the number of records to be returned by this query.
func (pssq *PartnerServiceStateQuery) Limit(limit int) *PartnerServiceStateQuery {
	pssq.ctx.Limit = &limit
	return pssq
}

// Offset to start from.
func (pssq *PartnerServiceStateQuery) Offset(offset int) *PartnerServiceStateQuery {
	pssq.ctx.Offset = &offset
	return pssq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pssq *PartnerServiceStateQuery) Unique(unique bool) *PartnerServiceStateQuery {
	pssq.ctx.Unique = &unique
	return pssq
}

// Order specifies how the records should be ordered.
func (pssq *PartnerServiceStateQuery) Order(o ...partnerservicestate.OrderOption) *PartnerServiceStateQuery {
	pssq.order = append(pssq.order, o...)
	return pssq
}

// QueryPartner chains the current query on the "partner" edge.
func (pssq *PartnerServiceStateQuery) QueryPartner() *PartnerQuery {
	query := (&PartnerClient{config: pssq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pssq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pssq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partnerservicestate.Table, partnerservicestate.FieldID, selector),
			sqlgraph.To(partner.Table, partner.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, partnerservicestate.PartnerTable, partnerservicestate.PartnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(pssq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PartnerServiceState entity from the query.
// Returns a *NotFoundError when no PartnerServiceState was found.
func (pssq *PartnerServiceStateQuery) First(ctx context.Context) (*PartnerServiceState, error) {
	nodes, err := pssq.Limit(1).All(setContextOp(ctx, pssq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{partnerservicestate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) FirstX(ctx context.Context) *PartnerServiceState {
	node, err := pssq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PartnerServiceState ID from the query.
// Returns a *NotFoundError when no PartnerServiceState ID was found.
func (pssq *PartnerServiceStateQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pssq.Limit(1).IDs(setContextOp(ctx, pssq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{partnerservicestate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) FirstIDX(ctx context.Context) string {
	id, err := pssq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PartnerServiceState entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PartnerServiceState entity is found.
// Returns a *NotFoundError when no PartnerServiceState entities are found.
func (pssq *PartnerServiceStateQuery) Only(ctx context.Context) (*PartnerServiceState, error) {
	nodes, err := pssq.Limit(2).All(setContextOp(ctx, pssq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{partnerservicestate.Label}
	default:
		return nil, &NotSingularError{partnerservicestate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) OnlyX(ctx context.Context) *PartnerServiceState {
	node, err := pssq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PartnerServiceState ID in the query.
// Returns a *NotSingularError when more than one PartnerServiceState ID is found.
// Returns a *NotFoundError when no entities are found.
func (pssq *PartnerServiceStateQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pssq.Limit(2).IDs(setContextOp(ctx, pssq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{partnerservicestate.Label}
	default:
		err = &NotSingularError{partnerservicestate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) OnlyIDX(ctx context.Context) string {
	id, err := pssq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PartnerServiceStates.
func (pssq *PartnerServiceStateQuery) All(ctx context.Context) ([]*PartnerServiceState, error) {
	ctx = setContextOp(ctx, pssq.ctx, "All")
	if err := pssq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PartnerServiceState, *PartnerServiceStateQuery]()
	return withInterceptors[[]*PartnerServiceState](ctx, pssq, qr, pssq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) AllX(ctx context.Context) []*PartnerServiceState {
	nodes, err := pssq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PartnerServiceState IDs.
func (pssq *PartnerServiceStateQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pssq.ctx.Unique == nil && pssq.path != nil {
		pssq.Unique(true)
	}
	ctx = setContextOp(ctx, pssq.ctx, "IDs")
	if err = pssq.Select(partnerservicestate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) IDsX(ctx context.Context) []string {
	ids, err := pssq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pssq *PartnerServiceStateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pssq.ctx, "Count")
	if err := pssq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pssq, querierCount[*PartnerServiceStateQuery](), pssq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) CountX(ctx context.Context) int {
	count, err := pssq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pssq *PartnerServiceStateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pssq.ctx, "Exist")
	switch _, err := pssq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pssq *PartnerServiceStateQuery) ExistX(ctx context.Context) bool {
	exist, err := pssq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PartnerServiceStateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pssq *PartnerServiceStateQuery) Clone() *PartnerServiceStateQuery {
	if pssq == nil {
		return nil
	}
	return &PartnerServiceStateQuery{
		config:      pssq.config,
		ctx:         pssq.ctx.Clone(),
		order:       append([]partnerservicestate.OrderOption{}, pssq.order...),
		inters:      append([]Interceptor{}, pssq.inters...),
		predicates:  append([]predicate.PartnerServiceState{}, pssq.predicates...),
		withPartner: pssq.withPartner.Clone(),
		// clone intermediate query.
		sql:  pssq.sql.Clone(),
		path: pssq.path,
	}
}

// WithPartner tells the query-builder to eager-load the nodes that are connected to
// the "partner" edge. The optional arguments are used to configure the query builder of the edge.
func (pssq *PartnerServiceStateQuery) WithPartner(opts ...func(*PartnerQuery)) *PartnerServiceStateQuery {
	query := (&PartnerClient{config: pssq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pssq.withPartner = query
	return pssq
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
//	client.PartnerServiceState.Query().
//		GroupBy(partnerservicestate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pssq *PartnerServiceStateQuery) GroupBy(field string, fields ...string) *PartnerServiceStateGroupBy {
	pssq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PartnerServiceStateGroupBy{build: pssq}
	grbuild.flds = &pssq.ctx.Fields
	grbuild.label = partnerservicestate.Label
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
//	client.PartnerServiceState.Query().
//		Select(partnerservicestate.FieldCreatedAt).
//		Scan(ctx, &v)
func (pssq *PartnerServiceStateQuery) Select(fields ...string) *PartnerServiceStateSelect {
	pssq.ctx.Fields = append(pssq.ctx.Fields, fields...)
	sbuild := &PartnerServiceStateSelect{PartnerServiceStateQuery: pssq}
	sbuild.label = partnerservicestate.Label
	sbuild.flds, sbuild.scan = &pssq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PartnerServiceStateSelect configured with the given aggregations.
func (pssq *PartnerServiceStateQuery) Aggregate(fns ...AggregateFunc) *PartnerServiceStateSelect {
	return pssq.Select().Aggregate(fns...)
}

func (pssq *PartnerServiceStateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pssq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pssq); err != nil {
				return err
			}
		}
	}
	for _, f := range pssq.ctx.Fields {
		if !partnerservicestate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pssq.path != nil {
		prev, err := pssq.path(ctx)
		if err != nil {
			return err
		}
		pssq.sql = prev
	}
	return nil
}

func (pssq *PartnerServiceStateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PartnerServiceState, error) {
	var (
		nodes       = []*PartnerServiceState{}
		withFKs     = pssq.withFKs
		_spec       = pssq.querySpec()
		loadedTypes = [1]bool{
			pssq.withPartner != nil,
		}
	)
	if pssq.withPartner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, partnerservicestate.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PartnerServiceState).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PartnerServiceState{config: pssq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pssq.modifiers) > 0 {
		_spec.Modifiers = pssq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pssq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pssq.withPartner; query != nil {
		if err := pssq.loadPartner(ctx, query, nodes, nil,
			func(n *PartnerServiceState, e *Partner) { n.Edges.Partner = e }); err != nil {
			return nil, err
		}
	}
	for i := range pssq.loadTotal {
		if err := pssq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pssq *PartnerServiceStateQuery) loadPartner(ctx context.Context, query *PartnerQuery, nodes []*PartnerServiceState, init func(*PartnerServiceState), assign func(*PartnerServiceState, *Partner)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PartnerServiceState)
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

func (pssq *PartnerServiceStateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pssq.querySpec()
	if len(pssq.modifiers) > 0 {
		_spec.Modifiers = pssq.modifiers
	}
	_spec.Node.Columns = pssq.ctx.Fields
	if len(pssq.ctx.Fields) > 0 {
		_spec.Unique = pssq.ctx.Unique != nil && *pssq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pssq.driver, _spec)
}

func (pssq *PartnerServiceStateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(partnerservicestate.Table, partnerservicestate.Columns, sqlgraph.NewFieldSpec(partnerservicestate.FieldID, field.TypeString))
	_spec.From = pssq.sql
	if unique := pssq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pssq.path != nil {
		_spec.Unique = true
	}
	if fields := pssq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partnerservicestate.FieldID)
		for i := range fields {
			if fields[i] != partnerservicestate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pssq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pssq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pssq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pssq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pssq *PartnerServiceStateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pssq.driver.Dialect())
	t1 := builder.Table(partnerservicestate.Table)
	columns := pssq.ctx.Fields
	if len(columns) == 0 {
		columns = partnerservicestate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pssq.sql != nil {
		selector = pssq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pssq.ctx.Unique != nil && *pssq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pssq.modifiers {
		m(selector)
	}
	for _, p := range pssq.predicates {
		p(selector)
	}
	for _, p := range pssq.order {
		p(selector)
	}
	if offset := pssq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pssq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pssq *PartnerServiceStateQuery) Modify(modifiers ...func(s *sql.Selector)) *PartnerServiceStateSelect {
	pssq.modifiers = append(pssq.modifiers, modifiers...)
	return pssq.Select()
}

// PartnerServiceStateGroupBy is the group-by builder for PartnerServiceState entities.
type PartnerServiceStateGroupBy struct {
	selector
	build *PartnerServiceStateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pssgb *PartnerServiceStateGroupBy) Aggregate(fns ...AggregateFunc) *PartnerServiceStateGroupBy {
	pssgb.fns = append(pssgb.fns, fns...)
	return pssgb
}

// Scan applies the selector query and scans the result into the given value.
func (pssgb *PartnerServiceStateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pssgb.build.ctx, "GroupBy")
	if err := pssgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerServiceStateQuery, *PartnerServiceStateGroupBy](ctx, pssgb.build, pssgb, pssgb.build.inters, v)
}

func (pssgb *PartnerServiceStateGroupBy) sqlScan(ctx context.Context, root *PartnerServiceStateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pssgb.fns))
	for _, fn := range pssgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pssgb.flds)+len(pssgb.fns))
		for _, f := range *pssgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pssgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pssgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PartnerServiceStateSelect is the builder for selecting fields of PartnerServiceState entities.
type PartnerServiceStateSelect struct {
	*PartnerServiceStateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (psss *PartnerServiceStateSelect) Aggregate(fns ...AggregateFunc) *PartnerServiceStateSelect {
	psss.fns = append(psss.fns, fns...)
	return psss
}

// Scan applies the selector query and scans the result into the given value.
func (psss *PartnerServiceStateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psss.ctx, "Select")
	if err := psss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerServiceStateQuery, *PartnerServiceStateSelect](ctx, psss.PartnerServiceStateQuery, psss, psss.inters, v)
}

func (psss *PartnerServiceStateSelect) sqlScan(ctx context.Context, root *PartnerServiceStateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(psss.fns))
	for _, fn := range psss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*psss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (psss *PartnerServiceStateSelect) Modify(modifiers ...func(s *sql.Selector)) *PartnerServiceStateSelect {
	psss.modifiers = append(psss.modifiers, modifiers...)
	return psss
}
