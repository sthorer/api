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
	"github.com/sthorer/api/ent/file"
	"github.com/sthorer/api/ent/predicate"
	"github.com/sthorer/api/ent/user"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks      []Hook
	mutation   *FileMutation
	predicates []predicate.File
}

// Where adds a new predicate for the builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetSize sets the size field.
func (fu *FileUpdate) SetSize(i int64) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// AddSize adds i to size.
func (fu *FileUpdate) AddSize(i int64) *FileUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// SetUnpinnedAt sets the unpinned_at field.
func (fu *FileUpdate) SetUnpinnedAt(t time.Time) *FileUpdate {
	fu.mutation.SetUnpinnedAt(t)
	return fu
}

// SetMetadata sets the metadata field.
func (fu *FileUpdate) SetMetadata(m map[string]interface{}) *FileUpdate {
	fu.mutation.SetMetadata(m)
	return fu
}

// ClearMetadata clears the value of metadata.
func (fu *FileUpdate) ClearMetadata() *FileUpdate {
	fu.mutation.ClearMetadata()
	return fu
}

// SetUserID sets the user edge to User by id.
func (fu *FileUpdate) SetUserID(id int) *FileUpdate {
	fu.mutation.SetUserID(id)
	return fu
}

// SetUser sets the user edge to User.
func (fu *FileUpdate) SetUser(u *User) *FileUpdate {
	return fu.SetUserID(u.ID)
}

// ClearUser clears the user edge to User.
func (fu *FileUpdate) ClearUser() *FileUpdate {
	fu.mutation.ClearUser()
	return fu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := fu.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}

	if _, ok := fu.mutation.UserID(); fu.mutation.UserCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: file.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fu.mutation.UnpinnedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldUnpinnedAt,
		})
	}
	if value, ok := fu.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: file.FieldMetadata,
		})
	}
	if fu.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: file.FieldMetadata,
		})
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
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
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// SetSize sets the size field.
func (fuo *FileUpdateOne) SetSize(i int64) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// AddSize adds i to size.
func (fuo *FileUpdateOne) AddSize(i int64) *FileUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// SetUnpinnedAt sets the unpinned_at field.
func (fuo *FileUpdateOne) SetUnpinnedAt(t time.Time) *FileUpdateOne {
	fuo.mutation.SetUnpinnedAt(t)
	return fuo
}

// SetMetadata sets the metadata field.
func (fuo *FileUpdateOne) SetMetadata(m map[string]interface{}) *FileUpdateOne {
	fuo.mutation.SetMetadata(m)
	return fuo
}

// ClearMetadata clears the value of metadata.
func (fuo *FileUpdateOne) ClearMetadata() *FileUpdateOne {
	fuo.mutation.ClearMetadata()
	return fuo
}

// SetUserID sets the user edge to User by id.
func (fuo *FileUpdateOne) SetUserID(id int) *FileUpdateOne {
	fuo.mutation.SetUserID(id)
	return fuo
}

// SetUser sets the user edge to User.
func (fuo *FileUpdateOne) SetUser(u *User) *FileUpdateOne {
	return fuo.SetUserID(u.ID)
}

// ClearUser clears the user edge to User.
func (fuo *FileUpdateOne) ClearUser() *FileUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// Save executes the query and returns the updated entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	if v, ok := fuo.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}

	if _, ok := fuo.mutation.UserID(); fuo.mutation.UserCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err  error
		node *File
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (f *File, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: file.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing File.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fuo.mutation.UnpinnedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldUnpinnedAt,
		})
	}
	if value, ok := fuo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: file.FieldMetadata,
		})
	}
	if fuo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: file.FieldMetadata,
		})
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
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
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
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
	f = &File{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
