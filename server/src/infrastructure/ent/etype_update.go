// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/etype"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/event"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/predicate"
)

// ETypeUpdate is the builder for updating EType entities.
type ETypeUpdate struct {
	config
	hooks    []Hook
	mutation *ETypeMutation
}

// Where appends a list predicates to the ETypeUpdate builder.
func (eu *ETypeUpdate) Where(ps ...predicate.EType) *ETypeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *ETypeUpdate) SetName(s string) *ETypeUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (eu *ETypeUpdate) SetEventID(id int) *ETypeUpdate {
	eu.mutation.SetEventID(id)
	return eu
}

// SetEvent sets the "event" edge to the Event entity.
func (eu *ETypeUpdate) SetEvent(e *Event) *ETypeUpdate {
	return eu.SetEventID(e.ID)
}

// Mutation returns the ETypeMutation object of the builder.
func (eu *ETypeUpdate) Mutation() *ETypeMutation {
	return eu.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (eu *ETypeUpdate) ClearEvent() *ETypeUpdate {
	eu.mutation.ClearEvent()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *ETypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ETypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ETypeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ETypeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ETypeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *ETypeUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := etype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EType.name": %w`, err)}
		}
	}
	if _, ok := eu.mutation.EventID(); eu.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EType.event"`)
	}
	return nil
}

func (eu *ETypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   etype.Table,
			Columns: etype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: etype.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(etype.FieldName, field.TypeString, value)
	}
	if eu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   etype.EventTable,
			Columns: []string{etype.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   etype.EventTable,
			Columns: []string{etype.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{etype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ETypeUpdateOne is the builder for updating a single EType entity.
type ETypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ETypeMutation
}

// SetName sets the "name" field.
func (euo *ETypeUpdateOne) SetName(s string) *ETypeUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (euo *ETypeUpdateOne) SetEventID(id int) *ETypeUpdateOne {
	euo.mutation.SetEventID(id)
	return euo
}

// SetEvent sets the "event" edge to the Event entity.
func (euo *ETypeUpdateOne) SetEvent(e *Event) *ETypeUpdateOne {
	return euo.SetEventID(e.ID)
}

// Mutation returns the ETypeMutation object of the builder.
func (euo *ETypeUpdateOne) Mutation() *ETypeMutation {
	return euo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (euo *ETypeUpdateOne) ClearEvent() *ETypeUpdateOne {
	euo.mutation.ClearEvent()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *ETypeUpdateOne) Select(field string, fields ...string) *ETypeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated EType entity.
func (euo *ETypeUpdateOne) Save(ctx context.Context) (*EType, error) {
	var (
		err  error
		node *EType
	)
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ETypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ETypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ETypeUpdateOne) SaveX(ctx context.Context) *EType {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *ETypeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ETypeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *ETypeUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := etype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EType.name": %w`, err)}
		}
	}
	if _, ok := euo.mutation.EventID(); euo.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EType.event"`)
	}
	return nil
}

func (euo *ETypeUpdateOne) sqlSave(ctx context.Context) (_node *EType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   etype.Table,
			Columns: etype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: etype.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, etype.FieldID)
		for _, f := range fields {
			if !etype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != etype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(etype.FieldName, field.TypeString, value)
	}
	if euo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   etype.EventTable,
			Columns: []string{etype.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   etype.EventTable,
			Columns: []string{etype.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EType{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{etype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
