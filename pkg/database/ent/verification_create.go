// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/user"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/verification"
)

// VerificationCreate is the builder for creating a Verification entity.
type VerificationCreate struct {
	config
	mutation *VerificationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vc *VerificationCreate) SetCreatedAt(t time.Time) *VerificationCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableCreatedAt(t *time.Time) *VerificationCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VerificationCreate) SetUpdatedAt(t time.Time) *VerificationCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableUpdatedAt(t *time.Time) *VerificationCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetClientIP sets the "client_ip" field.
func (vc *VerificationCreate) SetClientIP(s string) *VerificationCreate {
	vc.mutation.SetClientIP(s)
	return vc
}

// SetClientUserAgent sets the "clientUserAgent" field.
func (vc *VerificationCreate) SetClientUserAgent(s string) *VerificationCreate {
	vc.mutation.SetClientUserAgent(s)
	return vc
}

// SetLoginTime sets the "login_time" field.
func (vc *VerificationCreate) SetLoginTime(t time.Time) *VerificationCreate {
	vc.mutation.SetLoginTime(t)
	return vc
}

// SetNillableLoginTime sets the "login_time" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableLoginTime(t *time.Time) *VerificationCreate {
	if t != nil {
		vc.SetLoginTime(*t)
	}
	return vc
}

// SetLogoutTime sets the "logout_time" field.
func (vc *VerificationCreate) SetLogoutTime(t time.Time) *VerificationCreate {
	vc.mutation.SetLogoutTime(t)
	return vc
}

// SetNillableLogoutTime sets the "logout_time" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableLogoutTime(t *time.Time) *VerificationCreate {
	if t != nil {
		vc.SetLogoutTime(*t)
	}
	return vc
}

// SetRemember sets the "remember" field.
func (vc *VerificationCreate) SetRemember(b bool) *VerificationCreate {
	vc.mutation.SetRemember(b)
	return vc
}

// SetID sets the "id" field.
func (vc *VerificationCreate) SetID(s string) *VerificationCreate {
	vc.mutation.SetID(s)
	return vc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (vc *VerificationCreate) AddUserIDs(ids ...string) *VerificationCreate {
	vc.mutation.AddUserIDs(ids...)
	return vc
}

// AddUsers adds the "users" edges to the User entity.
func (vc *VerificationCreate) AddUsers(u ...*User) *VerificationCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return vc.AddUserIDs(ids...)
}

// Mutation returns the VerificationMutation object of the builder.
func (vc *VerificationCreate) Mutation() *VerificationMutation {
	return vc.mutation
}

// Save creates the Verification in the database.
func (vc *VerificationCreate) Save(ctx context.Context) (*Verification, error) {
	var (
		err  error
		node *Verification
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VerificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VerificationCreate) SaveX(ctx context.Context) *Verification {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VerificationCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VerificationCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VerificationCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := verification.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := verification.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vc.mutation.LoginTime(); !ok {
		v := verification.DefaultLoginTime()
		vc.mutation.SetLoginTime(v)
	}
	if _, ok := vc.mutation.LogoutTime(); !ok {
		v := verification.DefaultLogoutTime()
		vc.mutation.SetLogoutTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VerificationCreate) check() error {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := vc.mutation.ClientIP(); !ok {
		return &ValidationError{Name: "client_ip", err: errors.New(`ent: missing required field "client_ip"`)}
	}
	if _, ok := vc.mutation.ClientUserAgent(); !ok {
		return &ValidationError{Name: "clientUserAgent", err: errors.New(`ent: missing required field "clientUserAgent"`)}
	}
	if _, ok := vc.mutation.LoginTime(); !ok {
		return &ValidationError{Name: "login_time", err: errors.New(`ent: missing required field "login_time"`)}
	}
	if _, ok := vc.mutation.LogoutTime(); !ok {
		return &ValidationError{Name: "logout_time", err: errors.New(`ent: missing required field "logout_time"`)}
	}
	if _, ok := vc.mutation.Remember(); !ok {
		return &ValidationError{Name: "remember", err: errors.New(`ent: missing required field "remember"`)}
	}
	return nil
}

func (vc *VerificationCreate) sqlSave(ctx context.Context) (*Verification, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(string)
	}
	return _node, nil
}

func (vc *VerificationCreate) createSpec() (*Verification, *sqlgraph.CreateSpec) {
	var (
		_node = &Verification{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: verification.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: verification.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: verification.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: verification.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := vc.mutation.ClientIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: verification.FieldClientIP,
		})
		_node.ClientIP = value
	}
	if value, ok := vc.mutation.ClientUserAgent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: verification.FieldClientUserAgent,
		})
		_node.ClientUserAgent = value
	}
	if value, ok := vc.mutation.LoginTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: verification.FieldLoginTime,
		})
		_node.LoginTime = value
	}
	if value, ok := vc.mutation.LogoutTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: verification.FieldLogoutTime,
		})
		_node.LogoutTime = value
	}
	if value, ok := vc.mutation.Remember(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: verification.FieldRemember,
		})
		_node.Remember = value
	}
	if nodes := vc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   verification.UsersTable,
			Columns: []string{verification.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VerificationCreateBulk is the builder for creating many Verification entities in bulk.
type VerificationCreateBulk struct {
	config
	builders []*VerificationCreate
}

// Save creates the Verification entities in the database.
func (vcb *VerificationCreateBulk) Save(ctx context.Context) ([]*Verification, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Verification, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VerificationMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VerificationCreateBulk) SaveX(ctx context.Context) []*Verification {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VerificationCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VerificationCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
