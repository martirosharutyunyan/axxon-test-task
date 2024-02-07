package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/config"
	"log"
)

// Instance is database model
var Instance *sql.DB

func ConnectDB() {
	dsn := config.GetDBDsn()

	db, _ := sql.Open("postgres", dsn)
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected to database")
	}
	Instance = db
}
