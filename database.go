package main

import (
	"database/sql"
	"os"

	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func New(r bool) *Database {
	log.Debug("Connecting to database")

	db, err := sql.Open("postgres", connectionString.String())

	if err != nil {
		log.Errorf("Error connecting to database: %s", err.Error())
		os.Exit(1)
	}

	log.Debug("Connected to database")

	if r {
		log.Debug("Resetting database")
		reset(db)
		log.Debug("Database reset")
	}

	return &Database{
		db: db,
	}
}

func reset(db *sql.DB) {
	script, _ := os.ReadFile("reset.sql")
	sql := string(script)

	_, err := db.Exec(sql)

	if err != nil {
		log.Errorf("Error resetting database: %s", err.Error())
		os.Exit(1)
	}
}
