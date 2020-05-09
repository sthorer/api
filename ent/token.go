// github.com/sthorer/api

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"
)

// Token is the model entity for the Token schema.
type Token struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"-"`
	// Permissions holds the value of the "permissions" field.
	Permissions token.Permissions `json:"permissions,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LastUsed holds the value of the "last_used" field.
	LastUsed time.Time `json:"last_used,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TokenQuery when eager-loading is set.
	Edges       TokenEdges `json:"edges"`
	user_tokens *int
}

// TokenEdges holds the relations/edges for other nodes in the graph.
type TokenEdges struct {
	// User holds the value of the user edge.
	User *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TokenEdges) UserOrErr() (*User, error) {
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
func (*Token) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // name
		&sql.NullString{}, // secret
		&sql.NullString{}, // permissions
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // last_used
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Token) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_tokens
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Token fields.
func (t *Token) assignValues(values ...interface{}) error {
	if m, n := len(values), len(token.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		t.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		t.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field secret", values[1])
	} else if value.Valid {
		t.Secret = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field permissions", values[2])
	} else if value.Valid {
		t.Permissions = token.Permissions(value.String)
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		t.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field last_used", values[4])
	} else if value.Valid {
		t.LastUsed = value.Time
	}
	values = values[5:]
	if len(values) == len(token.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_tokens", value)
		} else if value.Valid {
			t.user_tokens = new(int)
			*t.user_tokens = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the Token.
func (t *Token) QueryUser() *UserQuery {
	return (&TokenClient{config: t.config}).QueryUser(t)
}

// Update returns a builder for updating this Token.
// Note that, you need to call Token.Unwrap() before calling this method, if this Token
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Token) Update() *TokenUpdateOne {
	return (&TokenClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Token) Unwrap() *Token {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Token is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Token) String() string {
	var builder strings.Builder
	builder.WriteString("Token(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", secret=<sensitive>")
	builder.WriteString(", permissions=")
	builder.WriteString(fmt.Sprintf("%v", t.Permissions))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", last_used=")
	builder.WriteString(t.LastUsed.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tokens is a parsable slice of Token.
type Tokens []*Token

func (t Tokens) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
