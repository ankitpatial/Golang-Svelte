// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/contactus"
	"roofix/ent/partner"
	"roofix/ent/predicate"
	"roofix/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContactUsQuery is the builder for querying ContactUs entities.
type ContactUsQuery struct {
	config
	ctx         *QueryContext
	order       []contactus.OrderOption
	inters      []Interceptor
	predicates  []predicate.ContactUs
	withPartner *PartnerQuery
	withCreator *UserQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*ContactUs) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContactUsQuery builder.
func (cuq *ContactUsQuery) Where(ps ...predicate.ContactUs) *ContactUsQuery {
	cuq.predicates = append(cuq.predicates, ps...)
	return cuq
}

// Limit the number of records to be returned by this query.
func (cuq *ContactUsQuery) Limit(limit int) *ContactUsQuery {
	cuq.ctx.Limit = &limit
	return cuq
}

// Offset to start from.
func (cuq *ContactUsQuery) Offset(offset int) *ContactUsQuery {
	cuq.ctx.Offset = &offset
	return cuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cuq *ContactUsQuery) Unique(unique bool) *ContactUsQuery {
	cuq.ctx.Unique = &unique
	return cuq
}

// Order specifies how the records should be ordered.
func (cuq *ContactUsQuery) Order(o ...contactus.OrderOption) *ContactUsQuery {
	cuq.order = append(cuq.order, o...)
	return cuq
}

// QueryPartner chains the current query on the "partner" edge.
func (cuq *ContactUsQuery) QueryPartner() *PartnerQuery {
	query := (&PartnerClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contactus.Table, contactus.FieldID, selector),
			sqlgraph.To(partner.Table, partner.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, contactus.PartnerTable, contactus.PartnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreator chains the current query on the "creator" edge.
func (cuq *ContactUsQuery) QueryCreator() *UserQuery {
	query := (&UserClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contactus.Table, contactus.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, contactus.CreatorTable, contactus.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ContactUs entity from the query.
// Returns a *NotFoundError when no ContactUs was found.
func (cuq *ContactUsQuery) First(ctx context.Context) (*ContactUs, error) {
	nodes, err := cuq.Limit(1).All(setContextOp(ctx, cuq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contactus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cuq *ContactUsQuery) FirstX(ctx context.Context) *ContactUs {
	node, err := cuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ContactUs ID from the query.
// Returns a *NotFoundError when no ContactUs ID was found.
func (cuq *ContactUsQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cuq.Limit(1).IDs(setContextOp(ctx, cuq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contactus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cuq *ContactUsQuery) FirstIDX(ctx context.Context) string {
	id, err := cuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ContactUs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ContactUs entity is found.
// Returns a *NotFoundError when no ContactUs entities are found.
func (cuq *ContactUsQuery) Only(ctx context.Context) (*ContactUs, error) {
	nodes, err := cuq.Limit(2).All(setContextOp(ctx, cuq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contactus.Label}
	default:
		return nil, &NotSingularError{contactus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cuq *ContactUsQuery) OnlyX(ctx context.Context) *ContactUs {
	node, err := cuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ContactUs ID in the query.
// Returns a *NotSingularError when more than one ContactUs ID is found.
// Returns a *NotFoundError when no entities are found.
func (cuq *ContactUsQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cuq.Limit(2).IDs(setContextOp(ctx, cuq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contactus.Label}
	default:
		err = &NotSingularError{contactus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cuq *ContactUsQuery) OnlyIDX(ctx context.Context) string {
	id, err := cuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ContactUsSlice.
func (cuq *ContactUsQuery) All(ctx context.Context) ([]*ContactUs, error) {
	ctx = setContextOp(ctx, cuq.ctx, "All")
	if err := cuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ContactUs, *ContactUsQuery]()
	return withInterceptors[[]*ContactUs](ctx, cuq, qr, cuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cuq *ContactUsQuery) AllX(ctx context.Context) []*ContactUs {
	nodes, err := cuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ContactUs IDs.
func (cuq *ContactUsQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cuq.ctx.Unique == nil && cuq.path != nil {
		cuq.Unique(true)
	}
	ctx = setContextOp(ctx, cuq.ctx, "IDs")
	if err = cuq.Select(contactus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cuq *ContactUsQuery) IDsX(ctx context.Context) []string {
	ids, err := cuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cuq *ContactUsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cuq.ctx, "Count")
	if err := cuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cuq, querierCount[*ContactUsQuery](), cuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cuq *ContactUsQuery) CountX(ctx context.Context) int {
	count, err := cuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cuq *ContactUsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cuq.ctx, "Exist")
	switch _, err := cuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cuq *ContactUsQuery) ExistX(ctx context.Context) bool {
	exist, err := cuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContactUsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cuq *ContactUsQuery) Clone() *ContactUsQuery {
	if cuq == nil {
		return nil
	}
	return &ContactUsQuery{
		config:      cuq.config,
		ctx:         cuq.ctx.Clone(),
		order:       append([]contactus.OrderOption{}, cuq.order...),
		inters:      append([]Interceptor{}, cuq.inters...),
		predicates:  append([]predicate.ContactUs{}, cuq.predicates...),
		withPartner: cuq.withPartner.Clone(),
		withCreator: cuq.withCreator.Clone(),
		// clone intermediate query.
		sql:  cuq.sql.Clone(),
		path: cuq.path,
	}
}

// WithPartner tells the query-builder to eager-load the nodes that are connected to
// the "partner" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *ContactUsQuery) WithPartner(opts ...func(*PartnerQuery)) *ContactUsQuery {
	query := (&PartnerClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withPartner = query
	return cuq
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *ContactUsQuery) WithCreator(opts ...func(*UserQuery)) *ContactUsQuery {
	query := (&UserClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withCreator = query
	return cuq
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
//	client.ContactUs.Query().
//		GroupBy(contactus.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cuq *ContactUsQuery) GroupBy(field string, fields ...string) *ContactUsGroupBy {
	cuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ContactUsGroupBy{build: cuq}
	grbuild.flds = &cuq.ctx.Fields
	grbuild.label = contactus.Label
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
//	client.ContactUs.Query().
//		Select(contactus.FieldCreatedAt).
//		Scan(ctx, &v)
func (cuq *ContactUsQuery) Select(fields ...string) *ContactUsSelect {
	cuq.ctx.Fields = append(cuq.ctx.Fields, fields...)
	sbuild := &ContactUsSelect{ContactUsQuery: cuq}
	sbuild.label = contactus.Label
	sbuild.flds, sbuild.scan = &cuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContactUsSelect configured with the given aggregations.
func (cuq *ContactUsQuery) Aggregate(fns ...AggregateFunc) *ContactUsSelect {
	return cuq.Select().Aggregate(fns...)
}

func (cuq *ContactUsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cuq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cuq); err != nil {
				return err
			}
		}
	}
	for _, f := range cuq.ctx.Fields {
		if !contactus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cuq.path != nil {
		prev, err := cuq.path(ctx)
		if err != nil {
			return err
		}
		cuq.sql = prev
	}
	return nil
}

func (cuq *ContactUsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ContactUs, error) {
	var (
		nodes       = []*ContactUs{}
		withFKs     = cuq.withFKs
		_spec       = cuq.querySpec()
		loadedTypes = [2]bool{
			cuq.withPartner != nil,
			cuq.withCreator != nil,
		}
	)
	if cuq.withPartner != nil || cuq.withCreator != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, contactus.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ContactUs).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ContactUs{config: cuq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cuq.modifiers) > 0 {
		_spec.Modifiers = cuq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cuq.withPartner; query != nil {
		if err := cuq.loadPartner(ctx, query, nodes, nil,
			func(n *ContactUs, e *Partner) { n.Edges.Partner = e }); err != nil {
			return nil, err
		}
	}
	if query := cuq.withCreator; query != nil {
		if err := cuq.loadCreator(ctx, query, nodes, nil,
			func(n *ContactUs, e *User) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	for i := range cuq.loadTotal {
		if err := cuq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cuq *ContactUsQuery) loadPartner(ctx context.Context, query *PartnerQuery, nodes []*ContactUs, init func(*ContactUs), assign func(*ContactUs, *Partner)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ContactUs)
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
func (cuq *ContactUsQuery) loadCreator(ctx context.Context, query *UserQuery, nodes []*ContactUs, init func(*ContactUs), assign func(*ContactUs, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ContactUs)
	for i := range nodes {
		if nodes[i].creator_id == nil {
			continue
		}
		fk := *nodes[i].creator_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "creator_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cuq *ContactUsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cuq.querySpec()
	if len(cuq.modifiers) > 0 {
		_spec.Modifiers = cuq.modifiers
	}
	_spec.Node.Columns = cuq.ctx.Fields
	if len(cuq.ctx.Fields) > 0 {
		_spec.Unique = cuq.ctx.Unique != nil && *cuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cuq.driver, _spec)
}

func (cuq *ContactUsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(contactus.Table, contactus.Columns, sqlgraph.NewFieldSpec(contactus.FieldID, field.TypeString))
	_spec.From = cuq.sql
	if unique := cuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cuq.path != nil {
		_spec.Unique = true
	}
	if fields := cuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contactus.FieldID)
		for i := range fields {
			if fields[i] != contactus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cuq *ContactUsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cuq.driver.Dialect())
	t1 := builder.Table(contactus.Table)
	columns := cuq.ctx.Fields
	if len(columns) == 0 {
		columns = contactus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cuq.sql != nil {
		selector = cuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cuq.ctx.Unique != nil && *cuq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cuq.modifiers {
		m(selector)
	}
	for _, p := range cuq.predicates {
		p(selector)
	}
	for _, p := range cuq.order {
		p(selector)
	}
	if offset := cuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cuq *ContactUsQuery) Modify(modifiers ...func(s *sql.Selector)) *ContactUsSelect {
	cuq.modifiers = append(cuq.modifiers, modifiers...)
	return cuq.Select()
}

// ContactUsGroupBy is the group-by builder for ContactUs entities.
type ContactUsGroupBy struct {
	selector
	build *ContactUsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cugb *ContactUsGroupBy) Aggregate(fns ...AggregateFunc) *ContactUsGroupBy {
	cugb.fns = append(cugb.fns, fns...)
	return cugb
}

// Scan applies the selector query and scans the result into the given value.
func (cugb *ContactUsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cugb.build.ctx, "GroupBy")
	if err := cugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContactUsQuery, *ContactUsGroupBy](ctx, cugb.build, cugb, cugb.build.inters, v)
}

func (cugb *ContactUsGroupBy) sqlScan(ctx context.Context, root *ContactUsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cugb.fns))
	for _, fn := range cugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cugb.flds)+len(cugb.fns))
		for _, f := range *cugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ContactUsSelect is the builder for selecting fields of ContactUs entities.
type ContactUsSelect struct {
	*ContactUsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cus *ContactUsSelect) Aggregate(fns ...AggregateFunc) *ContactUsSelect {
	cus.fns = append(cus.fns, fns...)
	return cus
}

// Scan applies the selector query and scans the result into the given value.
func (cus *ContactUsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cus.ctx, "Select")
	if err := cus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContactUsQuery, *ContactUsSelect](ctx, cus.ContactUsQuery, cus, cus.inters, v)
}

func (cus *ContactUsSelect) sqlScan(ctx context.Context, root *ContactUsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cus.fns))
	for _, fn := range cus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cus *ContactUsSelect) Modify(modifiers ...func(s *sql.Selector)) *ContactUsSelect {
	cus.modifiers = append(cus.modifiers, modifiers...)
	return cus
}
