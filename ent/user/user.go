// github.com/sthorer/api

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID        = "id"         // FieldEmail holds the string denoting the email vertex property in the database.
	FieldEmail     = "email"      // FieldPassword holds the string denoting the password vertex property in the database.
	FieldPassword  = "password"   // FieldActive holds the string denoting the active vertex property in the database.
	FieldActive    = "active"     // FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at" // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at" // FieldPlan holds the string denoting the plan vertex property in the database.
	FieldPlan      = "plan"

	// EdgeTokens holds the string denoting the tokens edge name in mutations.
	EdgeTokens = "tokens"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"

	// Table holds the table name of the user in the database.
	Table = "users"
	// TokensTable is the table the holds the tokens relation/edge.
	TokensTable = "tokens"
	// TokensInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokensInverseTable = "tokens"
	// TokensColumn is the table column denoting the tokens relation/edge.
	TokensColumn = "user_tokens"
	// FilesTable is the table the holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "user_files"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldActive,
	FieldUpdatedAt,
	FieldCreatedAt,
	FieldPlan,
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultActive holds the default value on creation for the active field.
	DefaultActive bool
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
)

// Plan defines the type for the plan enum field.
type Plan string

// PlanFree is the default Plan.
const DefaultPlan = PlanFree

// Plan values.
const (
	PlanFree    Plan = "Free"
	PlanPremium Plan = "Premium"
)

func (s Plan) String() string {
	return string(s)
}

// PlanValidator is a validator for the "pl" field enum values. It is called by the builders before save.
func PlanValidator(pl Plan) error {
	switch pl {
	case PlanFree, PlanPremium:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for plan field: %q", pl)
	}
}
