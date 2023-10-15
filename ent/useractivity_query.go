// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/apiuser"
	"roofix/ent/predicate"
	"roofix/ent/user"
	"roofix/ent/useractivity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserActivityQuery is the builder for querying UserActivity entities.
type UserActivityQuery struct {
	config
	ctx            *QueryContext
	order          []useractivity.OrderOption
	inters         []Interceptor
	predicates     []predicate.UserActivity
	withUser       *UserQuery
	withCreator    *UserQuery
	withCreatorAPI *ApiUserQuery
	withFKs        bool
	loadTotal      []func(context.Context, []*UserActivity) error
	modifiers      []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserActivityQuery builder.
func (uaq *UserActivityQuery) Where(ps ...predicate.UserActivity) *UserActivityQuery {
	uaq.predicates = append(uaq.predicates, ps...)
	return uaq
}

// Limit the number of records to be returned by this query.
func (uaq *UserActivityQuery) Limit(limit int) *UserActivityQuery {
	uaq.ctx.Limit = &limit
	return uaq
}

// Offset to start from.
func (uaq *UserActivityQuery) Offset(offset int) *UserActivityQuery {
	uaq.ctx.Offset = &offset
	return uaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uaq *UserActivityQuery) Unique(unique bool) *UserActivityQuery {
	uaq.ctx.Unique = &unique
	return uaq
}

// Order specifies how the records should be ordered.
func (uaq *UserActivityQuery) Order(o ...useractivity.OrderOption) *UserActivityQuery {
	uaq.order = append(uaq.order, o...)
	return uaq
}

// QueryUser chains the current query on the "user" edge.
func (uaq *UserActivityQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useractivity.Table, useractivity.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useractivity.UserTable, useractivity.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreator chains the current query on the "creator" edge.
func (uaq *UserActivityQuery) QueryCreator() *UserQuery {
	query := (&UserClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useractivity.Table, useractivity.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useractivity.CreatorTable, useractivity.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatorAPI chains the current query on the "creator_api" edge.
func (uaq *UserActivityQuery) QueryCreatorAPI() *ApiUserQuery {
	query := (&ApiUserClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useractivity.Table, useractivity.FieldID, selector),
			sqlgraph.To(apiuser.Table, apiuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useractivity.CreatorAPITable, useractivity.CreatorAPIColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserActivity entity from the query.
// Returns a *NotFoundError when no UserActivity was found.
func (uaq *UserActivityQuery) First(ctx context.Context) (*UserActivity, error) {
	nodes, err := uaq.Limit(1).All(setContextOp(ctx, uaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{useractivity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaq *UserActivityQuery) FirstX(ctx context.Context) *UserActivity {
	node, err := uaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserActivity ID from the query.
// Returns a *NotFoundError when no UserActivity ID was found.
func (uaq *UserActivityQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uaq.Limit(1).IDs(setContextOp(ctx, uaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{useractivity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uaq *UserActivityQuery) FirstIDX(ctx context.Context) string {
	id, err := uaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserActivity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserActivity entity is found.
// Returns a *NotFoundError when no UserActivity entities are found.
func (uaq *UserActivityQuery) Only(ctx context.Context) (*UserActivity, error) {
	nodes, err := uaq.Limit(2).All(setContextOp(ctx, uaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{useractivity.Label}
	default:
		return nil, &NotSingularError{useractivity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaq *UserActivityQuery) OnlyX(ctx context.Context) *UserActivity {
	node, err := uaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserActivity ID in the query.
// Returns a *NotSingularError when more than one UserActivity ID is found.
// Returns a *NotFoundError when no entities are found.
func (uaq *UserActivityQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uaq.Limit(2).IDs(setContextOp(ctx, uaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{useractivity.Label}
	default:
		err = &NotSingularError{useractivity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uaq *UserActivityQuery) OnlyIDX(ctx context.Context) string {
	id, err := uaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserActivities.
func (uaq *UserActivityQuery) All(ctx context.Context) ([]*UserActivity, error) {
	ctx = setContextOp(ctx, uaq.ctx, "All")
	if err := uaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserActivity, *UserActivityQuery]()
	return withInterceptors[[]*UserActivity](ctx, uaq, qr, uaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uaq *UserActivityQuery) AllX(ctx context.Context) []*UserActivity {
	nodes, err := uaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserActivity IDs.
func (uaq *UserActivityQuery) IDs(ctx context.Context) (ids []string, err error) {
	if uaq.ctx.Unique == nil && uaq.path != nil {
		uaq.Unique(true)
	}
	ctx = setContextOp(ctx, uaq.ctx, "IDs")
	if err = uaq.Select(useractivity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaq *UserActivityQuery) IDsX(ctx context.Context) []string {
	ids, err := uaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaq *UserActivityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uaq.ctx, "Count")
	if err := uaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uaq, querierCount[*UserActivityQuery](), uaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uaq *UserActivityQuery) CountX(ctx context.Context) int {
	count, err := uaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaq *UserActivityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uaq.ctx, "Exist")
	switch _, err := uaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uaq *UserActivityQuery) ExistX(ctx context.Context) bool {
	exist, err := uaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserActivityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaq *UserActivityQuery) Clone() *UserActivityQuery {
	if uaq == nil {
		return nil
	}
	return &UserActivityQuery{
		config:         uaq.config,
		ctx:            uaq.ctx.Clone(),
		order:          append([]useractivity.OrderOption{}, uaq.order...),
		inters:         append([]Interceptor{}, uaq.inters...),
		predicates:     append([]predicate.UserActivity{}, uaq.predicates...),
		withUser:       uaq.withUser.Clone(),
		withCreator:    uaq.withCreator.Clone(),
		withCreatorAPI: uaq.withCreatorAPI.Clone(),
		// clone intermediate query.
		sql:  uaq.sql.Clone(),
		path: uaq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserActivityQuery) WithUser(opts ...func(*UserQuery)) *UserActivityQuery {
	query := (&UserClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withUser = query
	return uaq
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserActivityQuery) WithCreator(opts ...func(*UserQuery)) *UserActivityQuery {
	query := (&UserClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withCreator = query
	return uaq
}

// WithCreatorAPI tells the query-builder to eager-load the nodes that are connected to
// the "creator_api" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserActivityQuery) WithCreatorAPI(opts ...func(*ApiUserQuery)) *UserActivityQuery {
	query := (&ApiUserClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withCreatorAPI = query
	return uaq
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
//	client.UserActivity.Query().
//		GroupBy(useractivity.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uaq *UserActivityQuery) GroupBy(field string, fields ...string) *UserActivityGroupBy {
	uaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserActivityGroupBy{build: uaq}
	grbuild.flds = &uaq.ctx.Fields
	grbuild.label = useractivity.Label
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
//	client.UserActivity.Query().
//		Select(useractivity.FieldCreatedAt).
//		Scan(ctx, &v)
func (uaq *UserActivityQuery) Select(fields ...string) *UserActivitySelect {
	uaq.ctx.Fields = append(uaq.ctx.Fields, fields...)
	sbuild := &UserActivitySelect{UserActivityQuery: uaq}
	sbuild.label = useractivity.Label
	sbuild.flds, sbuild.scan = &uaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserActivitySelect configured with the given aggregations.
func (uaq *UserActivityQuery) Aggregate(fns ...AggregateFunc) *UserActivitySelect {
	return uaq.Select().Aggregate(fns...)
}

func (uaq *UserActivityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uaq); err != nil {
				return err
			}
		}
	}
	for _, f := range uaq.ctx.Fields {
		if !useractivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uaq.path != nil {
		prev, err := uaq.path(ctx)
		if err != nil {
			return err
		}
		uaq.sql = prev
	}
	return nil
}

func (uaq *UserActivityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserActivity, error) {
	var (
		nodes       = []*UserActivity{}
		withFKs     = uaq.withFKs
		_spec       = uaq.querySpec()
		loadedTypes = [3]bool{
			uaq.withUser != nil,
			uaq.withCreator != nil,
			uaq.withCreatorAPI != nil,
		}
	)
	if uaq.withUser != nil || uaq.withCreator != nil || uaq.withCreatorAPI != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, useractivity.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserActivity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserActivity{config: uaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uaq.modifiers) > 0 {
		_spec.Modifiers = uaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uaq.withUser; query != nil {
		if err := uaq.loadUser(ctx, query, nodes, nil,
			func(n *UserActivity, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withCreator; query != nil {
		if err := uaq.loadCreator(ctx, query, nodes, nil,
			func(n *UserActivity, e *User) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withCreatorAPI; query != nil {
		if err := uaq.loadCreatorAPI(ctx, query, nodes, nil,
			func(n *UserActivity, e *ApiUser) { n.Edges.CreatorAPI = e }); err != nil {
			return nil, err
		}
	}
	for i := range uaq.loadTotal {
		if err := uaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uaq *UserActivityQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserActivity, init func(*UserActivity), assign func(*UserActivity, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserActivity)
	for i := range nodes {
		if nodes[i].user_id == nil {
			continue
		}
		fk := *nodes[i].user_id
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uaq *UserActivityQuery) loadCreator(ctx context.Context, query *UserQuery, nodes []*UserActivity, init func(*UserActivity), assign func(*UserActivity, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserActivity)
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
func (uaq *UserActivityQuery) loadCreatorAPI(ctx context.Context, query *ApiUserQuery, nodes []*UserActivity, init func(*UserActivity), assign func(*UserActivity, *ApiUser)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserActivity)
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

func (uaq *UserActivityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaq.querySpec()
	if len(uaq.modifiers) > 0 {
		_spec.Modifiers = uaq.modifiers
	}
	_spec.Node.Columns = uaq.ctx.Fields
	if len(uaq.ctx.Fields) > 0 {
		_spec.Unique = uaq.ctx.Unique != nil && *uaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uaq.driver, _spec)
}

func (uaq *UserActivityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(useractivity.Table, useractivity.Columns, sqlgraph.NewFieldSpec(useractivity.FieldID, field.TypeString))
	_spec.From = uaq.sql
	if unique := uaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uaq.path != nil {
		_spec.Unique = true
	}
	if fields := uaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useractivity.FieldID)
		for i := range fields {
			if fields[i] != useractivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaq *UserActivityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uaq.driver.Dialect())
	t1 := builder.Table(useractivity.Table)
	columns := uaq.ctx.Fields
	if len(columns) == 0 {
		columns = useractivity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uaq.sql != nil {
		selector = uaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uaq.ctx.Unique != nil && *uaq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range uaq.modifiers {
		m(selector)
	}
	for _, p := range uaq.predicates {
		p(selector)
	}
	for _, p := range uaq.order {
		p(selector)
	}
	if offset := uaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (uaq *UserActivityQuery) Modify(modifiers ...func(s *sql.Selector)) *UserActivitySelect {
	uaq.modifiers = append(uaq.modifiers, modifiers...)
	return uaq.Select()
}

// UserActivityGroupBy is the group-by builder for UserActivity entities.
type UserActivityGroupBy struct {
	selector
	build *UserActivityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uagb *UserActivityGroupBy) Aggregate(fns ...AggregateFunc) *UserActivityGroupBy {
	uagb.fns = append(uagb.fns, fns...)
	return uagb
}

// Scan applies the selector query and scans the result into the given value.
func (uagb *UserActivityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uagb.build.ctx, "GroupBy")
	if err := uagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserActivityQuery, *UserActivityGroupBy](ctx, uagb.build, uagb, uagb.build.inters, v)
}

func (uagb *UserActivityGroupBy) sqlScan(ctx context.Context, root *UserActivityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uagb.fns))
	for _, fn := range uagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uagb.flds)+len(uagb.fns))
		for _, f := range *uagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserActivitySelect is the builder for selecting fields of UserActivity entities.
type UserActivitySelect struct {
	*UserActivityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uas *UserActivitySelect) Aggregate(fns ...AggregateFunc) *UserActivitySelect {
	uas.fns = append(uas.fns, fns...)
	return uas
}

// Scan applies the selector query and scans the result into the given value.
func (uas *UserActivitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uas.ctx, "Select")
	if err := uas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserActivityQuery, *UserActivitySelect](ctx, uas.UserActivityQuery, uas, uas.inters, v)
}

func (uas *UserActivitySelect) sqlScan(ctx context.Context, root *UserActivityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uas.fns))
	for _, fn := range uas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (uas *UserActivitySelect) Modify(modifiers ...func(s *sql.Selector)) *UserActivitySelect {
	uas.modifiers = append(uas.modifiers, modifiers...)
	return uas
}