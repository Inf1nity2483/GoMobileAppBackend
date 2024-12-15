package db

import (
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func (d *DB) RunMigrations() {
	d.Connect()
	defer d.CloseConnection()

	driver, err := postgres.WithInstance(d.connection, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to apply migrations: %v", err)
	} else {
		log.Println("Migrations applied successfully!")
	}

	if err1, err2 := m.Close(); err1 != nil || err2 != nil {
		log.Fatalf("failed to close migrate instance: %v, %v", err1, err2)
	}
}
