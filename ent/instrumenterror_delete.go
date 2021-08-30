// Code generated by entc, DO NOT EDIT.

package ent

import (
	"backend/ent/instrumenterror"
	"backend/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InstrumentErrorDelete is the builder for deleting a InstrumentError entity.
type InstrumentErrorDelete struct {
	config
	hooks    []Hook
	mutation *InstrumentErrorMutation
}

// Where appends a list predicates to the InstrumentErrorDelete builder.
func (ied *InstrumentErrorDelete) Where(ps ...predicate.InstrumentError) *InstrumentErrorDelete {
	ied.mutation.Where(ps...)
	return ied
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ied *InstrumentErrorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ied.hooks) == 0 {
		affected, err = ied.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstrumentErrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ied.mutation = mutation
			affected, err = ied.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ied.hooks) - 1; i >= 0; i-- {
			if ied.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ied.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ied.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ied *InstrumentErrorDelete) ExecX(ctx context.Context) int {
	n, err := ied.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ied *InstrumentErrorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: instrumenterror.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instrumenterror.FieldID,
			},
		},
	}
	if ps := ied.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ied.driver, _spec)
}

// InstrumentErrorDeleteOne is the builder for deleting a single InstrumentError entity.
type InstrumentErrorDeleteOne struct {
	ied *InstrumentErrorDelete
}

// Exec executes the deletion query.
func (iedo *InstrumentErrorDeleteOne) Exec(ctx context.Context) error {
	n, err := iedo.ied.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{instrumenterror.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iedo *InstrumentErrorDeleteOne) ExecX(ctx context.Context) {
	iedo.ied.ExecX(ctx)
}
