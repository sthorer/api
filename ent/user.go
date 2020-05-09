// github.com/sthorer/api

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sthorer/api/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Plan holds the value of the "plan" field.
	Plan user.Plan `json:"plan,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Tokens holds the value of the tokens edge.
	Tokens []*Token
	// Files holds the value of the files edge.
	Files []*File
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TokensOrErr() ([]*Token, error) {
	if e.loadedTypes[0] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[1] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&sql.NullBool{},   // active
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // created_at
		&sql.NullString{}, // plan
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[0])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[1])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field active", values[2])
	} else if value.Valid {
		u.Active = value.Bool
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[3])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[4])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field plan", values[5])
	} else if value.Valid {
		u.Plan = user.Plan(value.String)
	}
	return nil
}

// QueryTokens queries the tokens edge of the User.
func (u *User) QueryTokens() *TokenQuery {
	return (&UserClient{config: u.config}).QueryTokens(u)
}

// QueryFiles queries the files edge of the User.
func (u *User) QueryFiles() *FileQuery {
	return (&UserClient{config: u.config}).QueryFiles(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", u.Active))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", plan=")
	builder.WriteString(fmt.Sprintf("%v", u.Plan))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
