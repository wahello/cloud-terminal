// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/user"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/verification"
)

// VerificationQuery is the builder for querying Verification entities.
type VerificationQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Verification
	// eager-loading edges.
	withUsers *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VerificationQuery builder.
func (vq *VerificationQuery) Where(ps ...predicate.Verification) *VerificationQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VerificationQuery) Limit(limit int) *VerificationQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VerificationQuery) Offset(offset int) *VerificationQuery {
	vq.offset = &offset
	return vq
}

// Order adds an order step to the query.
func (vq *VerificationQuery) Order(o ...OrderFunc) *VerificationQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryUsers chains the current query on the "users" edge.
func (vq *VerificationQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(verification.Table, verification.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, verification.UsersTable, verification.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Verification entity from the query.
// Returns a *NotFoundError when no Verification was found.
func (vq *VerificationQuery) First(ctx context.Context) (*Verification, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{verification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VerificationQuery) FirstX(ctx context.Context) *Verification {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Verification ID from the query.
// Returns a *NotFoundError when no Verification ID was found.
func (vq *VerificationQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{verification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VerificationQuery) FirstIDX(ctx context.Context) string {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Verification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Verification entity is not found.
// Returns a *NotFoundError when no Verification entities are found.
func (vq *VerificationQuery) Only(ctx context.Context) (*Verification, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{verification.Label}
	default:
		return nil, &NotSingularError{verification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VerificationQuery) OnlyX(ctx context.Context) *Verification {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Verification ID in the query.
// Returns a *NotSingularError when exactly one Verification ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vq *VerificationQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = &NotSingularError{verification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VerificationQuery) OnlyIDX(ctx context.Context) string {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Verifications.
func (vq *VerificationQuery) All(ctx context.Context) ([]*Verification, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VerificationQuery) AllX(ctx context.Context) []*Verification {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Verification IDs.
func (vq *VerificationQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := vq.Select(verification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VerificationQuery) IDsX(ctx context.Context) []string {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VerificationQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VerificationQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VerificationQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VerificationQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VerificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VerificationQuery) Clone() *VerificationQuery {
	if vq == nil {
		return nil
	}
	return &VerificationQuery{
		config:     vq.config,
		limit:      vq.limit,
		offset:     vq.offset,
		order:      append([]OrderFunc{}, vq.order...),
		predicates: append([]predicate.Verification{}, vq.predicates...),
		withUsers:  vq.withUsers.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VerificationQuery) WithUsers(opts ...func(*UserQuery)) *VerificationQuery {
	query := &UserQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withUsers = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClientIP string `json:"client_ip,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Verification.Query().
//		GroupBy(verification.FieldClientIP).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *VerificationQuery) GroupBy(field string, fields ...string) *VerificationGroupBy {
	group := &VerificationGroupBy{config: vq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClientIP string `json:"client_ip,omitempty"`
//	}
//
//	client.Verification.Query().
//		Select(verification.FieldClientIP).
//		Scan(ctx, &v)
//
func (vq *VerificationQuery) Select(field string, fields ...string) *VerificationSelect {
	vq.fields = append([]string{field}, fields...)
	return &VerificationSelect{VerificationQuery: vq}
}

func (vq *VerificationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !verification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VerificationQuery) sqlAll(ctx context.Context) ([]*Verification, error) {
	var (
		nodes       = []*Verification{}
		_spec       = vq.querySpec()
		loadedTypes = [1]bool{
			vq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Verification{config: vq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := vq.withUsers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*Verification)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Users = []*User{}
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(verification.UsersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.verification_users
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "verification_users" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "verification_users" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Users = append(node.Edges.Users, n)
		}
	}

	return nodes, nil
}

func (vq *VerificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VerificationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (vq *VerificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   verification.Table,
			Columns: verification.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: verification.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, verification.FieldID)
		for i := range fields {
			if fields[i] != verification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, verification.ValidColumn)
			}
		}
	}
	return _spec
}

func (vq *VerificationQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(verification.Table)
	selector := builder.Select(t1.Columns(verification.Columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(verification.Columns...)...)
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector, verification.ValidColumn)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VerificationGroupBy is the group-by builder for Verification entities.
type VerificationGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VerificationGroupBy) Aggregate(fns ...AggregateFunc) *VerificationGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VerificationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vgb *VerificationGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VerificationGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vgb *VerificationGroupBy) StringsX(ctx context.Context) []string {
	v, err := vgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vgb *VerificationGroupBy) StringX(ctx context.Context) string {
	v, err := vgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VerificationGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vgb *VerificationGroupBy) IntsX(ctx context.Context) []int {
	v, err := vgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vgb *VerificationGroupBy) IntX(ctx context.Context) int {
	v, err := vgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VerificationGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vgb *VerificationGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vgb *VerificationGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VerificationGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vgb *VerificationGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VerificationGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vgb *VerificationGroupBy) BoolX(ctx context.Context) bool {
	v, err := vgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vgb *VerificationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vgb.fields {
		if !verification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *VerificationGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql
	columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
	columns = append(columns, vgb.fields...)
	for _, fn := range vgb.fns {
		columns = append(columns, fn(selector, verification.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(vgb.fields...)
}

// VerificationSelect is the builder for selecting fields of Verification entities.
type VerificationSelect struct {
	*VerificationQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VerificationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VerificationQuery.sqlQuery()
	return vs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vs *VerificationSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VerificationSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vs *VerificationSelect) StringsX(ctx context.Context) []string {
	v, err := vs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vs *VerificationSelect) StringX(ctx context.Context) string {
	v, err := vs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VerificationSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vs *VerificationSelect) IntsX(ctx context.Context) []int {
	v, err := vs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vs *VerificationSelect) IntX(ctx context.Context) int {
	v, err := vs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VerificationSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vs *VerificationSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vs *VerificationSelect) Float64X(ctx context.Context) float64 {
	v, err := vs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VerificationSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vs *VerificationSelect) BoolsX(ctx context.Context) []bool {
	v, err := vs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vs *VerificationSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{verification.Label}
	default:
		err = fmt.Errorf("ent: VerificationSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vs *VerificationSelect) BoolX(ctx context.Context) bool {
	v, err := vs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vs *VerificationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vs.sqlQuery().Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vs *VerificationSelect) sqlQuery() sql.Querier {
	selector := vs.sql
	selector.Select(selector.Columns(vs.fields...)...)
	return selector
}