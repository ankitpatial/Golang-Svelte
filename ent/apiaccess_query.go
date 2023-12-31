// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"roofix/ent/apiaccess"
	"roofix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiAccessQuery is the builder for querying ApiAccess entities.
type ApiAccessQuery struct {
	config
	ctx        *QueryContext
	order      []apiaccess.OrderOption
	inters     []Interceptor
	predicates []predicate.ApiAccess
	loadTotal  []func(context.Context, []*ApiAccess) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApiAccessQuery builder.
func (aaq *ApiAccessQuery) Where(ps ...predicate.ApiAccess) *ApiAccessQuery {
	aaq.predicates = append(aaq.predicates, ps...)
	return aaq
}

// Limit the number of records to be returned by this query.
func (aaq *ApiAccessQuery) Limit(limit int) *ApiAccessQuery {
	aaq.ctx.Limit = &limit
	return aaq
}

// Offset to start from.
func (aaq *ApiAccessQuery) Offset(offset int) *ApiAccessQuery {
	aaq.ctx.Offset = &offset
	return aaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aaq *ApiAccessQuery) Unique(unique bool) *ApiAccessQuery {
	aaq.ctx.Unique = &unique
	return aaq
}

// Order specifies how the records should be ordered.
func (aaq *ApiAccessQuery) Order(o ...apiaccess.OrderOption) *ApiAccessQuery {
	aaq.order = append(aaq.order, o...)
	return aaq
}

// First returns the first ApiAccess entity from the query.
// Returns a *NotFoundError when no ApiAccess was found.
func (aaq *ApiAccessQuery) First(ctx context.Context) (*ApiAccess, error) {
	nodes, err := aaq.Limit(1).All(setContextOp(ctx, aaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apiaccess.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aaq *ApiAccessQuery) FirstX(ctx context.Context) *ApiAccess {
	node, err := aaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApiAccess ID from the query.
// Returns a *NotFoundError when no ApiAccess ID was found.
func (aaq *ApiAccessQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aaq.Limit(1).IDs(setContextOp(ctx, aaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apiaccess.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aaq *ApiAccessQuery) FirstIDX(ctx context.Context) string {
	id, err := aaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApiAccess entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ApiAccess entity is found.
// Returns a *NotFoundError when no ApiAccess entities are found.
func (aaq *ApiAccessQuery) Only(ctx context.Context) (*ApiAccess, error) {
	nodes, err := aaq.Limit(2).All(setContextOp(ctx, aaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apiaccess.Label}
	default:
		return nil, &NotSingularError{apiaccess.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aaq *ApiAccessQuery) OnlyX(ctx context.Context) *ApiAccess {
	node, err := aaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApiAccess ID in the query.
// Returns a *NotSingularError when more than one ApiAccess ID is found.
// Returns a *NotFoundError when no entities are found.
func (aaq *ApiAccessQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aaq.Limit(2).IDs(setContextOp(ctx, aaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apiaccess.Label}
	default:
		err = &NotSingularError{apiaccess.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aaq *ApiAccessQuery) OnlyIDX(ctx context.Context) string {
	id, err := aaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApiAccesses.
func (aaq *ApiAccessQuery) All(ctx context.Context) ([]*ApiAccess, error) {
	ctx = setContextOp(ctx, aaq.ctx, "All")
	if err := aaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ApiAccess, *ApiAccessQuery]()
	return withInterceptors[[]*ApiAccess](ctx, aaq, qr, aaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aaq *ApiAccessQuery) AllX(ctx context.Context) []*ApiAccess {
	nodes, err := aaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApiAccess IDs.
func (aaq *ApiAccessQuery) IDs(ctx context.Context) (ids []string, err error) {
	if aaq.ctx.Unique == nil && aaq.path != nil {
		aaq.Unique(true)
	}
	ctx = setContextOp(ctx, aaq.ctx, "IDs")
	if err = aaq.Select(apiaccess.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aaq *ApiAccessQuery) IDsX(ctx context.Context) []string {
	ids, err := aaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aaq *ApiAccessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aaq.ctx, "Count")
	if err := aaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aaq, querierCount[*ApiAccessQuery](), aaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aaq *ApiAccessQuery) CountX(ctx context.Context) int {
	count, err := aaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aaq *ApiAccessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aaq.ctx, "Exist")
	switch _, err := aaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aaq *ApiAccessQuery) ExistX(ctx context.Context) bool {
	exist, err := aaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApiAccessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aaq *ApiAccessQuery) Clone() *ApiAccessQuery {
	if aaq == nil {
		return nil
	}
	return &ApiAccessQuery{
		config:     aaq.config,
		ctx:        aaq.ctx.Clone(),
		order:      append([]apiaccess.OrderOption{}, aaq.order...),
		inters:     append([]Interceptor{}, aaq.inters...),
		predicates: append([]predicate.ApiAccess{}, aaq.predicates...),
		// clone intermediate query.
		sql:  aaq.sql.Clone(),
		path: aaq.path,
	}
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
//	client.ApiAccess.Query().
//		GroupBy(apiaccess.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aaq *ApiAccessQuery) GroupBy(field string, fields ...string) *ApiAccessGroupBy {
	aaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ApiAccessGroupBy{build: aaq}
	grbuild.flds = &aaq.ctx.Fields
	grbuild.label = apiaccess.Label
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
//	client.ApiAccess.Query().
//		Select(apiaccess.FieldCreatedAt).
//		Scan(ctx, &v)
func (aaq *ApiAccessQuery) Select(fields ...string) *ApiAccessSelect {
	aaq.ctx.Fields = append(aaq.ctx.Fields, fields...)
	sbuild := &ApiAccessSelect{ApiAccessQuery: aaq}
	sbuild.label = apiaccess.Label
	sbuild.flds, sbuild.scan = &aaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ApiAccessSelect configured with the given aggregations.
func (aaq *ApiAccessQuery) Aggregate(fns ...AggregateFunc) *ApiAccessSelect {
	return aaq.Select().Aggregate(fns...)
}

func (aaq *ApiAccessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aaq); err != nil {
				return err
			}
		}
	}
	for _, f := range aaq.ctx.Fields {
		if !apiaccess.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aaq.path != nil {
		prev, err := aaq.path(ctx)
		if err != nil {
			return err
		}
		aaq.sql = prev
	}
	return nil
}

func (aaq *ApiAccessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ApiAccess, error) {
	var (
		nodes = []*ApiAccess{}
		_spec = aaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ApiAccess).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ApiAccess{config: aaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(aaq.modifiers) > 0 {
		_spec.Modifiers = aaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range aaq.loadTotal {
		if err := aaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aaq *ApiAccessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aaq.querySpec()
	if len(aaq.modifiers) > 0 {
		_spec.Modifiers = aaq.modifiers
	}
	_spec.Node.Columns = aaq.ctx.Fields
	if len(aaq.ctx.Fields) > 0 {
		_spec.Unique = aaq.ctx.Unique != nil && *aaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aaq.driver, _spec)
}

func (aaq *ApiAccessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(apiaccess.Table, apiaccess.Columns, sqlgraph.NewFieldSpec(apiaccess.FieldID, field.TypeString))
	_spec.From = aaq.sql
	if unique := aaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aaq.path != nil {
		_spec.Unique = true
	}
	if fields := aaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apiaccess.FieldID)
		for i := range fields {
			if fields[i] != apiaccess.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aaq *ApiAccessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aaq.driver.Dialect())
	t1 := builder.Table(apiaccess.Table)
	columns := aaq.ctx.Fields
	if len(columns) == 0 {
		columns = apiaccess.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aaq.sql != nil {
		selector = aaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aaq.ctx.Unique != nil && *aaq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aaq.modifiers {
		m(selector)
	}
	for _, p := range aaq.predicates {
		p(selector)
	}
	for _, p := range aaq.order {
		p(selector)
	}
	if offset := aaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aaq *ApiAccessQuery) Modify(modifiers ...func(s *sql.Selector)) *ApiAccessSelect {
	aaq.modifiers = append(aaq.modifiers, modifiers...)
	return aaq.Select()
}

// ApiAccessGroupBy is the group-by builder for ApiAccess entities.
type ApiAccessGroupBy struct {
	selector
	build *ApiAccessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aagb *ApiAccessGroupBy) Aggregate(fns ...AggregateFunc) *ApiAccessGroupBy {
	aagb.fns = append(aagb.fns, fns...)
	return aagb
}

// Scan applies the selector query and scans the result into the given value.
func (aagb *ApiAccessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aagb.build.ctx, "GroupBy")
	if err := aagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiAccessQuery, *ApiAccessGroupBy](ctx, aagb.build, aagb, aagb.build.inters, v)
}

func (aagb *ApiAccessGroupBy) sqlScan(ctx context.Context, root *ApiAccessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aagb.fns))
	for _, fn := range aagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aagb.flds)+len(aagb.fns))
		for _, f := range *aagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ApiAccessSelect is the builder for selecting fields of ApiAccess entities.
type ApiAccessSelect struct {
	*ApiAccessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aas *ApiAccessSelect) Aggregate(fns ...AggregateFunc) *ApiAccessSelect {
	aas.fns = append(aas.fns, fns...)
	return aas
}

// Scan applies the selector query and scans the result into the given value.
func (aas *ApiAccessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aas.ctx, "Select")
	if err := aas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiAccessQuery, *ApiAccessSelect](ctx, aas.ApiAccessQuery, aas, aas.inters, v)
}

func (aas *ApiAccessSelect) sqlScan(ctx context.Context, root *ApiAccessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aas.fns))
	for _, fn := range aas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aas *ApiAccessSelect) Modify(modifiers ...func(s *sql.Selector)) *ApiAccessSelect {
	aas.modifiers = append(aas.modifiers, modifiers...)
	return aas
}
