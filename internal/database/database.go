package database

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		slog.Error("Ошибка при подключении БД", "error", err.Error())
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5)

	return db, nil
}
