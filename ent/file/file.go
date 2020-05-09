// github.com/sthorer/api

package file

import (
	"time"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID         = "id"          // FieldHash holds the string denoting the hash vertex property in the database.
	FieldHash       = "hash"        // FieldSize holds the string denoting the size vertex property in the database.
	FieldSize       = "size"        // FieldPinnedAt holds the string denoting the pinned_at vertex property in the database.
	FieldPinnedAt   = "pinned_at"   // FieldUnpinnedAt holds the string denoting the unpinned_at vertex property in the database.
	FieldUnpinnedAt = "unpinned_at" // FieldMetadata holds the string denoting the metadata vertex property in the database.
	FieldMetadata   = "metadata"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the file in the database.
	Table = "files"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "files"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_files"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldSize,
	FieldPinnedAt,
	FieldUnpinnedAt,
	FieldMetadata,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the File type.
var ForeignKeys = []string{
	"user_files",
}

var (
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func(string) error
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int64) error
	// DefaultPinnedAt holds the default value on creation for the pinned_at field.
	DefaultPinnedAt func() time.Time
)
