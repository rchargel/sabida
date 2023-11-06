package main

import (
	"embed"
	"log"

	"github.com/rchargel/sabida/models"

	"database/sql"

	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrations embed.FS

func connectToDatabase() *sql.DB {
	dbConfig := models.CreateDbConfig()
	connStr := dbConfig.GetConnectionStr()

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
