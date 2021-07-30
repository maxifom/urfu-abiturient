// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"urfu-abiturient-api/ent/lastupdated"
	"urfu-abiturient-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LastUpdatedDelete is the builder for deleting a LastUpdated entity.
type LastUpdatedDelete struct {
	config
	hooks    []Hook
	mutation *LastUpdatedMutation
}

// Where adds a new predicate to the LastUpdatedDelete builder.
func (lud *LastUpdatedDelete) Where(ps ...predicate.LastUpdated) *LastUpdatedDelete {
	lud.mutation.predicates = append(lud.mutation.predicates, ps...)
	return lud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lud *LastUpdatedDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lud.hooks) == 0 {
		affected, err = lud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LastUpdatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lud.mutation = mutation
			affected, err = lud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lud.hooks) - 1; i >= 0; i-- {
			mut = lud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lud *LastUpdatedDelete) ExecX(ctx context.Context) int {
	n, err := lud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lud *LastUpdatedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: lastupdated.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lastupdated.FieldID,
			},
		},
	}
	if ps := lud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, lud.driver, _spec)
}

// LastUpdatedDeleteOne is the builder for deleting a single LastUpdated entity.
type LastUpdatedDeleteOne struct {
	lud *LastUpdatedDelete
}

// Exec executes the deletion query.
func (ludo *LastUpdatedDeleteOne) Exec(ctx context.Context) error {
	n, err := ludo.lud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lastupdated.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ludo *LastUpdatedDeleteOne) ExecX(ctx context.Context) {
	ludo.lud.ExecX(ctx)
}
