package main

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := InitDB(os.Getenv("DATABASE_URL"))
	defer db.Close()

	router := mux.NewRouter()


	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

type Project struct {
	id uuid.UUID
	description string
	completed bool
}

type CurrentTask struct {
	id uuid.UUID
	task string

}

func InitDB(s string) *sql.DB {
	db, err := sql.Open("sqlite3", s)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	db = CreateProjectsTable(db)
	db = CreateCurrentsTable(db)
	db = CreateFuturesTable(db)
	db = CreateCompletedsTable(db)
	db = CreateBugsTable(db)


	return db
}

func CreateProjectsTable(db *sql.DB) *sql.DB {
	createProjectsStmt := `
	CREATE TABLE IF NOT EXISTS projects
	(
		id INTEGER PRIMARY KEY NOT NULL,
		description TEXT NOT NULL,
		completed BOOLEAN NOT NULL
	);
	`
	_, err := db.Exec(createProjectsStmt)
	if err != nil {
		log.Fatal(err)
	}


	return db
}

func CreateCurrentsTable(db *sql.DB) *sql.DB {
	createCurrentsStmt := `
	CREATE TABLE IF NOT EXISTS currents
	(
		id INTEGER PRIMARY KEY NOT NULL,
		project_id INTEGER NOT NULL,
		task TEXT NOT NULL
	);
	`
	_, err := db.Exec(createCurrentsStmt)
	if err != nil {
		log.Fatal(err)
	}


	return db
}

func CreateFuturesTable(db *sql.DB) *sql.DB {
	createFuturesStmt := `
	CREATE TABLE IF NOT EXISTS futures
	(
		id INTEGER PRIMARY KEY NOT NULL,
		project_id INTEGER NOT NULL,
		task TEXT NOT NULL
	);
	`
	_, err := db.Exec(createFuturesStmt)
	if err != nil {
		log.Fatal(err)
	}


	return db
}

func CreateCompletedsTable(db *sql.DB) *sql.DB {
	createCurrentsStmt := `
	CREATE TABLE IF NOT EXISTS completeds
	(
		id INTEGER PRIMARY KEY NOT NULL,
		project_id INTEGER NOT NULL,
		task TEXT NOT NULL
	);
	`
	_, err := db.Exec(createCurrentsStmt)
	if err != nil {
		log.Fatal(err)
	}


	return db
}

func CreateBugsTable(db *sql.DB) *sql.DB {
	createBugsStmt := `
	CREATE TABLE IF NOT EXISTS bugs
	(
		id INTEGER PRIMARY KEY NOT NULL,
		project_id INTEGER NOT NULL,
		task TEXT NOT NULL
	);
	`
	_, err := db.Exec(createBugsStmt)
	if err != nil {
		log.Fatal(err)
	}


	return db
}