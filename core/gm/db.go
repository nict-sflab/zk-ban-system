package gm

import (
	"context"
	"errors"
	"fmt"

	"github.com/akakou/zk-ban-system/core/gm/ent"
	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	Type   string
	Config string
}

type DB struct {
	Client *ent.Client
	Ctx    *context.Context
}

var errOpenDB = errors.New("Failed to open database")
var errCreateSchema = errors.New("Failed to create schema")

func NewDB(dbConfig *DBConfig) (*DB, error) {
	client, err := ent.Open(dbConfig.Type, dbConfig.Config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errOpenDB, err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("%v: %w", errCreateSchema, err)
	}

	return &DB{
		Client: client,
		Ctx:    &ctx,
	}, nil
}

func (db *DB) Close() {
	db.Client.Close()
}
