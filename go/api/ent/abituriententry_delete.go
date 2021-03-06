// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"urfu-abiturient-api/ent/abituriententry"
	"urfu-abiturient-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AbiturientEntryDelete is the builder for deleting a AbiturientEntry entity.
type AbiturientEntryDelete struct {
	config
	hooks    []Hook
	mutation *AbiturientEntryMutation
}

// Where adds a new predicate to the AbiturientEntryDelete builder.
func (aed *AbiturientEntryDelete) Where(ps ...predicate.AbiturientEntry) *AbiturientEntryDelete {
	aed.mutation.predicates = append(aed.mutation.predicates, ps...)
	return aed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aed *AbiturientEntryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aed.hooks) == 0 {
		affected, err = aed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AbiturientEntryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aed.mutation = mutation
			affected, err = aed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aed.hooks) - 1; i >= 0; i-- {
			mut = aed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aed *AbiturientEntryDelete) ExecX(ctx context.Context) int {
	n, err := aed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aed *AbiturientEntryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: abituriententry.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: abituriententry.FieldID,
			},
		},
	}
	if ps := aed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aed.driver, _spec)
}

// AbiturientEntryDeleteOne is the builder for deleting a single AbiturientEntry entity.
type AbiturientEntryDeleteOne struct {
	aed *AbiturientEntryDelete
}

// Exec executes the deletion query.
func (aedo *AbiturientEntryDeleteOne) Exec(ctx context.Context) error {
	n, err := aedo.aed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{abituriententry.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aedo *AbiturientEntryDeleteOne) ExecX(ctx context.Context) {
	aedo.aed.ExecX(ctx)
}
