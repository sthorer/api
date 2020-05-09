package schema

import (
	"time"

	"github.com/facebookincubator/ent/schema/edge"

	"github.com/google/uuid"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable(),
		field.String("hash").
			Unique().
			Immutable().
			NotEmpty(),
		field.Int64("size").
			Positive(),
		field.Time("pinned_at").
			Immutable().
			Default(time.Now),
		field.Time("unpinned_at"),
		field.JSON("metadata", map[string]interface{}{}).
			Optional(),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("files").
			Unique().
			Required(),
	}
}
