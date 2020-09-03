// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/zhanghongnian/brand-index/ent/predicate"
	"github.com/zhanghongnian/brand-index/ent/qccenterprisedata"
)

// QccEnterpriseDataQuery is the builder for querying QccEnterpriseData entities.
type QccEnterpriseDataQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.QccEnterpriseData
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (qedq *QccEnterpriseDataQuery) Where(ps ...predicate.QccEnterpriseData) *QccEnterpriseDataQuery {
	qedq.predicates = append(qedq.predicates, ps...)
	return qedq
}

// Limit adds a limit step to the query.
func (qedq *QccEnterpriseDataQuery) Limit(limit int) *QccEnterpriseDataQuery {
	qedq.limit = &limit
	return qedq
}

// Offset adds an offset step to the query.
func (qedq *QccEnterpriseDataQuery) Offset(offset int) *QccEnterpriseDataQuery {
	qedq.offset = &offset
	return qedq
}

// Order adds an order step to the query.
func (qedq *QccEnterpriseDataQuery) Order(o ...OrderFunc) *QccEnterpriseDataQuery {
	qedq.order = append(qedq.order, o...)
	return qedq
}

// First returns the first QccEnterpriseData entity in the query. Returns *NotFoundError when no qccenterprisedata was found.
func (qedq *QccEnterpriseDataQuery) First(ctx context.Context) (*QccEnterpriseData, error) {
	qeds, err := qedq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(qeds) == 0 {
		return nil, &NotFoundError{qccenterprisedata.Label}
	}
	return qeds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) FirstX(ctx context.Context) *QccEnterpriseData {
	qed, err := qedq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return qed
}

// FirstID returns the first QccEnterpriseData id in the query. Returns *NotFoundError when no id was found.
func (qedq *QccEnterpriseDataQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qedq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{qccenterprisedata.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) FirstXID(ctx context.Context) int {
	id, err := qedq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only QccEnterpriseData entity in the query, returns an error if not exactly one entity was returned.
func (qedq *QccEnterpriseDataQuery) Only(ctx context.Context) (*QccEnterpriseData, error) {
	qeds, err := qedq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(qeds) {
	case 1:
		return qeds[0], nil
	case 0:
		return nil, &NotFoundError{qccenterprisedata.Label}
	default:
		return nil, &NotSingularError{qccenterprisedata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) OnlyX(ctx context.Context) *QccEnterpriseData {
	qed, err := qedq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return qed
}

// OnlyID returns the only QccEnterpriseData id in the query, returns an error if not exactly one id was returned.
func (qedq *QccEnterpriseDataQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qedq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = &NotSingularError{qccenterprisedata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) OnlyIDX(ctx context.Context) int {
	id, err := qedq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of QccEnterpriseDataSlice.
func (qedq *QccEnterpriseDataQuery) All(ctx context.Context) ([]*QccEnterpriseData, error) {
	if err := qedq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qedq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) AllX(ctx context.Context) []*QccEnterpriseData {
	qeds, err := qedq.All(ctx)
	if err != nil {
		panic(err)
	}
	return qeds
}

// IDs executes the query and returns a list of QccEnterpriseData ids.
func (qedq *QccEnterpriseDataQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := qedq.Select(qccenterprisedata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) IDsX(ctx context.Context) []int {
	ids, err := qedq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qedq *QccEnterpriseDataQuery) Count(ctx context.Context) (int, error) {
	if err := qedq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qedq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) CountX(ctx context.Context) int {
	count, err := qedq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qedq *QccEnterpriseDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := qedq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qedq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qedq *QccEnterpriseDataQuery) ExistX(ctx context.Context) bool {
	exist, err := qedq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qedq *QccEnterpriseDataQuery) Clone() *QccEnterpriseDataQuery {
	return &QccEnterpriseDataQuery{
		config:     qedq.config,
		limit:      qedq.limit,
		offset:     qedq.offset,
		order:      append([]OrderFunc{}, qedq.order...),
		unique:     append([]string{}, qedq.unique...),
		predicates: append([]predicate.QccEnterpriseData{}, qedq.predicates...),
		// clone intermediate query.
		sql:  qedq.sql.Clone(),
		path: qedq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.QccEnterpriseData.Query().
//		GroupBy(qccenterprisedata.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (qedq *QccEnterpriseDataQuery) GroupBy(field string, fields ...string) *QccEnterpriseDataGroupBy {
	group := &QccEnterpriseDataGroupBy{config: qedq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qedq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qedq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.QccEnterpriseData.Query().
//		Select(qccenterprisedata.FieldName).
//		Scan(ctx, &v)
//
func (qedq *QccEnterpriseDataQuery) Select(field string, fields ...string) *QccEnterpriseDataSelect {
	selector := &QccEnterpriseDataSelect{config: qedq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qedq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qedq.sqlQuery(), nil
	}
	return selector
}

func (qedq *QccEnterpriseDataQuery) prepareQuery(ctx context.Context) error {
	if qedq.path != nil {
		prev, err := qedq.path(ctx)
		if err != nil {
			return err
		}
		qedq.sql = prev
	}
	return nil
}

func (qedq *QccEnterpriseDataQuery) sqlAll(ctx context.Context) ([]*QccEnterpriseData, error) {
	var (
		nodes = []*QccEnterpriseData{}
		_spec = qedq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &QccEnterpriseData{config: qedq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, qedq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (qedq *QccEnterpriseDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qedq.querySpec()
	return sqlgraph.CountNodes(ctx, qedq.driver, _spec)
}

func (qedq *QccEnterpriseDataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := qedq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (qedq *QccEnterpriseDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   qccenterprisedata.Table,
			Columns: qccenterprisedata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: qccenterprisedata.FieldID,
			},
		},
		From:   qedq.sql,
		Unique: true,
	}
	if ps := qedq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qedq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qedq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qedq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qedq *QccEnterpriseDataQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(qedq.driver.Dialect())
	t1 := builder.Table(qccenterprisedata.Table)
	selector := builder.Select(t1.Columns(qccenterprisedata.Columns...)...).From(t1)
	if qedq.sql != nil {
		selector = qedq.sql
		selector.Select(selector.Columns(qccenterprisedata.Columns...)...)
	}
	for _, p := range qedq.predicates {
		p(selector)
	}
	for _, p := range qedq.order {
		p(selector)
	}
	if offset := qedq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qedq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QccEnterpriseDataGroupBy is the builder for group-by QccEnterpriseData entities.
type QccEnterpriseDataGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qedgb *QccEnterpriseDataGroupBy) Aggregate(fns ...AggregateFunc) *QccEnterpriseDataGroupBy {
	qedgb.fns = append(qedgb.fns, fns...)
	return qedgb
}

// Scan applies the group-by query and scan the result into the given value.
func (qedgb *QccEnterpriseDataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := qedgb.path(ctx)
	if err != nil {
		return err
	}
	qedgb.sql = query
	return qedgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := qedgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(qedgb.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := qedgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) StringsX(ctx context.Context) []string {
	v, err := qedgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = qedgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) StringX(ctx context.Context) string {
	v, err := qedgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(qedgb.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := qedgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) IntsX(ctx context.Context) []int {
	v, err := qedgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = qedgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) IntX(ctx context.Context) int {
	v, err := qedgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(qedgb.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := qedgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := qedgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = qedgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) Float64X(ctx context.Context) float64 {
	v, err := qedgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(qedgb.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := qedgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := qedgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (qedgb *QccEnterpriseDataGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = qedgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (qedgb *QccEnterpriseDataGroupBy) BoolX(ctx context.Context) bool {
	v, err := qedgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qedgb *QccEnterpriseDataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := qedgb.sqlQuery().Query()
	if err := qedgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qedgb *QccEnterpriseDataGroupBy) sqlQuery() *sql.Selector {
	selector := qedgb.sql
	columns := make([]string, 0, len(qedgb.fields)+len(qedgb.fns))
	columns = append(columns, qedgb.fields...)
	for _, fn := range qedgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(qedgb.fields...)
}

// QccEnterpriseDataSelect is the builder for select fields of QccEnterpriseData entities.
type QccEnterpriseDataSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (qeds *QccEnterpriseDataSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := qeds.path(ctx)
	if err != nil {
		return err
	}
	qeds.sql = query
	return qeds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) ScanX(ctx context.Context, v interface{}) {
	if err := qeds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) Strings(ctx context.Context) ([]string, error) {
	if len(qeds.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := qeds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) StringsX(ctx context.Context) []string {
	v, err := qeds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = qeds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) StringX(ctx context.Context) string {
	v, err := qeds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) Ints(ctx context.Context) ([]int, error) {
	if len(qeds.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := qeds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) IntsX(ctx context.Context) []int {
	v, err := qeds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = qeds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) IntX(ctx context.Context) int {
	v, err := qeds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(qeds.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := qeds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) Float64sX(ctx context.Context) []float64 {
	v, err := qeds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = qeds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) Float64X(ctx context.Context) float64 {
	v, err := qeds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(qeds.fields) > 1 {
		return nil, errors.New("ent: QccEnterpriseDataSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := qeds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) BoolsX(ctx context.Context) []bool {
	v, err := qeds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (qeds *QccEnterpriseDataSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = qeds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{qccenterprisedata.Label}
	default:
		err = fmt.Errorf("ent: QccEnterpriseDataSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (qeds *QccEnterpriseDataSelect) BoolX(ctx context.Context) bool {
	v, err := qeds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qeds *QccEnterpriseDataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := qeds.sqlQuery().Query()
	if err := qeds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qeds *QccEnterpriseDataSelect) sqlQuery() sql.Querier {
	selector := qeds.sql
	selector.Select(selector.Columns(qeds.fields...)...)
	return selector
}