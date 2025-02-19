// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/job"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
)

// JobDelete is the builder for deleting a Job entity.
type JobDelete struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// Where appends a list predicates to the JobDelete builder.
func (jd *JobDelete) Where(ps ...predicate.Job) *JobDelete {
	jd.mutation.Where(ps...)
	return jd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jd *JobDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(jd.hooks) == 0 {
		affected, err = jd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jd.mutation = mutation
			affected, err = jd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(jd.hooks) - 1; i >= 0; i-- {
			if jd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (jd *JobDelete) ExecX(ctx context.Context) int {
	n, err := jd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jd *JobDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: job.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: job.FieldID,
			},
		},
	}
	if ps := jd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, jd.driver, _spec)
}

// JobDeleteOne is the builder for deleting a single Job entity.
type JobDeleteOne struct {
	jd *JobDelete
}

// Exec executes the deletion query.
func (jdo *JobDeleteOne) Exec(ctx context.Context) error {
	n, err := jdo.jd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{job.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jdo *JobDeleteOne) ExecX(ctx context.Context) {
	jdo.jd.ExecX(ctx)
}
