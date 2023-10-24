package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrations embed.FS

func connectToDatabase() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	scheme := os.Getenv("DB_SCHEME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		scheme,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print("Unable to connect to Postgres")
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}

func runMigrations(db *sql.DB) {
	goose.SetBaseFS(migrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Panic(err)
	}
}
