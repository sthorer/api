package database

import (
	"context"

	"github.com/google/uuid"

	"github.com/sthorer/api/ent"
	"github.com/sthorer/api/utils"
)

func (db *Database) NewToken(ctx context.Context, u *ent.User, name string) (*ent.Token, error) {
	secret, err := utils.GenerateSecret(40)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return db.Token.
		Create().
		SetID(id).
		SetSecret(secret).
		SetName(name).
		SetUser(u).
		Save(ctx)
}

func (db *Database) ResetToken(ctx context.Context, token *ent.Token) (*ent.Token, error) {
	secret, err := utils.GenerateSecret(40)
	if err != nil {
		return nil, err
	}

	return token.Update().
		SetSecret(secret).
		Save(ctx)
}

func (db *Database) RevokeToken(ctx context.Context, id uuid.UUID) error {
	return db.Client.Token.
		DeleteOneID(id).
		Exec(ctx)
}
