// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-manage/pkg/ent/personal"
	"go-manage/pkg/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PersonalCreate is the builder for creating a Personal entity.
type PersonalCreate struct {
	config
	mutation *PersonalMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PersonalCreate) SetName(s string) *PersonalCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetPhone sets the "phone" field.
func (pc *PersonalCreate) SetPhone(s string) *PersonalCreate {
	pc.mutation.SetPhone(s)
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *PersonalCreate) SetUserID(i int) *PersonalCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pc *PersonalCreate) SetNillableUserID(i *int) *PersonalCreate {
	if i != nil {
		pc.SetUserID(*i)
	}
	return pc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (pc *PersonalCreate) SetUsersID(id int) *PersonalCreate {
	pc.mutation.SetUsersID(id)
	return pc
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (pc *PersonalCreate) SetNillableUsersID(id *int) *PersonalCreate {
	if id != nil {
		pc = pc.SetUsersID(*id)
	}
	return pc
}

// SetUsers sets the "users" edge to the User entity.
func (pc *PersonalCreate) SetUsers(u *User) *PersonalCreate {
	return pc.SetUsersID(u.ID)
}

// Mutation returns the PersonalMutation object of the builder.
func (pc *PersonalCreate) Mutation() *PersonalMutation {
	return pc.mutation
}

// Save creates the Personal in the database.
func (pc *PersonalCreate) Save(ctx context.Context) (*Personal, error) {
	return withHooks[*Personal, PersonalMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonalCreate) SaveX(ctx context.Context) *Personal {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PersonalCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PersonalCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonalCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Personal.name"`)}
	}
	if _, ok := pc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Personal.phone"`)}
	}
	return nil
}

func (pc *PersonalCreate) sqlSave(ctx context.Context) (*Personal, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PersonalCreate) createSpec() (*Personal, *sqlgraph.CreateSpec) {
	var (
		_node = &Personal{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(personal.Table, sqlgraph.NewFieldSpec(personal.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(personal.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Phone(); ok {
		_spec.SetField(personal.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if nodes := pc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personal.UsersTable,
			Columns: []string{personal.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PersonalCreateBulk is the builder for creating many Personal entities in bulk.
type PersonalCreateBulk struct {
	config
	builders []*PersonalCreate
}

// Save creates the Personal entities in the database.
func (pcb *PersonalCreateBulk) Save(ctx context.Context) ([]*Personal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Personal, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonalMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonalCreateBulk) SaveX(ctx context.Context) []*Personal {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PersonalCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PersonalCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
