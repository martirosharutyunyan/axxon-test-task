package gooseDB

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/config"
	"github.com/pressly/goose/v3"
	"log"
)

func MigrationsUp() {
	dsn := config.GetDBDsn()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	dir := "../../pkg/migrations/"

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("migration-files: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.Up(db, dir); err != nil {
		log.Fatalf("migration-files run error: %v\n", err)
	}
}
