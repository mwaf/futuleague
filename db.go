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
	TEST_DB     = "./test.db"
)

var DB *sql.DB

func initDB(dbPath string) {
	dbFile, err := os.Open(dbPath)
	if os.IsNotExist(err) {
		fmt.Println("Creating Database at ", dbPath)
		createDB(dbPath)
	} else {
		dbFile.Close()
		openDB(dbPath)
		fmt.Println("Database found at ", dbPath)
	}
}

func openDB(dbPath string) {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	DB.Ping()
}

/*
Basically does the same as `sqlite3 dbPath < data/schema.sql && sqlite3 dbPath < data/seed.sql`
and has limitations (doesn't support multi-line SQL in schema and seed file) but we don't want
to require the user doing manual steps to get going.
*/
func createDB(dbPath string) {
	openDB(dbPath)
	execFile(SCHEMA_FILE)
	execFile(SEED_FILE)
}

func execFile(filePath string) {
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

		_, err := DB.Exec(sql)
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

func createTestDB() {
	os.Remove(TEST_DB)
	createDB(TEST_DB)
}
func removeTestDB() {
	os.Remove(TEST_DB)
}
