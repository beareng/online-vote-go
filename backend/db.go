package backend

import (
	"database/sql"
)

type ElectionsDb struct {
	db *sql.DB
}

func (edb ElectionsDb) StartTransaction() (*ElectionsTx, error) {
	if tx, err := edb.db.Begin(); nil != err {
		return nil, err
	} else {
		return &ElectionsTx{tx: tx}, nil
	}
}

func ConnectDatabase(db *sql.DB) (ElectionsDb, error) {
	db.Exec(`PRAGMA foreign_keys = ON`)
	if _, err := db.Exec(`
CREATE TABLE IF NOT EXISTS user (
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT UNIQUE,
	token TEXT UNIQUE,
	siteadmin BOOLEAN NOT NULL DEFAULT 0
)
`); nil != err {
		return ElectionsDb{}, err
	}

	//TODO: more DB schema

	return ElectionsDb{
		db: db,
	}, nil
}
