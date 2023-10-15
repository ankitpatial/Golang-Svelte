// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/predicate"
	"roofix/ent/usersession"
	"roofix/ent/usersessionsocket"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserSessionSocketQuery is the builder for querying UserSessionSocket entities.
type UserSessionSocketQuery struct {
	config
	ctx         *QueryContext
	order       []usersessionsocket.OrderOption
	inters      []Interceptor
	predicates  []predicate.UserSessionSocket
	withSession *UserSessionQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*UserSessionSocket) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserSessionSocketQuery builder.
func (ussq *UserSessionSocketQuery) Where(ps ...predicate.UserSessionSocket) *UserSessionSocketQuery {
	ussq.predicates = append(ussq.predicates, ps...)
	return ussq
}

// Limit the number of records to be returned by this query.
func (ussq *UserSessionSocketQuery) Limit(limit int) *UserSessionSocketQuery {
	ussq.ctx.Limit = &limit
	return ussq
}

// Offset to start from.
func (ussq *UserSessionSocketQuery) Offset(offset int) *UserSessionSocketQuery {
	ussq.ctx.Offset = &offset
	return ussq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ussq *UserSessionSocketQuery) Unique(unique bool) *UserSessionSocketQuery {
	ussq.ctx.Unique = &unique
	return ussq
}

// Order specifies how the records should be ordered.
func (ussq *UserSessionSocketQuery) Order(o ...usersessionsocket.OrderOption) *UserSessionSocketQuery {
	ussq.order = append(ussq.order, o...)
	return ussq
}

// QuerySession chains the current query on the "session" edge.
func (ussq *UserSessionSocketQuery) QuerySession() *UserSessionQuery {
	query := (&UserSessionClient{config: ussq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ussq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ussq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usersessionsocket.Table, usersessionsocket.FieldID, selector),
			sqlgraph.To(usersession.Table, usersession.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usersessionsocket.SessionTable, usersessionsocket.SessionColumn),
		)
		fromU = sqlgraph.SetNeighbors(ussq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserSessionSocket entity from the query.
// Returns a *NotFoundError when no UserSessionSocket was found.
func (ussq *UserSessionSocketQuery) First(ctx context.Context) (*UserSessionSocket, error) {
	nodes, err := ussq.Limit(1).All(setContextOp(ctx, ussq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usersessionsocket.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) FirstX(ctx context.Context) *UserSessionSocket {
	node, err := ussq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserSessionSocket ID from the query.
// Returns a *NotFoundError when no UserSessionSocket ID was found.
func (ussq *UserSessionSocketQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ussq.Limit(1).IDs(setContextOp(ctx, ussq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usersessionsocket.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) FirstIDX(ctx context.Context) string {
	id, err := ussq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserSessionSocket entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserSessionSocket entity is found.
// Returns a *NotFoundError when no UserSessionSocket entities are found.
func (ussq *UserSessionSocketQuery) Only(ctx context.Context) (*UserSessionSocket, error) {
	nodes, err := ussq.Limit(2).All(setContextOp(ctx, ussq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usersessionsocket.Label}
	default:
		return nil, &NotSingularError{usersessionsocket.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) OnlyX(ctx context.Context) *UserSessionSocket {
	node, err := ussq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserSessionSocket ID in the query.
// Returns a *NotSingularError when more than one UserSessionSocket ID is found.
// Returns a *NotFoundError when no entities are found.
func (ussq *UserSessionSocketQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ussq.Limit(2).IDs(setContextOp(ctx, ussq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usersessionsocket.Label}
	default:
		err = &NotSingularError{usersessionsocket.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) OnlyIDX(ctx context.Context) string {
	id, err := ussq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserSessionSockets.
func (ussq *UserSessionSocketQuery) All(ctx context.Context) ([]*UserSessionSocket, error) {
	ctx = setContextOp(ctx, ussq.ctx, "All")
	if err := ussq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserSessionSocket, *UserSessionSocketQuery]()
	return withInterceptors[[]*UserSessionSocket](ctx, ussq, qr, ussq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) AllX(ctx context.Context) []*UserSessionSocket {
	nodes, err := ussq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserSessionSocket IDs.
func (ussq *UserSessionSocketQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ussq.ctx.Unique == nil && ussq.path != nil {
		ussq.Unique(true)
	}
	ctx = setContextOp(ctx, ussq.ctx, "IDs")
	if err = ussq.Select(usersessionsocket.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) IDsX(ctx context.Context) []string {
	ids, err := ussq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ussq *UserSessionSocketQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ussq.ctx, "Count")
	if err := ussq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ussq, querierCount[*UserSessionSocketQuery](), ussq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) CountX(ctx context.Context) int {
	count, err := ussq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ussq *UserSessionSocketQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ussq.ctx, "Exist")
	switch _, err := ussq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ussq *UserSessionSocketQuery) ExistX(ctx context.Context) bool {
	exist, err := ussq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserSessionSocketQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ussq *UserSessionSocketQuery) Clone() *UserSessionSocketQuery {
	if ussq == nil {
		return nil
	}
	return &UserSessionSocketQuery{
		config:      ussq.config,
		ctx:         ussq.ctx.Clone(),
		order:       append([]usersessionsocket.OrderOption{}, ussq.order...),
		inters:      append([]Interceptor{}, ussq.inters...),
		predicates:  append([]predicate.UserSessionSocket{}, ussq.predicates...),
		withSession: ussq.withSession.Clone(),
		// clone intermediate query.
		sql:  ussq.sql.Clone(),
		path: ussq.path,
	}
}

// WithSession tells the query-builder to eager-load the nodes that are connected to
// the "session" edge. The optional arguments are used to configure the query builder of the edge.
func (ussq *UserSessionSocketQuery) WithSession(opts ...func(*UserSessionQuery)) *UserSessionSocketQuery {
	query := (&UserSessionClient{config: ussq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ussq.withSession = query
	return ussq
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
//	client.UserSessionSocket.Query().
//		GroupBy(usersessionsocket.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ussq *UserSessionSocketQuery) GroupBy(field string, fields ...string) *UserSessionSocketGroupBy {
	ussq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserSessionSocketGroupBy{build: ussq}
	grbuild.flds = &ussq.ctx.Fields
	grbuild.label = usersessionsocket.Label
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
//	client.UserSessionSocket.Query().
//		Select(usersessionsocket.FieldCreatedAt).
//		Scan(ctx, &v)
func (ussq *UserSessionSocketQuery) Select(fields ...string) *UserSessionSocketSelect {
	ussq.ctx.Fields = append(ussq.ctx.Fields, fields...)
	sbuild := &UserSessionSocketSelect{UserSessionSocketQuery: ussq}
	sbuild.label = usersessionsocket.Label
	sbuild.flds, sbuild.scan = &ussq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSessionSocketSelect configured with the given aggregations.
func (ussq *UserSessionSocketQuery) Aggregate(fns ...AggregateFunc) *UserSessionSocketSelect {
	return ussq.Select().Aggregate(fns...)
}

func (ussq *UserSessionSocketQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ussq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ussq); err != nil {
				return err
			}
		}
	}
	for _, f := range ussq.ctx.Fields {
		if !usersessionsocket.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ussq.path != nil {
		prev, err := ussq.path(ctx)
		if err != nil {
			return err
		}
		ussq.sql = prev
	}
	return nil
}

func (ussq *UserSessionSocketQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserSessionSocket, error) {
	var (
		nodes       = []*UserSessionSocket{}
		withFKs     = ussq.withFKs
		_spec       = ussq.querySpec()
		loadedTypes = [1]bool{
			ussq.withSession != nil,
		}
	)
	if ussq.withSession != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usersessionsocket.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserSessionSocket).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserSessionSocket{config: ussq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ussq.modifiers) > 0 {
		_spec.Modifiers = ussq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ussq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ussq.withSession; query != nil {
		if err := ussq.loadSession(ctx, query, nodes, nil,
			func(n *UserSessionSocket, e *UserSession) { n.Edges.Session = e }); err != nil {
			return nil, err
		}
	}
	for i := range ussq.loadTotal {
		if err := ussq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ussq *UserSessionSocketQuery) loadSession(ctx context.Context, query *UserSessionQuery, nodes []*UserSessionSocket, init func(*UserSessionSocket), assign func(*UserSessionSocket, *UserSession)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserSessionSocket)
	for i := range nodes {
		if nodes[i].sessions_id == nil {
			continue
		}
		fk := *nodes[i].sessions_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(usersession.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sessions_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ussq *UserSessionSocketQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ussq.querySpec()
	if len(ussq.modifiers) > 0 {
		_spec.Modifiers = ussq.modifiers
	}
	_spec.Node.Columns = ussq.ctx.Fields
	if len(ussq.ctx.Fields) > 0 {
		_spec.Unique = ussq.ctx.Unique != nil && *ussq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ussq.driver, _spec)
}

func (ussq *UserSessionSocketQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usersessionsocket.Table, usersessionsocket.Columns, sqlgraph.NewFieldSpec(usersessionsocket.FieldID, field.TypeString))
	_spec.From = ussq.sql
	if unique := ussq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ussq.path != nil {
		_spec.Unique = true
	}
	if fields := ussq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersessionsocket.FieldID)
		for i := range fields {
			if fields[i] != usersessionsocket.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ussq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ussq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ussq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ussq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ussq *UserSessionSocketQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ussq.driver.Dialect())
	t1 := builder.Table(usersessionsocket.Table)
	columns := ussq.ctx.Fields
	if len(columns) == 0 {
		columns = usersessionsocket.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ussq.sql != nil {
		selector = ussq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ussq.ctx.Unique != nil && *ussq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ussq.modifiers {
		m(selector)
	}
	for _, p := range ussq.predicates {
		p(selector)
	}
	for _, p := range ussq.order {
		p(selector)
	}
	if offset := ussq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ussq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ussq *UserSessionSocketQuery) Modify(modifiers ...func(s *sql.Selector)) *UserSessionSocketSelect {
	ussq.modifiers = append(ussq.modifiers, modifiers...)
	return ussq.Select()
}

// UserSessionSocketGroupBy is the group-by builder for UserSessionSocket entities.
type UserSessionSocketGroupBy struct {
	selector
	build *UserSessionSocketQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ussgb *UserSessionSocketGroupBy) Aggregate(fns ...AggregateFunc) *UserSessionSocketGroupBy {
	ussgb.fns = append(ussgb.fns, fns...)
	return ussgb
}

// Scan applies the selector query and scans the result into the given value.
func (ussgb *UserSessionSocketGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ussgb.build.ctx, "GroupBy")
	if err := ussgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSessionSocketQuery, *UserSessionSocketGroupBy](ctx, ussgb.build, ussgb, ussgb.build.inters, v)
}

func (ussgb *UserSessionSocketGroupBy) sqlScan(ctx context.Context, root *UserSessionSocketQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ussgb.fns))
	for _, fn := range ussgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ussgb.flds)+len(ussgb.fns))
		for _, f := range *ussgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ussgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ussgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSessionSocketSelect is the builder for selecting fields of UserSessionSocket entities.
type UserSessionSocketSelect struct {
	*UserSessionSocketQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (usss *UserSessionSocketSelect) Aggregate(fns ...AggregateFunc) *UserSessionSocketSelect {
	usss.fns = append(usss.fns, fns...)
	return usss
}

// Scan applies the selector query and scans the result into the given value.
func (usss *UserSessionSocketSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, usss.ctx, "Select")
	if err := usss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSessionSocketQuery, *UserSessionSocketSelect](ctx, usss.UserSessionSocketQuery, usss, usss.inters, v)
}

func (usss *UserSessionSocketSelect) sqlScan(ctx context.Context, root *UserSessionSocketQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(usss.fns))
	for _, fn := range usss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*usss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := usss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (usss *UserSessionSocketSelect) Modify(modifiers ...func(s *sql.Selector)) *UserSessionSocketSelect {
	usss.modifiers = append(usss.modifiers, modifiers...)
	return usss
}