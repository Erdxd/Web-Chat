package infrastructure

import (
	"database/sql"
)

var db *sql.DB

func InitDb(UrlDb string) (*sql.DB, error) {
	var err error
	PsqlInfo := UrlDb
	db, err = sql.Open("postrges", PsqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
