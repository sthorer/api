package schema

import (
	"time"

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
		field.Int64("id").
			Unique().
			Immutable().
			Positive(),
		field.String("name").
			MinLen(1).
			MaxLen(64).
			NotEmpty(),
		field.String("token").
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
		field.Time("last_used"),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tokens").
			Unique(),
	}
}
