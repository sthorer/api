package schema

import (
	"time"

	"github.com/google/uuid"

	"github.com/facebookincubator/ent/schema/edge"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable(),
		field.String("name").
			MinLen(1).
			MaxLen(64).
			NotEmpty(),
		field.String("secret").
			NotEmpty().
			MinLen(40).
			MaxLen(80).
			Sensitive(),
		field.Enum("permissions").
			Immutable().
			Values("Read", "Write", "ReadWrite").
			Default("ReadWrite"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("last_used").
			Optional(),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tokens").
			Unique().
			Required(),
	}
}
