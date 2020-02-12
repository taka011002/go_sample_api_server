package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"path"
	"path/filepath"
)

func Up(db *sql.DB) {
	m := connect(db)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated")
}

func Down(db *sql.DB) {
	m := connect(db)

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated")
}

func connect(db *sql.DB) *migrate.Migrate {
	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	p, err := filepath.Abs(".")
	p = filepath.ToSlash(p)
	p = path.Join(p, "app/infra/migrations")

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", p),
		"mysql",
		driver)

	if err != nil {
		log.Fatal(err)
	}

	return m
}
