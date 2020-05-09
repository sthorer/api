// github.com/sthorer/api

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
}

// SetName sets the name field.
func (tc *TokenCreate) SetName(s string) *TokenCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetSecret sets the secret field.
func (tc *TokenCreate) SetSecret(s string) *TokenCreate {
	tc.mutation.SetSecret(s)
	return tc
}

// SetPermissions sets the permissions field.
func (tc *TokenCreate) SetPermissions(t token.Permissions) *TokenCreate {
	tc.mutation.SetPermissions(t)
	return tc
}

// SetNillablePermissions sets the permissions field if the given value is not nil.
func (tc *TokenCreate) SetNillablePermissions(t *token.Permissions) *TokenCreate {
	if t != nil {
		tc.SetPermissions(*t)
	}
	return tc
}

// SetCreatedAt sets the created_at field.
func (tc *TokenCreate) SetCreatedAt(t time.Time) *TokenCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (tc *TokenCreate) SetNillableCreatedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetLastUsed sets the last_used field.
func (tc *TokenCreate) SetLastUsed(t time.Time) *TokenCreate {
	tc.mutation.SetLastUsed(t)
	return tc
}

// SetNillableLastUsed sets the last_used field if the given value is not nil.
func (tc *TokenCreate) SetNillableLastUsed(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetLastUsed(*t)
	}
	return tc
}

// SetID sets the id field.
func (tc *TokenCreate) SetID(u uuid.UUID) *TokenCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetUserID sets the user edge to User by id.
func (tc *TokenCreate) SetUserID(id int) *TokenCreate {
	tc.mutation.SetUserID(id)
	return tc
}

// SetUser sets the user edge to User.
func (tc *TokenCreate) SetUser(u *User) *TokenCreate {
	return tc.SetUserID(u.ID)
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	if _, ok := tc.mutation.Name(); !ok {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := token.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if _, ok := tc.mutation.Secret(); !ok {
		return nil, errors.New("ent: missing required field \"secret\"")
	}
	if v, ok := tc.mutation.Secret(); ok {
		if err := token.SecretValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"secret\": %v", err)
		}
	}
	if _, ok := tc.mutation.Permissions(); !ok {
		v := token.DefaultPermissions
		tc.mutation.SetPermissions(v)
	}
	if v, ok := tc.mutation.Permissions(); ok {
		if err := token.PermissionsValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"permissions\": %v", err)
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := token.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return nil, errors.New("ent: missing required edge \"user\"")
	}
	var (
		err  error
		node *Token
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	var (
		t     = &Token{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: token.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: token.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		t.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldName,
		})
		t.Name = value
	}
	if value, ok := tc.mutation.Secret(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldSecret,
		})
		t.Secret = value
	}
	if value, ok := tc.mutation.Permissions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: token.FieldPermissions,
		})
		t.Permissions = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldCreatedAt,
		})
		t.CreatedAt = value
	}
	if value, ok := tc.mutation.LastUsed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldLastUsed,
		})
		t.LastUsed = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.UserTable,
			Columns: []string{token.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
