// github.com/sthorer/api

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sthorer/api/ent/file"
	"github.com/sthorer/api/ent/user"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// PinnedAt holds the value of the "pinned_at" field.
	PinnedAt time.Time `json:"pinned_at,omitempty"`
	// UnpinnedAt holds the value of the "unpinned_at" field.
	UnpinnedAt time.Time `json:"unpinned_at,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges      FileEdges `json:"edges"`
	user_files *int
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// User holds the value of the user edge.
	User *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // hash
		&sql.NullInt64{},  // size
		&sql.NullTime{},   // pinned_at
		&sql.NullTime{},   // unpinned_at
		&[]byte{},         // metadata
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*File) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_files
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(values ...interface{}) error {
	if m, n := len(values), len(file.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		f.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field hash", values[0])
	} else if value.Valid {
		f.Hash = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field size", values[1])
	} else if value.Valid {
		f.Size = value.Int64
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field pinned_at", values[2])
	} else if value.Valid {
		f.PinnedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field unpinned_at", values[3])
	} else if value.Valid {
		f.UnpinnedAt = value.Time
	}

	if value, ok := values[4].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[4])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &f.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	values = values[5:]
	if len(values) == len(file.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_files", value)
		} else if value.Valid {
			f.user_files = new(int)
			*f.user_files = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the File.
func (f *File) QueryUser() *UserQuery {
	return (&FileClient{config: f.config}).QueryUser(f)
}

// Update returns a builder for updating this File.
// Note that, you need to call File.Unwrap() before calling this method, if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", hash=")
	builder.WriteString(f.Hash)
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", pinned_at=")
	builder.WriteString(f.PinnedAt.Format(time.ANSIC))
	builder.WriteString(", unpinned_at=")
	builder.WriteString(f.UnpinnedAt.Format(time.ANSIC))
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", f.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
