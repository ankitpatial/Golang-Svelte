// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"roofix/ent/postalcode"
	"roofix/ent/predicate"
	"roofix/ent/pricing"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostalCodeQuery is the builder for querying PostalCode entities.
type PostalCodeQuery struct {
	config
	ctx              *QueryContext
	order            []postalcode.OrderOption
	inters           []Interceptor
	predicates       []predicate.PostalCode
	withPricing      *PricingQuery
	loadTotal        []func(context.Context, []*PostalCode) error
	modifiers        []func(*sql.Selector)
	withNamedPricing map[string]*PricingQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostalCodeQuery builder.
func (pcq *PostalCodeQuery) Where(ps ...predicate.PostalCode) *PostalCodeQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *PostalCodeQuery) Limit(limit int) *PostalCodeQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *PostalCodeQuery) Offset(offset int) *PostalCodeQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PostalCodeQuery) Unique(unique bool) *PostalCodeQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *PostalCodeQuery) Order(o ...postalcode.OrderOption) *PostalCodeQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QueryPricing chains the current query on the "pricing" edge.
func (pcq *PostalCodeQuery) QueryPricing() *PricingQuery {
	query := (&PricingClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(postalcode.Table, postalcode.FieldID, selector),
			sqlgraph.To(pricing.Table, pricing.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, postalcode.PricingTable, postalcode.PricingColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PostalCode entity from the query.
// Returns a *NotFoundError when no PostalCode was found.
func (pcq *PostalCodeQuery) First(ctx context.Context) (*PostalCode, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{postalcode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PostalCodeQuery) FirstX(ctx context.Context) *PostalCode {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PostalCode ID from the query.
// Returns a *NotFoundError when no PostalCode ID was found.
func (pcq *PostalCodeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{postalcode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PostalCodeQuery) FirstIDX(ctx context.Context) string {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PostalCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PostalCode entity is found.
// Returns a *NotFoundError when no PostalCode entities are found.
func (pcq *PostalCodeQuery) Only(ctx context.Context) (*PostalCode, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{postalcode.Label}
	default:
		return nil, &NotSingularError{postalcode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PostalCodeQuery) OnlyX(ctx context.Context) *PostalCode {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PostalCode ID in the query.
// Returns a *NotSingularError when more than one PostalCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PostalCodeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{postalcode.Label}
	default:
		err = &NotSingularError{postalcode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PostalCodeQuery) OnlyIDX(ctx context.Context) string {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PostalCodes.
func (pcq *PostalCodeQuery) All(ctx context.Context) ([]*PostalCode, error) {
	ctx = setContextOp(ctx, pcq.ctx, "All")
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PostalCode, *PostalCodeQuery]()
	return withInterceptors[[]*PostalCode](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PostalCodeQuery) AllX(ctx context.Context) []*PostalCode {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PostalCode IDs.
func (pcq *PostalCodeQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, "IDs")
	if err = pcq.Select(postalcode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PostalCodeQuery) IDsX(ctx context.Context) []string {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PostalCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Count")
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*PostalCodeQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PostalCodeQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PostalCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Exist")
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PostalCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostalCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PostalCodeQuery) Clone() *PostalCodeQuery {
	if pcq == nil {
		return nil
	}
	return &PostalCodeQuery{
		config:      pcq.config,
		ctx:         pcq.ctx.Clone(),
		order:       append([]postalcode.OrderOption{}, pcq.order...),
		inters:      append([]Interceptor{}, pcq.inters...),
		predicates:  append([]predicate.PostalCode{}, pcq.predicates...),
		withPricing: pcq.withPricing.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithPricing tells the query-builder to eager-load the nodes that are connected to
// the "pricing" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *PostalCodeQuery) WithPricing(opts ...func(*PricingQuery)) *PostalCodeQuery {
	query := (&PricingClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withPricing = query
	return pcq
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
//	client.PostalCode.Query().
//		GroupBy(postalcode.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *PostalCodeQuery) GroupBy(field string, fields ...string) *PostalCodeGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PostalCodeGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = postalcode.Label
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
//	client.PostalCode.Query().
//		Select(postalcode.FieldCreatedAt).
//		Scan(ctx, &v)
func (pcq *PostalCodeQuery) Select(fields ...string) *PostalCodeSelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &PostalCodeSelect{PostalCodeQuery: pcq}
	sbuild.label = postalcode.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PostalCodeSelect configured with the given aggregations.
func (pcq *PostalCodeQuery) Aggregate(fns ...AggregateFunc) *PostalCodeSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *PostalCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !postalcode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *PostalCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PostalCode, error) {
	var (
		nodes       = []*PostalCode{}
		_spec       = pcq.querySpec()
		loadedTypes = [1]bool{
			pcq.withPricing != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PostalCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PostalCode{config: pcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pcq.modifiers) > 0 {
		_spec.Modifiers = pcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pcq.withPricing; query != nil {
		if err := pcq.loadPricing(ctx, query, nodes,
			func(n *PostalCode) { n.Edges.Pricing = []*Pricing{} },
			func(n *PostalCode, e *Pricing) { n.Edges.Pricing = append(n.Edges.Pricing, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pcq.withNamedPricing {
		if err := pcq.loadPricing(ctx, query, nodes,
			func(n *PostalCode) { n.appendNamedPricing(name) },
			func(n *PostalCode, e *Pricing) { n.appendNamedPricing(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pcq.loadTotal {
		if err := pcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pcq *PostalCodeQuery) loadPricing(ctx context.Context, query *PricingQuery, nodes []*PostalCode, init func(*PostalCode), assign func(*PostalCode, *Pricing)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*PostalCode)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Pricing(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(postalcode.PricingColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.postal_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "postal_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "postal_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pcq *PostalCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	if len(pcq.modifiers) > 0 {
		_spec.Modifiers = pcq.modifiers
	}
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PostalCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(postalcode.Table, postalcode.Columns, sqlgraph.NewFieldSpec(postalcode.FieldID, field.TypeString))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postalcode.FieldID)
		for i := range fields {
			if fields[i] != postalcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PostalCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(postalcode.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = postalcode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pcq.modifiers {
		m(selector)
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pcq *PostalCodeQuery) Modify(modifiers ...func(s *sql.Selector)) *PostalCodeSelect {
	pcq.modifiers = append(pcq.modifiers, modifiers...)
	return pcq.Select()
}

// WithNamedPricing tells the query-builder to eager-load the nodes that are connected to the "pricing"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pcq *PostalCodeQuery) WithNamedPricing(name string, opts ...func(*PricingQuery)) *PostalCodeQuery {
	query := (&PricingClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pcq.withNamedPricing == nil {
		pcq.withNamedPricing = make(map[string]*PricingQuery)
	}
	pcq.withNamedPricing[name] = query
	return pcq
}

// PostalCodeGroupBy is the group-by builder for PostalCode entities.
type PostalCodeGroupBy struct {
	selector
	build *PostalCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PostalCodeGroupBy) Aggregate(fns ...AggregateFunc) *PostalCodeGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *PostalCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, "GroupBy")
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostalCodeQuery, *PostalCodeGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *PostalCodeGroupBy) sqlScan(ctx context.Context, root *PostalCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PostalCodeSelect is the builder for selecting fields of PostalCode entities.
type PostalCodeSelect struct {
	*PostalCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *PostalCodeSelect) Aggregate(fns ...AggregateFunc) *PostalCodeSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PostalCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, "Select")
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostalCodeQuery, *PostalCodeSelect](ctx, pcs.PostalCodeQuery, pcs, pcs.inters, v)
}

func (pcs *PostalCodeSelect) sqlScan(ctx context.Context, root *PostalCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pcs *PostalCodeSelect) Modify(modifiers ...func(s *sql.Selector)) *PostalCodeSelect {
	pcs.modifiers = append(pcs.modifiers, modifiers...)
	return pcs
}
