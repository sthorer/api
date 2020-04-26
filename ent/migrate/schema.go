// github.com/sthorer/api

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "token", Type: field.TypeString, Size: 80},
		{Name: "permissions", Type: field.TypeEnum, Enums: []string{"Read", "Write", "ReadWrite"}, Default: "ReadWrite"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_used", Type: field.TypeTime},
		{Name: "user_tokens", Type: field.TypeInt, Nullable: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "tokens_users_tokens",
				Columns: []*schema.Column{TokensColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "password", Type: field.TypeString, Size: 72},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "plan", Type: field.TypeEnum, Enums: []string{"Free", "Premium"}, Default: "Free"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TokensTable,
		UsersTable,
	}
)

func init() {
	TokensTable.ForeignKeys[0].RefTable = UsersTable
}
