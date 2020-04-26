// github.com/sthorer/api

package predicate

import (
	"github.com/facebookincubator/ent/dialect/sql"
)

// Token is the predicate function for token builders.
type Token func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)
