package db

import (
	"database/sql"
	"fmt"
	cfg "myapp/config"

	_ "github.com/lib/pq"
)

type DB struct {
	connection *sql.DB
	db         *cfg.DB
}

func NewDB(db *cfg.Config) *DB {
	return &DB{
		db: &db.DB,
	}
}

func (d *DB) Connect() error {
	stmt := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.db.Host, d.db.Port, d.db.User, d.db.Password, d.db.Name)

	conn, err := sql.Open("postgres", stmt)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	d.connection = conn

	return nil
}

func (d *DB) CloseConnection() {
	_ = d.connection.Close()
}
