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
	"github.com/sthorer/api/ent/file"
	"github.com/sthorer/api/ent/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetHash sets the hash field.
func (fc *FileCreate) SetHash(s string) *FileCreate {
	fc.mutation.SetHash(s)
	return fc
}

// SetSize sets the size field.
func (fc *FileCreate) SetSize(i int64) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetPinnedAt sets the pinned_at field.
func (fc *FileCreate) SetPinnedAt(t time.Time) *FileCreate {
	fc.mutation.SetPinnedAt(t)
	return fc
}

// SetNillablePinnedAt sets the pinned_at field if the given value is not nil.
func (fc *FileCreate) SetNillablePinnedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetPinnedAt(*t)
	}
	return fc
}

// SetUnpinnedAt sets the unpinned_at field.
func (fc *FileCreate) SetUnpinnedAt(t time.Time) *FileCreate {
	fc.mutation.SetUnpinnedAt(t)
	return fc
}

// SetMetadata sets the metadata field.
func (fc *FileCreate) SetMetadata(m map[string]interface{}) *FileCreate {
	fc.mutation.SetMetadata(m)
	return fc
}

// SetID sets the id field.
func (fc *FileCreate) SetID(u uuid.UUID) *FileCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetUserID sets the user edge to User by id.
func (fc *FileCreate) SetUserID(id int) *FileCreate {
	fc.mutation.SetUserID(id)
	return fc
}

// SetUser sets the user edge to User.
func (fc *FileCreate) SetUser(u *User) *FileCreate {
	return fc.SetUserID(u.ID)
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	if _, ok := fc.mutation.Hash(); !ok {
		return nil, errors.New("ent: missing required field \"hash\"")
	}
	if v, ok := fc.mutation.Hash(); ok {
		if err := file.HashValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"hash\": %v", err)
		}
	}
	if _, ok := fc.mutation.Size(); !ok {
		return nil, errors.New("ent: missing required field \"size\"")
	}
	if v, ok := fc.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}
	if _, ok := fc.mutation.PinnedAt(); !ok {
		v := file.DefaultPinnedAt()
		fc.mutation.SetPinnedAt(v)
	}
	if _, ok := fc.mutation.UnpinnedAt(); !ok {
		return nil, errors.New("ent: missing required field \"unpinned_at\"")
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return nil, errors.New("ent: missing required edge \"user\"")
	}
	var (
		err  error
		node *File
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	var (
		f     = &File{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: file.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: file.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		f.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldHash,
		})
		f.Hash = value
	}
	if value, ok := fc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSize,
		})
		f.Size = value
	}
	if value, ok := fc.mutation.PinnedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldPinnedAt,
		})
		f.PinnedAt = value
	}
	if value, ok := fc.mutation.UnpinnedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldUnpinnedAt,
		})
		f.UnpinnedAt = value
	}
	if value, ok := fc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: file.FieldMetadata,
		})
		f.Metadata = value
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
