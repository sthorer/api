package database

import (
	"context"
	"time"

	"github.com/sthorer/api/ent"
	"github.com/sthorer/api/utils"
)

func (db *Database) NewToken(ctx context.Context, user *ent.User, name string) (*ent.Token, error) {
	token, err := utils.GenerateSecret(40)
	if err != nil {
		return nil, err
	}

	return db.Token.
		Create().
		SetID(time.Now().Unix()).
		SetToken(token).
		SetName(name).
		SetUser(user).
		Save(ctx)
}

func (db *Database) RevokeToken(ctx context.Context, id int64) error {
	return db.Token.
		DeleteOneID(id).
		Exec(ctx)
}
