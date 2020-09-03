// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/zhanghongnian/brand-index/ent/predicate"
	"github.com/zhanghongnian/brand-index/ent/qccenterprisedata"
)

// QccEnterpriseDataDelete is the builder for deleting a QccEnterpriseData entity.
type QccEnterpriseDataDelete struct {
	config
	hooks      []Hook
	mutation   *QccEnterpriseDataMutation
	predicates []predicate.QccEnterpriseData
}

// Where adds a new predicate to the delete builder.
func (qedd *QccEnterpriseDataDelete) Where(ps ...predicate.QccEnterpriseData) *QccEnterpriseDataDelete {
	qedd.predicates = append(qedd.predicates, ps...)
	return qedd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qedd *QccEnterpriseDataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qedd.hooks) == 0 {
		affected, err = qedd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QccEnterpriseDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qedd.mutation = mutation
			affected, err = qedd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qedd.hooks) - 1; i >= 0; i-- {
			mut = qedd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qedd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (qedd *QccEnterpriseDataDelete) ExecX(ctx context.Context) int {
	n, err := qedd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qedd *QccEnterpriseDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: qccenterprisedata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: qccenterprisedata.FieldID,
			},
		},
	}
	if ps := qedd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, qedd.driver, _spec)
}

// QccEnterpriseDataDeleteOne is the builder for deleting a single QccEnterpriseData entity.
type QccEnterpriseDataDeleteOne struct {
	qedd *QccEnterpriseDataDelete
}

// Exec executes the deletion query.
func (qeddo *QccEnterpriseDataDeleteOne) Exec(ctx context.Context) error {
	n, err := qeddo.qedd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{qccenterprisedata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qeddo *QccEnterpriseDataDeleteOne) ExecX(ctx context.Context) {
	qeddo.qedd.ExecX(ctx)
}
