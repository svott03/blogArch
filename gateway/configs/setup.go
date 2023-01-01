package configs

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB;

func SetupDB() {
	var connStr = "postgresql://" + EnvUsername() + ":" + EnvPassword() + "@localhost/blog?sslmode=disable";
	db_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB = db_
}