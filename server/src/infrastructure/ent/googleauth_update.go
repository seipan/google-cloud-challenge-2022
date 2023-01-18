// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/googleauth"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/predicate"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

// GoogleAuthUpdate is the builder for updating GoogleAuth entities.
type GoogleAuthUpdate struct {
	config
	hooks    []Hook
	mutation *GoogleAuthMutation
}

// Where appends a list predicates to the GoogleAuthUpdate builder.
func (gau *GoogleAuthUpdate) Where(ps ...predicate.GoogleAuth) *GoogleAuthUpdate {
	gau.mutation.Where(ps...)
	return gau
}

// SetUserID sets the "user_id" field.
func (gau *GoogleAuthUpdate) SetUserID(u uuid.UUID) *GoogleAuthUpdate {
	gau.mutation.SetUserID(u)
	return gau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (gau *GoogleAuthUpdate) SetNillableUserID(u *uuid.UUID) *GoogleAuthUpdate {
	if u != nil {
		gau.SetUserID(*u)
	}
	return gau
}

// SetAccessToken sets the "access_token" field.
func (gau *GoogleAuthUpdate) SetAccessToken(s string) *GoogleAuthUpdate {
	gau.mutation.SetAccessToken(s)
	return gau
}

// SetRefreshToken sets the "refresh_token" field.
func (gau *GoogleAuthUpdate) SetRefreshToken(s string) *GoogleAuthUpdate {
	gau.mutation.SetRefreshToken(s)
	return gau
}

// SetExpiry sets the "expiry" field.
func (gau *GoogleAuthUpdate) SetExpiry(t time.Time) *GoogleAuthUpdate {
	gau.mutation.SetExpiry(t)
	return gau
}

// SetUser sets the "user" edge to the User entity.
func (gau *GoogleAuthUpdate) SetUser(u *User) *GoogleAuthUpdate {
	return gau.SetUserID(u.ID)
}

// Mutation returns the GoogleAuthMutation object of the builder.
func (gau *GoogleAuthUpdate) Mutation() *GoogleAuthMutation {
	return gau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gau *GoogleAuthUpdate) ClearUser() *GoogleAuthUpdate {
	gau.mutation.ClearUser()
	return gau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gau *GoogleAuthUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GoogleAuthMutation](ctx, gau.sqlSave, gau.mutation, gau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gau *GoogleAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := gau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gau *GoogleAuthUpdate) Exec(ctx context.Context) error {
	_, err := gau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gau *GoogleAuthUpdate) ExecX(ctx context.Context) {
	if err := gau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gau *GoogleAuthUpdate) check() error {
	if v, ok := gau.mutation.AccessToken(); ok {
		if err := googleauth.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "GoogleAuth.access_token": %w`, err)}
		}
	}
	if v, ok := gau.mutation.RefreshToken(); ok {
		if err := googleauth.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "GoogleAuth.refresh_token": %w`, err)}
		}
	}
	if _, ok := gau.mutation.UserID(); gau.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GoogleAuth.user"`)
	}
	return nil
}

func (gau *GoogleAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gau.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   googleauth.Table,
			Columns: googleauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: googleauth.FieldID,
			},
		},
	}
	if ps := gau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gau.mutation.AccessToken(); ok {
		_spec.SetField(googleauth.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := gau.mutation.RefreshToken(); ok {
		_spec.SetField(googleauth.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := gau.mutation.Expiry(); ok {
		_spec.SetField(googleauth.FieldExpiry, field.TypeTime, value)
	}
	if gau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   googleauth.UserTable,
			Columns: []string{googleauth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   googleauth.UserTable,
			Columns: []string{googleauth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{googleauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gau.mutation.done = true
	return n, nil
}

// GoogleAuthUpdateOne is the builder for updating a single GoogleAuth entity.
type GoogleAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoogleAuthMutation
}

// SetUserID sets the "user_id" field.
func (gauo *GoogleAuthUpdateOne) SetUserID(u uuid.UUID) *GoogleAuthUpdateOne {
	gauo.mutation.SetUserID(u)
	return gauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (gauo *GoogleAuthUpdateOne) SetNillableUserID(u *uuid.UUID) *GoogleAuthUpdateOne {
	if u != nil {
		gauo.SetUserID(*u)
	}
	return gauo
}

// SetAccessToken sets the "access_token" field.
func (gauo *GoogleAuthUpdateOne) SetAccessToken(s string) *GoogleAuthUpdateOne {
	gauo.mutation.SetAccessToken(s)
	return gauo
}

// SetRefreshToken sets the "refresh_token" field.
func (gauo *GoogleAuthUpdateOne) SetRefreshToken(s string) *GoogleAuthUpdateOne {
	gauo.mutation.SetRefreshToken(s)
	return gauo
}

// SetExpiry sets the "expiry" field.
func (gauo *GoogleAuthUpdateOne) SetExpiry(t time.Time) *GoogleAuthUpdateOne {
	gauo.mutation.SetExpiry(t)
	return gauo
}

// SetUser sets the "user" edge to the User entity.
func (gauo *GoogleAuthUpdateOne) SetUser(u *User) *GoogleAuthUpdateOne {
	return gauo.SetUserID(u.ID)
}

// Mutation returns the GoogleAuthMutation object of the builder.
func (gauo *GoogleAuthUpdateOne) Mutation() *GoogleAuthMutation {
	return gauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gauo *GoogleAuthUpdateOne) ClearUser() *GoogleAuthUpdateOne {
	gauo.mutation.ClearUser()
	return gauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gauo *GoogleAuthUpdateOne) Select(field string, fields ...string) *GoogleAuthUpdateOne {
	gauo.fields = append([]string{field}, fields...)
	return gauo
}

// Save executes the query and returns the updated GoogleAuth entity.
func (gauo *GoogleAuthUpdateOne) Save(ctx context.Context) (*GoogleAuth, error) {
	return withHooks[*GoogleAuth, GoogleAuthMutation](ctx, gauo.sqlSave, gauo.mutation, gauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gauo *GoogleAuthUpdateOne) SaveX(ctx context.Context) *GoogleAuth {
	node, err := gauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gauo *GoogleAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := gauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gauo *GoogleAuthUpdateOne) ExecX(ctx context.Context) {
	if err := gauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gauo *GoogleAuthUpdateOne) check() error {
	if v, ok := gauo.mutation.AccessToken(); ok {
		if err := googleauth.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "GoogleAuth.access_token": %w`, err)}
		}
	}
	if v, ok := gauo.mutation.RefreshToken(); ok {
		if err := googleauth.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "GoogleAuth.refresh_token": %w`, err)}
		}
	}
	if _, ok := gauo.mutation.UserID(); gauo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GoogleAuth.user"`)
	}
	return nil
}

func (gauo *GoogleAuthUpdateOne) sqlSave(ctx context.Context) (_node *GoogleAuth, err error) {
	if err := gauo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   googleauth.Table,
			Columns: googleauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: googleauth.FieldID,
			},
		},
	}
	id, ok := gauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoogleAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, googleauth.FieldID)
		for _, f := range fields {
			if !googleauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != googleauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gauo.mutation.AccessToken(); ok {
		_spec.SetField(googleauth.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := gauo.mutation.RefreshToken(); ok {
		_spec.SetField(googleauth.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := gauo.mutation.Expiry(); ok {
		_spec.SetField(googleauth.FieldExpiry, field.TypeTime, value)
	}
	if gauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   googleauth.UserTable,
			Columns: []string{googleauth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   googleauth.UserTable,
			Columns: []string{googleauth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GoogleAuth{config: gauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{googleauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gauo.mutation.done = true
	return _node, nil
}
