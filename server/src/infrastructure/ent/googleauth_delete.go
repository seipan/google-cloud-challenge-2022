// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/googleauth"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/predicate"
)

// GoogleAuthDelete is the builder for deleting a GoogleAuth entity.
type GoogleAuthDelete struct {
	config
	hooks    []Hook
	mutation *GoogleAuthMutation
}

// Where appends a list predicates to the GoogleAuthDelete builder.
func (gad *GoogleAuthDelete) Where(ps ...predicate.GoogleAuth) *GoogleAuthDelete {
	gad.mutation.Where(ps...)
	return gad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gad *GoogleAuthDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GoogleAuthMutation](ctx, gad.sqlExec, gad.mutation, gad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gad *GoogleAuthDelete) ExecX(ctx context.Context) int {
	n, err := gad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gad *GoogleAuthDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: googleauth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: googleauth.FieldID,
			},
		},
	}
	if ps := gad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gad.mutation.done = true
	return affected, err
}

// GoogleAuthDeleteOne is the builder for deleting a single GoogleAuth entity.
type GoogleAuthDeleteOne struct {
	gad *GoogleAuthDelete
}

// Exec executes the deletion query.
func (gado *GoogleAuthDeleteOne) Exec(ctx context.Context) error {
	n, err := gado.gad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{googleauth.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gado *GoogleAuthDeleteOne) ExecX(ctx context.Context) {
	gado.gad.ExecX(ctx)
}
