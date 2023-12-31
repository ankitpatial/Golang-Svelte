// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/apiuser"
	"roofix/ent/partner"
	"roofix/ent/partneractivity"
	"roofix/ent/predicate"
	"roofix/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PartnerActivityQuery is the builder for querying PartnerActivity entities.
type PartnerActivityQuery struct {
	config
	ctx            *QueryContext
	order          []partneractivity.OrderOption
	inters         []Interceptor
	predicates     []predicate.PartnerActivity
	withPartner    *PartnerQuery
	withCreator    *UserQuery
	withCreatorAPI *ApiUserQuery
	withFKs        bool
	loadTotal      []func(context.Context, []*PartnerActivity) error
	modifiers      []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PartnerActivityQuery builder.
func (paq *PartnerActivityQuery) Where(ps ...predicate.PartnerActivity) *PartnerActivityQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit the number of records to be returned by this query.
func (paq *PartnerActivityQuery) Limit(limit int) *PartnerActivityQuery {
	paq.ctx.Limit = &limit
	return paq
}

// Offset to start from.
func (paq *PartnerActivityQuery) Offset(offset int) *PartnerActivityQuery {
	paq.ctx.Offset = &offset
	return paq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (paq *PartnerActivityQuery) Unique(unique bool) *PartnerActivityQuery {
	paq.ctx.Unique = &unique
	return paq
}

// Order specifies how the records should be ordered.
func (paq *PartnerActivityQuery) Order(o ...partneractivity.OrderOption) *PartnerActivityQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// QueryPartner chains the current query on the "partner" edge.
func (paq *PartnerActivityQuery) QueryPartner() *PartnerQuery {
	query := (&PartnerClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partneractivity.Table, partneractivity.FieldID, selector),
			sqlgraph.To(partner.Table, partner.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, partneractivity.PartnerTable, partneractivity.PartnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreator chains the current query on the "creator" edge.
func (paq *PartnerActivityQuery) QueryCreator() *UserQuery {
	query := (&UserClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partneractivity.Table, partneractivity.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, partneractivity.CreatorTable, partneractivity.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatorAPI chains the current query on the "creator_api" edge.
func (paq *PartnerActivityQuery) QueryCreatorAPI() *ApiUserQuery {
	query := (&ApiUserClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partneractivity.Table, partneractivity.FieldID, selector),
			sqlgraph.To(apiuser.Table, apiuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, partneractivity.CreatorAPITable, partneractivity.CreatorAPIColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PartnerActivity entity from the query.
// Returns a *NotFoundError when no PartnerActivity was found.
func (paq *PartnerActivityQuery) First(ctx context.Context) (*PartnerActivity, error) {
	nodes, err := paq.Limit(1).All(setContextOp(ctx, paq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{partneractivity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *PartnerActivityQuery) FirstX(ctx context.Context) *PartnerActivity {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PartnerActivity ID from the query.
// Returns a *NotFoundError when no PartnerActivity ID was found.
func (paq *PartnerActivityQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = paq.Limit(1).IDs(setContextOp(ctx, paq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{partneractivity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (paq *PartnerActivityQuery) FirstIDX(ctx context.Context) string {
	id, err := paq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PartnerActivity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PartnerActivity entity is found.
// Returns a *NotFoundError when no PartnerActivity entities are found.
func (paq *PartnerActivityQuery) Only(ctx context.Context) (*PartnerActivity, error) {
	nodes, err := paq.Limit(2).All(setContextOp(ctx, paq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{partneractivity.Label}
	default:
		return nil, &NotSingularError{partneractivity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *PartnerActivityQuery) OnlyX(ctx context.Context) *PartnerActivity {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PartnerActivity ID in the query.
// Returns a *NotSingularError when more than one PartnerActivity ID is found.
// Returns a *NotFoundError when no entities are found.
func (paq *PartnerActivityQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = paq.Limit(2).IDs(setContextOp(ctx, paq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{partneractivity.Label}
	default:
		err = &NotSingularError{partneractivity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (paq *PartnerActivityQuery) OnlyIDX(ctx context.Context) string {
	id, err := paq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PartnerActivities.
func (paq *PartnerActivityQuery) All(ctx context.Context) ([]*PartnerActivity, error) {
	ctx = setContextOp(ctx, paq.ctx, "All")
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PartnerActivity, *PartnerActivityQuery]()
	return withInterceptors[[]*PartnerActivity](ctx, paq, qr, paq.inters)
}

// AllX is like All, but panics if an error occurs.
func (paq *PartnerActivityQuery) AllX(ctx context.Context) []*PartnerActivity {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PartnerActivity IDs.
func (paq *PartnerActivityQuery) IDs(ctx context.Context) (ids []string, err error) {
	if paq.ctx.Unique == nil && paq.path != nil {
		paq.Unique(true)
	}
	ctx = setContextOp(ctx, paq.ctx, "IDs")
	if err = paq.Select(partneractivity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (paq *PartnerActivityQuery) IDsX(ctx context.Context) []string {
	ids, err := paq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (paq *PartnerActivityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, paq.ctx, "Count")
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, paq, querierCount[*PartnerActivityQuery](), paq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (paq *PartnerActivityQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *PartnerActivityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, paq.ctx, "Exist")
	switch _, err := paq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *PartnerActivityQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PartnerActivityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *PartnerActivityQuery) Clone() *PartnerActivityQuery {
	if paq == nil {
		return nil
	}
	return &PartnerActivityQuery{
		config:         paq.config,
		ctx:            paq.ctx.Clone(),
		order:          append([]partneractivity.OrderOption{}, paq.order...),
		inters:         append([]Interceptor{}, paq.inters...),
		predicates:     append([]predicate.PartnerActivity{}, paq.predicates...),
		withPartner:    paq.withPartner.Clone(),
		withCreator:    paq.withCreator.Clone(),
		withCreatorAPI: paq.withCreatorAPI.Clone(),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
	}
}

// WithPartner tells the query-builder to eager-load the nodes that are connected to
// the "partner" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *PartnerActivityQuery) WithPartner(opts ...func(*PartnerQuery)) *PartnerActivityQuery {
	query := (&PartnerClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withPartner = query
	return paq
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *PartnerActivityQuery) WithCreator(opts ...func(*UserQuery)) *PartnerActivityQuery {
	query := (&UserClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withCreator = query
	return paq
}

// WithCreatorAPI tells the query-builder to eager-load the nodes that are connected to
// the "creator_api" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *PartnerActivityQuery) WithCreatorAPI(opts ...func(*ApiUserQuery)) *PartnerActivityQuery {
	query := (&ApiUserClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withCreatorAPI = query
	return paq
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
//	client.PartnerActivity.Query().
//		GroupBy(partneractivity.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (paq *PartnerActivityQuery) GroupBy(field string, fields ...string) *PartnerActivityGroupBy {
	paq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PartnerActivityGroupBy{build: paq}
	grbuild.flds = &paq.ctx.Fields
	grbuild.label = partneractivity.Label
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
//	client.PartnerActivity.Query().
//		Select(partneractivity.FieldCreatedAt).
//		Scan(ctx, &v)
func (paq *PartnerActivityQuery) Select(fields ...string) *PartnerActivitySelect {
	paq.ctx.Fields = append(paq.ctx.Fields, fields...)
	sbuild := &PartnerActivitySelect{PartnerActivityQuery: paq}
	sbuild.label = partneractivity.Label
	sbuild.flds, sbuild.scan = &paq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PartnerActivitySelect configured with the given aggregations.
func (paq *PartnerActivityQuery) Aggregate(fns ...AggregateFunc) *PartnerActivitySelect {
	return paq.Select().Aggregate(fns...)
}

func (paq *PartnerActivityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range paq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, paq); err != nil {
				return err
			}
		}
	}
	for _, f := range paq.ctx.Fields {
		if !partneractivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	return nil
}

func (paq *PartnerActivityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PartnerActivity, error) {
	var (
		nodes       = []*PartnerActivity{}
		withFKs     = paq.withFKs
		_spec       = paq.querySpec()
		loadedTypes = [3]bool{
			paq.withPartner != nil,
			paq.withCreator != nil,
			paq.withCreatorAPI != nil,
		}
	)
	if paq.withPartner != nil || paq.withCreator != nil || paq.withCreatorAPI != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, partneractivity.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PartnerActivity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PartnerActivity{config: paq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(paq.modifiers) > 0 {
		_spec.Modifiers = paq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := paq.withPartner; query != nil {
		if err := paq.loadPartner(ctx, query, nodes, nil,
			func(n *PartnerActivity, e *Partner) { n.Edges.Partner = e }); err != nil {
			return nil, err
		}
	}
	if query := paq.withCreator; query != nil {
		if err := paq.loadCreator(ctx, query, nodes, nil,
			func(n *PartnerActivity, e *User) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := paq.withCreatorAPI; query != nil {
		if err := paq.loadCreatorAPI(ctx, query, nodes, nil,
			func(n *PartnerActivity, e *ApiUser) { n.Edges.CreatorAPI = e }); err != nil {
			return nil, err
		}
	}
	for i := range paq.loadTotal {
		if err := paq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (paq *PartnerActivityQuery) loadPartner(ctx context.Context, query *PartnerQuery, nodes []*PartnerActivity, init func(*PartnerActivity), assign func(*PartnerActivity, *Partner)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PartnerActivity)
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
func (paq *PartnerActivityQuery) loadCreator(ctx context.Context, query *UserQuery, nodes []*PartnerActivity, init func(*PartnerActivity), assign func(*PartnerActivity, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PartnerActivity)
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
func (paq *PartnerActivityQuery) loadCreatorAPI(ctx context.Context, query *ApiUserQuery, nodes []*PartnerActivity, init func(*PartnerActivity), assign func(*PartnerActivity, *ApiUser)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PartnerActivity)
	for i := range nodes {
		if nodes[i].api_user_id == nil {
			continue
		}
		fk := *nodes[i].api_user_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(apiuser.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "api_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (paq *PartnerActivityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	if len(paq.modifiers) > 0 {
		_spec.Modifiers = paq.modifiers
	}
	_spec.Node.Columns = paq.ctx.Fields
	if len(paq.ctx.Fields) > 0 {
		_spec.Unique = paq.ctx.Unique != nil && *paq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *PartnerActivityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(partneractivity.Table, partneractivity.Columns, sqlgraph.NewFieldSpec(partneractivity.FieldID, field.TypeString))
	_spec.From = paq.sql
	if unique := paq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if paq.path != nil {
		_spec.Unique = true
	}
	if fields := paq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partneractivity.FieldID)
		for i := range fields {
			if fields[i] != partneractivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (paq *PartnerActivityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(partneractivity.Table)
	columns := paq.ctx.Fields
	if len(columns) == 0 {
		columns = partneractivity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if paq.ctx.Unique != nil && *paq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range paq.modifiers {
		m(selector)
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector)
	}
	if offset := paq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (paq *PartnerActivityQuery) Modify(modifiers ...func(s *sql.Selector)) *PartnerActivitySelect {
	paq.modifiers = append(paq.modifiers, modifiers...)
	return paq.Select()
}

// PartnerActivityGroupBy is the group-by builder for PartnerActivity entities.
type PartnerActivityGroupBy struct {
	selector
	build *PartnerActivityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *PartnerActivityGroupBy) Aggregate(fns ...AggregateFunc) *PartnerActivityGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the selector query and scans the result into the given value.
func (pagb *PartnerActivityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pagb.build.ctx, "GroupBy")
	if err := pagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerActivityQuery, *PartnerActivityGroupBy](ctx, pagb.build, pagb, pagb.build.inters, v)
}

func (pagb *PartnerActivityGroupBy) sqlScan(ctx context.Context, root *PartnerActivityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pagb.fns))
	for _, fn := range pagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pagb.flds)+len(pagb.fns))
		for _, f := range *pagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PartnerActivitySelect is the builder for selecting fields of PartnerActivity entities.
type PartnerActivitySelect struct {
	*PartnerActivityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pas *PartnerActivitySelect) Aggregate(fns ...AggregateFunc) *PartnerActivitySelect {
	pas.fns = append(pas.fns, fns...)
	return pas
}

// Scan applies the selector query and scans the result into the given value.
func (pas *PartnerActivitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pas.ctx, "Select")
	if err := pas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerActivityQuery, *PartnerActivitySelect](ctx, pas.PartnerActivityQuery, pas, pas.inters, v)
}

func (pas *PartnerActivitySelect) sqlScan(ctx context.Context, root *PartnerActivityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pas.fns))
	for _, fn := range pas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pas *PartnerActivitySelect) Modify(modifiers ...func(s *sql.Selector)) *PartnerActivitySelect {
	pas.modifiers = append(pas.modifiers, modifiers...)
	return pas
}
