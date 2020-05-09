// github.com/sthorer/api

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sthorer/api/ent/predicate"
	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks      []Hook
	mutation   *TokenMutation
	predicates []predicate.Token
}

// Where adds a new predicate for the builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetName sets the name field.
func (tu *TokenUpdate) SetName(s string) *TokenUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetSecret sets the secret field.
func (tu *TokenUpdate) SetSecret(s string) *TokenUpdate {
	tu.mutation.SetSecret(s)
	return tu
}

// SetLastUsed sets the last_used field.
func (tu *TokenUpdate) SetLastUsed(t time.Time) *TokenUpdate {
	tu.mutation.SetLastUsed(t)
	return tu
}

// SetNillableLastUsed sets the last_used field if the given value is not nil.
func (tu *TokenUpdate) SetNillableLastUsed(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetLastUsed(*t)
	}
	return tu
}

// ClearLastUsed clears the value of last_used.
func (tu *TokenUpdate) ClearLastUsed() *TokenUpdate {
	tu.mutation.ClearLastUsed()
	return tu
}

// SetUserID sets the user edge to User by id.
func (tu *TokenUpdate) SetUserID(id int) *TokenUpdate {
	tu.mutation.SetUserID(id)
	return tu
}

// SetUser sets the user edge to User.
func (tu *TokenUpdate) SetUser(u *User) *TokenUpdate {
	return tu.SetUserID(u.ID)
}

// ClearUser clears the user edge to User.
func (tu *TokenUpdate) ClearUser() *TokenUpdate {
	tu.mutation.ClearUser()
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tu.mutation.Name(); ok {
		if err := token.NameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if v, ok := tu.mutation.Secret(); ok {
		if err := token.SecretValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"secret\": %v", err)
		}
	}

	if _, ok := tu.mutation.UserID(); tu.mutation.UserCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: token.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldName,
		})
	}
	if value, ok := tu.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldSecret,
		})
	}
	if value, ok := tu.mutation.LastUsed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldLastUsed,
		})
	}
	if tu.mutation.LastUsedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: token.FieldLastUsed,
		})
	}
	if tu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	hooks    []Hook
	mutation *TokenMutation
}

// SetName sets the name field.
func (tuo *TokenUpdateOne) SetName(s string) *TokenUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetSecret sets the secret field.
func (tuo *TokenUpdateOne) SetSecret(s string) *TokenUpdateOne {
	tuo.mutation.SetSecret(s)
	return tuo
}

// SetLastUsed sets the last_used field.
func (tuo *TokenUpdateOne) SetLastUsed(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetLastUsed(t)
	return tuo
}

// SetNillableLastUsed sets the last_used field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableLastUsed(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetLastUsed(*t)
	}
	return tuo
}

// ClearLastUsed clears the value of last_used.
func (tuo *TokenUpdateOne) ClearLastUsed() *TokenUpdateOne {
	tuo.mutation.ClearLastUsed()
	return tuo
}

// SetUserID sets the user edge to User by id.
func (tuo *TokenUpdateOne) SetUserID(id int) *TokenUpdateOne {
	tuo.mutation.SetUserID(id)
	return tuo
}

// SetUser sets the user edge to User.
func (tuo *TokenUpdateOne) SetUser(u *User) *TokenUpdateOne {
	return tuo.SetUserID(u.ID)
}

// ClearUser clears the user edge to User.
func (tuo *TokenUpdateOne) ClearUser() *TokenUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	if v, ok := tuo.mutation.Name(); ok {
		if err := token.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if v, ok := tuo.mutation.Secret(); ok {
		if err := token.SecretValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"secret\": %v", err)
		}
	}

	if _, ok := tuo.mutation.UserID(); tuo.mutation.UserCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err  error
		node *Token
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (t *Token, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: token.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Token.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldName,
		})
	}
	if value, ok := tuo.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldSecret,
		})
	}
	if value, ok := tuo.mutation.LastUsed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldLastUsed,
		})
	}
	if tuo.mutation.LastUsedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: token.FieldLastUsed,
		})
	}
	if tuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Token{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
