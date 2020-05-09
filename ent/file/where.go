// github.com/sthorer/api

package file

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/sthorer/api/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// PinnedAt applies equality check predicate on the "pinned_at" field. It's identical to PinnedAtEQ.
func PinnedAt(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPinnedAt), v))
	})
}

// UnpinnedAt applies equality check predicate on the "unpinned_at" field. It's identical to UnpinnedAtEQ.
func UnpinnedAt(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnpinnedAt), v))
	})
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSize), v))
	})
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSize), v...))
	})
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSize), v...))
	})
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSize), v))
	})
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSize), v))
	})
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSize), v))
	})
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSize), v))
	})
}

// PinnedAtEQ applies the EQ predicate on the "pinned_at" field.
func PinnedAtEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPinnedAt), v))
	})
}

// PinnedAtNEQ applies the NEQ predicate on the "pinned_at" field.
func PinnedAtNEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPinnedAt), v))
	})
}

// PinnedAtIn applies the In predicate on the "pinned_at" field.
func PinnedAtIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPinnedAt), v...))
	})
}

// PinnedAtNotIn applies the NotIn predicate on the "pinned_at" field.
func PinnedAtNotIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPinnedAt), v...))
	})
}

// PinnedAtGT applies the GT predicate on the "pinned_at" field.
func PinnedAtGT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPinnedAt), v))
	})
}

// PinnedAtGTE applies the GTE predicate on the "pinned_at" field.
func PinnedAtGTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPinnedAt), v))
	})
}

// PinnedAtLT applies the LT predicate on the "pinned_at" field.
func PinnedAtLT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPinnedAt), v))
	})
}

// PinnedAtLTE applies the LTE predicate on the "pinned_at" field.
func PinnedAtLTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPinnedAt), v))
	})
}

// UnpinnedAtEQ applies the EQ predicate on the "unpinned_at" field.
func UnpinnedAtEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnpinnedAt), v))
	})
}

// UnpinnedAtNEQ applies the NEQ predicate on the "unpinned_at" field.
func UnpinnedAtNEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnpinnedAt), v))
	})
}

// UnpinnedAtIn applies the In predicate on the "unpinned_at" field.
func UnpinnedAtIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUnpinnedAt), v...))
	})
}

// UnpinnedAtNotIn applies the NotIn predicate on the "unpinned_at" field.
func UnpinnedAtNotIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUnpinnedAt), v...))
	})
}

// UnpinnedAtGT applies the GT predicate on the "unpinned_at" field.
func UnpinnedAtGT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnpinnedAt), v))
	})
}

// UnpinnedAtGTE applies the GTE predicate on the "unpinned_at" field.
func UnpinnedAtGTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnpinnedAt), v))
	})
}

// UnpinnedAtLT applies the LT predicate on the "unpinned_at" field.
func UnpinnedAtLT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnpinnedAt), v))
	})
}

// UnpinnedAtLTE applies the LTE predicate on the "unpinned_at" field.
func UnpinnedAtLTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnpinnedAt), v))
	})
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMetadata)))
	})
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMetadata)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		p(s.Not())
	})
}
