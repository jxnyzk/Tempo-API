package main 

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DailyClaims = make(map[string]int)
)

func AddClaim(id string) {
	var currentCount int
	err := db.QueryRow("SELECT count FROM claims WHERE id = ?", id).Scan(&currentCount)
	if err == sql.ErrNoRows {
		_, err = db.Exec("INSERT INTO claims (id, count) VALUES (?, 1)", id)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		_, err = db.Exec("UPDATE claims SET count = count + 1 WHERE id = ?", id)
		if err != nil {
			log.Fatal(err)
		}
	}
	DailyClaims[id]++
	fmt.Println("[+] New claim for ", id)
}

func GetAllClaims() (map[string]int, error) {
	results := make(map[string]int)

	rows, err := db.Query("SELECT id, count FROM claims")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var count int
		if err := rows.Scan(&id, &count); err != nil {
			return nil, err
		}
		results[id] = count
	}

	return results, nil
}

