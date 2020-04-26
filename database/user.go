package database

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/sthorer/api/ent"
	"github.com/sthorer/api/ent/user"
)

func (db *Database) UserLogin(ctx context.Context, email, password string) (*ent.User, error) {
	u, err := db.User.
		Query().
		Where(user.Email(strings.ToLower(email))).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}

	return u, nil
}

func (db *Database) UserRegister(ctx context.Context, email, password string) (*ent.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return db.User.
		Create().
		SetEmail(strings.ToLower(email)).
		SetPassword(string(hash)).
		Save(ctx)
}
