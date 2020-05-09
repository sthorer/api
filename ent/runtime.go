// github.com/sthorer/api

package ent

import (
	"time"

	"github.com/sthorer/api/ent/file"
	"github.com/sthorer/api/ent/schema"
	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescHash is the schema descriptor for hash field.
	fileDescHash := fileFields[1].Descriptor()
	// file.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	file.HashValidator = fileDescHash.Validators[0].(func(string) error)
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[2].Descriptor()
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int64) error)
	// fileDescPinnedAt is the schema descriptor for pinned_at field.
	fileDescPinnedAt := fileFields[3].Descriptor()
	// file.DefaultPinnedAt holds the default value on creation for the pinned_at field.
	file.DefaultPinnedAt = fileDescPinnedAt.Default.(func() time.Time)
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescName is the schema descriptor for name field.
	tokenDescName := tokenFields[1].Descriptor()
	// token.NameValidator is a validator for the "name" field. It is called by the builders before save.
	token.NameValidator = func() func(string) error {
		validators := tokenDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tokenDescSecret is the schema descriptor for secret field.
	tokenDescSecret := tokenFields[2].Descriptor()
	// token.SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	token.SecretValidator = func() func(string) error {
		validators := tokenDescSecret.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(secret string) error {
			for _, fn := range fns {
				if err := fn(secret); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenFields[4].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
			validators[3].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[2].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
