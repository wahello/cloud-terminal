// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/loginlog"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
)

// LoginLogQuery is the builder for querying LoginLog entities.
type LoginLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LoginLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoginLogQuery builder.
func (llq *LoginLogQuery) Where(ps ...predicate.LoginLog) *LoginLogQuery {
	llq.predicates = append(llq.predicates, ps...)
	return llq
}

// Limit adds a limit step to the query.
func (llq *LoginLogQuery) Limit(limit int) *LoginLogQuery {
	llq.limit = &limit
	return llq
}

// Offset adds an offset step to the query.
func (llq *LoginLogQuery) Offset(offset int) *LoginLogQuery {
	llq.offset = &offset
	return llq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (llq *LoginLogQuery) Unique(unique bool) *LoginLogQuery {
	llq.unique = &unique
	return llq
}

// Order adds an order step to the query.
func (llq *LoginLogQuery) Order(o ...OrderFunc) *LoginLogQuery {
	llq.order = append(llq.order, o...)
	return llq
}

// First returns the first LoginLog entity from the query.
// Returns a *NotFoundError when no LoginLog was found.
func (llq *LoginLogQuery) First(ctx context.Context) (*LoginLog, error) {
	nodes, err := llq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loginlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (llq *LoginLogQuery) FirstX(ctx context.Context) *LoginLog {
	node, err := llq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoginLog ID from the query.
// Returns a *NotFoundError when no LoginLog ID was found.
func (llq *LoginLogQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = llq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loginlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (llq *LoginLogQuery) FirstIDX(ctx context.Context) string {
	id, err := llq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoginLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one LoginLog entity is not found.
// Returns a *NotFoundError when no LoginLog entities are found.
func (llq *LoginLogQuery) Only(ctx context.Context) (*LoginLog, error) {
	nodes, err := llq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loginlog.Label}
	default:
		return nil, &NotSingularError{loginlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (llq *LoginLogQuery) OnlyX(ctx context.Context) *LoginLog {
	node, err := llq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoginLog ID in the query.
// Returns a *NotSingularError when exactly one LoginLog ID is not found.
// Returns a *NotFoundError when no entities are found.
func (llq *LoginLogQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = llq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = &NotSingularError{loginlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (llq *LoginLogQuery) OnlyIDX(ctx context.Context) string {
	id, err := llq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoginLogs.
func (llq *LoginLogQuery) All(ctx context.Context) ([]*LoginLog, error) {
	if err := llq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return llq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (llq *LoginLogQuery) AllX(ctx context.Context) []*LoginLog {
	nodes, err := llq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoginLog IDs.
func (llq *LoginLogQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := llq.Select(loginlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (llq *LoginLogQuery) IDsX(ctx context.Context) []string {
	ids, err := llq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (llq *LoginLogQuery) Count(ctx context.Context) (int, error) {
	if err := llq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return llq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (llq *LoginLogQuery) CountX(ctx context.Context) int {
	count, err := llq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (llq *LoginLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := llq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return llq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (llq *LoginLogQuery) ExistX(ctx context.Context) bool {
	exist, err := llq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoginLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (llq *LoginLogQuery) Clone() *LoginLogQuery {
	if llq == nil {
		return nil
	}
	return &LoginLogQuery{
		config:     llq.config,
		limit:      llq.limit,
		offset:     llq.offset,
		order:      append([]OrderFunc{}, llq.order...),
		predicates: append([]predicate.LoginLog{}, llq.predicates...),
		// clone intermediate query.
		sql:  llq.sql.Clone(),
		path: llq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LoginLog.Query().
//		GroupBy(loginlog.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (llq *LoginLogQuery) GroupBy(field string, fields ...string) *LoginLogGroupBy {
	group := &LoginLogGroupBy{config: llq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := llq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return llq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.LoginLog.Query().
//		Select(loginlog.FieldUserID).
//		Scan(ctx, &v)
//
func (llq *LoginLogQuery) Select(field string, fields ...string) *LoginLogSelect {
	llq.fields = append([]string{field}, fields...)
	return &LoginLogSelect{LoginLogQuery: llq}
}

func (llq *LoginLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range llq.fields {
		if !loginlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if llq.path != nil {
		prev, err := llq.path(ctx)
		if err != nil {
			return err
		}
		llq.sql = prev
	}
	return nil
}

func (llq *LoginLogQuery) sqlAll(ctx context.Context) ([]*LoginLog, error) {
	var (
		nodes = []*LoginLog{}
		_spec = llq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &LoginLog{config: llq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, llq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (llq *LoginLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := llq.querySpec()
	return sqlgraph.CountNodes(ctx, llq.driver, _spec)
}

func (llq *LoginLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := llq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (llq *LoginLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   loginlog.Table,
			Columns: loginlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: loginlog.FieldID,
			},
		},
		From:   llq.sql,
		Unique: true,
	}
	if unique := llq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := llq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginlog.FieldID)
		for i := range fields {
			if fields[i] != loginlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := llq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := llq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := llq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := llq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (llq *LoginLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(llq.driver.Dialect())
	t1 := builder.Table(loginlog.Table)
	selector := builder.Select(t1.Columns(loginlog.Columns...)...).From(t1)
	if llq.sql != nil {
		selector = llq.sql
		selector.Select(selector.Columns(loginlog.Columns...)...)
	}
	for _, p := range llq.predicates {
		p(selector)
	}
	for _, p := range llq.order {
		p(selector)
	}
	if offset := llq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := llq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoginLogGroupBy is the group-by builder for LoginLog entities.
type LoginLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (llgb *LoginLogGroupBy) Aggregate(fns ...AggregateFunc) *LoginLogGroupBy {
	llgb.fns = append(llgb.fns, fns...)
	return llgb
}

// Scan applies the group-by query and scans the result into the given value.
func (llgb *LoginLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := llgb.path(ctx)
	if err != nil {
		return err
	}
	llgb.sql = query
	return llgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (llgb *LoginLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := llgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(llgb.fields) > 1 {
		return nil, errors.New("ent: LoginLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := llgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (llgb *LoginLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := llgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = llgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (llgb *LoginLogGroupBy) StringX(ctx context.Context) string {
	v, err := llgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(llgb.fields) > 1 {
		return nil, errors.New("ent: LoginLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := llgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (llgb *LoginLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := llgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = llgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (llgb *LoginLogGroupBy) IntX(ctx context.Context) int {
	v, err := llgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(llgb.fields) > 1 {
		return nil, errors.New("ent: LoginLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := llgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (llgb *LoginLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := llgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = llgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (llgb *LoginLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := llgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(llgb.fields) > 1 {
		return nil, errors.New("ent: LoginLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := llgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (llgb *LoginLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := llgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (llgb *LoginLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = llgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (llgb *LoginLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := llgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (llgb *LoginLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range llgb.fields {
		if !loginlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := llgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := llgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (llgb *LoginLogGroupBy) sqlQuery() *sql.Selector {
	selector := llgb.sql
	columns := make([]string, 0, len(llgb.fields)+len(llgb.fns))
	columns = append(columns, llgb.fields...)
	for _, fn := range llgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(llgb.fields...)
}

// LoginLogSelect is the builder for selecting fields of LoginLog entities.
type LoginLogSelect struct {
	*LoginLogQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lls *LoginLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := lls.prepareQuery(ctx); err != nil {
		return err
	}
	lls.sql = lls.LoginLogQuery.sqlQuery(ctx)
	return lls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lls *LoginLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := lls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(lls.fields) > 1 {
		return nil, errors.New("ent: LoginLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := lls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lls *LoginLogSelect) StringsX(ctx context.Context) []string {
	v, err := lls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lls *LoginLogSelect) StringX(ctx context.Context) string {
	v, err := lls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(lls.fields) > 1 {
		return nil, errors.New("ent: LoginLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := lls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lls *LoginLogSelect) IntsX(ctx context.Context) []int {
	v, err := lls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lls *LoginLogSelect) IntX(ctx context.Context) int {
	v, err := lls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(lls.fields) > 1 {
		return nil, errors.New("ent: LoginLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := lls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lls *LoginLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := lls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lls *LoginLogSelect) Float64X(ctx context.Context) float64 {
	v, err := lls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(lls.fields) > 1 {
		return nil, errors.New("ent: LoginLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := lls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lls *LoginLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := lls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (lls *LoginLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loginlog.Label}
	default:
		err = fmt.Errorf("ent: LoginLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lls *LoginLogSelect) BoolX(ctx context.Context) bool {
	v, err := lls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lls *LoginLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lls.sqlQuery().Query()
	if err := lls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lls *LoginLogSelect) sqlQuery() sql.Querier {
	selector := lls.sql
	selector.Select(selector.Columns(lls.fields...)...)
	return selector
}