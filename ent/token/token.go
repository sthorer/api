// github.com/sthorer/api

package token

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the token type in the database.
	Label = "token"
	// FieldID holds the string denoting the id field in the database.
	FieldID          = "id"          // FieldName holds the string denoting the name vertex property in the database.
	FieldName        = "name"        // FieldSecret holds the string denoting the secret vertex property in the database.
	FieldSecret      = "secret"      // FieldPermissions holds the string denoting the permissions vertex property in the database.
	FieldPermissions = "permissions" // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt   = "created_at"  // FieldLastUsed holds the string denoting the last_used vertex property in the database.
	FieldLastUsed    = "last_used"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the token in the database.
	Table = "tokens"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_tokens"
)

// Columns holds all SQL columns for token fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSecret,
	FieldPermissions,
	FieldCreatedAt,
	FieldLastUsed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Token type.
var ForeignKeys = []string{
	"user_tokens",
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	SecretValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
)

// Permissions defines the type for the permissions enum field.
type Permissions string

// PermissionsReadWrite is the default Permissions.
const DefaultPermissions = PermissionsReadWrite

// Permissions values.
const (
	PermissionsRead      Permissions = "Read"
	PermissionsWrite     Permissions = "Write"
	PermissionsReadWrite Permissions = "ReadWrite"
)

func (s Permissions) String() string {
	return string(s)
}

// PermissionsValidator is a validator for the "pe" field enum values. It is called by the builders before save.
func PermissionsValidator(pe Permissions) error {
	switch pe {
	case PermissionsRead, PermissionsWrite, PermissionsReadWrite:
		return nil
	default:
		return fmt.Errorf("token: invalid enum value for permissions field: %q", pe)
	}
}
