package db

import (
	"database/sql"
	"log"
	"os"

	"cognisafe.com/b/db/sqlc"
	_ "github.com/lib/pq"
)

func Instantiate(l *log.Logger) (*sql.DB, *sqlc.Queries) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	ssl := ""
	if DB_HOST == "localhost:5432" {
		ssl = "?sslmode=disable"
	}

	// connection_string := "postgresql://" + DB_NAME + "?host=" + DB_HOST + "&user=" + DB_USER + "&password=" + DB_PASS + ssl
	connection_string := "postgresql://" + DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + ssl

	db, err1 := sql.Open(os.Getenv("DB_DRIVER"), connection_string)

	if err1 != nil {
		l.Println("Error creating DB connection", err1)
		return nil, nil
	}

	err2 := db.Ping()
	if err2 != nil {
		l.Println("Error connecting to DB ", err2)
		return nil, nil
	}

	l.Println("🛢️  DB Connected")
	return db, sqlc.New(db)
}
