// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/apiuser"
	"roofix/ent/apiusertoken"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiUserTokenQuery is the builder for querying ApiUserToken entities.
type ApiUserTokenQuery struct {
	config
	ctx         *QueryContext
	order       []apiusertoken.OrderOption
	inters      []Interceptor
	predicates  []predicate.ApiUserToken
	withAPIUser *ApiUserQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*ApiUserToken) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApiUserTokenQuery builder.
func (autq *ApiUserTokenQuery) Where(ps ...predicate.ApiUserToken) *ApiUserTokenQuery {
	autq.predicates = append(autq.predicates, ps...)
	return autq
}

// Limit the number of records to be returned by this query.
func (autq *ApiUserTokenQuery) Limit(limit int) *ApiUserTokenQuery {
	autq.ctx.Limit = &limit
	return autq
}

// Offset to start from.
func (autq *ApiUserTokenQuery) Offset(offset int) *ApiUserTokenQuery {
	autq.ctx.Offset = &offset
	return autq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (autq *ApiUserTokenQuery) Unique(unique bool) *ApiUserTokenQuery {
	autq.ctx.Unique = &unique
	return autq
}

// Order specifies how the records should be ordered.
func (autq *ApiUserTokenQuery) Order(o ...apiusertoken.OrderOption) *ApiUserTokenQuery {
	autq.order = append(autq.order, o...)
	return autq
}

// QueryAPIUser chains the current query on the "api_user" edge.
func (autq *ApiUserTokenQuery) QueryAPIUser() *ApiUserQuery {
	query := (&ApiUserClient{config: autq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := autq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := autq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(apiusertoken.Table, apiusertoken.FieldID, selector),
			sqlgraph.To(apiuser.Table, apiuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, apiusertoken.APIUserTable, apiusertoken.APIUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(autq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ApiUserToken entity from the query.
// Returns a *NotFoundError when no ApiUserToken was found.
func (autq *ApiUserTokenQuery) First(ctx context.Context) (*ApiUserToken, error) {
	nodes, err := autq.Limit(1).All(setContextOp(ctx, autq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apiusertoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (autq *ApiUserTokenQuery) FirstX(ctx context.Context) *ApiUserToken {
	node, err := autq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApiUserToken ID from the query.
// Returns a *NotFoundError when no ApiUserToken ID was found.
func (autq *ApiUserTokenQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = autq.Limit(1).IDs(setContextOp(ctx, autq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apiusertoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (autq *ApiUserTokenQuery) FirstIDX(ctx context.Context) string {
	id, err := autq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApiUserToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ApiUserToken entity is found.
// Returns a *NotFoundError when no ApiUserToken entities are found.
func (autq *ApiUserTokenQuery) Only(ctx context.Context) (*ApiUserToken, error) {
	nodes, err := autq.Limit(2).All(setContextOp(ctx, autq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apiusertoken.Label}
	default:
		return nil, &NotSingularError{apiusertoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (autq *ApiUserTokenQuery) OnlyX(ctx context.Context) *ApiUserToken {
	node, err := autq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApiUserToken ID in the query.
// Returns a *NotSingularError when more than one ApiUserToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (autq *ApiUserTokenQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = autq.Limit(2).IDs(setContextOp(ctx, autq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apiusertoken.Label}
	default:
		err = &NotSingularError{apiusertoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (autq *ApiUserTokenQuery) OnlyIDX(ctx context.Context) string {
	id, err := autq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApiUserTokens.
func (autq *ApiUserTokenQuery) All(ctx context.Context) ([]*ApiUserToken, error) {
	ctx = setContextOp(ctx, autq.ctx, "All")
	if err := autq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ApiUserToken, *ApiUserTokenQuery]()
	return withInterceptors[[]*ApiUserToken](ctx, autq, qr, autq.inters)
}

// AllX is like All, but panics if an error occurs.
func (autq *ApiUserTokenQuery) AllX(ctx context.Context) []*ApiUserToken {
	nodes, err := autq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApiUserToken IDs.
func (autq *ApiUserTokenQuery) IDs(ctx context.Context) (ids []string, err error) {
	if autq.ctx.Unique == nil && autq.path != nil {
		autq.Unique(true)
	}
	ctx = setContextOp(ctx, autq.ctx, "IDs")
	if err = autq.Select(apiusertoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (autq *ApiUserTokenQuery) IDsX(ctx context.Context) []string {
	ids, err := autq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (autq *ApiUserTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, autq.ctx, "Count")
	if err := autq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, autq, querierCount[*ApiUserTokenQuery](), autq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (autq *ApiUserTokenQuery) CountX(ctx context.Context) int {
	count, err := autq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (autq *ApiUserTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, autq.ctx, "Exist")
	switch _, err := autq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (autq *ApiUserTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := autq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApiUserTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (autq *ApiUserTokenQuery) Clone() *ApiUserTokenQuery {
	if autq == nil {
		return nil
	}
	return &ApiUserTokenQuery{
		config:      autq.config,
		ctx:         autq.ctx.Clone(),
		order:       append([]apiusertoken.OrderOption{}, autq.order...),
		inters:      append([]Interceptor{}, autq.inters...),
		predicates:  append([]predicate.ApiUserToken{}, autq.predicates...),
		withAPIUser: autq.withAPIUser.Clone(),
		// clone intermediate query.
		sql:  autq.sql.Clone(),
		path: autq.path,
	}
}

// WithAPIUser tells the query-builder to eager-load the nodes that are connected to
// the "api_user" edge. The optional arguments are used to configure the query builder of the edge.
func (autq *ApiUserTokenQuery) WithAPIUser(opts ...func(*ApiUserQuery)) *ApiUserTokenQuery {
	query := (&ApiUserClient{config: autq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	autq.withAPIUser = query
	return autq
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
//	client.ApiUserToken.Query().
//		GroupBy(apiusertoken.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (autq *ApiUserTokenQuery) GroupBy(field string, fields ...string) *ApiUserTokenGroupBy {
	autq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ApiUserTokenGroupBy{build: autq}
	grbuild.flds = &autq.ctx.Fields
	grbuild.label = apiusertoken.Label
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
//	client.ApiUserToken.Query().
//		Select(apiusertoken.FieldCreatedAt).
//		Scan(ctx, &v)
func (autq *ApiUserTokenQuery) Select(fields ...string) *ApiUserTokenSelect {
	autq.ctx.Fields = append(autq.ctx.Fields, fields...)
	sbuild := &ApiUserTokenSelect{ApiUserTokenQuery: autq}
	sbuild.label = apiusertoken.Label
	sbuild.flds, sbuild.scan = &autq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ApiUserTokenSelect configured with the given aggregations.
func (autq *ApiUserTokenQuery) Aggregate(fns ...AggregateFunc) *ApiUserTokenSelect {
	return autq.Select().Aggregate(fns...)
}

func (autq *ApiUserTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range autq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, autq); err != nil {
				return err
			}
		}
	}
	for _, f := range autq.ctx.Fields {
		if !apiusertoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if autq.path != nil {
		prev, err := autq.path(ctx)
		if err != nil {
			return err
		}
		autq.sql = prev
	}
	return nil
}

func (autq *ApiUserTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ApiUserToken, error) {
	var (
		nodes       = []*ApiUserToken{}
		withFKs     = autq.withFKs
		_spec       = autq.querySpec()
		loadedTypes = [1]bool{
			autq.withAPIUser != nil,
		}
	)
	if autq.withAPIUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, apiusertoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ApiUserToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ApiUserToken{config: autq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(autq.modifiers) > 0 {
		_spec.Modifiers = autq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, autq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := autq.withAPIUser; query != nil {
		if err := autq.loadAPIUser(ctx, query, nodes, nil,
			func(n *ApiUserToken, e *ApiUser) { n.Edges.APIUser = e }); err != nil {
			return nil, err
		}
	}
	for i := range autq.loadTotal {
		if err := autq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (autq *ApiUserTokenQuery) loadAPIUser(ctx context.Context, query *ApiUserQuery, nodes []*ApiUserToken, init func(*ApiUserToken), assign func(*ApiUserToken, *ApiUser)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ApiUserToken)
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

func (autq *ApiUserTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := autq.querySpec()
	if len(autq.modifiers) > 0 {
		_spec.Modifiers = autq.modifiers
	}
	_spec.Node.Columns = autq.ctx.Fields
	if len(autq.ctx.Fields) > 0 {
		_spec.Unique = autq.ctx.Unique != nil && *autq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, autq.driver, _spec)
}

func (autq *ApiUserTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(apiusertoken.Table, apiusertoken.Columns, sqlgraph.NewFieldSpec(apiusertoken.FieldID, field.TypeString))
	_spec.From = autq.sql
	if unique := autq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if autq.path != nil {
		_spec.Unique = true
	}
	if fields := autq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apiusertoken.FieldID)
		for i := range fields {
			if fields[i] != apiusertoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := autq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := autq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := autq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := autq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (autq *ApiUserTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(autq.driver.Dialect())
	t1 := builder.Table(apiusertoken.Table)
	columns := autq.ctx.Fields
	if len(columns) == 0 {
		columns = apiusertoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if autq.sql != nil {
		selector = autq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if autq.ctx.Unique != nil && *autq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range autq.modifiers {
		m(selector)
	}
	for _, p := range autq.predicates {
		p(selector)
	}
	for _, p := range autq.order {
		p(selector)
	}
	if offset := autq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := autq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (autq *ApiUserTokenQuery) Modify(modifiers ...func(s *sql.Selector)) *ApiUserTokenSelect {
	autq.modifiers = append(autq.modifiers, modifiers...)
	return autq.Select()
}

// ApiUserTokenGroupBy is the group-by builder for ApiUserToken entities.
type ApiUserTokenGroupBy struct {
	selector
	build *ApiUserTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (autgb *ApiUserTokenGroupBy) Aggregate(fns ...AggregateFunc) *ApiUserTokenGroupBy {
	autgb.fns = append(autgb.fns, fns...)
	return autgb
}

// Scan applies the selector query and scans the result into the given value.
func (autgb *ApiUserTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, autgb.build.ctx, "GroupBy")
	if err := autgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiUserTokenQuery, *ApiUserTokenGroupBy](ctx, autgb.build, autgb, autgb.build.inters, v)
}

func (autgb *ApiUserTokenGroupBy) sqlScan(ctx context.Context, root *ApiUserTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(autgb.fns))
	for _, fn := range autgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*autgb.flds)+len(autgb.fns))
		for _, f := range *autgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*autgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := autgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ApiUserTokenSelect is the builder for selecting fields of ApiUserToken entities.
type ApiUserTokenSelect struct {
	*ApiUserTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (auts *ApiUserTokenSelect) Aggregate(fns ...AggregateFunc) *ApiUserTokenSelect {
	auts.fns = append(auts.fns, fns...)
	return auts
}

// Scan applies the selector query and scans the result into the given value.
func (auts *ApiUserTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, auts.ctx, "Select")
	if err := auts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiUserTokenQuery, *ApiUserTokenSelect](ctx, auts.ApiUserTokenQuery, auts, auts.inters, v)
}

func (auts *ApiUserTokenSelect) sqlScan(ctx context.Context, root *ApiUserTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(auts.fns))
	for _, fn := range auts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*auts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := auts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (auts *ApiUserTokenSelect) Modify(modifiers ...func(s *sql.Selector)) *ApiUserTokenSelect {
	auts.modifiers = append(auts.modifiers, modifiers...)
	return auts
}