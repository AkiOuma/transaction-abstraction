// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AkiOuma/transaction-abstraction/adapt-ent/datasource/ent/userrole"
)

// UserRoleCreate is the builder for creating a UserRole entity.
type UserRoleCreate struct {
	config
	mutation *UserRoleMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (urc *UserRoleCreate) SetUserID(i int) *UserRoleCreate {
	urc.mutation.SetUserID(i)
	return urc
}

// SetRoleID sets the "role_id" field.
func (urc *UserRoleCreate) SetRoleID(i int) *UserRoleCreate {
	urc.mutation.SetRoleID(i)
	return urc
}

// Mutation returns the UserRoleMutation object of the builder.
func (urc *UserRoleCreate) Mutation() *UserRoleMutation {
	return urc.mutation
}

// Save creates the UserRole in the database.
func (urc *UserRoleCreate) Save(ctx context.Context) (*UserRole, error) {
	var (
		err  error
		node *UserRole
	)
	if len(urc.hooks) == 0 {
		if err = urc.check(); err != nil {
			return nil, err
		}
		node, err = urc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = urc.check(); err != nil {
				return nil, err
			}
			urc.mutation = mutation
			if node, err = urc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(urc.hooks) - 1; i >= 0; i-- {
			if urc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = urc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, urc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRoleCreate) SaveX(ctx context.Context) *UserRole {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRoleCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRoleCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRoleCreate) check() error {
	if _, ok := urc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserRole.user_id"`)}
	}
	if _, ok := urc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "UserRole.role_id"`)}
	}
	return nil
}

func (urc *UserRoleCreate) sqlSave(ctx context.Context) (*UserRole, error) {
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (urc *UserRoleCreate) createSpec() (*UserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRole{config: urc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userrole.FieldID,
			},
		}
	)
	if value, ok := urc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := urc.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldRoleID,
		})
		_node.RoleID = value
	}
	return _node, _spec
}

// UserRoleCreateBulk is the builder for creating many UserRole entities in bulk.
type UserRoleCreateBulk struct {
	config
	builders []*UserRoleCreate
}

// Save creates the UserRole entities in the database.
func (urcb *UserRoleCreateBulk) Save(ctx context.Context) ([]*UserRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRole, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) SaveX(ctx context.Context) []*UserRole {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
