package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"regexp"
)

const (
	SCHEMA_FILE = "./data/schema.sql"
	SEED_FILE   = "./data/seed.sql"
)

func initDB(dbPath string) {
	dbFile, err := os.Open(dbPath)
	if os.IsNotExist(err) {
		fmt.Println("Creating Database at ", dbPath)
		createDB(dbPath)
	} else {
		fmt.Println("Database found at ", dbPath)
		dbFile.Close()
	}
}

/*
Basically does the same as `sqlite3 dbPath < data/schema.sql && sqlite3 dbPath < data/seed.sql`
and has limitations (doesn't support multi-line SQL in schema and seed file) but we don't want
to require the user doing manual steps to get going.
*/
func createDB(dbPath string) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	execFile(db, SCHEMA_FILE)
	execFile(db, SEED_FILE)
}

func execFile(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sql := scanner.Text()
		if isCommentOrEmptyLine(sql) {
			continue
		}

		_, err := db.Exec(sql)
		if err != nil {
			log.Printf("%q: %s\n", err, sql)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isCommentOrEmptyLine(sql string) bool {
	comment := regexp.MustCompile("^\\s*--.*$")
	emptyLine := regexp.MustCompile("^\\s*$")

	return comment.MatchString(sql) || emptyLine.MatchString(sql)
}
