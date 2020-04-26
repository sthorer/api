package database

import (
	"context"
	"fmt"
	"os"

	"github.com/sthorer/api/ent"
)

type Database struct {
	*ent.Client
}

const (
	defaultDatabaseDriver = "sqlite3"
	defaultDatabaseURL    = "db.sqlite?_fk=1"
)

func Initialize() (*Database, error) {
	databaseDriver := os.Getenv("STHORER_DB_DRIVER")
	if databaseDriver == "" {
		databaseDriver = defaultDatabaseDriver
	}

	databaseURL := os.Getenv("STHORER_DB_URL")
	if databaseURL == "" {
		databaseURL = defaultDatabaseURL
	}

	client, err := ent.Open(databaseDriver, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v", databaseDriver, err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return &Database{Client: client}, err
}
