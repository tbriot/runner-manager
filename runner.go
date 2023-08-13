package main

import (
	"database/sql"
	"fmt"
	
	_ "github.com/mattn/go-sqlite3"
)

type runner struct {
	id        int
	firstname string
	lastname  string
	country   string
}

func addRunner(db *sql.DB, newRunner runner) {
	stmt, _ := db.Prepare("INSERT INTO runner (id, firstname, lastname, country) VALUES (?, ?, ?, ?)")
	stmt.Exec(nil, newRunner.firstname, newRunner.lastname, newRunner.country)
	defer stmt.Close()

	fmt.Printf("Add %v %v\n", newRunner.firstname, newRunner.lastname)
}

func getAllRunners(db *sql.DB) []runner {
	stmt, err := db.Prepare("SELECT id, firstname, lastname, country FROM runner")
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)

	var runners []runner

	for rows.Next() {
		var currentRunner runner
		rows.Scan(&currentRunner.id, &currentRunner.firstname, &currentRunner.lastname, &currentRunner.country)
		runners = append(runners, currentRunner)
	}
	fmt.Println("Select all runners")

	return runners
}


func deleteRunner(db *sql.DB, idToDelete int) int64 {
	stmt, err := db.Prepare("DELETE FROM runner where id = ?")
	checkError(err)
	defer stmt.Close()

	res, err := stmt.Exec(idToDelete)
	checkError(err)

	affected, err := res.RowsAffected()
	checkError(err)

	fmt.Printf("Delete runner with id=%d\n", idToDelete)

	return affected
}	
